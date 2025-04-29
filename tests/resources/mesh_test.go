package tests

import (
	"github.com/Kong/shared-speakeasy/tfbuilder"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
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
		mesh := tfbuilder.NewMeshBuilder("m2", "m2")
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

	meshRetryAllSpec := `
  spec = {
    target_ref = {
      kind = "Mesh"
    }
    to = [
      {
        target_ref = {
          kind = "Mesh"
        }
        default = {
          grpc = {
            num_retries = 5
            per_try_timeout = "16s"
          }
        }
      }
    ]
  }
`

	// Check for fix "provider produced invalid plan for retry_on"
	t.Run("create mesh_retry with retry_on", func(t *testing.T) {
		builder := tfbuilder.NewBuilder(tfbuilder.KongMesh, "http", "localhost", port.Int())
		mesh := tfbuilder.NewMeshBuilder("m3", "m3")
		mr := tfbuilder.NewPolicyBuilder("mesh_retry", "retry_on", "retry-on", "MeshRetry").
			WithMeshRef(builder.ResourceAddress("mesh", mesh.ResourceName) + ".name").
			WithDependsOn(builder.ResourceAddress("mesh", mesh.ResourceName)).
			WithSpec(meshRetryAllSpec).
			WithLabels(map[string]string{
				"kuma.io/mesh":   mesh.MeshName,
				"kuma.io/env":    "universal",
				"kuma.io/origin": "zone",
				"kuma.io/zone":   "default",
			})
		builder.AddMesh(mesh)

		resource.ParallelTest(t, resource.TestCase{
			ProtoV6ProviderFactories: providerFactory,
			Steps: []resource.TestStep{
				{
					Config: builder.Build(),
				},
				{
					Config: builder.AddPolicy(mr).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress(mr.ResourceType, mr.ResourceName), plancheck.ResourceActionCreate),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
				{
					Config: builder.AddPolicy(mr.AddToSpec(`per_try_timeout = "16s"`, `retry_on = []`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress(mr.ResourceType, mr.ResourceName), plancheck.ResourceActionNoop),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
				{
					Config: builder.AddPolicy(mr.UpdateSpec(`retry_on = []`, `retry_on = ["Canceled"]`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress(mr.ResourceType, mr.ResourceName), plancheck.ResourceActionUpdate),
							plancheck.ExpectKnownValue(builder.ResourceAddress(mr.ResourceType, mr.ResourceName), tfjsonpath.New("spec").AtMapKey("to").AtSliceIndex(0).AtMapKey("default").AtMapKey("grpc").AtMapKey("retry_on"), knownvalue.ListExact([]knownvalue.Check{knownvalue.StringExact("Canceled")})),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
				{
					Config: builder.AddPolicy(mr.RemoveFromSpec(`retry_on = ["Canceled"]`)).Build(),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectResourceAction(builder.ResourceAddress(mr.ResourceType, mr.ResourceName), plancheck.ResourceActionUpdate),
							plancheck.ExpectKnownValue(builder.ResourceAddress(mr.ResourceType, mr.ResourceName), tfjsonpath.New("spec").AtMapKey("to").AtSliceIndex(0).AtMapKey("default").AtMapKey("grpc").AtMapKey("retry_on"), knownvalue.Null()),
						},
					},
				},
				tfbuilder.CheckReapplyPlanEmpty(builder),
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
