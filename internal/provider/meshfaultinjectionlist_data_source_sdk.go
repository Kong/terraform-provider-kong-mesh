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

func (r *MeshFaultInjectionListDataSourceModel) ToOperationsGetMeshFaultInjectionListRequest(ctx context.Context) (*operations.GetMeshFaultInjectionListRequest, diag.Diagnostics) {
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

	out := operations.GetMeshFaultInjectionListRequest{
		Offset: offset,
		Size:   size,
		Mesh:   mesh,
	}

	return &out, diags
}

func (r *MeshFaultInjectionListDataSourceModel) RefreshFromSharedMeshFaultInjectionList(ctx context.Context, resp *shared.MeshFaultInjectionList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Items = []tfTypes.MeshFaultInjectionItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items tfTypes.MeshFaultInjectionItem
			items.CreationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.CreationTime))
			labelsValue, labelsDiags := types.MapValueFrom(ctx, types.StringType, itemsItem.Labels)
			diags.Append(labelsDiags...)
			labelsValuable, labelsDiags := kumalabels.KumaLabelsMapType{MapType: types.MapType{ElemType: types.StringType}}.ValueFromMap(ctx, labelsValue)
			diags.Append(labelsDiags...)
			items.Labels, _ = labelsValuable.(kumalabels.KumaLabelsMapValue)
			items.Mesh = types.StringPointerValue(itemsItem.Mesh)
			items.ModificationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.ModificationTime))
			items.Name = types.StringValue(itemsItem.Name)
			items.Spec.From = []tfTypes.MeshFaultInjectionItemFrom{}
			for fromCount, fromItem := range itemsItem.Spec.From {
				var from tfTypes.MeshFaultInjectionItemFrom
				if fromItem.Default == nil {
					from.Default = nil
				} else {
					from.Default = &tfTypes.MeshFaultInjectionItemDefault{}
					from.Default.HTTP = []tfTypes.HTTP{}
					for httpCount, httpItem := range fromItem.Default.HTTP {
						var http tfTypes.HTTP
						if httpItem.Abort == nil {
							http.Abort = nil
						} else {
							http.Abort = &tfTypes.Abort{}
							http.Abort.HTTPStatus = types.Int32Value(int32(httpItem.Abort.HTTPStatus))
							if httpItem.Abort.Percentage.Integer != nil {
								http.Abort.Percentage.Integer = types.Int64PointerValue(httpItem.Abort.Percentage.Integer)
							}
							if httpItem.Abort.Percentage.Str != nil {
								http.Abort.Percentage.Str = types.StringPointerValue(httpItem.Abort.Percentage.Str)
							}
						}
						if httpItem.Delay == nil {
							http.Delay = nil
						} else {
							http.Delay = &tfTypes.Delay{}
							if httpItem.Delay.Percentage.Integer != nil {
								http.Delay.Percentage.Integer = types.Int64PointerValue(httpItem.Delay.Percentage.Integer)
							}
							if httpItem.Delay.Percentage.Str != nil {
								http.Delay.Percentage.Str = types.StringPointerValue(httpItem.Delay.Percentage.Str)
							}
							http.Delay.Value = types.StringValue(httpItem.Delay.Value)
						}
						if httpItem.ResponseBandwidth == nil {
							http.ResponseBandwidth = nil
						} else {
							http.ResponseBandwidth = &tfTypes.ResponseBandwidth{}
							http.ResponseBandwidth.Limit = types.StringValue(httpItem.ResponseBandwidth.Limit)
							if httpItem.ResponseBandwidth.Percentage.Integer != nil {
								http.ResponseBandwidth.Percentage.Integer = types.Int64PointerValue(httpItem.ResponseBandwidth.Percentage.Integer)
							}
							if httpItem.ResponseBandwidth.Percentage.Str != nil {
								http.ResponseBandwidth.Percentage.Str = types.StringPointerValue(httpItem.ResponseBandwidth.Percentage.Str)
							}
						}
						if httpCount+1 > len(from.Default.HTTP) {
							from.Default.HTTP = append(from.Default.HTTP, http)
						} else {
							from.Default.HTTP[httpCount].Abort = http.Abort
							from.Default.HTTP[httpCount].Delay = http.Delay
							from.Default.HTTP[httpCount].ResponseBandwidth = http.ResponseBandwidth
						}
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
			items.Spec.To = []tfTypes.MeshFaultInjectionItemFrom{}
			for toCount, toItem := range itemsItem.Spec.To {
				var to tfTypes.MeshFaultInjectionItemFrom
				if toItem.Default == nil {
					to.Default = nil
				} else {
					to.Default = &tfTypes.MeshFaultInjectionItemDefault{}
					to.Default.HTTP = []tfTypes.HTTP{}
					for httpCount1, httpItem1 := range toItem.Default.HTTP {
						var http1 tfTypes.HTTP
						if httpItem1.Abort == nil {
							http1.Abort = nil
						} else {
							http1.Abort = &tfTypes.Abort{}
							http1.Abort.HTTPStatus = types.Int32Value(int32(httpItem1.Abort.HTTPStatus))
							if httpItem1.Abort.Percentage.Integer != nil {
								http1.Abort.Percentage.Integer = types.Int64PointerValue(httpItem1.Abort.Percentage.Integer)
							}
							if httpItem1.Abort.Percentage.Str != nil {
								http1.Abort.Percentage.Str = types.StringPointerValue(httpItem1.Abort.Percentage.Str)
							}
						}
						if httpItem1.Delay == nil {
							http1.Delay = nil
						} else {
							http1.Delay = &tfTypes.Delay{}
							if httpItem1.Delay.Percentage.Integer != nil {
								http1.Delay.Percentage.Integer = types.Int64PointerValue(httpItem1.Delay.Percentage.Integer)
							}
							if httpItem1.Delay.Percentage.Str != nil {
								http1.Delay.Percentage.Str = types.StringPointerValue(httpItem1.Delay.Percentage.Str)
							}
							http1.Delay.Value = types.StringValue(httpItem1.Delay.Value)
						}
						if httpItem1.ResponseBandwidth == nil {
							http1.ResponseBandwidth = nil
						} else {
							http1.ResponseBandwidth = &tfTypes.ResponseBandwidth{}
							http1.ResponseBandwidth.Limit = types.StringValue(httpItem1.ResponseBandwidth.Limit)
							if httpItem1.ResponseBandwidth.Percentage.Integer != nil {
								http1.ResponseBandwidth.Percentage.Integer = types.Int64PointerValue(httpItem1.ResponseBandwidth.Percentage.Integer)
							}
							if httpItem1.ResponseBandwidth.Percentage.Str != nil {
								http1.ResponseBandwidth.Percentage.Str = types.StringPointerValue(httpItem1.ResponseBandwidth.Percentage.Str)
							}
						}
						if httpCount1+1 > len(to.Default.HTTP) {
							to.Default.HTTP = append(to.Default.HTTP, http1)
						} else {
							to.Default.HTTP[httpCount1].Abort = http1.Abort
							to.Default.HTTP[httpCount1].Delay = http1.Delay
							to.Default.HTTP[httpCount1].ResponseBandwidth = http1.ResponseBandwidth
						}
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
