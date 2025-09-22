package tests

import (
	"fmt"
	"github.com/Kong/shared-speakeasy/tfbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"net"
	"os"
	"testing"
)

type TestLogConsumer struct{}

func (g *TestLogConsumer) Accept(l testcontainers.Log) {
	fmt.Printf("cpLog: %s", l.Content)
}

func TestMesh(t *testing.T) {
	ctx := t.Context()
	req := testcontainers.ContainerRequest{
		Image:        "kong/kuma-cp:2.10.1",
		ExposedPorts: []string{"5681/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForLog("default AccessRoleBinding created"),
			wait.ForLog("default AccessRole created"),
			wait.ForLog("saving generated Admin User Token"),
			wait.ForListeningPort("5681/tcp"),
		),
		Cmd: []string{"run"},
	}
	if os.Getenv("RUNNER_DEBUG") == "1" {
		req.Cmd = []string{"run", "--log-level", "debug"}
		req.LogConsumerCfg = &testcontainers.LogConsumerConfig{
			Consumers: []testcontainers.LogConsumer{&TestLogConsumer{}},
		}
	}
	cpContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer testcontainers.CleanupContainer(t, cpContainer)
	port, err := cpContainer.MappedPort(ctx, "5681/tcp")
	require.NoError(t, err)

	t.Run("should create a mesh without initial policies", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("m0", "m0")

		// if this grows move this to shared-speakeasy
		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.AddMesh(mesh).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", mesh.ResourceName), plancheck.ResourceActionCreate),
						},
					},
					ExpectNonEmptyPlan: true, // skip_creating_initial_policies was set by the hook
				},
				{
					Config: builder.AddMesh(mesh.WithSpec(`skip_creating_initial_policies = [ "*" ]`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", mesh.ResourceName), plancheck.ResourceActionNoop),
						},
					},
				},
			},
		})
	})

	t.Run("create a mesh and modify fields on it", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("m1", "m1").
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)

		resource.ParallelTest(t, tfbuilder.CreateMeshAndModifyFieldsOnIt(providerFactory, builder, mesh))
	})

	t.Run("create a policy and modify fields on it", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("default", "terraform-provider-kong-mesh").
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)
		mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", "allow-all", "MeshTrafficPermission").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		builder.AddMesh(mesh)

		resource.ParallelTest(t, tfbuilder.CreatePolicyAndModifyFieldsOnIt(providerFactory, builder, mtp))
	})

	t.Run("not imported resource should error out with meaningful message", func(t *testing.T) {
		meshName := "m3"
		mtpName := "allow-all"

		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("default", meshName).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)
		mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", mtpName, "MeshTrafficPermission").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		builder.AddMesh(mesh)

		resource.ParallelTest(t, tfbuilder.NotImportedResourceShouldErrorOutWithMeaningfulMessage(providerFactory, builder, mtp, func() { createAnMTP(t, "http://"+net.JoinHostPort("localhost", port.Port()), meshName, mtpName) }))
	})

}

func createAnMTP(t *testing.T, url string, meshName string, mtpName string) {
	ctx := t.Context()
	client := sdk.New(sdk.WithServerURL(url))
	action := shared.ActionAllow
	resp, err := client.MeshTrafficPermission.CreateMeshTrafficPermission(ctx, operations.CreateMeshTrafficPermissionRequest{
		Mesh: meshName,
		Name: mtpName,
		MeshTrafficPermissionItem: shared.MeshTrafficPermissionItemInput{
			Mesh: &meshName,
			Name: mtpName,
			Type: shared.MeshTrafficPermissionItemTypeMeshTrafficPermission,
			Spec: shared.MeshTrafficPermissionItemSpec{
				From: []shared.MeshTrafficPermissionItemFrom{
					{
						TargetRef: shared.MeshTrafficPermissionItemSpecTargetRef{Kind: shared.MeshTrafficPermissionItemSpecKindMesh},
						Default:   &shared.MeshTrafficPermissionItemDefault{Action: &action},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
}
