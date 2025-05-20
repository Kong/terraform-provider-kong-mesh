package tests

import (
	"github.com/Kong/shared-speakeasy/tfbuilder"
	"github.com/docker/go-connections/nat"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"net"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

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
	cpContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer testcontainers.CleanupContainer(t, cpContainer)
	port, err := cpContainer.MappedPort(ctx, "5681/tcp")
	require.NoError(t, err)

	t.Run("create a mesh and modify fields on it", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("m1", "m1")

		resource.ParallelTest(t, tfbuilder.CreateMeshAndModifyFieldsOnIt(providerFactory, builder, mesh))
	})

	t.Run("create a policy and modify fields on it", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("default", "terraform-provider-kong-mesh")
		mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", "allow-all", "MeshTrafficPermission").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName)).
			WithLabels(map[string]string{
				"kuma.io/mesh":   mesh.MeshName,
				"kuma.io/env":    "universal",
				"kuma.io/origin": "zone",
				"kuma.io/zone":   "default",
			})
		builder.AddMesh(mesh)

		resource.ParallelTest(t, tfbuilder.CreatePolicyAndModifyFieldsOnIt(providerFactory, builder, mtp))
	})

	t.Run("not imported resource should error out with meaningful message", func(t *testing.T) {
		meshName := "m3"
		mtpName := "allow-all"

		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("default", meshName)
		mtp := tfbuilder.NewPolicyBuilder("mesh_traffic_permission", "allow_all", mtpName, "MeshTrafficPermission").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName)).
			WithLabels(map[string]string{
				"kuma.io/mesh":   mesh.MeshName,
				"kuma.io/env":    "universal",
				"kuma.io/origin": "zone",
				"kuma.io/zone":   "default",
			})
		builder.AddMesh(mesh)

		expectedErr, err := regexp.Compile(`MeshTrafficPermission already exists`)
		require.NoError(t, err)
		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Build(),
				},
				{
					PreConfig: func() {
						createAnMTP(t, port, meshName, mtpName)
					},
					Config:      builder.AddPolicy(mtp.WithSpec(tfbuilder.AllowAllTrafficPermissionSpec)).Build(),
					ExpectError: expectedErr,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress(mtp.ResourceType, mtp.ResourceName), plancheck.ResourceActionCreate),
						},
					},
				},
			},
		})
	})

	if t.Failed() {
		logs, err := cpContainer.Logs(ctx)
		require.NoError(t, err)
		defer logs.Close()
		logContent, err := io.ReadAll(logs)
		require.NoError(t, err)
		t.Logf("Container logs: %s", logContent)
	}
}

func createAnMTP(t *testing.T, port nat.Port, meshName string, mtpName string) {
	ctx := t.Context()
	client := sdk.New("http://" + net.JoinHostPort("localhost", port.Port()))
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
