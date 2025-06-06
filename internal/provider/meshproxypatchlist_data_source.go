// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/Kong/shared-speakeasy/customtypes/kumalabels"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MeshProxyPatchListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshProxyPatchListDataSource{}

func NewMeshProxyPatchListDataSource() datasource.DataSource {
	return &MeshProxyPatchListDataSource{}
}

// MeshProxyPatchListDataSource is the data source implementation.
type MeshProxyPatchListDataSource struct {
	client *sdk.KongMesh
}

// MeshProxyPatchListDataSourceModel describes the data model.
type MeshProxyPatchListDataSourceModel struct {
	Items  []tfTypes.MeshProxyPatchItem `tfsdk:"items"`
	Key    types.String                 `queryParam:"name=key" tfsdk:"key"`
	Mesh   types.String                 `tfsdk:"mesh"`
	Next   types.String                 `tfsdk:"next"`
	Offset types.Int64                  `queryParam:"style=form,explode=true,name=offset" tfsdk:"offset"`
	Size   types.Int64                  `queryParam:"style=form,explode=true,name=size" tfsdk:"size"`
	Total  types.Float64                `tfsdk:"total"`
	Value  types.String                 `queryParam:"name=value" tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshProxyPatchListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_proxy_patch_list"
}

// Schema defines the schema for the data source.
func (r *MeshProxyPatchListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshProxyPatchList DataSource",

		Attributes: map[string]schema.Attribute{
			"items": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"creation_time": schema.StringAttribute{
							Computed:    true,
							Description: `Time at which the resource was created`,
						},
						"labels": schema.MapAttribute{
							CustomType:  kumalabels.KumaLabelsMapType{MapType: types.MapType{ElemType: types.StringType}},
							Computed:    true,
							ElementType: types.StringType,
							Description: `The labels to help identity resources`,
						},
						"mesh": schema.StringAttribute{
							Computed:    true,
							Description: `Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.`,
						},
						"modification_time": schema.StringAttribute{
							Computed:    true,
							Description: `Time at which the resource was updated`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `Name of the Kuma resource`,
						},
						"spec": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"append_modifications": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cluster": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"json_patches": schema.ListNestedAttribute{
																Computed: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"from": schema.StringAttribute{
																			Computed:    true,
																			Description: `From is a jsonpatch from string, used by move and copy operations.`,
																		},
																		"op": schema.StringAttribute{
																			Computed:    true,
																			Description: `Op is a jsonpatch operation string.`,
																		},
																		"path": schema.StringAttribute{
																			Computed:    true,
																			Description: `Path is a jsonpatch path string.`,
																		},
																		"value": schema.StringAttribute{
																			Computed:    true,
																			Description: `Value must be a valid json value used by replace and add operations. Parsed as JSON.`,
																		},
																	},
																},
																MarkdownDescription: `JsonPatches specifies list of jsonpatches to apply to on Envoy's Cluster` + "\n" +
																	`resource`,
															},
															"match": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the cluster to match.`,
																	},
																	"origin": schema.StringAttribute{
																		Computed: true,
																		MarkdownDescription: `Origin is the name of the component or plugin that generated the resource.` + "\n" +
																			`` + "\n" +
																			`Here is the list of well-known origins:` + "\n" +
																			`inbound - resources generated for handling incoming traffic.` + "\n" +
																			`outbound - resources generated for handling outgoing traffic.` + "\n" +
																			`transparent - resources generated for transparent proxy functionality.` + "\n" +
																			`prometheus - resources generated when Prometheus metrics are enabled.` + "\n" +
																			`direct-access - resources generated for Direct Access functionality.` + "\n" +
																			`ingress - resources generated for Zone Ingress.` + "\n" +
																			`egress - resources generated for Zone Egress.` + "\n" +
																			`gateway - resources generated for MeshGateway.` + "\n" +
																			`` + "\n" +
																			`The list is not complete, because policy plugins can introduce new resources.` + "\n" +
																			`For example MeshTrace plugin can create Cluster with "mesh-trace" origin.`,
																	},
																},
																Description: `Match is a set of conditions that have to be matched for modification operation to happen.`,
															},
															"operation": schema.StringAttribute{
																Computed:    true,
																Description: `Operation to execute on matched cluster.`,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value of xDS resource in YAML format to add or patch.`,
															},
														},
														Description: `Cluster is a modification of Envoy's Cluster resource.`,
													},
													"http_filter": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"json_patches": schema.ListNestedAttribute{
																Computed: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"from": schema.StringAttribute{
																			Computed:    true,
																			Description: `From is a jsonpatch from string, used by move and copy operations.`,
																		},
																		"op": schema.StringAttribute{
																			Computed:    true,
																			Description: `Op is a jsonpatch operation string.`,
																		},
																		"path": schema.StringAttribute{
																			Computed:    true,
																			Description: `Path is a jsonpatch path string.`,
																		},
																		"value": schema.StringAttribute{
																			Computed:    true,
																			Description: `Value must be a valid json value used by replace and add operations. Parsed as JSON.`,
																		},
																	},
																},
																MarkdownDescription: `JsonPatches specifies list of jsonpatches to apply to on Envoy's` + "\n" +
																	`HTTP Filter available in HTTP Connection Manager in a Listener resource.`,
															},
															"match": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"listener_name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the listener to match.`,
																	},
																	"listener_tags": schema.MapAttribute{
																		Computed:    true,
																		ElementType: types.StringType,
																		Description: `Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]`,
																	},
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the HTTP filter. For example "envoy.filters.http.local_ratelimit"`,
																	},
																	"origin": schema.StringAttribute{
																		Computed: true,
																		MarkdownDescription: `Origin is the name of the component or plugin that generated the resource.` + "\n" +
																			`` + "\n" +
																			`Here is the list of well-known origins:` + "\n" +
																			`inbound - resources generated for handling incoming traffic.` + "\n" +
																			`outbound - resources generated for handling outgoing traffic.` + "\n" +
																			`transparent - resources generated for transparent proxy functionality.` + "\n" +
																			`prometheus - resources generated when Prometheus metrics are enabled.` + "\n" +
																			`direct-access - resources generated for Direct Access functionality.` + "\n" +
																			`ingress - resources generated for Zone Ingress.` + "\n" +
																			`egress - resources generated for Zone Egress.` + "\n" +
																			`gateway - resources generated for MeshGateway.` + "\n" +
																			`` + "\n" +
																			`The list is not complete, because policy plugins can introduce new resources.` + "\n" +
																			`For example MeshTrace plugin can create Cluster with "mesh-trace" origin.`,
																	},
																},
																Description: `Match is a set of conditions that have to be matched for modification operation to happen.`,
															},
															"operation": schema.StringAttribute{
																Computed:    true,
																Description: `Operation to execute on matched listener.`,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value of xDS resource in YAML format to add or patch.`,
															},
														},
														MarkdownDescription: `HTTPFilter is a modification of Envoy HTTP Filter` + "\n" +
															`available in HTTP Connection Manager in a Listener resource.`,
													},
													"listener": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"json_patches": schema.ListNestedAttribute{
																Computed: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"from": schema.StringAttribute{
																			Computed:    true,
																			Description: `From is a jsonpatch from string, used by move and copy operations.`,
																		},
																		"op": schema.StringAttribute{
																			Computed:    true,
																			Description: `Op is a jsonpatch operation string.`,
																		},
																		"path": schema.StringAttribute{
																			Computed:    true,
																			Description: `Path is a jsonpatch path string.`,
																		},
																		"value": schema.StringAttribute{
																			Computed:    true,
																			Description: `Value must be a valid json value used by replace and add operations. Parsed as JSON.`,
																		},
																	},
																},
																MarkdownDescription: `JsonPatches specifies list of jsonpatches to apply to on Envoy's Listener` + "\n" +
																	`resource`,
															},
															"match": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the listener to match.`,
																	},
																	"origin": schema.StringAttribute{
																		Computed: true,
																		MarkdownDescription: `Origin is the name of the component or plugin that generated the resource.` + "\n" +
																			`` + "\n" +
																			`Here is the list of well-known origins:` + "\n" +
																			`inbound - resources generated for handling incoming traffic.` + "\n" +
																			`outbound - resources generated for handling outgoing traffic.` + "\n" +
																			`transparent - resources generated for transparent proxy functionality.` + "\n" +
																			`prometheus - resources generated when Prometheus metrics are enabled.` + "\n" +
																			`direct-access - resources generated for Direct Access functionality.` + "\n" +
																			`ingress - resources generated for Zone Ingress.` + "\n" +
																			`egress - resources generated for Zone Egress.` + "\n" +
																			`gateway - resources generated for MeshGateway.` + "\n" +
																			`` + "\n" +
																			`The list is not complete, because policy plugins can introduce new resources.` + "\n" +
																			`For example MeshTrace plugin can create Cluster with "mesh-trace" origin.`,
																	},
																	"tags": schema.MapAttribute{
																		Computed:    true,
																		ElementType: types.StringType,
																		Description: `Tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]`,
																	},
																},
																Description: `Match is a set of conditions that have to be matched for modification operation to happen.`,
															},
															"operation": schema.StringAttribute{
																Computed:    true,
																Description: `Operation to execute on matched listener.`,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value of xDS resource in YAML format to add or patch.`,
															},
														},
														Description: `Listener is a modification of Envoy's Listener resource.`,
													},
													"network_filter": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"json_patches": schema.ListNestedAttribute{
																Computed: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"from": schema.StringAttribute{
																			Computed:    true,
																			Description: `From is a jsonpatch from string, used by move and copy operations.`,
																		},
																		"op": schema.StringAttribute{
																			Computed:    true,
																			Description: `Op is a jsonpatch operation string.`,
																		},
																		"path": schema.StringAttribute{
																			Computed:    true,
																			Description: `Path is a jsonpatch path string.`,
																		},
																		"value": schema.StringAttribute{
																			Computed:    true,
																			Description: `Value must be a valid json value used by replace and add operations. Parsed as JSON.`,
																		},
																	},
																},
																MarkdownDescription: `JsonPatches specifies list of jsonpatches to apply to on Envoy Listener's` + "\n" +
																	`filter.`,
															},
															"match": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"listener_name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the listener to match.`,
																	},
																	"listener_tags": schema.MapAttribute{
																		Computed:    true,
																		ElementType: types.StringType,
																		Description: `Listener tags available in Listener#Metadata#FilterMetadata[io.kuma.tags]`,
																	},
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the network filter. For example "envoy.filters.network.ratelimit"`,
																	},
																	"origin": schema.StringAttribute{
																		Computed: true,
																		MarkdownDescription: `Origin is the name of the component or plugin that generated the resource.` + "\n" +
																			`` + "\n" +
																			`Here is the list of well-known origins:` + "\n" +
																			`inbound - resources generated for handling incoming traffic.` + "\n" +
																			`outbound - resources generated for handling outgoing traffic.` + "\n" +
																			`transparent - resources generated for transparent proxy functionality.` + "\n" +
																			`prometheus - resources generated when Prometheus metrics are enabled.` + "\n" +
																			`direct-access - resources generated for Direct Access functionality.` + "\n" +
																			`ingress - resources generated for Zone Ingress.` + "\n" +
																			`egress - resources generated for Zone Egress.` + "\n" +
																			`gateway - resources generated for MeshGateway.` + "\n" +
																			`` + "\n" +
																			`The list is not complete, because policy plugins can introduce new resources.` + "\n" +
																			`For example MeshTrace plugin can create Cluster with "mesh-trace" origin.`,
																	},
																},
																Description: `Match is a set of conditions that have to be matched for modification operation to happen.`,
															},
															"operation": schema.StringAttribute{
																Computed:    true,
																Description: `Operation to execute on matched listener.`,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value of xDS resource in YAML format to add or patch.`,
															},
														},
														Description: `NetworkFilter is a modification of Envoy Listener's filter.`,
													},
													"virtual_host": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"json_patches": schema.ListNestedAttribute{
																Computed: true,
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"from": schema.StringAttribute{
																			Computed:    true,
																			Description: `From is a jsonpatch from string, used by move and copy operations.`,
																		},
																		"op": schema.StringAttribute{
																			Computed:    true,
																			Description: `Op is a jsonpatch operation string.`,
																		},
																		"path": schema.StringAttribute{
																			Computed:    true,
																			Description: `Path is a jsonpatch path string.`,
																		},
																		"value": schema.StringAttribute{
																			Computed:    true,
																			Description: `Value must be a valid json value used by replace and add operations. Parsed as JSON.`,
																		},
																	},
																},
																MarkdownDescription: `JsonPatches specifies list of jsonpatches to apply to on Envoy's` + "\n" +
																	`VirtualHost resource`,
															},
															"match": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the VirtualHost to match.`,
																	},
																	"origin": schema.StringAttribute{
																		Computed: true,
																		MarkdownDescription: `Origin is the name of the component or plugin that generated the resource.` + "\n" +
																			`` + "\n" +
																			`Here is the list of well-known origins:` + "\n" +
																			`inbound - resources generated for handling incoming traffic.` + "\n" +
																			`outbound - resources generated for handling outgoing traffic.` + "\n" +
																			`transparent - resources generated for transparent proxy functionality.` + "\n" +
																			`prometheus - resources generated when Prometheus metrics are enabled.` + "\n" +
																			`direct-access - resources generated for Direct Access functionality.` + "\n" +
																			`ingress - resources generated for Zone Ingress.` + "\n" +
																			`egress - resources generated for Zone Egress.` + "\n" +
																			`gateway - resources generated for MeshGateway.` + "\n" +
																			`` + "\n" +
																			`The list is not complete, because policy plugins can introduce new resources.` + "\n" +
																			`For example MeshTrace plugin can create Cluster with "mesh-trace" origin.`,
																	},
																	"route_configuration_name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the RouteConfiguration resource to match.`,
																	},
																},
																Description: `Match is a set of conditions that have to be matched for modification operation to happen.`,
															},
															"operation": schema.StringAttribute{
																Computed:    true,
																Description: `Operation to execute on matched listener.`,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value of xDS resource in YAML format to add or patch.`,
															},
														},
														MarkdownDescription: `VirtualHost is a modification of Envoy's VirtualHost` + "\n" +
															`referenced in HTTP Connection Manager in a Listener resource.`,
													},
												},
											},
											Description: `AppendModifications is a list of modifications applied on the selected proxy.`,
										},
									},
									MarkdownDescription: `Default is a configuration specific to the group of destinations` + "\n" +
										`referenced in 'targetRef'.`,
								},
								"target_ref": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Computed:    true,
											Description: `Kind of the referenced resource`,
										},
										"labels": schema.MapAttribute{
											Computed:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
												`Name and Namespace can be used.`,
										},
										"mesh": schema.StringAttribute{
											Computed:    true,
											Description: `Mesh is reserved for future use to identify cross mesh resources.`,
										},
										"name": schema.StringAttribute{
											Computed: true,
											MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
												`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
										},
										"namespace": schema.StringAttribute{
											Computed: true,
											MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
												`will be targeted.`,
										},
										"proxy_types": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
											MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
												`all data plane types are targeted by the policy.`,
										},
										"section_name": schema.StringAttribute{
											Computed: true,
											MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
												`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
										},
										"tags": schema.MapAttribute{
											Computed:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
												`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
										},
									},
									MarkdownDescription: `TargetRef is a reference to the resource the policy takes an effect on.` + "\n" +
										`The resource could be either a real store object or virtual resource` + "\n" +
										`defined inplace.`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshProxyPatch resource.`,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: `the type of the resource`,
						},
					},
				},
			},
			"key": schema.StringAttribute{
				Optional: true,
			},
			"mesh": schema.StringAttribute{
				Required:    true,
				Description: `name of the mesh`,
			},
			"next": schema.StringAttribute{
				Computed:    true,
				Description: `URL to the next page`,
			},
			"offset": schema.Int64Attribute{
				Optional:    true,
				Description: `offset in the list of entities`,
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Description: `the number of items per page`,
			},
			"total": schema.Float64Attribute{
				Computed:    true,
				Description: `The total number of entities`,
			},
			"value": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *MeshProxyPatchListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KongMesh)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.KongMesh, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *MeshProxyPatchListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshProxyPatchListDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsGetMeshProxyPatchListRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.MeshProxyPatch.GetMeshProxyPatchList(ctx, *request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MeshProxyPatchList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedMeshProxyPatchList(ctx, res.MeshProxyPatchList)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
