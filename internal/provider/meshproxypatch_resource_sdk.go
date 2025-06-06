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

func (r *MeshProxyPatchResourceModel) ToSharedMeshProxyPatchItemInput(ctx context.Context) (*shared.MeshProxyPatchItemInput, diag.Diagnostics) {
	var diags diag.Diagnostics

	typeVar := shared.MeshProxyPatchItemType(r.Type.ValueString())
	mesh := new(string)
	if !r.Mesh.IsUnknown() && !r.Mesh.IsNull() {
		*mesh = r.Mesh.ValueString()
	} else {
		mesh = nil
	}
	var name string
	name = r.Name.ValueString()

	var labels map[string]string
	if !r.Labels.IsUnknown() && !r.Labels.IsNull() {
		diags.Append(r.Labels.ElementsAs(ctx, &labels, true)...)
	}
	appendModifications := make([]shared.AppendModifications, 0, len(r.Spec.Default.AppendModifications))
	for _, appendModificationsItem := range r.Spec.Default.AppendModifications {
		var cluster *shared.Cluster
		if appendModificationsItem.Cluster != nil {
			jsonPatches := make([]shared.JSONPatches, 0, len(appendModificationsItem.Cluster.JSONPatches))
			for _, jsonPatchesItem := range appendModificationsItem.Cluster.JSONPatches {
				from := new(string)
				if !jsonPatchesItem.From.IsUnknown() && !jsonPatchesItem.From.IsNull() {
					*from = jsonPatchesItem.From.ValueString()
				} else {
					from = nil
				}
				op := shared.MeshProxyPatchItemOp(jsonPatchesItem.Op.ValueString())
				var path string
				path = jsonPatchesItem.Path.ValueString()

				var value interface{}
				if !jsonPatchesItem.Value.IsUnknown() && !jsonPatchesItem.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem.Value.ValueString()), &value)
				}
				jsonPatches = append(jsonPatches, shared.JSONPatches{
					From:  from,
					Op:    op,
					Path:  path,
					Value: value,
				})
			}
			var match *shared.MeshProxyPatchItemMatch
			if appendModificationsItem.Cluster.Match != nil {
				name1 := new(string)
				if !appendModificationsItem.Cluster.Match.Name.IsUnknown() && !appendModificationsItem.Cluster.Match.Name.IsNull() {
					*name1 = appendModificationsItem.Cluster.Match.Name.ValueString()
				} else {
					name1 = nil
				}
				origin := new(string)
				if !appendModificationsItem.Cluster.Match.Origin.IsUnknown() && !appendModificationsItem.Cluster.Match.Origin.IsNull() {
					*origin = appendModificationsItem.Cluster.Match.Origin.ValueString()
				} else {
					origin = nil
				}
				match = &shared.MeshProxyPatchItemMatch{
					Name:   name1,
					Origin: origin,
				}
			}
			operation := shared.Operation(appendModificationsItem.Cluster.Operation.ValueString())
			value1 := new(string)
			if !appendModificationsItem.Cluster.Value.IsUnknown() && !appendModificationsItem.Cluster.Value.IsNull() {
				*value1 = appendModificationsItem.Cluster.Value.ValueString()
			} else {
				value1 = nil
			}
			cluster = &shared.Cluster{
				JSONPatches: jsonPatches,
				Match:       match,
				Operation:   operation,
				Value:       value1,
			}
		}
		var httpFilter *shared.HTTPFilter
		if appendModificationsItem.HTTPFilter != nil {
			jsonPatches1 := make([]shared.MeshProxyPatchItemJSONPatches, 0, len(appendModificationsItem.HTTPFilter.JSONPatches))
			for _, jsonPatchesItem1 := range appendModificationsItem.HTTPFilter.JSONPatches {
				from1 := new(string)
				if !jsonPatchesItem1.From.IsUnknown() && !jsonPatchesItem1.From.IsNull() {
					*from1 = jsonPatchesItem1.From.ValueString()
				} else {
					from1 = nil
				}
				op1 := shared.MeshProxyPatchItemSpecOp(jsonPatchesItem1.Op.ValueString())
				var path1 string
				path1 = jsonPatchesItem1.Path.ValueString()

				var value2 interface{}
				if !jsonPatchesItem1.Value.IsUnknown() && !jsonPatchesItem1.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem1.Value.ValueString()), &value2)
				}
				jsonPatches1 = append(jsonPatches1, shared.MeshProxyPatchItemJSONPatches{
					From:  from1,
					Op:    op1,
					Path:  path1,
					Value: value2,
				})
			}
			var match1 *shared.MeshProxyPatchItemSpecMatch
			if appendModificationsItem.HTTPFilter.Match != nil {
				listenerName := new(string)
				if !appendModificationsItem.HTTPFilter.Match.ListenerName.IsUnknown() && !appendModificationsItem.HTTPFilter.Match.ListenerName.IsNull() {
					*listenerName = appendModificationsItem.HTTPFilter.Match.ListenerName.ValueString()
				} else {
					listenerName = nil
				}
				listenerTags := make(map[string]string)
				for listenerTagsKey, listenerTagsValue := range appendModificationsItem.HTTPFilter.Match.ListenerTags {
					var listenerTagsInst string
					listenerTagsInst = listenerTagsValue.ValueString()

					listenerTags[listenerTagsKey] = listenerTagsInst
				}
				name2 := new(string)
				if !appendModificationsItem.HTTPFilter.Match.Name.IsUnknown() && !appendModificationsItem.HTTPFilter.Match.Name.IsNull() {
					*name2 = appendModificationsItem.HTTPFilter.Match.Name.ValueString()
				} else {
					name2 = nil
				}
				origin1 := new(string)
				if !appendModificationsItem.HTTPFilter.Match.Origin.IsUnknown() && !appendModificationsItem.HTTPFilter.Match.Origin.IsNull() {
					*origin1 = appendModificationsItem.HTTPFilter.Match.Origin.ValueString()
				} else {
					origin1 = nil
				}
				match1 = &shared.MeshProxyPatchItemSpecMatch{
					ListenerName: listenerName,
					ListenerTags: listenerTags,
					Name:         name2,
					Origin:       origin1,
				}
			}
			operation1 := shared.MeshProxyPatchItemOperation(appendModificationsItem.HTTPFilter.Operation.ValueString())
			value3 := new(string)
			if !appendModificationsItem.HTTPFilter.Value.IsUnknown() && !appendModificationsItem.HTTPFilter.Value.IsNull() {
				*value3 = appendModificationsItem.HTTPFilter.Value.ValueString()
			} else {
				value3 = nil
			}
			httpFilter = &shared.HTTPFilter{
				JSONPatches: jsonPatches1,
				Match:       match1,
				Operation:   operation1,
				Value:       value3,
			}
		}
		var listener *shared.Listener
		if appendModificationsItem.Listener != nil {
			jsonPatches2 := make([]shared.MeshProxyPatchItemSpecJSONPatches, 0, len(appendModificationsItem.Listener.JSONPatches))
			for _, jsonPatchesItem2 := range appendModificationsItem.Listener.JSONPatches {
				from2 := new(string)
				if !jsonPatchesItem2.From.IsUnknown() && !jsonPatchesItem2.From.IsNull() {
					*from2 = jsonPatchesItem2.From.ValueString()
				} else {
					from2 = nil
				}
				op2 := shared.MeshProxyPatchItemSpecDefaultOp(jsonPatchesItem2.Op.ValueString())
				var path2 string
				path2 = jsonPatchesItem2.Path.ValueString()

				var value4 interface{}
				if !jsonPatchesItem2.Value.IsUnknown() && !jsonPatchesItem2.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem2.Value.ValueString()), &value4)
				}
				jsonPatches2 = append(jsonPatches2, shared.MeshProxyPatchItemSpecJSONPatches{
					From:  from2,
					Op:    op2,
					Path:  path2,
					Value: value4,
				})
			}
			var match2 *shared.MeshProxyPatchItemSpecDefaultMatch
			if appendModificationsItem.Listener.Match != nil {
				name3 := new(string)
				if !appendModificationsItem.Listener.Match.Name.IsUnknown() && !appendModificationsItem.Listener.Match.Name.IsNull() {
					*name3 = appendModificationsItem.Listener.Match.Name.ValueString()
				} else {
					name3 = nil
				}
				origin2 := new(string)
				if !appendModificationsItem.Listener.Match.Origin.IsUnknown() && !appendModificationsItem.Listener.Match.Origin.IsNull() {
					*origin2 = appendModificationsItem.Listener.Match.Origin.ValueString()
				} else {
					origin2 = nil
				}
				tags := make(map[string]string)
				for tagsKey, tagsValue := range appendModificationsItem.Listener.Match.Tags {
					var tagsInst string
					tagsInst = tagsValue.ValueString()

					tags[tagsKey] = tagsInst
				}
				match2 = &shared.MeshProxyPatchItemSpecDefaultMatch{
					Name:   name3,
					Origin: origin2,
					Tags:   tags,
				}
			}
			operation2 := shared.MeshProxyPatchItemSpecOperation(appendModificationsItem.Listener.Operation.ValueString())
			value5 := new(string)
			if !appendModificationsItem.Listener.Value.IsUnknown() && !appendModificationsItem.Listener.Value.IsNull() {
				*value5 = appendModificationsItem.Listener.Value.ValueString()
			} else {
				value5 = nil
			}
			listener = &shared.Listener{
				JSONPatches: jsonPatches2,
				Match:       match2,
				Operation:   operation2,
				Value:       value5,
			}
		}
		var networkFilter *shared.NetworkFilter
		if appendModificationsItem.NetworkFilter != nil {
			jsonPatches3 := make([]shared.MeshProxyPatchItemSpecDefaultJSONPatches, 0, len(appendModificationsItem.NetworkFilter.JSONPatches))
			for _, jsonPatchesItem3 := range appendModificationsItem.NetworkFilter.JSONPatches {
				from3 := new(string)
				if !jsonPatchesItem3.From.IsUnknown() && !jsonPatchesItem3.From.IsNull() {
					*from3 = jsonPatchesItem3.From.ValueString()
				} else {
					from3 = nil
				}
				op3 := shared.MeshProxyPatchItemSpecDefaultAppendModificationsOp(jsonPatchesItem3.Op.ValueString())
				var path3 string
				path3 = jsonPatchesItem3.Path.ValueString()

				var value6 interface{}
				if !jsonPatchesItem3.Value.IsUnknown() && !jsonPatchesItem3.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem3.Value.ValueString()), &value6)
				}
				jsonPatches3 = append(jsonPatches3, shared.MeshProxyPatchItemSpecDefaultJSONPatches{
					From:  from3,
					Op:    op3,
					Path:  path3,
					Value: value6,
				})
			}
			var match3 *shared.MeshProxyPatchItemSpecDefaultAppendModificationsMatch
			if appendModificationsItem.NetworkFilter.Match != nil {
				listenerName1 := new(string)
				if !appendModificationsItem.NetworkFilter.Match.ListenerName.IsUnknown() && !appendModificationsItem.NetworkFilter.Match.ListenerName.IsNull() {
					*listenerName1 = appendModificationsItem.NetworkFilter.Match.ListenerName.ValueString()
				} else {
					listenerName1 = nil
				}
				listenerTags1 := make(map[string]string)
				for listenerTagsKey1, listenerTagsValue1 := range appendModificationsItem.NetworkFilter.Match.ListenerTags {
					var listenerTagsInst1 string
					listenerTagsInst1 = listenerTagsValue1.ValueString()

					listenerTags1[listenerTagsKey1] = listenerTagsInst1
				}
				name4 := new(string)
				if !appendModificationsItem.NetworkFilter.Match.Name.IsUnknown() && !appendModificationsItem.NetworkFilter.Match.Name.IsNull() {
					*name4 = appendModificationsItem.NetworkFilter.Match.Name.ValueString()
				} else {
					name4 = nil
				}
				origin3 := new(string)
				if !appendModificationsItem.NetworkFilter.Match.Origin.IsUnknown() && !appendModificationsItem.NetworkFilter.Match.Origin.IsNull() {
					*origin3 = appendModificationsItem.NetworkFilter.Match.Origin.ValueString()
				} else {
					origin3 = nil
				}
				match3 = &shared.MeshProxyPatchItemSpecDefaultAppendModificationsMatch{
					ListenerName: listenerName1,
					ListenerTags: listenerTags1,
					Name:         name4,
					Origin:       origin3,
				}
			}
			operation3 := shared.MeshProxyPatchItemSpecDefaultOperation(appendModificationsItem.NetworkFilter.Operation.ValueString())
			value7 := new(string)
			if !appendModificationsItem.NetworkFilter.Value.IsUnknown() && !appendModificationsItem.NetworkFilter.Value.IsNull() {
				*value7 = appendModificationsItem.NetworkFilter.Value.ValueString()
			} else {
				value7 = nil
			}
			networkFilter = &shared.NetworkFilter{
				JSONPatches: jsonPatches3,
				Match:       match3,
				Operation:   operation3,
				Value:       value7,
			}
		}
		var virtualHost *shared.VirtualHost
		if appendModificationsItem.VirtualHost != nil {
			jsonPatches4 := make([]shared.MeshProxyPatchItemSpecDefaultAppendModificationsJSONPatches, 0, len(appendModificationsItem.VirtualHost.JSONPatches))
			for _, jsonPatchesItem4 := range appendModificationsItem.VirtualHost.JSONPatches {
				from4 := new(string)
				if !jsonPatchesItem4.From.IsUnknown() && !jsonPatchesItem4.From.IsNull() {
					*from4 = jsonPatchesItem4.From.ValueString()
				} else {
					from4 = nil
				}
				op4 := shared.MeshProxyPatchItemSpecDefaultAppendModificationsVirtualHostOp(jsonPatchesItem4.Op.ValueString())
				var path4 string
				path4 = jsonPatchesItem4.Path.ValueString()

				var value8 interface{}
				if !jsonPatchesItem4.Value.IsUnknown() && !jsonPatchesItem4.Value.IsNull() {
					_ = json.Unmarshal([]byte(jsonPatchesItem4.Value.ValueString()), &value8)
				}
				jsonPatches4 = append(jsonPatches4, shared.MeshProxyPatchItemSpecDefaultAppendModificationsJSONPatches{
					From:  from4,
					Op:    op4,
					Path:  path4,
					Value: value8,
				})
			}
			name5 := new(string)
			if !appendModificationsItem.VirtualHost.Match.Name.IsUnknown() && !appendModificationsItem.VirtualHost.Match.Name.IsNull() {
				*name5 = appendModificationsItem.VirtualHost.Match.Name.ValueString()
			} else {
				name5 = nil
			}
			origin4 := new(string)
			if !appendModificationsItem.VirtualHost.Match.Origin.IsUnknown() && !appendModificationsItem.VirtualHost.Match.Origin.IsNull() {
				*origin4 = appendModificationsItem.VirtualHost.Match.Origin.ValueString()
			} else {
				origin4 = nil
			}
			routeConfigurationName := new(string)
			if !appendModificationsItem.VirtualHost.Match.RouteConfigurationName.IsUnknown() && !appendModificationsItem.VirtualHost.Match.RouteConfigurationName.IsNull() {
				*routeConfigurationName = appendModificationsItem.VirtualHost.Match.RouteConfigurationName.ValueString()
			} else {
				routeConfigurationName = nil
			}
			match4 := shared.MeshProxyPatchItemSpecDefaultAppendModificationsVirtualHostMatch{
				Name:                   name5,
				Origin:                 origin4,
				RouteConfigurationName: routeConfigurationName,
			}
			operation4 := shared.MeshProxyPatchItemSpecDefaultAppendModificationsOperation(appendModificationsItem.VirtualHost.Operation.ValueString())
			value9 := new(string)
			if !appendModificationsItem.VirtualHost.Value.IsUnknown() && !appendModificationsItem.VirtualHost.Value.IsNull() {
				*value9 = appendModificationsItem.VirtualHost.Value.ValueString()
			} else {
				value9 = nil
			}
			virtualHost = &shared.VirtualHost{
				JSONPatches: jsonPatches4,
				Match:       match4,
				Operation:   operation4,
				Value:       value9,
			}
		}
		appendModifications = append(appendModifications, shared.AppendModifications{
			Cluster:       cluster,
			HTTPFilter:    httpFilter,
			Listener:      listener,
			NetworkFilter: networkFilter,
			VirtualHost:   virtualHost,
		})
	}
	defaultVar := shared.MeshProxyPatchItemDefault{
		AppendModifications: appendModifications,
	}
	var targetRef *shared.MeshProxyPatchItemTargetRef
	if r.Spec.TargetRef != nil {
		kind := shared.MeshProxyPatchItemKind(r.Spec.TargetRef.Kind.ValueString())
		labels1 := make(map[string]string)
		for labelsKey, labelsValue := range r.Spec.TargetRef.Labels {
			var labelsInst string
			labelsInst = labelsValue.ValueString()

			labels1[labelsKey] = labelsInst
		}
		mesh1 := new(string)
		if !r.Spec.TargetRef.Mesh.IsUnknown() && !r.Spec.TargetRef.Mesh.IsNull() {
			*mesh1 = r.Spec.TargetRef.Mesh.ValueString()
		} else {
			mesh1 = nil
		}
		name6 := new(string)
		if !r.Spec.TargetRef.Name.IsUnknown() && !r.Spec.TargetRef.Name.IsNull() {
			*name6 = r.Spec.TargetRef.Name.ValueString()
		} else {
			name6 = nil
		}
		namespace := new(string)
		if !r.Spec.TargetRef.Namespace.IsUnknown() && !r.Spec.TargetRef.Namespace.IsNull() {
			*namespace = r.Spec.TargetRef.Namespace.ValueString()
		} else {
			namespace = nil
		}
		proxyTypes := make([]shared.MeshProxyPatchItemProxyTypes, 0, len(r.Spec.TargetRef.ProxyTypes))
		for _, proxyTypesItem := range r.Spec.TargetRef.ProxyTypes {
			proxyTypes = append(proxyTypes, shared.MeshProxyPatchItemProxyTypes(proxyTypesItem.ValueString()))
		}
		sectionName := new(string)
		if !r.Spec.TargetRef.SectionName.IsUnknown() && !r.Spec.TargetRef.SectionName.IsNull() {
			*sectionName = r.Spec.TargetRef.SectionName.ValueString()
		} else {
			sectionName = nil
		}
		tags1 := make(map[string]string)
		for tagsKey1, tagsValue1 := range r.Spec.TargetRef.Tags {
			var tagsInst1 string
			tagsInst1 = tagsValue1.ValueString()

			tags1[tagsKey1] = tagsInst1
		}
		targetRef = &shared.MeshProxyPatchItemTargetRef{
			Kind:        kind,
			Labels:      labels1,
			Mesh:        mesh1,
			Name:        name6,
			Namespace:   namespace,
			ProxyTypes:  proxyTypes,
			SectionName: sectionName,
			Tags:        tags1,
		}
	}
	spec := shared.MeshProxyPatchItemSpec{
		Default:   defaultVar,
		TargetRef: targetRef,
	}
	out := shared.MeshProxyPatchItemInput{
		Type:   typeVar,
		Mesh:   mesh,
		Name:   name,
		Labels: labels,
		Spec:   spec,
	}

	return &out, diags
}

