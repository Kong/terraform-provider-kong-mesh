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

func (r *MeshMultiZoneServiceDataSourceModel) ToOperationsGetMeshMultiZoneServiceRequest(ctx context.Context) (*operations.GetMeshMultiZoneServiceRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	out := operations.GetMeshMultiZoneServiceRequest{
		Mesh: mesh,
		Name: name,
	}

	return &out, diags
}

func (r *MeshMultiZoneServiceDataSourceModel) RefreshFromSharedMeshMultiZoneServiceItem(ctx context.Context, resp *shared.MeshMultiZoneServiceItem) diag.Diagnostics {
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
		r.Spec.Ports = []tfTypes.Ports{}
		if len(r.Spec.Ports) > len(resp.Spec.Ports) {
			r.Spec.Ports = r.Spec.Ports[:len(resp.Spec.Ports)]
		}
		for portsCount, portsItem := range resp.Spec.Ports {
			var ports tfTypes.Ports
			ports.AppProtocol = types.StringPointerValue(portsItem.AppProtocol)
			ports.Name = types.StringPointerValue(portsItem.Name)
			ports.Port = types.Int32Value(int32(portsItem.Port))
			if portsCount+1 > len(r.Spec.Ports) {
				r.Spec.Ports = append(r.Spec.Ports, ports)
			} else {
				r.Spec.Ports[portsCount].AppProtocol = ports.AppProtocol
				r.Spec.Ports[portsCount].Name = ports.Name
				r.Spec.Ports[portsCount].Port = ports.Port
			}
		}
		if len(resp.Spec.Selector.MeshService.MatchLabels) > 0 {
			r.Spec.Selector.MeshService.MatchLabels = make(map[string]types.String, len(resp.Spec.Selector.MeshService.MatchLabels))
			for key, value := range resp.Spec.Selector.MeshService.MatchLabels {
				r.Spec.Selector.MeshService.MatchLabels[key] = types.StringValue(value)
			}
		}
		if resp.Status == nil {
			r.Status = nil
		} else {
			r.Status = &tfTypes.MeshMultiZoneServiceItemStatus{}
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
			r.Status.MeshServices = []tfTypes.MeshMultiZoneServiceItemMeshServices{}
			if len(r.Status.MeshServices) > len(resp.Status.MeshServices) {
				r.Status.MeshServices = r.Status.MeshServices[:len(resp.Status.MeshServices)]
			}
			for meshServicesCount, meshServicesItem := range resp.Status.MeshServices {
				var meshServices tfTypes.MeshMultiZoneServiceItemMeshServices
				meshServices.Mesh = types.StringValue(meshServicesItem.Mesh)
				meshServices.Name = types.StringValue(meshServicesItem.Name)
				meshServices.Namespace = types.StringValue(meshServicesItem.Namespace)
				meshServices.Zone = types.StringValue(meshServicesItem.Zone)
				if meshServicesCount+1 > len(r.Status.MeshServices) {
					r.Status.MeshServices = append(r.Status.MeshServices, meshServices)
				} else {
					r.Status.MeshServices[meshServicesCount].Mesh = meshServices.Mesh
					r.Status.MeshServices[meshServicesCount].Name = meshServices.Name
					r.Status.MeshServices[meshServicesCount].Namespace = meshServices.Namespace
					r.Status.MeshServices[meshServicesCount].Zone = meshServices.Zone
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
