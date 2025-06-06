// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Matches struct {
	Headers     []Headers     `tfsdk:"headers"`
	Method      types.String  `tfsdk:"method"`
	Path        *Path         `tfsdk:"path"`
	QueryParams []QueryParams `tfsdk:"query_params"`
}
