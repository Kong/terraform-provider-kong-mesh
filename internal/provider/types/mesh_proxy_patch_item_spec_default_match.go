// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MeshProxyPatchItemSpecDefaultMatch struct {
	Name   types.String            `tfsdk:"name"`
	Origin types.String            `tfsdk:"origin"`
	Tags   map[string]types.String `tfsdk:"tags"`
}
