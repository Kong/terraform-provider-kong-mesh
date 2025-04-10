package stringplanmodifier

import (
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

    shared_speakeasy "github.com/Kong/shared-speakeasy/planmodifiers/arbitrary_json/stringplanmodifier"
)

func ArbitraryJSONModifier() planmodifier.String {
    return shared_speakeasy.ArbitraryJSONModifier()
}
