package tests

import (
	"fmt"
	"net"
	"os"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/Kong/shared-speakeasy/tfbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	resourcesshared "github.com/kong/terraform-provider-kong-mesh/tests/resources/shared"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
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
		Env: map[string]string{
			"KUMA_MODE": "global",
		},
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
		serverURL := fmt.Sprintf("http://localhost:%d", port.Int())
		builder := hclbuilder.NewTestBuilder(hclbuilder.KongMesh)
		builder.SetAttribute("provider.kong-mesh", "server_url", serverURL)

		meshName := "m0"
		meshResourceName := "m0"

		// if this grows move this to shared-speakeasy
		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.AddMeshFromHCL(meshName, meshResourceName, "").Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", meshResourceName), plancheck.ResourceActionCreate),
						},
					},
					ExpectNonEmptyPlan: true, // skip_creating_initial_policies was set by the hook
				},
				{
					Config: builder.AddMeshFromHCL(meshName, meshResourceName, `skip_creating_initial_policies = [ "*" ]`).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", meshResourceName), plancheck.ResourceActionNoop),
						},
					},
				},
			},
		})
	})

	t.Run("create a mesh and modify fields on it", func(t *testing.T) {
		serverURL := fmt.Sprintf("http://localhost:%d", port.Int())
		meshConfig := hclbuilder.MeshConfig{
			MeshName:     "m1",
			ResourceName: "m1",
			Provider:     hclbuilder.KongMesh,
			ServerURL:    serverURL,
		}

		resource.ParallelTest(t, hclbuilder.CreateMeshAndModifyFields(providerFactory, meshConfig))
	})

	t.Run("create a policy and modify fields on it", func(t *testing.T) {
		serverURL := fmt.Sprintf("http://localhost:%d", port.Int())
		policyConfig := hclbuilder.PolicyConfig{
			PolicyType:   "mesh_traffic_permission",
			PolicyName:   "allow-all",
			ResourceName: "allow_all",
			MeshRef:      "default",
			Provider:     hclbuilder.KongMesh,
			ServerURL:    serverURL,
		}

		resource.ParallelTest(t, hclbuilder.CreatePolicyAndModifyFields(providerFactory, policyConfig))
	})

	t.Run("not imported resource should error out with meaningful message", func(t *testing.T) {
		meshName := "policy-test-mesh-2"
		mtpName := "allow-all"
		serverURL := fmt.Sprintf("http://localhost:%d", port.Int())

		policyConfig := hclbuilder.PolicyConfig{
			PolicyType:   "mesh_traffic_permission",
			PolicyName:   mtpName,
			ResourceName: "allow_all",
			MeshRef:      meshName,
			Provider:     hclbuilder.KongMesh,
			ServerURL:    serverURL,
		}

		resource.ParallelTest(t, hclbuilder.NotImportedResourceShouldError(providerFactory, policyConfig, func() { createAnMTP(t, "http://"+net.JoinHostPort("localhost", port.Port()), meshName, mtpName) }))
	})

	t.Run("should be able to store secrets", func(t *testing.T) {
		// This test uses tfbuilder because the shared test function hasn't been migrated to hclbuilder yet
		meshName := "m4"

		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("default", meshName).
			WithSpec(`skip_creating_initial_policies = [ "*" ]`)
		skey := tfbuilder.NewPolicyBuilder("mesh_secret", "skey", "skey", "Secret").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		scert := tfbuilder.NewPolicyBuilder("mesh_secret", "scert", "scert", "Secret").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName))
		builder.AddMesh(mesh)
		resource.ParallelTest(t, resourcesshared.ShouldBeAbleToStoreSecrets(providerFactory, builder, scert, skey, mesh))
	})
}

func createAnMTP(t *testing.T, url string, meshName string, mtpName string) {
	ctx := t.Context()
	opts := []sdk.SDKOption{
		sdk.WithServerURL(url),
	}
	client := sdk.New(opts...)
	action := shared.ActionAllow
	resp, err := client.MeshTrafficPermission.PutMeshTrafficPermission(ctx, operations.PutMeshTrafficPermissionRequest{
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
	require.True(t, resp.StatusCode == 200 || resp.StatusCode == 201, "Expected 200 or 201, got %d", resp.StatusCode)
}
