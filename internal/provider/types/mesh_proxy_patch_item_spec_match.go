// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MeshProxyPatchItemSpecMatch struct {
	ListenerName types.String            `tfsdk:"listener_name"`
	ListenerTags map[string]types.String `tfsdk:"listener_tags"`
	Name         types.String            `tfsdk:"name"`
	Origin       types.String            `tfsdk:"origin"`
}
