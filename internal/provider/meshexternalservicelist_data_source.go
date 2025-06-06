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
var _ datasource.DataSource = &MeshExternalServiceListDataSource{}
var _ datasource.DataSourceWithConfigure = &MeshExternalServiceListDataSource{}

func NewMeshExternalServiceListDataSource() datasource.DataSource {
	return &MeshExternalServiceListDataSource{}
}

// MeshExternalServiceListDataSource is the data source implementation.
type MeshExternalServiceListDataSource struct {
	client *sdk.KongMesh
}

// MeshExternalServiceListDataSourceModel describes the data model.
type MeshExternalServiceListDataSourceModel struct {
	Items  []tfTypes.MeshExternalServiceItem `tfsdk:"items"`
	Key    types.String                      `queryParam:"name=key" tfsdk:"key"`
	Mesh   types.String                      `tfsdk:"mesh"`
	Next   types.String                      `tfsdk:"next"`
	Offset types.Int64                       `queryParam:"style=form,explode=true,name=offset" tfsdk:"offset"`
	Size   types.Int64                       `queryParam:"style=form,explode=true,name=size" tfsdk:"size"`
	Total  types.Float64                     `tfsdk:"total"`
	Value  types.String                      `queryParam:"name=value" tfsdk:"value"`
}

// Metadata returns the data source type name.
func (r *MeshExternalServiceListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_external_service_list"
}

