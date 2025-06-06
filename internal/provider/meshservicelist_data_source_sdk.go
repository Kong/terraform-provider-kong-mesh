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

func (r *MeshServiceListDataSourceModel) ToOperationsGetMeshServiceListRequest(ctx context.Context) (*operations.GetMeshServiceListRequest, diag.Diagnostics) {
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

	out := operations.GetMeshServiceListRequest{
		Offset: offset,
		Size:   size,
		Mesh:   mesh,
	}

	return &out, diags
}

func (r *MeshServiceListDataSourceModel) RefreshFromSharedMeshServiceList(ctx context.Context, resp *shared.MeshServiceList) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Items = []tfTypes.MeshServiceItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items tfTypes.MeshServiceItem
			items.CreationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.CreationTime))
			labelsValue, labelsDiags := types.MapValueFrom(ctx, types.StringType, itemsItem.Labels)
			diags.Append(labelsDiags...)
			labelsValuable, labelsDiags := kumalabels.KumaLabelsMapType{MapType: types.MapType{ElemType: types.StringType}}.ValueFromMap(ctx, labelsValue)
			diags.Append(labelsDiags...)
			items.Labels, _ = labelsValuable.(kumalabels.KumaLabelsMapValue)
			items.Mesh = types.StringPointerValue(itemsItem.Mesh)
			items.ModificationTime = types.StringPointerValue(typeconvert.TimePointerToStringPointer(itemsItem.ModificationTime))
			items.Name = types.StringValue(itemsItem.Name)
			items.Spec.Identities = []tfTypes.Path{}
			for identitiesCount, identitiesItem := range itemsItem.Spec.Identities {
				var identities tfTypes.Path
				identities.Type = types.StringValue(string(identitiesItem.Type))
				identities.Value = types.StringValue(identitiesItem.Value)
				if identitiesCount+1 > len(items.Spec.Identities) {
					items.Spec.Identities = append(items.Spec.Identities, identities)
				} else {
					items.Spec.Identities[identitiesCount].Type = identities.Type
					items.Spec.Identities[identitiesCount].Value = identities.Value
				}
			}
			items.Spec.Ports = []tfTypes.MeshServiceItemPorts{}
			for portsCount, portsItem := range itemsItem.Spec.Ports {
				var ports tfTypes.MeshServiceItemPorts
				ports.AppProtocol = types.StringPointerValue(portsItem.AppProtocol)
				ports.Name = types.StringPointerValue(portsItem.Name)
				ports.Port = types.Int32Value(int32(portsItem.Port))
				if portsItem.TargetPort != nil {
					ports.TargetPort = &tfTypes.Mode{}
					if portsItem.TargetPort.Integer != nil {
						ports.TargetPort.Integer = types.Int64PointerValue(portsItem.TargetPort.Integer)
					}
					if portsItem.TargetPort.Str != nil {
						ports.TargetPort.Str = types.StringPointerValue(portsItem.TargetPort.Str)
					}
				}
				if portsCount+1 > len(items.Spec.Ports) {
					items.Spec.Ports = append(items.Spec.Ports, ports)
				} else {
					items.Spec.Ports[portsCount].AppProtocol = ports.AppProtocol
					items.Spec.Ports[portsCount].Name = ports.Name
					items.Spec.Ports[portsCount].Port = ports.Port
					items.Spec.Ports[portsCount].TargetPort = ports.TargetPort
				}
			}
			if itemsItem.Spec.Selector == nil {
				items.Spec.Selector = nil
			} else {
				items.Spec.Selector = &tfTypes.MeshServiceItemSelector{}
				if itemsItem.Spec.Selector.DataplaneRef == nil {
					items.Spec.Selector.DataplaneRef = nil
				} else {
					items.Spec.Selector.DataplaneRef = &tfTypes.DataplaneRef{}
					items.Spec.Selector.DataplaneRef.Name = types.StringPointerValue(itemsItem.Spec.Selector.DataplaneRef.Name)
				}
				if len(itemsItem.Spec.Selector.DataplaneTags) > 0 {
					items.Spec.Selector.DataplaneTags = make(map[string]types.String, len(itemsItem.Spec.Selector.DataplaneTags))
					for key, value := range itemsItem.Spec.Selector.DataplaneTags {
						items.Spec.Selector.DataplaneTags[key] = types.StringValue(value)
					}
				}
			}
			if itemsItem.Spec.State != nil {
				items.Spec.State = types.StringValue(string(*itemsItem.Spec.State))
			} else {
				items.Spec.State = types.StringNull()
			}
			if itemsItem.Status == nil {
				items.Status = nil
			} else {
				items.Status = &tfTypes.MeshServiceItemStatus{}
				items.Status.Addresses = []tfTypes.Addresses{}
				for addressesCount, addressesItem := range itemsItem.Status.Addresses {
					var addresses tfTypes.Addresses
					addresses.Hostname = types.StringPointerValue(addressesItem.Hostname)
					if addressesItem.HostnameGeneratorRef == nil {
						addresses.HostnameGeneratorRef = nil
					} else {
						addresses.HostnameGeneratorRef = &tfTypes.HostnameGeneratorRef{}
						addresses.HostnameGeneratorRef.CoreName = types.StringValue(addressesItem.HostnameGeneratorRef.CoreName)
					}
					addresses.Origin = types.StringPointerValue(addressesItem.Origin)
					if addressesCount+1 > len(items.Status.Addresses) {
						items.Status.Addresses = append(items.Status.Addresses, addresses)
					} else {
						items.Status.Addresses[addressesCount].Hostname = addresses.Hostname
						items.Status.Addresses[addressesCount].HostnameGeneratorRef = addresses.HostnameGeneratorRef
						items.Status.Addresses[addressesCount].Origin = addresses.Origin
					}
				}
				if itemsItem.Status.DataplaneProxies == nil {
					items.Status.DataplaneProxies = nil
				} else {
					items.Status.DataplaneProxies = &tfTypes.DataplaneProxies{}
					items.Status.DataplaneProxies.Connected = types.Int64PointerValue(itemsItem.Status.DataplaneProxies.Connected)
					items.Status.DataplaneProxies.Healthy = types.Int64PointerValue(itemsItem.Status.DataplaneProxies.Healthy)
					items.Status.DataplaneProxies.Total = types.Int64PointerValue(itemsItem.Status.DataplaneProxies.Total)
				}
				items.Status.HostnameGenerators = []tfTypes.HostnameGenerators{}
				for hostnameGeneratorsCount, hostnameGeneratorsItem := range itemsItem.Status.HostnameGenerators {
					var hostnameGenerators tfTypes.HostnameGenerators
					hostnameGenerators.Conditions = []tfTypes.Conditions{}
					for conditionsCount, conditionsItem := range hostnameGeneratorsItem.Conditions {
						var conditions tfTypes.Conditions
						conditions.Message = types.StringValue(conditionsItem.Message)
						conditions.Reason = types.StringValue(conditionsItem.Reason)
						conditions.Status = types.StringValue(string(conditionsItem.Status))
						conditions.Type = types.StringValue(conditionsItem.Type)
						if conditionsCount+1 > len(hostnameGenerators.Conditions) {
							hostnameGenerators.Conditions = append(hostnameGenerators.Conditions, conditions)
						} else {
							hostnameGenerators.Conditions[conditionsCount].Message = conditions.Message
							hostnameGenerators.Conditions[conditionsCount].Reason = conditions.Reason
							hostnameGenerators.Conditions[conditionsCount].Status = conditions.Status
							hostnameGenerators.Conditions[conditionsCount].Type = conditions.Type
						}
					}
					hostnameGenerators.HostnameGeneratorRef.CoreName = types.StringValue(hostnameGeneratorsItem.HostnameGeneratorRef.CoreName)
					if hostnameGeneratorsCount+1 > len(items.Status.HostnameGenerators) {
						items.Status.HostnameGenerators = append(items.Status.HostnameGenerators, hostnameGenerators)
					} else {
						items.Status.HostnameGenerators[hostnameGeneratorsCount].Conditions = hostnameGenerators.Conditions
						items.Status.HostnameGenerators[hostnameGeneratorsCount].HostnameGeneratorRef = hostnameGenerators.HostnameGeneratorRef
					}
				}
				if itemsItem.Status.TLS == nil {
					items.Status.TLS = nil
				} else {
					items.Status.TLS = &tfTypes.MeshServiceItemTLS{}
					if itemsItem.Status.TLS.Status != nil {
						items.Status.TLS.Status = types.StringValue(string(*itemsItem.Status.TLS.Status))
					} else {
						items.Status.TLS.Status = types.StringNull()
					}
				}
				items.Status.Vips = []tfTypes.Vip{}
				for vipsCount, vipsItem := range itemsItem.Status.Vips {
					var vips tfTypes.Vip
					vips.IP = types.StringPointerValue(vipsItem.IP)
					if vipsCount+1 > len(items.Status.Vips) {
						items.Status.Vips = append(items.Status.Vips, vips)
					} else {
						items.Status.Vips[vipsCount].IP = vips.IP
					}
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
				r.Items[itemsCount].Status = items.Status
				r.Items[itemsCount].Type = items.Type
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		r.Total = types.Float64PointerValue(resp.Total)
	}

	return diags
}
