package objectplanmodifier

import (
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

    shared_speakeasy "github.com/Kong/shared-speakeasy/planmodifiers/suppress_zero_null/objectplanmodifier"
)

func SupressZeroNullModifier() planmodifier.Object {
    return shared_speakeasy.SupressZeroNullModifier()
}
