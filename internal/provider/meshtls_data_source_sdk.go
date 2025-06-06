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

func (r *MeshTLSDataSourceModel) ToOperationsGetMeshTLSRequest(ctx context.Context) (*operations.GetMeshTLSRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	out := operations.GetMeshTLSRequest{
		Mesh: mesh,
		Name: name,
	}

	return &out, diags
}

func (r *MeshTLSDataSourceModel) RefreshFromSharedMeshTLSItem(ctx context.Context, resp *shared.MeshTLSItem) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.CreationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.CreationTime))
		labelsValue, labelsDiags := types.MapValueFrom(ctx, types.StringType, resp.Labels)
		diags.Append(labelsDiags...)
		labelsValuable, labelsDiags := kumalabels.KumaLabelsMapType{MapType: types.MapType{ElemType: types.StringType}}.ValueFromMap(ctx, labelsValue)
		diags.Append(labelsDiags...)
		r.Labels, _ = labelsValuable.(kumalabels.KumaLabelsMapValue)
		r.Mesh = types.StringPointerValue(resp.Mesh)
		r.ModificationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(resp.ModificationTime))
		r.Name = types.StringValue(resp.Name)
		r.Spec.From = []tfTypes.MeshTLSItemFrom{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from tfTypes.MeshTLSItemFrom
			if fromItem.Default == nil {
				from.Default = nil
			} else {
				from.Default = &tfTypes.MeshTLSItemDefault{}
				if fromItem.Default.Mode != nil {
					from.Default.Mode = types.StringValue(string(*fromItem.Default.Mode))
				} else {
					from.Default.Mode = types.StringNull()
				}
				from.Default.TLSCiphers = make([]types.String, 0, len(fromItem.Default.TLSCiphers))
				for _, v := range fromItem.Default.TLSCiphers {
					from.Default.TLSCiphers = append(from.Default.TLSCiphers, types.StringValue(string(v)))
				}
				if fromItem.Default.TLSVersion == nil {
					from.Default.TLSVersion = nil
				} else {
					from.Default.TLSVersion = &tfTypes.MeshExternalServiceItemVersion{}
					if fromItem.Default.TLSVersion.Max != nil {
						from.Default.TLSVersion.Max = types.StringValue(string(*fromItem.Default.TLSVersion.Max))
					} else {
						from.Default.TLSVersion.Max = types.StringNull()
					}
					if fromItem.Default.TLSVersion.Min != nil {
						from.Default.TLSVersion.Min = types.StringValue(string(*fromItem.Default.TLSVersion.Min))
					} else {
						from.Default.TLSVersion.Min = types.StringNull()
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
			if fromCount+1 > len(r.Spec.From) {
				r.Spec.From = append(r.Spec.From, from)
			} else {
				r.Spec.From[fromCount].Default = from.Default
				r.Spec.From[fromCount].TargetRef = from.TargetRef
			}
		}
		r.Spec.Rules = []tfTypes.MeshTLSItemRules{}
		if len(r.Spec.Rules) > len(resp.Spec.Rules) {
			r.Spec.Rules = r.Spec.Rules[:len(resp.Spec.Rules)]
		}
		for rulesCount, rulesItem := range resp.Spec.Rules {
			var rules tfTypes.MeshTLSItemRules
			if rulesItem.Default == nil {
				rules.Default = nil
			} else {
				rules.Default = &tfTypes.MeshTLSItemDefault{}
				if rulesItem.Default.Mode != nil {
					rules.Default.Mode = types.StringValue(string(*rulesItem.Default.Mode))
				} else {
					rules.Default.Mode = types.StringNull()
				}
				rules.Default.TLSCiphers = make([]types.String, 0, len(rulesItem.Default.TLSCiphers))
				for _, v := range rulesItem.Default.TLSCiphers {
					rules.Default.TLSCiphers = append(rules.Default.TLSCiphers, types.StringValue(string(v)))
				}
				if rulesItem.Default.TLSVersion == nil {
					rules.Default.TLSVersion = nil
				} else {
					rules.Default.TLSVersion = &tfTypes.MeshExternalServiceItemVersion{}
					if rulesItem.Default.TLSVersion.Max != nil {
						rules.Default.TLSVersion.Max = types.StringValue(string(*rulesItem.Default.TLSVersion.Max))
					} else {
						rules.Default.TLSVersion.Max = types.StringNull()
					}
					if rulesItem.Default.TLSVersion.Min != nil {
						rules.Default.TLSVersion.Min = types.StringValue(string(*rulesItem.Default.TLSVersion.Min))
					} else {
						rules.Default.TLSVersion.Min = types.StringNull()
					}
				}
			}
			if rulesCount+1 > len(r.Spec.Rules) {
				r.Spec.Rules = append(r.Spec.Rules, rules)
			} else {
				r.Spec.Rules[rulesCount].Default = rules.Default
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
				for key2, value2 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key2] = types.StringValue(value2)
				}
			}
			r.Spec.TargetRef.Mesh = types.StringPointerValue(resp.Spec.TargetRef.Mesh)
			r.Spec.TargetRef.Name = types.StringPointerValue(resp.Spec.TargetRef.Name)
			r.Spec.TargetRef.Namespace = types.StringPointerValue(resp.Spec.TargetRef.Namespace)
			r.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(resp.Spec.TargetRef.ProxyTypes))
			for _, v := range resp.Spec.TargetRef.ProxyTypes {
				r.Spec.TargetRef.ProxyTypes = append(r.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
			}
			r.Spec.TargetRef.SectionName = types.StringPointerValue(resp.Spec.TargetRef.SectionName)
			if len(resp.Spec.TargetRef.Tags) > 0 {
				r.Spec.TargetRef.Tags = make(map[string]types.String, len(resp.Spec.TargetRef.Tags))
				for key3, value3 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key3] = types.StringValue(value3)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}

	return diags
}
