// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"time"
)

func (r *MeshPassthroughDataSourceModel) RefreshFromSharedMeshPassthroughItem(resp *shared.MeshPassthroughItem) {
	if resp != nil {
		if resp.CreationTime != nil {
			r.CreationTime = types.StringValue(resp.CreationTime.Format(time.RFC3339Nano))
		} else {
			r.CreationTime = types.StringNull()
		}
		if len(resp.Labels) > 0 {
			r.Labels = make(map[string]types.String, len(resp.Labels))
			for key, value := range resp.Labels {
				r.Labels[key] = types.StringValue(value)
			}
		}
		r.Mesh = types.StringPointerValue(resp.Mesh)
		if resp.ModificationTime != nil {
			r.ModificationTime = types.StringValue(resp.ModificationTime.Format(time.RFC3339Nano))
		} else {
			r.ModificationTime = types.StringNull()
		}
		r.Name = types.StringValue(resp.Name)
		if resp.Spec.Default == nil {
			r.Spec.Default = nil
		} else {
			r.Spec.Default = &tfTypes.MeshPassthroughItemDefault{}
			r.Spec.Default.AppendMatch = []tfTypes.AppendMatch{}
			if len(r.Spec.Default.AppendMatch) > len(resp.Spec.Default.AppendMatch) {
				r.Spec.Default.AppendMatch = r.Spec.Default.AppendMatch[:len(resp.Spec.Default.AppendMatch)]
			}
			for appendMatchCount, appendMatchItem := range resp.Spec.Default.AppendMatch {
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
				if appendMatchCount+1 > len(r.Spec.Default.AppendMatch) {
					r.Spec.Default.AppendMatch = append(r.Spec.Default.AppendMatch, appendMatch1)
				} else {
					r.Spec.Default.AppendMatch[appendMatchCount].Port = appendMatch1.Port
					r.Spec.Default.AppendMatch[appendMatchCount].Protocol = appendMatch1.Protocol
					r.Spec.Default.AppendMatch[appendMatchCount].Type = appendMatch1.Type
					r.Spec.Default.AppendMatch[appendMatchCount].Value = appendMatch1.Value
				}
			}
			if resp.Spec.Default.PassthroughMode != nil {
				r.Spec.Default.PassthroughMode = types.StringValue(string(*resp.Spec.Default.PassthroughMode))
			} else {
				r.Spec.Default.PassthroughMode = types.StringNull()
			}
		}
		if resp.Spec.TargetRef == nil {
			r.Spec.TargetRef = nil
		} else {
			r.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
			r.Spec.TargetRef.Kind = types.StringValue(string(resp.Spec.TargetRef.Kind))
			if len(resp.Spec.TargetRef.Labels) > 0 {
				r.Spec.TargetRef.Labels = make(map[string]types.String, len(resp.Spec.TargetRef.Labels))
				for key1, value2 := range resp.Spec.TargetRef.Labels {
					r.Spec.TargetRef.Labels[key1] = types.StringValue(value2)
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
				for key2, value3 := range resp.Spec.TargetRef.Tags {
					r.Spec.TargetRef.Tags[key2] = types.StringValue(value3)
				}
			}
		}
		r.Type = types.StringValue(string(resp.Type))
	}
}
