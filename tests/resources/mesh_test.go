package tests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"testing"

	"github.com/Kong/shared-speakeasy/hclbuilder"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	tcexec "github.com/testcontainers/testcontainers-go/exec"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestLogConsumer struct{}

func (g *TestLogConsumer) Accept(l testcontainers.Log) {
	fmt.Printf("cpLog: %s", l.Content)
}

func TestMesh(t *testing.T) {
	ctx := t.Context()
	req := testcontainers.ContainerRequest{
		Image:        "kong/kuma-cp:3.0.0-preview.v6c3b256d2",
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
	token := adminToken(t, cpContainer)

	t.Run("create a mesh and modify labels on it", func(t *testing.T) {
		serverURL := fmt.Sprintf("http://localhost:%d", port.Num())
		builder := newBuilder(serverURL, token)

		meshName := "m1"
		meshResourceName := "m1"

		mesh, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh" "%s" {
  type = "Mesh"
  name = "%s"
}
`, meshResourceName, meshName))

		resource.ParallelTest(t, hclbuilder.CreateMeshAndModifyLabels(providerFactory, builder, mesh))
	})

	t.Run("create a policy and modify fields on it", func(t *testing.T) {
		serverURL := fmt.Sprintf("http://localhost:%d", port.Num())
		builder := newBuilder(serverURL, token)

		meshName := "policy-test-mesh"
		meshResourceName := "test_mesh"

		mesh, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh" "%s" {
  type = "Mesh"
  name = "%s"
}
`, meshResourceName, meshName))

		policyResourceName := "allow_all"
		policyName := "allow-all"

		policy, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh_traffic_permission" "%s" {
  type = "MeshTrafficPermission"
  name = "%s"
  mesh = "%s"
}
`, policyResourceName, policyName, meshName))

		resource.ParallelTest(t, hclbuilder.CreatePolicyWithRulesAndModifyFields(providerFactory, builder, mesh, policy))
	})

	t.Run("not imported resource should error out with meaningful message", func(t *testing.T) {
		meshName := "policy-test-mesh-2"
		meshResourceName := "test_mesh"
		mtpName := "allow-all"
		serverURL := fmt.Sprintf("http://localhost:%d", port.Num())

		builder := newBuilder(serverURL, token)

		mesh, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh" "%s" {
  type = "Mesh"
  name = "%s"
}
`, meshResourceName, meshName))

		policyResourceName := "allow_all"

		policy, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh_traffic_permission" "%s" {
  type = "MeshTrafficPermission"
  name = "%s"
  mesh = "%s"
}
`, policyResourceName, mtpName, meshName))

		resource.ParallelTest(t, hclbuilder.NotImportedResourceWithRulesShouldError(providerFactory, builder, mesh, policy, func() { createAnMTP(t, "http://"+net.JoinHostPort("localhost", port.Port()), token, meshName, mtpName) }))
	})

	t.Run("should be able to store secrets", func(t *testing.T) {
		meshName := "m4"
		meshResourceName := "test_mesh"
		serverURL := fmt.Sprintf("http://localhost:%d", port.Num())

		builder := newBuilder(serverURL, token)

		mesh, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh" "%s" {
  type = "Mesh"
  name = "%s"
}
`, meshResourceName, meshName))

		scertResourceName := "scert"
		scertName := "scert"

		scert, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh_secret" "%s" {
  type = "Secret"
  name = "%s"
  mesh = "%s"
}
`, scertResourceName, scertName, meshName))

		skeyResourceName := "skey"
		skeyName := "skey"

		skey, _ := hclbuilder.FromString(fmt.Sprintf(`
resource "kong-mesh_mesh_secret" "%s" {
  type = "Secret"
  name = "%s"
  mesh = "%s"
}
`, skeyResourceName, skeyName, meshName))

		resource.ParallelTest(t, hclbuilder.ShouldBeAbleToStoreAndUpdateSecrets(providerFactory, builder, mesh, scert, skey))
	})
}

// adminToken reads the admin user token from inside the container,
// because only requests coming from the CP's loopback are treated as admin.
func adminToken(t *testing.T, c testcontainers.Container) string {
	t.Helper()
	code, out, err := c.Exec(t.Context(), []string{"wget", "-qO-", "http://localhost:5681/global-secrets/admin-user-token"}, tcexec.Multiplexed())
	require.NoError(t, err)
	body, err := io.ReadAll(out)
	require.NoError(t, err)
	require.Equal(t, 0, code, string(body))
	var secret struct {
		Data string `json:"data"`
	}
	require.NoError(t, json.Unmarshal(body, &secret))
	token, err := base64.StdEncoding.DecodeString(secret.Data)
	require.NoError(t, err)
	return string(token)
}

func newBuilder(serverURL string, token string) *hclbuilder.Builder {
	builder := hclbuilder.NewWithProvider(hclbuilder.KongMesh, serverURL)
	builder.SetAttribute(fmt.Sprintf("provider.%s.bearer_auth", hclbuilder.KongMesh), token)
	return builder
}

func createAnMTP(t *testing.T, url string, token string, meshName string, mtpName string) {
	ctx := t.Context()
	opts := []sdk.SDKOption{
		sdk.WithServerURL(url),
		sdk.WithSecurity(shared.Security{BearerAuth: &token}),
	}
	client := sdk.New(opts...)
	resp, err := client.MeshTrafficPermission.PutMeshTrafficPermission(ctx, operations.PutMeshTrafficPermissionRequest{
		Mesh: meshName,
		Name: mtpName,
		MeshTrafficPermissionItem: shared.MeshTrafficPermissionItemInput{
			Mesh: &meshName,
			Name: mtpName,
			Type: shared.MeshTrafficPermissionItemTypeMeshTrafficPermission,
			Spec: shared.MeshTrafficPermissionItemSpec{
				Rules: []shared.MeshTrafficPermissionItemRules{
					{
						Default: shared.MeshTrafficPermissionItemDefault{
							Allow: []shared.Allow{
								{
									SpiffeID: &shared.MeshTrafficPermissionItemSpiffeID{
										Type:  shared.MeshTrafficPermissionItemSpecRulesTypePrefix,
										Value: "spiffe://example.org",
									},
								},
							},
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
}
