package shared

import (
	"github.com/Kong/shared-speakeasy/tfbuilder"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func ShouldBeAbleToStoreSecrets(
	providerFactory map[string]func() (tfprotov6.ProviderServer, error),
	builder tfbuilder.ModifyPolicyBuilder,
	secret *tfbuilder.PolicyBuilder,
) resource.TestCase {
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
		},
	}
}
