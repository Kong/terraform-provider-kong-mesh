// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Addresses struct {
	Hostname             types.String          `tfsdk:"hostname"`
	HostnameGeneratorRef *HostnameGeneratorRef `tfsdk:"hostname_generator_ref"`
	Origin               types.String          `tfsdk:"origin"`
}
