// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type MeshLoadBalancingStrategyItemSpec struct {
	TargetRef *MeshAccessLogItemTargetRef       `tfsdk:"target_ref"`
	To        []MeshLoadBalancingStrategyItemTo `tfsdk:"to"`
}
