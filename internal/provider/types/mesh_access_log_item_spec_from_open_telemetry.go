// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshAccessLogItemSpecFromOpenTelemetry struct {
	Attributes []JSON       `tfsdk:"attributes"`
	Body       types.String `tfsdk:"body"`
	Endpoint   types.String `tfsdk:"endpoint"`
}
