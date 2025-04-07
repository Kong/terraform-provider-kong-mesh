// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshPassthroughListDataSourceModel) RefreshFromSharedMeshPassthroughList(resp *shared.MeshPassthroughList) {
	if resp != nil {
		r.Items = []tfTypes.MeshPassthroughItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshPassthroughItem
			if itemsItem.CreationTime != nil {
				items1.CreationTime = types.StringValue(itemsItem.CreationTime.Format(time.RFC3339Nano))
			} else {
				items1.CreationTime = types.StringNull()
			}
			if len(itemsItem.Labels) > 0 {
				items1.Labels = make(map[string]types.String, len(itemsItem.Labels))
				for key, value := range itemsItem.Labels {
					items1.Labels[key] = types.StringValue(value)
				}
			}
			items1.Mesh = types.StringPointerValue(itemsItem.Mesh)
			if itemsItem.ModificationTime != nil {
				items1.ModificationTime = types.StringValue(itemsItem.ModificationTime.Format(time.RFC3339Nano))
			} else {
				items1.ModificationTime = types.StringNull()
			}
			items1.Name = types.StringValue(itemsItem.Name)
			if itemsItem.Spec.Default == nil {
				items1.Spec.Default = nil
			} else {
				items1.Spec.Default = &tfTypes.MeshPassthroughItemDefault{}
				items1.Spec.Default.AppendMatch = []tfTypes.AppendMatch{}
				for appendMatchCount, appendMatchItem := range itemsItem.Spec.Default.AppendMatch {
					var appendMatch1 tfTypes.AppendMatch
					if appendMatchItem.Port != nil {
						appendMatch1.Port = types.Int32Value(int32(*appendMatchItem.Port))
					} else {
						appendMatch1.Port = types.Int32Null()
					}
					if appendMatchItem.Protocol != nil {
						appendMatch1.Protocol = types.StringValue(string(*appendMatchItem.Protocol))
					} else {
						appendMatch1.Protocol = types.StringNull()
					}
					appendMatch1.Type = types.StringValue(string(appendMatchItem.Type))
					appendMatch1.Value = types.StringValue(appendMatchItem.Value)
					if appendMatchCount+1 > len(items1.Spec.Default.AppendMatch) {
						items1.Spec.Default.AppendMatch = append(items1.Spec.Default.AppendMatch, appendMatch1)
					} else {
						items1.Spec.Default.AppendMatch[appendMatchCount].Port = appendMatch1.Port
						items1.Spec.Default.AppendMatch[appendMatchCount].Protocol = appendMatch1.Protocol
						items1.Spec.Default.AppendMatch[appendMatchCount].Type = appendMatch1.Type
						items1.Spec.Default.AppendMatch[appendMatchCount].Value = appendMatch1.Value
					}
				}
				if itemsItem.Spec.Default.PassthroughMode != nil {
					items1.Spec.Default.PassthroughMode = types.StringValue(string(*itemsItem.Spec.Default.PassthroughMode))
				} else {
					items1.Spec.Default.PassthroughMode = types.StringNull()
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items1.Spec.TargetRef = nil
			} else {
				items1.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				items1.Spec.TargetRef.Kind = types.StringValue(string(itemsItem.Spec.TargetRef.Kind))
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items1.Spec.TargetRef.Labels = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Labels))
					for key1, value2 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key1] = types.StringValue(value2)
					}
				}
				items1.Spec.TargetRef.Mesh = types.StringPointerValue(itemsItem.Spec.TargetRef.Mesh)
				items1.Spec.TargetRef.Name = types.StringPointerValue(itemsItem.Spec.TargetRef.Name)
				items1.Spec.TargetRef.Namespace = types.StringPointerValue(itemsItem.Spec.TargetRef.Namespace)
				items1.Spec.TargetRef.ProxyTypes = make([]types.String, 0, len(itemsItem.Spec.TargetRef.ProxyTypes))
				for _, v := range itemsItem.Spec.TargetRef.ProxyTypes {
					items1.Spec.TargetRef.ProxyTypes = append(items1.Spec.TargetRef.ProxyTypes, types.StringValue(string(v)))
				}
				items1.Spec.TargetRef.SectionName = types.StringPointerValue(itemsItem.Spec.TargetRef.SectionName)
				if len(itemsItem.Spec.TargetRef.Tags) > 0 {
					items1.Spec.TargetRef.Tags = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Tags))
					for key2, value3 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key2] = types.StringValue(value3)
					}
				}
			}
			items1.Type = types.StringValue(string(itemsItem.Type))
			if itemsCount+1 > len(r.Items) {
				r.Items = append(r.Items, items1)
			} else {
				r.Items[itemsCount].CreationTime = items1.CreationTime
				r.Items[itemsCount].Labels = items1.Labels
				r.Items[itemsCount].Mesh = items1.Mesh
				r.Items[itemsCount].ModificationTime = items1.ModificationTime
				r.Items[itemsCount].Name = items1.Name
				r.Items[itemsCount].Spec = items1.Spec
				r.Items[itemsCount].Type = items1.Type
			}
		}
		r.Next = types.StringPointerValue(resp.Next)
		if resp.Total != nil {
			r.Total = types.NumberValue(big.NewFloat(float64(*resp.Total)))
		} else {
			r.Total = types.NumberNull()
		}
	}
}