// Schema defines the schema for the data source.
func (r *MeshExternalServiceListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshExternalServiceList DataSource",

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
								"endpoints": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Computed:    true,
												Description: `Address defines an address to which a user want to send a request. Is possible to provide ` + "`" + `domain` + "`" + `, ` + "`" + `ip` + "`" + `.`,
											},
											"port": schema.Int64Attribute{
												Computed:    true,
												Description: `Port of the endpoint`,
											},
										},
									},
									Description: `Endpoints defines a list of destinations to send traffic to.`,
								},
								"extension": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"config": schema.StringAttribute{
											Computed:    true,
											Description: `Config freeform configuration for the extension. Parsed as JSON.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `Type of the extension.`,
										},
									},
									Description: `Extension struct for a plugin configuration, in the presence of an extension ` + "`" + `endpoints` + "`" + ` and ` + "`" + `tls` + "`" + ` are not required anymore - it's up to the extension to validate them independently.`,
								},
								"match": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"port": schema.Int64Attribute{
											Computed:    true,
											Description: `Port defines a port to which a user does request.`,
										},
										"protocol": schema.StringAttribute{
											Computed:    true,
											Description: `Protocol defines a protocol of the communication. Possible values: ` + "`" + `tcp` + "`" + `, ` + "`" + `grpc` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `http2` + "`" + `.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `Type of the match, only ` + "`" + `HostnameGenerator` + "`" + ` is available at the moment.`,
										},
									},
									Description: `Match defines traffic that should be routed through the sidecar.`,
								},
								"tls": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"allow_renegotiation": schema.BoolAttribute{
											Computed: true,
											MarkdownDescription: `AllowRenegotiation defines if TLS sessions will allow renegotiation.` + "\n" +
												`Setting this to true is not recommended for security reasons.`,
										},
										"enabled": schema.BoolAttribute{
											Computed:    true,
											Description: `Enabled defines if proxy should originate TLS.`,
										},
										"verification": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"ca_cert": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"inline": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is inline bytes.`,
														},
														"inline_string": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is inline string` + "`" + ``,
														},
														"secret": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is a secret with given Secret key.`,
														},
													},
													Description: `CaCert defines a certificate of CA.`,
												},
												"client_cert": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"inline": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is inline bytes.`,
														},
														"inline_string": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is inline string` + "`" + ``,
														},
														"secret": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is a secret with given Secret key.`,
														},
													},
													Description: `ClientCert defines a certificate of a client.`,
												},
												"client_key": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"inline": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is inline bytes.`,
														},
														"inline_string": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is inline string` + "`" + ``,
														},
														"secret": schema.StringAttribute{
															Computed:    true,
															Description: `Data source is a secret with given Secret key.`,
														},
													},
													Description: `ClientKey defines a client private key.`,
												},
												"mode": schema.StringAttribute{
													Computed:    true,
													Description: `Mode defines if proxy should skip verification, one of ` + "`" + `SkipSAN` + "`" + `, ` + "`" + `SkipCA` + "`" + `, ` + "`" + `Secured` + "`" + `, ` + "`" + `SkipAll` + "`" + `. Default ` + "`" + `Secured` + "`" + `.`,
												},
												"server_name": schema.StringAttribute{
													Computed:    true,
													Description: `ServerName overrides the default Server Name Indicator set by Kuma.`,
												},
												"subject_alt_names": schema.ListNestedAttribute{
													Computed: true,
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"type": schema.StringAttribute{
																Computed:    true,
																Description: `Type specifies matching type, one of ` + "`" + `Exact` + "`" + `, ` + "`" + `Prefix` + "`" + `. Default: ` + "`" + `Exact` + "`" + ``,
															},
															"value": schema.StringAttribute{
																Computed:    true,
																Description: `Value to match.`,
															},
														},
													},
													Description: `SubjectAltNames list of names to verify in the certificate.`,
												},
											},
											Description: `Verification section for providing TLS verification details.`,
										},
										"version": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"max": schema.StringAttribute{
													Computed:    true,
													Description: `Max defines maximum supported version. One of ` + "`" + `TLSAuto` + "`" + `, ` + "`" + `TLS10` + "`" + `, ` + "`" + `TLS11` + "`" + `, ` + "`" + `TLS12` + "`" + `, ` + "`" + `TLS13` + "`" + `.`,
												},
												"min": schema.StringAttribute{
													Computed:    true,
													Description: `Min defines minimum supported version. One of ` + "`" + `TLSAuto` + "`" + `, ` + "`" + `TLS10` + "`" + `, ` + "`" + `TLS11` + "`" + `, ` + "`" + `TLS12` + "`" + `, ` + "`" + `TLS13` + "`" + `.`,
												},
											},
											Description: `Version section for providing version specification.`,
										},
									},
									Description: `Tls provides a TLS configuration when proxy is resposible for a TLS origination`,
								},
							},
							Description: `Spec is the specification of the Kuma MeshExternalService resource.`,
						},
						"status": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"addresses": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"hostname": schema.StringAttribute{
												Computed: true,
											},
											"hostname_generator_ref": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"core_name": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											"origin": schema.StringAttribute{
												Computed: true,
											},
										},
									},
									Description: `Addresses section for generated domains`,
								},
								"hostname_generators": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"conditions": schema.ListNestedAttribute{
												Computed: true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"message": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `message is a human readable message indicating details about the transition.` + "\n" +
																`This may be an empty string.`,
														},
														"reason": schema.StringAttribute{
															Computed: true,
															MarkdownDescription: `reason contains a programmatic identifier indicating the reason for the condition's last transition.` + "\n" +
																`Producers of specific condition types may define expected values and meanings for this field,` + "\n" +
																`and whether the values are considered a guaranteed API.` + "\n" +
																`The value should be a CamelCase string.` + "\n" +
																`This field may not be empty.`,
														},
														"status": schema.StringAttribute{
															Computed:    true,
															Description: `status of the condition, one of True, False, Unknown.`,
														},
														"type": schema.StringAttribute{
															Computed:    true,
															Description: `type of condition in CamelCase or in foo.example.com/CamelCase.`,
														},
													},
												},
												Description: `Conditions is an array of hostname generator conditions.`,
											},
											"hostname_generator_ref": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"core_name": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
								"vip": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Computed:    true,
											Description: `Value allocated IP for a provided domain with ` + "`" + `HostnameGenerator` + "`" + ` type in a match section.`,
										},
									},
									Description: `Vip section for allocated IP`,
								},
							},
							Description: `Status is the current status of the Kuma MeshExternalService resource.`,
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

func (r *MeshExternalServiceListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *MeshExternalServiceListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *MeshExternalServiceListDataSourceModel
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

	request, requestDiags := data.ToOperationsGetMeshExternalServiceListRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.MeshExternalService.GetMeshExternalServiceList(ctx, *request)
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
	if !(res.MeshExternalServiceList != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedMeshExternalServiceList(ctx, res.MeshExternalServiceList)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
