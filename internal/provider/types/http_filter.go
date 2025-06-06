// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type HTTPFilter struct {
	JSONPatches []JSONPatches                `tfsdk:"json_patches"`
	Match       *MeshProxyPatchItemSpecMatch `tfsdk:"match"`
	Operation   types.String                 `tfsdk:"operation"`
	Value       types.String                 `tfsdk:"value"`
}
