package shared

import (
	"github.com/Kong/shared-speakeasy/tfbuilder"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func ShouldBeAbleToStoreSecrets(providerFactory map[string]func() (tfprotov6.ProviderServer, error), builder *tfbuilder.Builder, secret *tfbuilder.PolicyBuilder, mesh *tfbuilder.MeshBuilder) resource.TestCase {
	return resource.TestCase{
		ProtoV6ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: builder.Build(),
			},
			{
				Config: builder.AddPolicy(secret.WithSpec("data = \"dGVzdAo=\"")).Build(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(builder.ResourceAddress(secret.ResourceType, secret.ResourceName), plancheck.ResourceActionCreate),
					},
				},
			},
			{
				Config: builder.AddMesh(mesh.WithSpec(`skip_creating_initial_policies = [ "*" ]
  mtls = {
    backends = [
      {
        conf = {
          provided_certificate_authority_config = {
            cert = {
              data_source_secret = {
                secret = "mysecret"
              }
            }
            key = {
              data_source_secret = {
                secret = "mysecret"
              }
            }
          }
        }
        name = "provided-1"
        type = "provided"
      }
    ]
    enabled_backend = "provided-1"
  }
`)).Build(),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectResourceAction(builder.ResourceAddress("mesh", mesh.ResourceName), plancheck.ResourceActionUpdate),
					},
				},
			},
			{
				Config: builder.AddMesh(mesh.WithSpec(`skip_creating_initial_policies = [ "*" ]`)).Build(),
			},
			tfbuilder.CheckReapplyPlanEmpty(builder),
		},
	}
}
