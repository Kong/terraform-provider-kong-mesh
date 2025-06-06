// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MeshHealthCheckItemHTTP struct {
	Disabled            types.Bool                      `tfsdk:"disabled"`
	ExpectedStatuses    []types.Int64                   `tfsdk:"expected_statuses"`
	Path                types.String                    `tfsdk:"path"`
	RequestHeadersToAdd *MeshGlobalRateLimitItemHeaders `tfsdk:"request_headers_to_add"`
}
