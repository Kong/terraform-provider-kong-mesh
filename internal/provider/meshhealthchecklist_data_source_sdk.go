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

func (r *MeshHealthCheckListDataSourceModel) ToOperationsGetMeshHealthCheckListRequest(ctx context.Context) (*operations.GetMeshHealthCheckListRequest, diag.Diagnostics) {
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

	out := operations.GetMeshHealthCheckListRequest{
		Offset: offset,
		Size:   size,
		Mesh:   mesh,
	}

	return &out, diags
}

func (r *MeshHealthCheckListDataSourceModel) RefreshFromSharedMeshHealthCheckList(ctx context.Context, resp *shared.MeshHealthCheckList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Items = []tfTypes.MeshHealthCheckItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items tfTypes.MeshHealthCheckItem
			items.CreationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.CreationTime))
			labelsValue, labelsDiags := types.MapValueFrom(ctx, types.StringType, itemsItem.Labels)
			diags.Append(labelsDiags...)
			labelsValuable, labelsDiags := kumalabels.KumaLabelsMapType{MapType: types.MapType{ElemType: types.StringType}}.ValueFromMap(ctx, labelsValue)
			diags.Append(labelsDiags...)
			items.Labels, _ = labelsValuable.(kumalabels.KumaLabelsMapValue)
			items.Mesh = types.StringPointerValue(itemsItem.Mesh)
			items.ModificationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.ModificationTime))
			items.Name = types.StringValue(itemsItem.Name)
			if itemsItem.Spec.TargetRef == nil {
				items.Spec.TargetRef = nil
			} else {
				items.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				items.Spec.TargetRef.Kind = types.StringValue(string(itemsItem.Spec.TargetRef.Kind))
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items.Spec.TargetRef.Labels = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Labels))
					for key, value := range itemsItem.Spec.TargetRef.Labels {
						items.Spec.TargetRef.Labels[key] = types.StringValue(value)
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
					for key1, value1 := range itemsItem.Spec.TargetRef.Tags {
						items.Spec.TargetRef.Tags[key1] = types.StringValue(value1)
					}
				}
			}
			items.Spec.To = []tfTypes.MeshHealthCheckItemTo{}
			for toCount, toItem := range itemsItem.Spec.To {
				var to tfTypes.MeshHealthCheckItemTo
				if toItem.Default == nil {
					to.Default = nil
				} else {
					to.Default = &tfTypes.MeshHealthCheckItemDefault{}
					to.Default.AlwaysLogHealthCheckFailures = types.BoolPointerValue(toItem.Default.AlwaysLogHealthCheckFailures)
					to.Default.EventLogPath = types.StringPointerValue(toItem.Default.EventLogPath)
					to.Default.FailTrafficOnPanic = types.BoolPointerValue(toItem.Default.FailTrafficOnPanic)
					if toItem.Default.Grpc == nil {
						to.Default.Grpc = nil
					} else {
						to.Default.Grpc = &tfTypes.Grpc{}
						to.Default.Grpc.Authority = types.StringPointerValue(toItem.Default.Grpc.Authority)
						to.Default.Grpc.Disabled = types.BoolPointerValue(toItem.Default.Grpc.Disabled)
						to.Default.Grpc.ServiceName = types.StringPointerValue(toItem.Default.Grpc.ServiceName)
					}
					if toItem.Default.HealthyPanicThreshold != nil {
						to.Default.HealthyPanicThreshold = &tfTypes.Mode{}
						if toItem.Default.HealthyPanicThreshold.Integer != nil {
							to.Default.HealthyPanicThreshold.Integer = types.Int64PointerValue(toItem.Default.HealthyPanicThreshold.Integer)
						}
						if toItem.Default.HealthyPanicThreshold.Str != nil {
							to.Default.HealthyPanicThreshold.Str = types.StringPointerValue(toItem.Default.HealthyPanicThreshold.Str)
						}
					}
					to.Default.HealthyThreshold = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.HealthyThreshold))
					if toItem.Default.HTTP == nil {
						to.Default.HTTP = nil
					} else {
						to.Default.HTTP = &tfTypes.MeshHealthCheckItemHTTP{}
						to.Default.HTTP.Disabled = types.BoolPointerValue(toItem.Default.HTTP.Disabled)
						to.Default.HTTP.ExpectedStatuses = make([]types.Int64, 0, len(toItem.Default.HTTP.ExpectedStatuses))
						for _, v := range toItem.Default.HTTP.ExpectedStatuses {
							to.Default.HTTP.ExpectedStatuses = append(to.Default.HTTP.ExpectedStatuses, types.Int64Value(v))
						}
						to.Default.HTTP.Path = types.StringPointerValue(toItem.Default.HTTP.Path)
						if toItem.Default.HTTP.RequestHeadersToAdd == nil {
							to.Default.HTTP.RequestHeadersToAdd = nil
						} else {
							to.Default.HTTP.RequestHeadersToAdd = &tfTypes.MeshGlobalRateLimitItemHeaders{}
							to.Default.HTTP.RequestHeadersToAdd.Add = []tfTypes.MeshGlobalRateLimitItemAdd{}
							for addCount, addItem := range toItem.Default.HTTP.RequestHeadersToAdd.Add {
								var add tfTypes.MeshGlobalRateLimitItemAdd
								add.Name = types.StringValue(addItem.Name)
								add.Value = types.StringValue(addItem.Value)
								if addCount+1 > len(to.Default.HTTP.RequestHeadersToAdd.Add) {
									to.Default.HTTP.RequestHeadersToAdd.Add = append(to.Default.HTTP.RequestHeadersToAdd.Add, add)
								} else {
									to.Default.HTTP.RequestHeadersToAdd.Add[addCount].Name = add.Name
									to.Default.HTTP.RequestHeadersToAdd.Add[addCount].Value = add.Value
								}
							}
							to.Default.HTTP.RequestHeadersToAdd.Set = []tfTypes.MeshGlobalRateLimitItemAdd{}
							for setCount, setItem := range toItem.Default.HTTP.RequestHeadersToAdd.Set {
								var set tfTypes.MeshGlobalRateLimitItemAdd
								set.Name = types.StringValue(setItem.Name)
								set.Value = types.StringValue(setItem.Value)
								if setCount+1 > len(to.Default.HTTP.RequestHeadersToAdd.Set) {
									to.Default.HTTP.RequestHeadersToAdd.Set = append(to.Default.HTTP.RequestHeadersToAdd.Set, set)
								} else {
									to.Default.HTTP.RequestHeadersToAdd.Set[setCount].Name = set.Name
									to.Default.HTTP.RequestHeadersToAdd.Set[setCount].Value = set.Value
								}
							}
						}
					}
					to.Default.InitialJitter = types.StringPointerValue(toItem.Default.InitialJitter)
					to.Default.Interval = types.StringPointerValue(toItem.Default.Interval)
					to.Default.IntervalJitter = types.StringPointerValue(toItem.Default.IntervalJitter)
					to.Default.IntervalJitterPercent = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.IntervalJitterPercent))
					to.Default.NoTrafficInterval = types.StringPointerValue(toItem.Default.NoTrafficInterval)
					to.Default.ReuseConnection = types.BoolPointerValue(toItem.Default.ReuseConnection)
					if toItem.Default.TCP == nil {
						to.Default.TCP = nil
					} else {
						to.Default.TCP = &tfTypes.TCP{}
						to.Default.TCP.Disabled = types.BoolPointerValue(toItem.Default.TCP.Disabled)
						to.Default.TCP.Receive = make([]types.String, 0, len(toItem.Default.TCP.Receive))
						for _, v := range toItem.Default.TCP.Receive {
							to.Default.TCP.Receive = append(to.Default.TCP.Receive, types.StringValue(v))
						}
						to.Default.TCP.Send = types.StringPointerValue(toItem.Default.TCP.Send)
					}
					to.Default.Timeout = types.StringPointerValue(toItem.Default.Timeout)
					to.Default.UnhealthyThreshold = types.Int32PointerValue(typeconvert.IntPointerToInt32Pointer(toItem.Default.UnhealthyThreshold))
				}
				to.TargetRef.Kind = types.StringValue(string(toItem.TargetRef.Kind))
				if len(toItem.TargetRef.Labels) > 0 {
					to.TargetRef.Labels = make(map[string]types.String, len(toItem.TargetRef.Labels))
					for key2, value2 := range toItem.TargetRef.Labels {
						to.TargetRef.Labels[key2] = types.StringValue(value2)
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
					for key3, value3 := range toItem.TargetRef.Tags {
						to.TargetRef.Tags[key3] = types.StringValue(value3)
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
