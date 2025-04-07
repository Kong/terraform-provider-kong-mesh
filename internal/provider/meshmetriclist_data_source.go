// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MeshMetricListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshMetricListDataSource{}

func NewMeshMetricListDataSource() datasource.DataSource {
	return &MeshMetricListDataSource{}
}

// MeshMetricListDataSource is the data source implementation.
type MeshMetricListDataSource struct {
	client *sdk.KongMesh
}

// MeshMetricListDataSourceModel describes the data model.
type MeshMetricListDataSourceModel struct {
	Items  []tfTypes.MeshMetricItem `tfsdk:"items"`
	Key    types.String             `queryParam:"name=key" tfsdk:"key"`
	Mesh   types.String             `tfsdk:"mesh"`
	Next   types.String             `tfsdk:"next"`
	Offset types.Int64              `queryParam:"style=form,explode=true,name=offset" tfsdk:"offset"`
	Size   types.Int64              `queryParam:"style=form,explode=true,name=size" tfsdk:"size"`
	Total  types.Number             `tfsdk:"total"`
	Value  types.String             `queryParam:"name=value" tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshMetricListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_metric_list"
}

// Schema defines the schema for the data source.
func (r *MeshMetricListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshMetricList DataSource",

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
										"applications": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"address": schema.StringAttribute{
														Computed:    true,
														Description: `Address on which an application listens.`,
													},
													"name": schema.StringAttribute{
														Computed:    true,
														Description: `Name of the application to scrape`,
													},
													"path": schema.StringAttribute{
														Computed:    true,
														Description: `Path on which an application expose HTTP endpoint with metrics.`,
													},
													"port": schema.Int32Attribute{
														Computed:    true,
														Description: `Port on which an application expose HTTP endpoint with metrics.`,
													},
												},
											},
											Description: `Applications is a list of application that Dataplane Proxy will scrape`,
										},
										"backends": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"open_telemetry": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"endpoint": schema.StringAttribute{
																Computed:    true,
																Description: `Endpoint for OpenTelemetry collector`,
															},
															"refresh_interval": schema.StringAttribute{
																Computed:    true,
																Description: `RefreshInterval defines how frequent metrics should be pushed to collector`,
															},
														},
														Description: `OpenTelemetry backend configuration`,
													},
													"prometheus": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"client_id": schema.StringAttribute{
																Computed:    true,
																Description: `ClientId of the Prometheus backend. Needed when using MADS for DP discovery.`,
															},
															"path": schema.StringAttribute{
																Computed:    true,
																Description: `Path on which a dataplane should expose HTTP endpoint with Prometheus metrics.`,
															},
															"port": schema.Int32Attribute{
																Computed:    true,
																Description: `Port on which a dataplane should expose HTTP endpoint with Prometheus metrics.`,
															},
															"tls": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"mode": schema.StringAttribute{
																		Computed:    true,
																		Description: `Configuration of TLS for Prometheus listener.`,
																	},
																},
																Description: `Configuration of TLS for prometheus listener.`,
															},
														},
														Description: `Prometheus backend configuration.`,
													},
													"type": schema.StringAttribute{
														Computed:    true,
														Description: `Type of the backend that will be used to collect metrics. At the moment only Prometheus backend is available.`,
													},
												},
											},
											Description: `Backends list that will be used to collect metrics.`,
										},
										"sidecar": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"include_unused": schema.BoolAttribute{
													Computed: true,
													MarkdownDescription: `IncludeUnused if false will scrape only metrics that has been by sidecar (counters incremented` + "\n" +
														`at least once, gauges changed at least once, and histograms added to at` + "\n" +
														`least once). If true will scrape all metrics (even the ones with zeros).` + "\n" +
														`If not specified then the default value is false.`,
												},
												"profiles": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"append_profiles": schema.ListNestedAttribute{
															Computed: true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Computed:    true,
																		Description: `Name of the predefined profile, one of: all, basic, none`,
																	},
																},
															},
															Description: `AppendProfiles allows to combine the metrics from multiple predefined profiles.`,
														},
														"exclude": schema.ListNestedAttribute{
															Computed: true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"match": schema.StringAttribute{
																		Computed:    true,
																		Description: `Match is the value used to match using particular Type`,
																	},
																	"type": schema.StringAttribute{
																		Computed:    true,
																		Description: `Type defined the type of selector, one of: prefix, regex, exact`,
																	},
																},
															},
															MarkdownDescription: `Exclude makes it possible to exclude groups of metrics from a resulting profile.` + "\n" +
																`Exclude is subordinate to Include.`,
														},
														"include": schema.ListNestedAttribute{
															Computed: true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"match": schema.StringAttribute{
																		Computed:    true,
																		Description: `Match is the value used to match using particular Type`,
																	},
																	"type": schema.StringAttribute{
																		Computed:    true,
																		Description: `Type defined the type of selector, one of: prefix, regex, exact`,
																	},
																},
															},
															MarkdownDescription: `Include makes it possible to include additional metrics in a selected profiles.` + "\n" +
																`Include takes precedence over Exclude.`,
														},
													},
													Description: `Profiles allows to customize which metrics are published.`,
												},
											},
											Description: `Sidecar metrics collection configuration`,
										},
									},
									Description: `MeshMetric configuration.`,
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
										`defined in-place.`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshMetric resource.`,
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
			"total": schema.NumberAttribute{
				Computed:    true,
				Description: `The total number of entities`,
			},
			"value": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *MeshMetricListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshMetricListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshMetricListDataSourceModel
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

	offset := new(int64)
	if !data.Offset.IsUnknown() && !data.Offset.IsNull() {
		*offset = data.Offset.ValueInt64()
	} else {
		offset = nil
	}
	size := new(int64)
	if !data.Size.IsUnknown() && !data.Size.IsNull() {
		*size = data.Size.ValueInt64()
	} else {
		size = nil
	}
	var filter *operations.GetMeshMetricListQueryParamFilter
	key := new(string)
	if !data.Key.IsUnknown() && !data.Key.IsNull() {
		*key = data.Key.ValueString()
	} else {
		key = nil
	}
	value := new(string)
	if !data.Value.IsUnknown() && !data.Value.IsNull() {
		*value = data.Value.ValueString()
	} else {
		value = nil
	}
	filter = &operations.GetMeshMetricListQueryParamFilter{
		Key:   key,
		Value: value,
	}
	var mesh string
	mesh = data.Mesh.ValueString()

	request := operations.GetMeshMetricListRequest{
		Offset: offset,
		Size:   size,
		Filter: filter,
		Mesh:   mesh,
	}
	res, err := r.client.MeshMetric.GetMeshMetricList(ctx, request)
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
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MeshMetricList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMeshMetricList(res.MeshMetricList)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
