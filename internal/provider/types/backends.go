// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Backends struct {
	Conf   *MeshItemLoggingConf `tfsdk:"conf"`
	Format types.String         `tfsdk:"format"`
	Name   types.String         `tfsdk:"name"`
	Type   types.String         `tfsdk:"type"`
}