func (r *MeshProxyPatchResourceModel) ToOperationsCreateMeshProxyPatchRequest(ctx context.Context) (*operations.CreateMeshProxyPatchRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	meshProxyPatchItem, meshProxyPatchItemDiags := r.ToSharedMeshProxyPatchItemInput(ctx)
	diags.Append(meshProxyPatchItemDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.CreateMeshProxyPatchRequest{
		Mesh:               mesh,
		Name:               name,
		MeshProxyPatchItem: *meshProxyPatchItem,
	}

	return &out, diags
}

func (r *MeshProxyPatchResourceModel) ToOperationsUpdateMeshProxyPatchRequest(ctx context.Context) (*operations.UpdateMeshProxyPatchRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	meshProxyPatchItem, meshProxyPatchItemDiags := r.ToSharedMeshProxyPatchItemInput(ctx)
	diags.Append(meshProxyPatchItemDiags...)

	if diags.HasError() {
		return nil, diags
	}

	out := operations.UpdateMeshProxyPatchRequest{
		Mesh:               mesh,
		Name:               name,
		MeshProxyPatchItem: *meshProxyPatchItem,
	}

	return &out, diags
}

func (r *MeshProxyPatchResourceModel) ToOperationsGetMeshProxyPatchRequest(ctx context.Context) (*operations.GetMeshProxyPatchRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	out := operations.GetMeshProxyPatchRequest{
		Mesh: mesh,
		Name: name,
	}

	return &out, diags
}

func (r *MeshProxyPatchResourceModel) ToOperationsDeleteMeshProxyPatchRequest(ctx context.Context) (*operations.DeleteMeshProxyPatchRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	var mesh string
	mesh = r.Mesh.ValueString()

	var name string
	name = r.Name.ValueString()

	out := operations.DeleteMeshProxyPatchRequest{
		Mesh: mesh,
		Name: name,
	}

	return &out, diags
}

func (r *MeshProxyPatchResourceModel) RefreshFromSharedMeshProxyPatchCreateOrUpdateSuccessResponse(ctx context.Context, resp *shared.MeshProxyPatchCreateOrUpdateSuccessResponse) diag.Diagnostics {
	var diags diag.Diagnostics

	if resp != nil {
		r.Warnings = make([]types.String, 0, len(resp.Warnings))
		for _, v := range resp.Warnings {
			r.Warnings = append(r.Warnings, types.StringValue(v))
		}
	}

	return diags
}

func (r *MeshProxyPatchResourceModel) RefreshFromSharedMeshProxyPatchItem(ctx context.Context, resp *shared.MeshProxyPatchItem) diag.Diagnostics {
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
		r.Spec.Default.AppendModifications = []tfTypes.AppendModifications{}
		if len(r.Spec.Default.AppendModifications) > len(resp.Spec.Default.AppendModifications) {
			r.Spec.Default.AppendModifications = r.Spec.Default.AppendModifications[:len(resp.Spec.Default.AppendModifications)]
		}
		for appendModificationsCount, appendModificationsItem := range resp.Spec.Default.AppendModifications {
			var appendModifications tfTypes.AppendModifications
			if appendModificationsItem.Cluster == nil {
				appendModifications.Cluster = nil
			} else {
				appendModifications.Cluster = &tfTypes.Cluster{}
				appendModifications.Cluster.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount, jsonPatchesItem := range appendModificationsItem.Cluster.JSONPatches {
					var jsonPatches tfTypes.JSONPatches
					jsonPatches.From = types.StringPointerValue(jsonPatchesItem.From)
					jsonPatches.Op = types.StringValue(string(jsonPatchesItem.Op))
					jsonPatches.Path = types.StringValue(jsonPatchesItem.Path)
					if jsonPatchesItem.Value == nil {
						jsonPatches.Value = types.StringNull()
					} else {
						valueResult, _ := json.Marshal(jsonPatchesItem.Value)
						jsonPatches.Value = types.StringValue(string(valueResult))
					}
					if jsonPatchesCount+1 > len(appendModifications.Cluster.JSONPatches) {
						appendModifications.Cluster.JSONPatches = append(appendModifications.Cluster.JSONPatches, jsonPatches)
					} else {
						appendModifications.Cluster.JSONPatches[jsonPatchesCount].From = jsonPatches.From
						appendModifications.Cluster.JSONPatches[jsonPatchesCount].Op = jsonPatches.Op
						appendModifications.Cluster.JSONPatches[jsonPatchesCount].Path = jsonPatches.Path
						appendModifications.Cluster.JSONPatches[jsonPatchesCount].Value = jsonPatches.Value
					}
				}
				if appendModificationsItem.Cluster.Match == nil {
					appendModifications.Cluster.Match = nil
				} else {
					appendModifications.Cluster.Match = &tfTypes.MeshProxyPatchItemMatch{}
					appendModifications.Cluster.Match.Name = types.StringPointerValue(appendModificationsItem.Cluster.Match.Name)
					appendModifications.Cluster.Match.Origin = types.StringPointerValue(appendModificationsItem.Cluster.Match.Origin)
				}
				appendModifications.Cluster.Operation = types.StringValue(string(appendModificationsItem.Cluster.Operation))
				appendModifications.Cluster.Value = types.StringPointerValue(appendModificationsItem.Cluster.Value)
			}
			if appendModificationsItem.HTTPFilter == nil {
				appendModifications.HTTPFilter = nil
			} else {
				appendModifications.HTTPFilter = &tfTypes.HTTPFilter{}
				appendModifications.HTTPFilter.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount1, jsonPatchesItem1 := range appendModificationsItem.HTTPFilter.JSONPatches {
					var jsonPatches1 tfTypes.JSONPatches
					jsonPatches1.From = types.StringPointerValue(jsonPatchesItem1.From)
					jsonPatches1.Op = types.StringValue(string(jsonPatchesItem1.Op))
					jsonPatches1.Path = types.StringValue(jsonPatchesItem1.Path)
					if jsonPatchesItem1.Value == nil {
						jsonPatches1.Value = types.StringNull()
					} else {
						valueResult1, _ := json.Marshal(jsonPatchesItem1.Value)
						jsonPatches1.Value = types.StringValue(string(valueResult1))
					}
					if jsonPatchesCount1+1 > len(appendModifications.HTTPFilter.JSONPatches) {
						appendModifications.HTTPFilter.JSONPatches = append(appendModifications.HTTPFilter.JSONPatches, jsonPatches1)
					} else {
						appendModifications.HTTPFilter.JSONPatches[jsonPatchesCount1].From = jsonPatches1.From
						appendModifications.HTTPFilter.JSONPatches[jsonPatchesCount1].Op = jsonPatches1.Op
						appendModifications.HTTPFilter.JSONPatches[jsonPatchesCount1].Path = jsonPatches1.Path
						appendModifications.HTTPFilter.JSONPatches[jsonPatchesCount1].Value = jsonPatches1.Value
					}
				}
				if appendModificationsItem.HTTPFilter.Match == nil {
					appendModifications.HTTPFilter.Match = nil
				} else {
					appendModifications.HTTPFilter.Match = &tfTypes.MeshProxyPatchItemSpecMatch{}
					appendModifications.HTTPFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.ListenerName)
					if len(appendModificationsItem.HTTPFilter.Match.ListenerTags) > 0 {
						appendModifications.HTTPFilter.Match.ListenerTags = make(map[string]types.String, len(appendModificationsItem.HTTPFilter.Match.ListenerTags))
						for key, value := range appendModificationsItem.HTTPFilter.Match.ListenerTags {
							appendModifications.HTTPFilter.Match.ListenerTags[key] = types.StringValue(value)
						}
					}
					appendModifications.HTTPFilter.Match.Name = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Name)
					appendModifications.HTTPFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.HTTPFilter.Match.Origin)
				}
				appendModifications.HTTPFilter.Operation = types.StringValue(string(appendModificationsItem.HTTPFilter.Operation))
				appendModifications.HTTPFilter.Value = types.StringPointerValue(appendModificationsItem.HTTPFilter.Value)
			}
			if appendModificationsItem.Listener == nil {
				appendModifications.Listener = nil
			} else {
				appendModifications.Listener = &tfTypes.Listener{}
				appendModifications.Listener.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount2, jsonPatchesItem2 := range appendModificationsItem.Listener.JSONPatches {
					var jsonPatches2 tfTypes.JSONPatches
					jsonPatches2.From = types.StringPointerValue(jsonPatchesItem2.From)
					jsonPatches2.Op = types.StringValue(string(jsonPatchesItem2.Op))
					jsonPatches2.Path = types.StringValue(jsonPatchesItem2.Path)
					if jsonPatchesItem2.Value == nil {
						jsonPatches2.Value = types.StringNull()
					} else {
						valueResult2, _ := json.Marshal(jsonPatchesItem2.Value)
						jsonPatches2.Value = types.StringValue(string(valueResult2))
					}
					if jsonPatchesCount2+1 > len(appendModifications.Listener.JSONPatches) {
						appendModifications.Listener.JSONPatches = append(appendModifications.Listener.JSONPatches, jsonPatches2)
					} else {
						appendModifications.Listener.JSONPatches[jsonPatchesCount2].From = jsonPatches2.From
						appendModifications.Listener.JSONPatches[jsonPatchesCount2].Op = jsonPatches2.Op
						appendModifications.Listener.JSONPatches[jsonPatchesCount2].Path = jsonPatches2.Path
						appendModifications.Listener.JSONPatches[jsonPatchesCount2].Value = jsonPatches2.Value
					}
				}
				if appendModificationsItem.Listener.Match == nil {
					appendModifications.Listener.Match = nil
				} else {
					appendModifications.Listener.Match = &tfTypes.MeshProxyPatchItemSpecDefaultMatch{}
					appendModifications.Listener.Match.Name = types.StringPointerValue(appendModificationsItem.Listener.Match.Name)
					appendModifications.Listener.Match.Origin = types.StringPointerValue(appendModificationsItem.Listener.Match.Origin)
					if len(appendModificationsItem.Listener.Match.Tags) > 0 {
						appendModifications.Listener.Match.Tags = make(map[string]types.String, len(appendModificationsItem.Listener.Match.Tags))
						for key1, value1 := range appendModificationsItem.Listener.Match.Tags {
							appendModifications.Listener.Match.Tags[key1] = types.StringValue(value1)
						}
					}
				}
				appendModifications.Listener.Operation = types.StringValue(string(appendModificationsItem.Listener.Operation))
				appendModifications.Listener.Value = types.StringPointerValue(appendModificationsItem.Listener.Value)
			}
			if appendModificationsItem.NetworkFilter == nil {
				appendModifications.NetworkFilter = nil
			} else {
				appendModifications.NetworkFilter = &tfTypes.HTTPFilter{}
				appendModifications.NetworkFilter.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount3, jsonPatchesItem3 := range appendModificationsItem.NetworkFilter.JSONPatches {
					var jsonPatches3 tfTypes.JSONPatches
					jsonPatches3.From = types.StringPointerValue(jsonPatchesItem3.From)
					jsonPatches3.Op = types.StringValue(string(jsonPatchesItem3.Op))
					jsonPatches3.Path = types.StringValue(jsonPatchesItem3.Path)
					if jsonPatchesItem3.Value == nil {
						jsonPatches3.Value = types.StringNull()
					} else {
						valueResult3, _ := json.Marshal(jsonPatchesItem3.Value)
						jsonPatches3.Value = types.StringValue(string(valueResult3))
					}
					if jsonPatchesCount3+1 > len(appendModifications.NetworkFilter.JSONPatches) {
						appendModifications.NetworkFilter.JSONPatches = append(appendModifications.NetworkFilter.JSONPatches, jsonPatches3)
					} else {
						appendModifications.NetworkFilter.JSONPatches[jsonPatchesCount3].From = jsonPatches3.From
						appendModifications.NetworkFilter.JSONPatches[jsonPatchesCount3].Op = jsonPatches3.Op
						appendModifications.NetworkFilter.JSONPatches[jsonPatchesCount3].Path = jsonPatches3.Path
						appendModifications.NetworkFilter.JSONPatches[jsonPatchesCount3].Value = jsonPatches3.Value
					}
				}
				if appendModificationsItem.NetworkFilter.Match == nil {
					appendModifications.NetworkFilter.Match = nil
				} else {
					appendModifications.NetworkFilter.Match = &tfTypes.MeshProxyPatchItemSpecMatch{}
					appendModifications.NetworkFilter.Match.ListenerName = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.ListenerName)
					if len(appendModificationsItem.NetworkFilter.Match.ListenerTags) > 0 {
						appendModifications.NetworkFilter.Match.ListenerTags = make(map[string]types.String, len(appendModificationsItem.NetworkFilter.Match.ListenerTags))
						for key2, value2 := range appendModificationsItem.NetworkFilter.Match.ListenerTags {
							appendModifications.NetworkFilter.Match.ListenerTags[key2] = types.StringValue(value2)
						}
					}
					appendModifications.NetworkFilter.Match.Name = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Name)
					appendModifications.NetworkFilter.Match.Origin = types.StringPointerValue(appendModificationsItem.NetworkFilter.Match.Origin)
				}
				appendModifications.NetworkFilter.Operation = types.StringValue(string(appendModificationsItem.NetworkFilter.Operation))
				appendModifications.NetworkFilter.Value = types.StringPointerValue(appendModificationsItem.NetworkFilter.Value)
			}
			if appendModificationsItem.VirtualHost == nil {
				appendModifications.VirtualHost = nil
			} else {
				appendModifications.VirtualHost = &tfTypes.VirtualHost{}
				appendModifications.VirtualHost.JSONPatches = []tfTypes.JSONPatches{}
				for jsonPatchesCount4, jsonPatchesItem4 := range appendModificationsItem.VirtualHost.JSONPatches {
					var jsonPatches4 tfTypes.JSONPatches
					jsonPatches4.From = types.StringPointerValue(jsonPatchesItem4.From)
					jsonPatches4.Op = types.StringValue(string(jsonPatchesItem4.Op))
					jsonPatches4.Path = types.StringValue(jsonPatchesItem4.Path)
					if jsonPatchesItem4.Value == nil {
						jsonPatches4.Value = types.StringNull()
					} else {
						valueResult4, _ := json.Marshal(jsonPatchesItem4.Value)
						jsonPatches4.Value = types.StringValue(string(valueResult4))
					}
					if jsonPatchesCount4+1 > len(appendModifications.VirtualHost.JSONPatches) {
						appendModifications.VirtualHost.JSONPatches = append(appendModifications.VirtualHost.JSONPatches, jsonPatches4)
					} else {
						appendModifications.VirtualHost.JSONPatches[jsonPatchesCount4].From = jsonPatches4.From
						appendModifications.VirtualHost.JSONPatches[jsonPatchesCount4].Op = jsonPatches4.Op
						appendModifications.VirtualHost.JSONPatches[jsonPatchesCount4].Path = jsonPatches4.Path
						appendModifications.VirtualHost.JSONPatches[jsonPatchesCount4].Value = jsonPatches4.Value
					}
				}
				appendModifications.VirtualHost.Match.Name = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Name)
				appendModifications.VirtualHost.Match.Origin = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.Origin)
				appendModifications.VirtualHost.Match.RouteConfigurationName = types.StringPointerValue(appendModificationsItem.VirtualHost.Match.RouteConfigurationName)
				appendModifications.VirtualHost.Operation = types.StringValue(string(appendModificationsItem.VirtualHost.Operation))
				appendModifications.VirtualHost.Value = types.StringPointerValue(appendModificationsItem.VirtualHost.Value)
			}
			if appendModificationsCount+1 > len(r.Spec.Default.AppendModifications) {
				r.Spec.Default.AppendModifications = append(r.Spec.Default.AppendModifications, appendModifications)
			} else {
				r.Spec.Default.AppendModifications[appendModificationsCount].Cluster = appendModifications.Cluster
				r.Spec.Default.AppendModifications[appendModificationsCount].HTTPFilter = appendModifications.HTTPFilter
				r.Spec.Default.AppendModifications[appendModificationsCount].Listener = appendModifications.Listener
				r.Spec.Default.AppendModifications[appendModificationsCount].NetworkFilter = appendModifications.NetworkFilter
				r.Spec.Default.AppendModifications[appendModificationsCount].VirtualHost = appendModifications.VirtualHost
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
				for key3, value3 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key3] = types.StringValue(value3)
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
				for key4, value4 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key4] = types.StringValue(value4)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}

	return diags
}
