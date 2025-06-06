// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"github.com/Kong/shared-speakeasy/customtypes/kumalabels"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/provider/typeconvert"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
)

func (r *MeshCircuitBreakerListDataSourceModel) ToOperationsGetMeshCircuitBreakerListRequest(ctx context.Context) (*operations.GetMeshCircuitBreakerListRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	offset := new(int64)
	if !r.Offset.IsUnknown() && !r.Offset.IsNull() {
		*offset = r.Offset.ValueInt64()
	} else {
		offset = nil
	}
	size := new(int64)
	if !r.Size.IsUnknown() && !r.Size.IsNull() {
		*size = r.Size.ValueInt64()
	} else {
		size = nil
	}
	var mesh string
	mesh = r.Mesh.ValueString()

	out := operations.GetMeshCircuitBreakerListRequest{
		Offset: offset,
		Size:   size,
		Mesh:   mesh,
	}

	return &out, diags
}

func (r *MeshCircuitBreakerListDataSourceModel) RefreshFromSharedMeshCircuitBreakerList(ctx context.Context, resp *shared.MeshCircuitBreakerList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Items = []tfTypes.MeshCircuitBreakerItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items tfTypes.MeshCircuitBreakerItem
			items.CreationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.CreationTime))
			labelsValue, labelsDiags := types.MapValueFrom(ctx, types.StringType, itemsItem.Labels)
			diags.Append(labelsDiags...)
			labelsValuable, labelsDiags := kumalabels.KumaLabelsMapType{MapType: types.MapType{ElemType: types.StringType}}.ValueFromMap(ctx, labelsValue)
			diags.Append(labelsDiags...)
			items.Labels, _ = labelsValuable.(kumalabels.KumaLabelsMapValue)
			items.Mesh = types.StringPointerValue(itemsItem.Mesh)
			items.ModificationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.ModificationTime))
			items.Name = types.StringValue(itemsItem.Name)
			items.Spec.From = []tfTypes.MeshCircuitBreakerItemFrom{}
			for fromCount, fromItem := range itemsItem.Spec.From {
				var from tfTypes.MeshCircuitBreakerItemFrom
				if fromItem.Default == nil {
					from.Default = nil
				} else {
					from.Default = &tfTypes.MeshCircuitBreakerItemDefault{}
					if fromItem.Default.ConnectionLimits == nil {
						from.Default.ConnectionLimits = nil
					} else {
						from.Default.ConnectionLimits = &tfTypes.ConnectionLimits{}
						from.Default.ConnectionLimits.MaxConnectionPools = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.ConnectionLimits.MaxConnectionPools))
						from.Default.ConnectionLimits.MaxConnections = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.ConnectionLimits.MaxConnections))
						from.Default.ConnectionLimits.MaxPendingRequests = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.ConnectionLimits.MaxPendingRequests))
						from.Default.ConnectionLimits.MaxRequests = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.ConnectionLimits.MaxRequests))
						from.Default.ConnectionLimits.MaxRetries = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.ConnectionLimits.MaxRetries))
					}
					if fromItem.Default.OutlierDetection == nil {
						from.Default.OutlierDetection = nil
					} else {
						from.Default.OutlierDetection = &tfTypes.OutlierDetection{}
						from.Default.OutlierDetection.BaseEjectionTime = types.StringPointerValue(fromItem.Default.OutlierDetection.BaseEjectionTime)
						if fromItem.Default.OutlierDetection.Detectors == nil {
							from.Default.OutlierDetection.Detectors = nil
						} else {
							from.Default.OutlierDetection.Detectors = &tfTypes.Detectors{}
							if fromItem.Default.OutlierDetection.Detectors.FailurePercentage == nil {
								from.Default.OutlierDetection.Detectors.FailurePercentage = nil
							} else {
								from.Default.OutlierDetection.Detectors.FailurePercentage = &tfTypes.FailurePercentage{}
								from.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts))
								from.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume))
								from.Default.OutlierDetection.Detectors.FailurePercentage.Threshold = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.FailurePercentage.Threshold))
							}
							if fromItem.Default.OutlierDetection.Detectors.GatewayFailures == nil {
								from.Default.OutlierDetection.Detectors.GatewayFailures = nil
							} else {
								from.Default.OutlierDetection.Detectors.GatewayFailures = &tfTypes.GatewayFailures{}
								from.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive))
							}
							if fromItem.Default.OutlierDetection.Detectors.LocalOriginFailures == nil {
								from.Default.OutlierDetection.Detectors.LocalOriginFailures = nil
							} else {
								from.Default.OutlierDetection.Detectors.LocalOriginFailures = &tfTypes.GatewayFailures{}
								from.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive))
							}
							if fromItem.Default.OutlierDetection.Detectors.SuccessRate == nil {
								from.Default.OutlierDetection.Detectors.SuccessRate = nil
							} else {
								from.Default.OutlierDetection.Detectors.SuccessRate = &tfTypes.SuccessRate{}
								from.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts))
								from.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume))
								if fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor != nil {
									from.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor = &tfTypes.Mode{}
									if fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer != nil {
										from.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer = types.Int64PointerValue(fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer)
									}
									if fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str != nil {
										from.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str = types.StringPointerValue(fromItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str)
									}
								}
							}
							if fromItem.Default.OutlierDetection.Detectors.TotalFailures == nil {
								from.Default.OutlierDetection.Detectors.TotalFailures = nil
							} else {
								from.Default.OutlierDetection.Detectors.TotalFailures = &tfTypes.GatewayFailures{}
								from.Default.OutlierDetection.Detectors.TotalFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.Detectors.TotalFailures.Consecutive))
							}
						}
						from.Default.OutlierDetection.Disabled = types.BoolPointerValue(fromItem.Default.OutlierDetection.Disabled)
						if fromItem.Default.OutlierDetection.HealthyPanicThreshold != nil {
							from.Default.OutlierDetection.HealthyPanicThreshold = &tfTypes.Mode{}
							if fromItem.Default.OutlierDetection.HealthyPanicThreshold.Integer != nil {
								from.Default.OutlierDetection.HealthyPanicThreshold.Integer = types.Int64PointerValue(fromItem.Default.OutlierDetection.HealthyPanicThreshold.Integer)
							}
							if fromItem.Default.OutlierDetection.HealthyPanicThreshold.Str != nil {
								from.Default.OutlierDetection.HealthyPanicThreshold.Str = types.StringPointerValue(fromItem.Default.OutlierDetection.HealthyPanicThreshold.Str)
							}
						}
						from.Default.OutlierDetection.Interval = types.StringPointerValue(fromItem.Default.OutlierDetection.Interval)
						from.Default.OutlierDetection.MaxEjectionPercent = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(fromItem.Default.OutlierDetection.MaxEjectionPercent))
						from.Default.OutlierDetection.SplitExternalAndLocalErrors = types.BoolPointerValue(fromItem.Default.OutlierDetection.SplitExternalAndLocalErrors)
					}
				}
				from.TargetRef.Kind = types.StringValue(string(fromItem.TargetRef.Kind))
				if len(fromItem.TargetRef.Labels) > 0 {
					from.TargetRef.Labels = make(map[string]types.String, len(fromItem.TargetRef.Labels))
					for key, value := range fromItem.TargetRef.Labels {
						from.TargetRef.Labels[key] = types.StringValue(value)
					}
				}
				from.TargetRef.Mesh = types.StringPointerValue(fromItem.TargetRef.Mesh)
				from.TargetRef.Name = types.StringPointerValue(fromItem.TargetRef.Name)
				from.TargetRef.Namespace = types.StringPointerValue(fromItem.TargetRef.Namespace)
				from.TargetRef.ProxyTypes = make([]types.String, 0, len(fromItem.TargetRef.ProxyTypes))
				for _, v := range fromItem.TargetRef.ProxyTypes {
					from.TargetRef.ProxyTypes = append(from.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				from.TargetRef.SectionName = types.StringPointerValue(fromItem.TargetRef.SectionName)
				if len(fromItem.TargetRef.Tags) > 0 {
					from.TargetRef.Tags = make(map[string]types.String, len(fromItem.TargetRef.Tags))
					for key1, value1 := range fromItem.TargetRef.Tags {
						from.TargetRef.Tags[key1] = types.StringValue(value1)
					}
				}
				if fromCount+1 > len(items.Spec.From) {
					items.Spec.From = append(items.Spec.From, from)
				} else {
					items.Spec.From[fromCount].Default = from.Default
					items.Spec.From[fromCount].TargetRef = from.TargetRef
				}
			}
			items.Spec.Rules = []tfTypes.MeshCircuitBreakerItemRules{}
			for rulesCount, rulesItem := range itemsItem.Spec.Rules {
				var rules tfTypes.MeshCircuitBreakerItemRules
				if rulesItem.Default == nil {
					rules.Default = nil
				} else {
					rules.Default = &tfTypes.MeshCircuitBreakerItemDefault{}
					if rulesItem.Default.ConnectionLimits == nil {
						rules.Default.ConnectionLimits = nil
					} else {
						rules.Default.ConnectionLimits = &tfTypes.ConnectionLimits{}
						rules.Default.ConnectionLimits.MaxConnectionPools = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.ConnectionLimits.MaxConnectionPools))
						rules.Default.ConnectionLimits.MaxConnections = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.ConnectionLimits.MaxConnections))
						rules.Default.ConnectionLimits.MaxPendingRequests = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.ConnectionLimits.MaxPendingRequests))
						rules.Default.ConnectionLimits.MaxRequests = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.ConnectionLimits.MaxRequests))
						rules.Default.ConnectionLimits.MaxRetries = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.ConnectionLimits.MaxRetries))
					}
					if rulesItem.Default.OutlierDetection == nil {
						rules.Default.OutlierDetection = nil
					} else {
						rules.Default.OutlierDetection = &tfTypes.OutlierDetection{}
						rules.Default.OutlierDetection.BaseEjectionTime = types.StringPointerValue(rulesItem.Default.OutlierDetection.BaseEjectionTime)
						if rulesItem.Default.OutlierDetection.Detectors == nil {
							rules.Default.OutlierDetection.Detectors = nil
						} else {
							rules.Default.OutlierDetection.Detectors = &tfTypes.Detectors{}
							if rulesItem.Default.OutlierDetection.Detectors.FailurePercentage == nil {
								rules.Default.OutlierDetection.Detectors.FailurePercentage = nil
							} else {
								rules.Default.OutlierDetection.Detectors.FailurePercentage = &tfTypes.FailurePercentage{}
								rules.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts))
								rules.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume))
								rules.Default.OutlierDetection.Detectors.FailurePercentage.Threshold = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.FailurePercentage.Threshold))
							}
							if rulesItem.Default.OutlierDetection.Detectors.GatewayFailures == nil {
								rules.Default.OutlierDetection.Detectors.GatewayFailures = nil
							} else {
								rules.Default.OutlierDetection.Detectors.GatewayFailures = &tfTypes.GatewayFailures{}
								rules.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive))
							}
							if rulesItem.Default.OutlierDetection.Detectors.LocalOriginFailures == nil {
								rules.Default.OutlierDetection.Detectors.LocalOriginFailures = nil
							} else {
								rules.Default.OutlierDetection.Detectors.LocalOriginFailures = &tfTypes.GatewayFailures{}
								rules.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive))
							}
							if rulesItem.Default.OutlierDetection.Detectors.SuccessRate == nil {
								rules.Default.OutlierDetection.Detectors.SuccessRate = nil
							} else {
								rules.Default.OutlierDetection.Detectors.SuccessRate = &tfTypes.SuccessRate{}
								rules.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts))
								rules.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume))
								if rulesItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor != nil {
									rules.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor = &tfTypes.Mode{}
									if rulesItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer != nil {
										rules.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer = types.Int64PointerValue(rulesItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer)
									}
									if rulesItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str != nil {
										rules.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str = types.StringPointerValue(rulesItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str)
									}
								}
							}
							if rulesItem.Default.OutlierDetection.Detectors.TotalFailures == nil {
								rules.Default.OutlierDetection.Detectors.TotalFailures = nil
							} else {
								rules.Default.OutlierDetection.Detectors.TotalFailures = &tfTypes.GatewayFailures{}
								rules.Default.OutlierDetection.Detectors.TotalFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.Detectors.TotalFailures.Consecutive))
							}
						}
						rules.Default.OutlierDetection.Disabled = types.BoolPointerValue(rulesItem.Default.OutlierDetection.Disabled)
						if rulesItem.Default.OutlierDetection.HealthyPanicThreshold != nil {
							rules.Default.OutlierDetection.HealthyPanicThreshold = &tfTypes.Mode{}
							if rulesItem.Default.OutlierDetection.HealthyPanicThreshold.Integer != nil {
								rules.Default.OutlierDetection.HealthyPanicThreshold.Integer = types.Int64PointerValue(rulesItem.Default.OutlierDetection.HealthyPanicThreshold.Integer)
							}
							if rulesItem.Default.OutlierDetection.HealthyPanicThreshold.Str != nil {
								rules.Default.OutlierDetection.HealthyPanicThreshold.Str = types.StringPointerValue(rulesItem.Default.OutlierDetection.HealthyPanicThreshold.Str)
							}
						}
						rules.Default.OutlierDetection.Interval = types.StringPointerValue(rulesItem.Default.OutlierDetection.Interval)
						rules.Default.OutlierDetection.MaxEjectionPercent = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(rulesItem.Default.OutlierDetection.MaxEjectionPercent))
						rules.Default.OutlierDetection.SplitExternalAndLocalErrors = types.BoolPointerValue(rulesItem.Default.OutlierDetection.SplitExternalAndLocalErrors)
					}
				}
				if rulesCount+1 > len(items.Spec.Rules) {
					items.Spec.Rules = append(items.Spec.Rules, rules)
				} else {
					items.Spec.Rules[rulesCount].Default = rules.Default
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items.Spec.TargetRef = nil
			} else {
				items.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				items.Spec.TargetRef.Kind = types.StringValue(string(itemsItem.Spec.TargetRef.Kind))
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items.Spec.TargetRef.Labels = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Labels))
					for key2, value2 := range itemsItem.Spec.TargetRef.Labels {
						items.Spec.TargetRef.Labels[key2] = types.StringValue(value2)
					}
				}
				items.Spec.TargetRef.Mesh = types.StringPointerValue(itemsItem.Spec.TargetRef.Mesh)
				items.Spec.TargetRef.Name = types.StringPointerValue(itemsItem.Spec.TargetRef.Name)
				items.Spec.TargetRef.Namespace = types.StringPointerValue(itemsItem.Spec.TargetRef.Namespace)
				items.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(itemsItem.Spec.TargetRef.ProxyTypes))
				for _, v := range itemsItem.Spec.TargetRef.ProxyTypes {
					items.Spec.TargetRef.ProxyTypes = append(items.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				items.Spec.TargetRef.SectionName = types.StringPointerValue(itemsItem.Spec.TargetRef.SectionName)
				if len(itemsItem.Spec.TargetRef.Tags) > 0 {
					items.Spec.TargetRef.Tags = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Tags))
					for key3, value3 := range itemsItem.Spec.TargetRef.Tags {
						items.Spec.TargetRef.Tags[key3] = types.StringValue(value3)
					}
				}
			}
			items.Spec.To = []tfTypes.MeshCircuitBreakerItemFrom{}
			for toCount, toItem := range itemsItem.Spec.To {
				var to tfTypes.MeshCircuitBreakerItemFrom
				if toItem.Default == nil {
					to.Default = nil
				} else {
					to.Default = &tfTypes.MeshCircuitBreakerItemDefault{}
					if toItem.Default.ConnectionLimits == nil {
						to.Default.ConnectionLimits = nil
					} else {
						to.Default.ConnectionLimits = &tfTypes.ConnectionLimits{}
						to.Default.ConnectionLimits.MaxConnectionPools = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.ConnectionLimits.MaxConnectionPools))
						to.Default.ConnectionLimits.MaxConnections = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.ConnectionLimits.MaxConnections))
						to.Default.ConnectionLimits.MaxPendingRequests = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.ConnectionLimits.MaxPendingRequests))
						to.Default.ConnectionLimits.MaxRequests = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.ConnectionLimits.MaxRequests))
						to.Default.ConnectionLimits.MaxRetries = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.ConnectionLimits.MaxRetries))
					}
					if toItem.Default.OutlierDetection == nil {
						to.Default.OutlierDetection = nil
					} else {
						to.Default.OutlierDetection = &tfTypes.OutlierDetection{}
						to.Default.OutlierDetection.BaseEjectionTime = types.StringPointerValue(toItem.Default.OutlierDetection.BaseEjectionTime)
						if toItem.Default.OutlierDetection.Detectors == nil {
							to.Default.OutlierDetection.Detectors = nil
						} else {
							to.Default.OutlierDetection.Detectors = &tfTypes.Detectors{}
							if toItem.Default.OutlierDetection.Detectors.FailurePercentage == nil {
								to.Default.OutlierDetection.Detectors.FailurePercentage = nil
							} else {
								to.Default.OutlierDetection.Detectors.FailurePercentage = &tfTypes.FailurePercentage{}
								to.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.FailurePercentage.MinimumHosts))
								to.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.FailurePercentage.RequestVolume))
								to.Default.OutlierDetection.Detectors.FailurePercentage.Threshold = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.FailurePercentage.Threshold))
							}
							if toItem.Default.OutlierDetection.Detectors.GatewayFailures == nil {
								to.Default.OutlierDetection.Detectors.GatewayFailures = nil
							} else {
								to.Default.OutlierDetection.Detectors.GatewayFailures = &tfTypes.GatewayFailures{}
								to.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.GatewayFailures.Consecutive))
							}
							if toItem.Default.OutlierDetection.Detectors.LocalOriginFailures == nil {
								to.Default.OutlierDetection.Detectors.LocalOriginFailures = nil
							} else {
								to.Default.OutlierDetection.Detectors.LocalOriginFailures = &tfTypes.GatewayFailures{}
								to.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.LocalOriginFailures.Consecutive))
							}
							if toItem.Default.OutlierDetection.Detectors.SuccessRate == nil {
								to.Default.OutlierDetection.Detectors.SuccessRate = nil
							} else {
								to.Default.OutlierDetection.Detectors.SuccessRate = &tfTypes.SuccessRate{}
								to.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.SuccessRate.MinimumHosts))
								to.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.SuccessRate.RequestVolume))
								if toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor != nil {
									to.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor = &tfTypes.Mode{}
									if toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer != nil {
										to.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer = types.Int64PointerValue(toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Integer)
									}
									if toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str != nil {
										to.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str = types.StringPointerValue(toItem.Default.OutlierDetection.Detectors.SuccessRate.StandardDeviationFactor.Str)
									}
								}
							}
							if toItem.Default.OutlierDetection.Detectors.TotalFailures == nil {
								to.Default.OutlierDetection.Detectors.TotalFailures = nil
							} else {
								to.Default.OutlierDetection.Detectors.TotalFailures = &tfTypes.GatewayFailures{}
								to.Default.OutlierDetection.Detectors.TotalFailures.Consecutive = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.Detectors.TotalFailures.Consecutive))
							}
						}
						to.Default.OutlierDetection.Disabled = types.BoolPointerValue(toItem.Default.OutlierDetection.Disabled)
						if toItem.Default.OutlierDetection.HealthyPanicThreshold != nil {
							to.Default.OutlierDetection.HealthyPanicThreshold = &tfTypes.Mode{}
							if toItem.Default.OutlierDetection.HealthyPanicThreshold.Integer != nil {
								to.Default.OutlierDetection.HealthyPanicThreshold.Integer = types.Int64PointerValue(toItem.Default.OutlierDetection.HealthyPanicThreshold.Integer)
							}
							if toItem.Default.OutlierDetection.HealthyPanicThreshold.Str != nil {
								to.Default.OutlierDetection.HealthyPanicThreshold.Str = types.StringPointerValue(toItem.Default.OutlierDetection.HealthyPanicThreshold.Str)
							}
						}
						to.Default.OutlierDetection.Interval = types.StringPointerValue(toItem.Default.OutlierDetection.Interval)
						to.Default.OutlierDetection.MaxEjectionPercent = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.OutlierDetection.MaxEjectionPercent))
						to.Default.OutlierDetection.SplitExternalAndLocalErrors = types.BoolPointerValue(toItem.Default.OutlierDetection.SplitExternalAndLocalErrors)
					}
				}
				to.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
				if len(toItem.TargetRef.Labels) > 0 {
					to.TargetRef.Labels = make(map[string]types.String, len(toItem.TargetRef.Labels))
					for key4, value4 := range toItem.TargetRef.Labels {
						to.TargetRef.Labels[key4] = types.StringValue(value4)
					}
				}
				to.TargetRef.Mesh = types.StringPointerValue(toItem.TargetRef.Mesh)
				to.TargetRef.Name = types.StringPointerValue(toItem.TargetRef.Name)
				to.TargetRef.Namespace = types.StringPointerValue(toItem.TargetRef.Namespace)
				to.TargetRef.ProxyTypes = make([]types.String, 0, len(toItem.TargetRef.ProxyTypes))
				for _, v := range toItem.TargetRef.ProxyTypes {
					to.TargetRef.ProxyTypes = append(to.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				to.TargetRef.SectionName = types.StringPointerValue(toItem.TargetRef.SectionName)
				if len(toItem.TargetRef.Tags) > 0 {
					to.TargetRef.Tags = make(map[string]types.String, len(toItem.TargetRef.Tags))
					for key5, value5 := range toItem.TargetRef.Tags {
						to.TargetRef.Tags[key5] = types.StringValue(value5)
					}
				}
				if toCount+1 > len(items.Spec.To) {
					items.Spec.To = append(items.Spec.To, to)
				} else {
					items.Spec.To[toCount].Default = to.Default
					items.Spec.To[toCount].TargetRef = to.TargetRef
				}
			}
			items.Type = types.StringValue(string(itemsItem.Type))
			if itemsCount+1 > len(r.Items) {
				r.Items = append(r.Items, items)
			} else {
				r.Items[itemsCount].CreationTime = items.CreationTime
				r.Items[itemsCount].Labels = items.Labels
				r.Items[itemsCount].Mesh = items.Mesh
				r.Items[itemsCount].ModificationTime = items.ModificationTime
				r.Items[itemsCount].Name = items.Name
				r.Items[itemsCount].Spec = items.Spec
				r.Items[itemsCount].Type = items.Type
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		r.Total = types.Float64PointerValue(resp.Total)
	}

	return diags
}
