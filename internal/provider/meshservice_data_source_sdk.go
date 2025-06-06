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

func (r *MeshServiceDataSourceModel) ToOperationsGetMeshServiceRequest(ctx context.Context) (*operations.GetMeshServiceRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	out := operations.GetMeshServiceRequest{
		Mesh: mesh,
		Name: name,
	}

	return &out, diags
}

func (r *MeshServiceDataSourceModel) RefreshFromSharedMeshServiceItem(ctx context.Context, resp *shared.MeshServiceItem) diag.Diagnostics {
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
		r.Spec.Identities = []tfTypes.Path{}
		if len(r.Spec.Identities) > len(resp.Spec.Identities) {
			r.Spec.Identities = r.Spec.Identities[:len(resp.Spec.Identities)]
		}
		for identitiesCount, identitiesItem := range resp.Spec.Identities {
			var identities tfTypes.Path
			identities.Type = types.StringValue(string(identitiesItem.Type))
			identities.Value = types.StringValue(identitiesItem.Value)
			if identitiesCount+1 > len(r.Spec.Identities) {
				r.Spec.Identities = append(r.Spec.Identities, identities)
			} else {
				r.Spec.Identities[identitiesCount].Type = identities.Type
				r.Spec.Identities[identitiesCount].Value = identities.Value
			}
		}
		r.Spec.Ports = []tfTypes.MeshServiceItemPorts{}
		if len(r.Spec.Ports) > len(resp.Spec.Ports) {
			r.Spec.Ports = r.Spec.Ports[:len(resp.Spec.Ports)]
		}
		for portsCount, portsItem := range resp.Spec.Ports {
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
			if portsCount+1 > len(r.Spec.Ports) {
				r.Spec.Ports = append(r.Spec.Ports, ports)
			} else {
				r.Spec.Ports[portsCount].AppProtocol = ports.AppProtocol
				r.Spec.Ports[portsCount].Name = ports.Name
				r.Spec.Ports[portsCount].Port = ports.Port
				r.Spec.Ports[portsCount].TargetPort = ports.TargetPort
			}
		}
		if resp.Spec.Selector == nil {
			r.Spec.Selector = nil
		} else {
			r.Spec.Selector = &tfTypes.MeshServiceItemSelector{}
			if resp.Spec.Selector.DataplaneRef == nil {
				r.Spec.Selector.DataplaneRef = nil
			} else {
				r.Spec.Selector.DataplaneRef = &tfTypes.DataplaneRef{}
				r.Spec.Selector.DataplaneRef.Name = types.StringPointerValue(resp.Spec.Selector.DataplaneRef.Name)
			}
			if len(resp.Spec.Selector.DataplaneTags) > 0 {
				r.Spec.Selector.DataplaneTags = make(map[string]types.String, len(resp.Spec.Selector.DataplaneTags))
				for key, value := range resp.Spec.Selector.DataplaneTags {
					r.Spec.Selector.DataplaneTags[key] = types.StringValue(value)
				}
			}
		}
		if resp.Spec.State != nil {
			r.Spec.State = types.StringValue(string(*resp.Spec.State))
		} else {
			r.Spec.State = types.StringNull()
		}
		if resp.Status == nil {
			r.Status = nil
		} else {
			r.Status = &tfTypes.MeshServiceItemStatus{}
			r.Status.Addresses = []tfTypes.Addresses{}
			if len(r.Status.Addresses) > len(resp.Status.Addresses) {
				r.Status.Addresses = r.Status.Addresses[:len(resp.Status.Addresses)]
			}
			for addressesCount, addressesItem := range resp.Status.Addresses {
				var addresses tfTypes.Addresses
				addresses.Hostname = types.StringPointerValue(addressesItem.Hostname)
				if addressesItem.HostnameGeneratorRef == nil {
					addresses.HostnameGeneratorRef = nil
				} else {
					addresses.HostnameGeneratorRef = &tfTypes.HostnameGeneratorRef{}
					addresses.HostnameGeneratorRef.CoreName = types.StringValue(addressesItem.HostnameGeneratorRef.CoreName)
				}
				addresses.Origin = types.StringPointerValue(addressesItem.Origin)
				if addressesCount+1 > len(r.Status.Addresses) {
					r.Status.Addresses = append(r.Status.Addresses, addresses)
				} else {
					r.Status.Addresses[addressesCount].Hostname = addresses.Hostname
					r.Status.Addresses[addressesCount].HostnameGeneratorRef = addresses.HostnameGeneratorRef
					r.Status.Addresses[addressesCount].Origin = addresses.Origin
				}
			}
			if resp.Status.DataplaneProxies == nil {
				r.Status.DataplaneProxies = nil
			} else {
				r.Status.DataplaneProxies = &tfTypes.DataplaneProxies{}
				r.Status.DataplaneProxies.Connected = types.Int64PointerValue(resp.Status.DataplaneProxies.Connected)
				r.Status.DataplaneProxies.Healthy = types.Int64PointerValue(resp.Status.DataplaneProxies.Healthy)
				r.Status.DataplaneProxies.Total = types.Int64PointerValue(resp.Status.DataplaneProxies.Total)
			}
			r.Status.HostnameGenerators = []tfTypes.HostnameGenerators{}
			if len(r.Status.HostnameGenerators) > len(resp.Status.HostnameGenerators) {
				r.Status.HostnameGenerators = r.Status.HostnameGenerators[:len(resp.Status.HostnameGenerators)]
			}
			for hostnameGeneratorsCount, hostnameGeneratorsItem := range resp.Status.HostnameGenerators {
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
				if hostnameGeneratorsCount+1 > len(r.Status.HostnameGenerators) {
					r.Status.HostnameGenerators = append(r.Status.HostnameGenerators, hostnameGenerators)
				} else {
					r.Status.HostnameGenerators[hostnameGeneratorsCount].Conditions = hostnameGenerators.Conditions
					r.Status.HostnameGenerators[hostnameGeneratorsCount].HostnameGeneratorRef = hostnameGenerators.HostnameGeneratorRef
				}
			}
			if resp.Status.TLS == nil {
				r.Status.TLS = nil
			} else {
				r.Status.TLS = &tfTypes.MeshServiceItemTLS{}
				if resp.Status.TLS.Status != nil {
					r.Status.TLS.Status = types.StringValue(string(*resp.Status.TLS.Status))
				} else {
					r.Status.TLS.Status = types.StringNull()
				}
			}
			r.Status.Vips = []tfTypes.Vip{}
			if len(r.Status.Vips) > len(resp.Status.Vips) {
				r.Status.Vips = r.Status.Vips[:len(resp.Status.Vips)]
			}
			for vipsCount, vipsItem := range resp.Status.Vips {
				var vips tfTypes.Vip
				vips.IP = types.StringPointerValue(vipsItem.IP)
				if vipsCount+1 > len(r.Status.Vips) {
					r.Status.Vips = append(r.Status.Vips, vips)
				} else {
					r.Status.Vips[vipsCount].IP = vips.IP
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}

	return diags
}
