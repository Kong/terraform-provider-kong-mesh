// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type MeshRetryItemGrpc struct {
	BackOff            *BackOff            `tfsdk:"back_off"`
	NumRetries         types.Int32         `tfsdk:"num_retries"`
	PerTryTimeout      types.String        `tfsdk:"per_try_timeout"`
	RateLimitedBackOff *RateLimitedBackOff `tfsdk:"rate_limited_back_off"`
	RetryOn            []types.String      `tfsdk:"retry_on"`
}
