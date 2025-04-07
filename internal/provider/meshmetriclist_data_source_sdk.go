// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"math/big"
	"time"
)

func (r *MeshMetricListDataSourceModel) RefreshFromSharedMeshMetricList(resp *shared.MeshMetricList) {
	if resp != nil {
		r.Items = []tfTypes.MeshMetricItem{}
		if len(r.Items) > len(resp.Items) {
			r.Items = r.Items[:len(resp.Items)]
		}
		for itemsCount, itemsItem := range resp.Items {
			var items1 tfTypes.MeshMetricItem
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
				items1.Spec.Default = &tfTypes.Default{}
				items1.Spec.Default.Applications = []tfTypes.Applications{}
				for applicationsCount, applicationsItem := range itemsItem.Spec.Default.Applications {
					var applications1 tfTypes.Applications
					applications1.Address = types.StringPointerValue(applicationsItem.Address)
					applications1.Name = types.StringPointerValue(applicationsItem.Name)
					applications1.Path = types.StringPointerValue(applicationsItem.Path)
					applications1.Port = types.Int32Value(int32(applicationsItem.Port))
					if applicationsCount+1 > len(items1.Spec.Default.Applications) {
						items1.Spec.Default.Applications = append(items1.Spec.Default.Applications, applications1)
					} else {
						items1.Spec.Default.Applications[applicationsCount].Address = applications1.Address
						items1.Spec.Default.Applications[applicationsCount].Name = applications1.Name
						items1.Spec.Default.Applications[applicationsCount].Path = applications1.Path
						items1.Spec.Default.Applications[applicationsCount].Port = applications1.Port
					}
				}
				items1.Spec.Default.Backends = []tfTypes.MeshMetricItemBackends{}
				for backendsCount, backendsItem := range itemsItem.Spec.Default.Backends {
					var backends1 tfTypes.MeshMetricItemBackends
					if backendsItem.OpenTelemetry == nil {
						backends1.OpenTelemetry = nil
					} else {
						backends1.OpenTelemetry = &tfTypes.OpenTelemetry{}
						backends1.OpenTelemetry.Endpoint = types.StringValue(backendsItem.OpenTelemetry.Endpoint)
						backends1.OpenTelemetry.RefreshInterval = types.StringPointerValue(backendsItem.OpenTelemetry.RefreshInterval)
					}
					if backendsItem.Prometheus == nil {
						backends1.Prometheus = nil
					} else {
						backends1.Prometheus = &tfTypes.Prometheus{}
						backends1.Prometheus.ClientID = types.StringPointerValue(backendsItem.Prometheus.ClientID)
						backends1.Prometheus.Path = types.StringPointerValue(backendsItem.Prometheus.Path)
						if backendsItem.Prometheus.Port != nil {
							backends1.Prometheus.Port = types.Int32Value(int32(*backendsItem.Prometheus.Port))
						} else {
							backends1.Prometheus.Port = types.Int32Null()
						}
						if backendsItem.Prometheus.TLS == nil {
							backends1.Prometheus.TLS = nil
						} else {
							backends1.Prometheus.TLS = &tfTypes.MeshMetricItemTLS{}
							if backendsItem.Prometheus.TLS.Mode != nil {
								backends1.Prometheus.TLS.Mode = types.StringValue(string(*backendsItem.Prometheus.TLS.Mode))
							} else {
								backends1.Prometheus.TLS.Mode = types.StringNull()
							}
						}
					}
					backends1.Type = types.StringValue(string(backendsItem.Type))
					if backendsCount+1 > len(items1.Spec.Default.Backends) {
						items1.Spec.Default.Backends = append(items1.Spec.Default.Backends, backends1)
					} else {
						items1.Spec.Default.Backends[backendsCount].OpenTelemetry = backends1.OpenTelemetry
						items1.Spec.Default.Backends[backendsCount].Prometheus = backends1.Prometheus
						items1.Spec.Default.Backends[backendsCount].Type = backends1.Type
					}
				}
				if itemsItem.Spec.Default.Sidecar == nil {
					items1.Spec.Default.Sidecar = nil
				} else {
					items1.Spec.Default.Sidecar = &tfTypes.Sidecar{}
					items1.Spec.Default.Sidecar.IncludeUnused = types.BoolPointerValue(itemsItem.Spec.Default.Sidecar.IncludeUnused)
					if itemsItem.Spec.Default.Sidecar.Profiles == nil {
						items1.Spec.Default.Sidecar.Profiles = nil
					} else {
						items1.Spec.Default.Sidecar.Profiles = &tfTypes.Profiles{}
						items1.Spec.Default.Sidecar.Profiles.AppendProfiles = []tfTypes.MeshLoadBalancingStrategyItemSpecHeader{}
						for appendProfilesCount, appendProfilesItem := range itemsItem.Spec.Default.Sidecar.Profiles.AppendProfiles {
							var appendProfiles1 tfTypes.MeshLoadBalancingStrategyItemSpecHeader
							appendProfiles1.Name = types.StringValue(string(appendProfilesItem.Name))
							if appendProfilesCount+1 > len(items1.Spec.Default.Sidecar.Profiles.AppendProfiles) {
								items1.Spec.Default.Sidecar.Profiles.AppendProfiles = append(items1.Spec.Default.Sidecar.Profiles.AppendProfiles, appendProfiles1)
							} else {
								items1.Spec.Default.Sidecar.Profiles.AppendProfiles[appendProfilesCount].Name = appendProfiles1.Name
							}
						}
						items1.Spec.Default.Sidecar.Profiles.Exclude = []tfTypes.Exclude{}
						for excludeCount, excludeItem := range itemsItem.Spec.Default.Sidecar.Profiles.Exclude {
							var exclude1 tfTypes.Exclude
							exclude1.Match = types.StringValue(excludeItem.Match)
							exclude1.Type = types.StringValue(string(excludeItem.Type))
							if excludeCount+1 > len(items1.Spec.Default.Sidecar.Profiles.Exclude) {
								items1.Spec.Default.Sidecar.Profiles.Exclude = append(items1.Spec.Default.Sidecar.Profiles.Exclude, exclude1)
							} else {
								items1.Spec.Default.Sidecar.Profiles.Exclude[excludeCount].Match = exclude1.Match
								items1.Spec.Default.Sidecar.Profiles.Exclude[excludeCount].Type = exclude1.Type
							}
						}
						items1.Spec.Default.Sidecar.Profiles.Include = []tfTypes.Exclude{}
						for includeCount, includeItem := range itemsItem.Spec.Default.Sidecar.Profiles.Include {
							var include1 tfTypes.Exclude
							include1.Match = types.StringValue(includeItem.Match)
							include1.Type = types.StringValue(string(includeItem.Type))
							if includeCount+1 > len(items1.Spec.Default.Sidecar.Profiles.Include) {
								items1.Spec.Default.Sidecar.Profiles.Include = append(items1.Spec.Default.Sidecar.Profiles.Include, include1)
							} else {
								items1.Spec.Default.Sidecar.Profiles.Include[includeCount].Match = include1.Match
								items1.Spec.Default.Sidecar.Profiles.Include[includeCount].Type = include1.Type
							}
						}
					}
				}
			}
			if itemsItem.Spec.TargetRef == nil {
				items1.Spec.TargetRef = nil
			} else {
				items1.Spec.TargetRef = &tfTypes.MeshAccessLogItemTargetRef{}
				items1.Spec.TargetRef.Kind = types.StringValue(string(itemsItem.Spec.TargetRef.Kind))
				if len(itemsItem.Spec.TargetRef.Labels) > 0 {
					items1.Spec.TargetRef.Labels = make(map[string]types.String, len(itemsItem.Spec.TargetRef.Labels))
					for key1, value1 := range itemsItem.Spec.TargetRef.Labels {
						items1.Spec.TargetRef.Labels[key1] = types.StringValue(value1)
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
					for key2, value2 := range itemsItem.Spec.TargetRef.Tags {
						items1.Spec.TargetRef.Tags[key2] = types.StringValue(value2)
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
