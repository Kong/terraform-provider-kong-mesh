// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Applications struct {
	Address types.String `tfsdk:"address"`
	Name    types.String `tfsdk:"name"`
	Path    types.String `tfsdk:"path"`
	Port    types.Int32  `tfsdk:"port"`
}
