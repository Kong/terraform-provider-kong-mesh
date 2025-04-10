package listplanmodifier

import (
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

    shared_speakeasy "github.com/Kong/shared-speakeasy/planmodifiers/suppress_zero_null/listplanmodifier"
)

func SupressZeroNullModifier() planmodifier.List {
    return shared_speakeasy.SupressZeroNullModifier()
}
