// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"github.com/Kong/shared-speakeasy/customtypes/kumalabels"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/provider/typeconvert"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
)

func (r *MeshAccessLogDataSourceModel) ToOperationsGetMeshAccessLogRequest(ctx context.Context) (*operations.GetMeshAccessLogRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	out := operations.GetMeshAccessLogRequest{
		Mesh: mesh,
		Name: name,
	}

	return &out, diags
}

func (r *MeshAccessLogDataSourceModel) RefreshFromSharedMeshAccessLogItem(ctx context.Context, resp *shared.MeshAccessLogItem) diag.Diagnostics {
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
		r.Spec.From = []tfTypes.From{}
		if len(r.Spec.From) > len(resp.Spec.From) {
			r.Spec.From = r.Spec.From[:len(resp.Spec.From)]
		}
		for fromCount, fromItem := range resp.Spec.From {
			var from tfTypes.From
			from.Default.Backends = []tfTypes.MeshAccessLogItemSpecFromBackends{}
			for backendsCount, backendsItem := range fromItem.Default.Backends {
				var backends tfTypes.MeshAccessLogItemSpecFromBackends
				if backendsItem.File == nil {
					backends.File = nil
				} else {
					backends.File = &tfTypes.File{}
					if backendsItem.File.Format == nil {
						backends.File.Format = nil
					} else {
						backends.File.Format = &tfTypes.Format{}
						backends.File.Format.JSON = []tfTypes.JSON{}
						for jsonVarCount, jsonVarItem := range backendsItem.File.Format.JSON {
							var jsonVar tfTypes.JSON
							jsonVar.Key = types.StringValue(jsonVarItem.Key)
							jsonVar.Value = types.StringValue(jsonVarItem.Value)
							if jsonVarCount+1 > len(backends.File.Format.JSON) {
								backends.File.Format.JSON = append(backends.File.Format.JSON, jsonVar)
							} else {
								backends.File.Format.JSON[jsonVarCount].Key = jsonVar.Key
								backends.File.Format.JSON[jsonVarCount].Value = jsonVar.Value
							}
						}
						backends.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem.File.Format.OmitEmptyValues)
						backends.File.Format.Plain = types.StringPointerValue(backendsItem.File.Format.Plain)
						backends.File.Format.Type = types.StringValue(string(backendsItem.File.Format.Type))
					}
					backends.File.Path = types.StringValue(backendsItem.File.Path)
				}
				if backendsItem.OpenTelemetry == nil {
					backends.OpenTelemetry = nil
				} else {
					backends.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecFromOpenTelemetry{}
					backends.OpenTelemetry.Attributes = []tfTypes.JSON{}
					for attributesCount, attributesItem := range backendsItem.OpenTelemetry.Attributes {
						var attributes tfTypes.JSON
						attributes.Key = types.StringValue(attributesItem.Key)
						attributes.Value = types.StringValue(attributesItem.Value)
						if attributesCount+1 > len(backends.OpenTelemetry.Attributes) {
							backends.OpenTelemetry.Attributes = append(backends.OpenTelemetry.Attributes, attributes)
						} else {
							backends.OpenTelemetry.Attributes[attributesCount].Key = attributes.Key
							backends.OpenTelemetry.Attributes[attributesCount].Value = attributes.Value
						}
					}
					if backendsItem.OpenTelemetry.Body == nil {
						backends.OpenTelemetry.Body = types.StringNull()
					} else {
						bodyResult, _ := json.Marshal(backendsItem.OpenTelemetry.Body)
						backends.OpenTelemetry.Body = types.StringValue(string(bodyResult))
					}
					backends.OpenTelemetry.Endpoint = types.StringValue(backendsItem.OpenTelemetry.Endpoint)
				}
				if backendsItem.TCP == nil {
					backends.TCP = nil
				} else {
					backends.TCP = &tfTypes.MeshAccessLogItemTCP{}
					backends.TCP.Address = types.StringValue(backendsItem.TCP.Address)
					if backendsItem.TCP.Format == nil {
						backends.TCP.Format = nil
					} else {
						backends.TCP.Format = &tfTypes.Format{}
						backends.TCP.Format.JSON = []tfTypes.JSON{}
						for jsonVarCount1, jsonVarItem1 := range backendsItem.TCP.Format.JSON {
							var jsonVar1 tfTypes.JSON
							jsonVar1.Key = types.StringValue(jsonVarItem1.Key)
							jsonVar1.Value = types.StringValue(jsonVarItem1.Value)
							if jsonVarCount1+1 > len(backends.TCP.Format.JSON) {
								backends.TCP.Format.JSON = append(backends.TCP.Format.JSON, jsonVar1)
							} else {
								backends.TCP.Format.JSON[jsonVarCount1].Key = jsonVar1.Key
								backends.TCP.Format.JSON[jsonVarCount1].Value = jsonVar1.Value
							}
						}
						backends.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem.TCP.Format.OmitEmptyValues)
						backends.TCP.Format.Plain = types.StringPointerValue(backendsItem.TCP.Format.Plain)
						backends.TCP.Format.Type = types.StringValue(string(backendsItem.TCP.Format.Type))
					}
				}
				backends.Type = types.StringValue(string(backendsItem.Type))
				if backendsCount+1 > len(from.Default.Backends) {
					from.Default.Backends = append(from.Default.Backends, backends)
				} else {
					from.Default.Backends[backendsCount].File = backends.File
					from.Default.Backends[backendsCount].OpenTelemetry = backends.OpenTelemetry
					from.Default.Backends[backendsCount].TCP = backends.TCP
					from.Default.Backends[backendsCount].Type = backends.Type
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
		r.Spec.Rules = []tfTypes.Rules{}
		if len(r.Spec.Rules) > len(resp.Spec.Rules) {
			r.Spec.Rules = r.Spec.Rules[:len(resp.Spec.Rules)]
		}
		for rulesCount, rulesItem := range resp.Spec.Rules {
			var rules tfTypes.Rules
			rules.Default.Backends = []tfTypes.MeshAccessLogItemSpecFromBackends{}
			for backendsCount1, backendsItem1 := range rulesItem.Default.Backends {
				var backends1 tfTypes.MeshAccessLogItemSpecFromBackends
				if backendsItem1.File == nil {
					backends1.File = nil
				} else {
					backends1.File = &tfTypes.File{}
					if backendsItem1.File.Format == nil {
						backends1.File.Format = nil
					} else {
						backends1.File.Format = &tfTypes.Format{}
						backends1.File.Format.JSON = []tfTypes.JSON{}
						for jsonVarCount2, jsonVarItem2 := range backendsItem1.File.Format.JSON {
							var jsonVar2 tfTypes.JSON
							jsonVar2.Key = types.StringValue(jsonVarItem2.Key)
							jsonVar2.Value = types.StringValue(jsonVarItem2.Value)
							if jsonVarCount2+1 > len(backends1.File.Format.JSON) {
								backends1.File.Format.JSON = append(backends1.File.Format.JSON, jsonVar2)
							} else {
								backends1.File.Format.JSON[jsonVarCount2].Key = jsonVar2.Key
								backends1.File.Format.JSON[jsonVarCount2].Value = jsonVar2.Value
							}
						}
						backends1.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem1.File.Format.OmitEmptyValues)
						backends1.File.Format.Plain = types.StringPointerValue(backendsItem1.File.Format.Plain)
						backends1.File.Format.Type = types.StringValue(string(backendsItem1.File.Format.Type))
					}
					backends1.File.Path = types.StringValue(backendsItem1.File.Path)
				}
				if backendsItem1.OpenTelemetry == nil {
					backends1.OpenTelemetry = nil
				} else {
					backends1.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecFromOpenTelemetry{}
					backends1.OpenTelemetry.Attributes = []tfTypes.JSON{}
					for attributesCount1, attributesItem1 := range backendsItem1.OpenTelemetry.Attributes {
						var attributes1 tfTypes.JSON
						attributes1.Key = types.StringValue(attributesItem1.Key)
						attributes1.Value = types.StringValue(attributesItem1.Value)
						if attributesCount1+1 > len(backends1.OpenTelemetry.Attributes) {
							backends1.OpenTelemetry.Attributes = append(backends1.OpenTelemetry.Attributes, attributes1)
						} else {
							backends1.OpenTelemetry.Attributes[attributesCount1].Key = attributes1.Key
							backends1.OpenTelemetry.Attributes[attributesCount1].Value = attributes1.Value
						}
					}
					if backendsItem1.OpenTelemetry.Body == nil {
						backends1.OpenTelemetry.Body = types.StringNull()
					} else {
						bodyResult1, _ := json.Marshal(backendsItem1.OpenTelemetry.Body)
						backends1.OpenTelemetry.Body = types.StringValue(string(bodyResult1))
					}
					backends1.OpenTelemetry.Endpoint = types.StringValue(backendsItem1.OpenTelemetry.Endpoint)
				}
				if backendsItem1.TCP == nil {
					backends1.TCP = nil
				} else {
					backends1.TCP = &tfTypes.MeshAccessLogItemTCP{}
					backends1.TCP.Address = types.StringValue(backendsItem1.TCP.Address)
					if backendsItem1.TCP.Format == nil {
						backends1.TCP.Format = nil
					} else {
						backends1.TCP.Format = &tfTypes.Format{}
						backends1.TCP.Format.JSON = []tfTypes.JSON{}
						for jsonVarCount3, jsonVarItem3 := range backendsItem1.TCP.Format.JSON {
							var jsonVar3 tfTypes.JSON
							jsonVar3.Key = types.StringValue(jsonVarItem3.Key)
							jsonVar3.Value = types.StringValue(jsonVarItem3.Value)
							if jsonVarCount3+1 > len(backends1.TCP.Format.JSON) {
								backends1.TCP.Format.JSON = append(backends1.TCP.Format.JSON, jsonVar3)
							} else {
								backends1.TCP.Format.JSON[jsonVarCount3].Key = jsonVar3.Key
								backends1.TCP.Format.JSON[jsonVarCount3].Value = jsonVar3.Value
							}
						}
						backends1.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem1.TCP.Format.OmitEmptyValues)
						backends1.TCP.Format.Plain = types.StringPointerValue(backendsItem1.TCP.Format.Plain)
						backends1.TCP.Format.Type = types.StringValue(string(backendsItem1.TCP.Format.Type))
					}
				}
				backends1.Type = types.StringValue(string(backendsItem1.Type))
				if backendsCount1+1 > len(rules.Default.Backends) {
					rules.Default.Backends = append(rules.Default.Backends, backends1)
				} else {
					rules.Default.Backends[backendsCount1].File = backends1.File
					rules.Default.Backends[backendsCount1].OpenTelemetry = backends1.OpenTelemetry
					rules.Default.Backends[backendsCount1].TCP = backends1.TCP
					rules.Default.Backends[backendsCount1].Type = backends1.Type
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
		r.Spec.To = []tfTypes.From{}
		if len(r.Spec.To) > len(resp.Spec.To) {
			r.Spec.To = r.Spec.To[:len(resp.Spec.To)]
		}
		for toCount, toItem := range resp.Spec.To {
			var to tfTypes.From
			to.Default.Backends = []tfTypes.MeshAccessLogItemSpecFromBackends{}
			for backendsCount2, backendsItem2 := range toItem.Default.Backends {
				var backends2 tfTypes.MeshAccessLogItemSpecFromBackends
				if backendsItem2.File == nil {
					backends2.File = nil
				} else {
					backends2.File = &tfTypes.File{}
					if backendsItem2.File.Format == nil {
						backends2.File.Format = nil
					} else {
						backends2.File.Format = &tfTypes.Format{}
						backends2.File.Format.JSON = []tfTypes.JSON{}
						for jsonVarCount4, jsonVarItem4 := range backendsItem2.File.Format.JSON {
							var jsonVar4 tfTypes.JSON
							jsonVar4.Key = types.StringValue(jsonVarItem4.Key)
							jsonVar4.Value = types.StringValue(jsonVarItem4.Value)
							if jsonVarCount4+1 > len(backends2.File.Format.JSON) {
								backends2.File.Format.JSON = append(backends2.File.Format.JSON, jsonVar4)
							} else {
								backends2.File.Format.JSON[jsonVarCount4].Key = jsonVar4.Key
								backends2.File.Format.JSON[jsonVarCount4].Value = jsonVar4.Value
							}
						}
						backends2.File.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem2.File.Format.OmitEmptyValues)
						backends2.File.Format.Plain = types.StringPointerValue(backendsItem2.File.Format.Plain)
						backends2.File.Format.Type = types.StringValue(string(backendsItem2.File.Format.Type))
					}
					backends2.File.Path = types.StringValue(backendsItem2.File.Path)
				}
				if backendsItem2.OpenTelemetry == nil {
					backends2.OpenTelemetry = nil
				} else {
					backends2.OpenTelemetry = &tfTypes.MeshAccessLogItemSpecFromOpenTelemetry{}
					backends2.OpenTelemetry.Attributes = []tfTypes.JSON{}
					for attributesCount2, attributesItem2 := range backendsItem2.OpenTelemetry.Attributes {
						var attributes2 tfTypes.JSON
						attributes2.Key = types.StringValue(attributesItem2.Key)
						attributes2.Value = types.StringValue(attributesItem2.Value)
						if attributesCount2+1 > len(backends2.OpenTelemetry.Attributes) {
							backends2.OpenTelemetry.Attributes = append(backends2.OpenTelemetry.Attributes, attributes2)
						} else {
							backends2.OpenTelemetry.Attributes[attributesCount2].Key = attributes2.Key
							backends2.OpenTelemetry.Attributes[attributesCount2].Value = attributes2.Value
						}
					}
					if backendsItem2.OpenTelemetry.Body == nil {
						backends2.OpenTelemetry.Body = types.StringNull()
					} else {
						bodyResult2, _ := json.Marshal(backendsItem2.OpenTelemetry.Body)
						backends2.OpenTelemetry.Body = types.StringValue(string(bodyResult2))
					}
					backends2.OpenTelemetry.Endpoint = types.StringValue(backendsItem2.OpenTelemetry.Endpoint)
				}
				if backendsItem2.TCP == nil {
					backends2.TCP = nil
				} else {
					backends2.TCP = &tfTypes.MeshAccessLogItemTCP{}
					backends2.TCP.Address = types.StringValue(backendsItem2.TCP.Address)
					if backendsItem2.TCP.Format == nil {
						backends2.TCP.Format = nil
					} else {
						backends2.TCP.Format = &tfTypes.Format{}
						backends2.TCP.Format.JSON = []tfTypes.JSON{}
						for jsonVarCount5, jsonVarItem5 := range backendsItem2.TCP.Format.JSON {
							var jsonVar5 tfTypes.JSON
							jsonVar5.Key = types.StringValue(jsonVarItem5.Key)
							jsonVar5.Value = types.StringValue(jsonVarItem5.Value)
							if jsonVarCount5+1 > len(backends2.TCP.Format.JSON) {
								backends2.TCP.Format.JSON = append(backends2.TCP.Format.JSON, jsonVar5)
							} else {
								backends2.TCP.Format.JSON[jsonVarCount5].Key = jsonVar5.Key
								backends2.TCP.Format.JSON[jsonVarCount5].Value = jsonVar5.Value
							}
						}
						backends2.TCP.Format.OmitEmptyValues = types.BoolPointerValue(backendsItem2.TCP.Format.OmitEmptyValues)
						backends2.TCP.Format.Plain = types.StringPointerValue(backendsItem2.TCP.Format.Plain)
						backends2.TCP.Format.Type = types.StringValue(string(backendsItem2.TCP.Format.Type))
					}
				}
				backends2.Type = types.StringValue(string(backendsItem2.Type))
				if backendsCount2+1 > len(to.Default.Backends) {
					to.Default.Backends = append(to.Default.Backends, backends2)
				} else {
					to.Default.Backends[backendsCount2].File = backends2.File
					to.Default.Backends[backendsCount2].OpenTelemetry = backends2.OpenTelemetry
					to.Default.Backends[backendsCount2].TCP = backends2.TCP
					to.Default.Backends[backendsCount2].Type = backends2.Type
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
			if toCount+1 > len(r.Spec.To) {
				r.Spec.To = append(r.Spec.To, to)
			} else {
				r.Spec.To[toCount].Default = to.Default
				r.Spec.To[toCount].TargetRef = to.TargetRef
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}

	return diags
}
