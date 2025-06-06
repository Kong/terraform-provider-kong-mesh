// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Kong/shared-speakeasy/customtypes/kumalabels"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	custom_listplanmodifier "github.com/kong/terraform-provider-kong-mesh/internal/planmodifiers/listplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-kong-mesh/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
	"github.com/kong/terraform-provider-kong-mesh/internal/validators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-kong-mesh/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-kong-mesh/internal/validators/stringvalidators"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshHealthCheckResource{}
var _ resource.ResourceWithImportState = &MeshHealthCheckResource{}

func NewMeshHealthCheckResource() resource.Resource {
	return &MeshHealthCheckResource{}
}

// MeshHealthCheckResource defines the resource implementation.
type MeshHealthCheckResource struct {
	client *sdk.KongMesh
}

// MeshHealthCheckResourceModel describes the resource data model.
type MeshHealthCheckResourceModel struct {
	CreationTime     types.String                    `tfsdk:"creation_time"`
	Labels           kumalabels.KumaLabelsMapValue   `tfsdk:"labels"`
	Mesh             types.String                    `tfsdk:"mesh"`
	ModificationTime types.String                    `tfsdk:"modification_time"`
	Name             types.String                    `tfsdk:"name"`
	Spec             tfTypes.MeshHealthCheckItemSpec `tfsdk:"spec"`
	Type             types.String                    `tfsdk:"type"`
	Warnings         []types.String                  `tfsdk:"warnings"`
}

func (r *MeshHealthCheckResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_health_check"
}

func (r *MeshHealthCheckResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshHealthCheck Resource",
		Attributes: map[string]schema.Attribute{
			"creation_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Time at which the resource was created`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"labels": schema.MapAttribute{
				CustomType:  kumalabels.KumaLabelsMapType{MapType: types.MapType{ElemType: types.StringType}},
				Optional:    true,
				ElementType: types.StringType,
				Description: `The labels to help identity resources`,
			},
			"mesh": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `name of the mesh. Requires replacement if changed.`,
			},
			"modification_time": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Time at which the resource was updated`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `name of the MeshHealthCheck. Requires replacement if changed.`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"target_ref": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Required:    true,
								Description: `Kind of the referenced resource. must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]`,
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Mesh",
										"MeshSubset",
										"MeshGateway",
										"MeshService",
										"MeshExternalService",
										"MeshMultiZoneService",
										"MeshServiceSubset",
										"MeshHTTPRoute",
										"Dataplane",
									),
								},
							},
							"labels": schema.MapAttribute{
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
									`Name and Namespace can be used.`,
							},
							"mesh": schema.StringAttribute{
								Optional:    true,
								Description: `Mesh is reserved for future use to identify cross mesh resources.`,
							},
							"name": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
									`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
							},
							"namespace": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
									`will be targeted.`,
							},
							"proxy_types": schema.ListAttribute{
								Optional: true,
								PlanModifiers: []planmodifier.List{
									custom_listplanmodifier.SupressZeroNullModifier(),
								},
								ElementType: types.StringType,
								MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
									`all data plane types are targeted by the policy.`,
							},
							"section_name": schema.StringAttribute{
								Optional: true,
								MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
									`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
							},
							"tags": schema.MapAttribute{
								Optional:    true,
								ElementType: types.StringType,
								MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
									`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
							},
						},
						MarkdownDescription: `TargetRef is a reference to the resource the policy takes an effect on.` + "\n" +
							`The resource could be either a real store object or virtual resource` + "\n" +
							`defined inplace.`,
					},
					"to": schema.ListNestedAttribute{
						Optional: true,
						PlanModifiers: []planmodifier.List{
							custom_listplanmodifier.SupressZeroNullModifier(),
						},
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"default": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"always_log_health_check_failures": schema.BoolAttribute{
											Optional: true,
											MarkdownDescription: `If set to true, health check failure events will always be logged. If set` + "\n" +
												`to false, only the initial health check failure event will be logged. The` + "\n" +
												`default value is false.`,
										},
										"event_log_path": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `Specifies the path to the file where Envoy can log health check events.` + "\n" +
												`If empty, no event log will be written.`,
										},
										"fail_traffic_on_panic": schema.BoolAttribute{
											Optional: true,
											MarkdownDescription: `If set to true, Envoy will not consider any hosts when the cluster is in` + "\n" +
												`'panic mode'. Instead, the cluster will fail all requests as if all hosts` + "\n" +
												`are unhealthy. This can help avoid potentially overwhelming a failing` + "\n" +
												`service.`,
										},
										"grpc": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"authority": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `The value of the :authority header in the gRPC health check request,` + "\n" +
														`by default name of the cluster this health check is associated with`,
												},
												"disabled": schema.BoolAttribute{
													Optional:    true,
													Description: `If true the GrpcHealthCheck is disabled`,
												},
												"service_name": schema.StringAttribute{
													Optional:    true,
													Description: `Service name parameter which will be sent to gRPC service`,
												},
											},
											MarkdownDescription: `GrpcHealthCheck defines gRPC configuration which will instruct the service` + "\n" +
												`the health check will be made for is a gRPC service.`,
										},
										"healthy_panic_threshold": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"integer": schema.Int64Attribute{
													Optional: true,
													Validators: []validator.Int64{
														int64validator.ConflictsWith(path.Expressions{
															path.MatchRelative().AtParent().AtName("str"),
														}...),
													},
												},
												"str": schema.StringAttribute{
													Optional: true,
													Validators: []validator.String{
														stringvalidator.ConflictsWith(path.Expressions{
															path.MatchRelative().AtParent().AtName("integer"),
														}...),
													},
												},
											},
											MarkdownDescription: `Allows to configure panic threshold for Envoy cluster. If not specified,` + "\n" +
												`the default is 50%. To disable panic mode, set to 0%.` + "\n" +
												`Either int or decimal represented as string.` + "\n" +
												`Deprecated: the setting has been moved to MeshCircuitBreaker policy,` + "\n" +
												`please use MeshCircuitBreaker policy instead.`,
										},
										"healthy_threshold": schema.Int32Attribute{
											Optional: true,
											MarkdownDescription: `Number of consecutive healthy checks before considering a host healthy.` + "\n" +
												`If not specified then the default value is 1`,
										},
										"http": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"disabled": schema.BoolAttribute{
													Optional:    true,
													Description: `If true the HttpHealthCheck is disabled`,
												},
												"expected_statuses": schema.ListAttribute{
													Optional: true,
													PlanModifiers: []planmodifier.List{
														custom_listplanmodifier.SupressZeroNullModifier(),
													},
													ElementType: types.Int64Type,
													Description: `List of HTTP response statuses which are considered healthy`,
												},
												"path": schema.StringAttribute{
													Optional: true,
													MarkdownDescription: `The HTTP path which will be requested during the health check` + "\n" +
														`(ie. /health)` + "\n" +
														`If not specified then the default value is "/"`,
												},
												"request_headers_to_add": schema.SingleNestedAttribute{
													Optional: true,
													Attributes: map[string]schema.Attribute{
														"add": schema.ListNestedAttribute{
															Optional: true,
															PlanModifiers: []planmodifier.List{
																custom_listplanmodifier.SupressZeroNullModifier(),
															},
															NestedObject: schema.NestedAttributeObject{
																Validators: []validator.Object{
																	speakeasy_objectvalidators.NotNull(),
																},
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Optional:    true,
																		Description: `Not Null`,
																		Validators: []validator.String{
																			speakeasy_stringvalidators.NotNull(),
																			stringvalidator.UTF8LengthBetween(1, 256),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), "must match pattern "+regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`).String()),
																		},
																	},
																	"value": schema.StringAttribute{
																		Optional:    true,
																		Description: `Not Null`,
																		Validators: []validator.String{
																			speakeasy_stringvalidators.NotNull(),
																		},
																	},
																},
															},
															Validators: []validator.List{
																listvalidator.SizeAtMost(16),
															},
														},
														"set": schema.ListNestedAttribute{
															Optional: true,
															PlanModifiers: []planmodifier.List{
																custom_listplanmodifier.SupressZeroNullModifier(),
															},
															NestedObject: schema.NestedAttributeObject{
																Validators: []validator.Object{
																	speakeasy_objectvalidators.NotNull(),
																},
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Optional:    true,
																		Description: `Not Null`,
																		Validators: []validator.String{
																			speakeasy_stringvalidators.NotNull(),
																			stringvalidator.UTF8LengthBetween(1, 256),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`), "must match pattern "+regexp.MustCompile(`^[a-z0-9!#$%&'*+\-.^_\x60|~]+$`).String()),
																		},
																	},
																	"value": schema.StringAttribute{
																		Optional:    true,
																		Description: `Not Null`,
																		Validators: []validator.String{
																			speakeasy_stringvalidators.NotNull(),
																		},
																	},
																},
															},
															Validators: []validator.List{
																listvalidator.SizeAtMost(16),
															},
														},
													},
													MarkdownDescription: `The list of HTTP headers which should be added to each health check` + "\n" +
														`request`,
												},
											},
											MarkdownDescription: `HttpHealthCheck defines HTTP configuration which will instruct the service` + "\n" +
												`the health check will be made for is an HTTP service.`,
										},
										"initial_jitter": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `If specified, Envoy will start health checking after a random time in` + "\n" +
												`ms between 0 and initialJitter. This only applies to the first health` + "\n" +
												`check.`,
										},
										"interval": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `Interval between consecutive health checks.` + "\n" +
												`If not specified then the default value is 1m`,
										},
										"interval_jitter": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `If specified, during every interval Envoy will add IntervalJitter to the` + "\n" +
												`wait time.`,
										},
										"interval_jitter_percent": schema.Int32Attribute{
											Optional: true,
											MarkdownDescription: `If specified, during every interval Envoy will add IntervalJitter *` + "\n" +
												`IntervalJitterPercent / 100 to the wait time. If IntervalJitter and` + "\n" +
												`IntervalJitterPercent are both set, both of them will be used to` + "\n" +
												`increase the wait time.`,
										},
										"no_traffic_interval": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `The "no traffic interval" is a special health check interval that is used` + "\n" +
												`when a cluster has never had traffic routed to it. This lower interval` + "\n" +
												`allows cluster information to be kept up to date, without sending a` + "\n" +
												`potentially large amount of active health checking traffic for no reason.` + "\n" +
												`Once a cluster has been used for traffic routing, Envoy will shift back` + "\n" +
												`to using the standard health check interval that is defined. Note that` + "\n" +
												`this interval takes precedence over any other. The default value for "no` + "\n" +
												`traffic interval" is 60 seconds.`,
										},
										"reuse_connection": schema.BoolAttribute{
											Optional:    true,
											Description: `Reuse health check connection between health checks. Default is true.`,
										},
										"tcp": schema.SingleNestedAttribute{
											Optional: true,
											Attributes: map[string]schema.Attribute{
												"disabled": schema.BoolAttribute{
													Optional:    true,
													Description: `If true the TcpHealthCheck is disabled`,
												},
												"receive": schema.ListAttribute{
													Optional: true,
													PlanModifiers: []planmodifier.List{
														custom_listplanmodifier.SupressZeroNullModifier(),
													},
													ElementType: types.StringType,
													MarkdownDescription: `List of Base64 encoded blocks of strings expected as a response. When checking the response,` + "\n" +
														`"fuzzy" matching is performed such that each block must be found, and` + "\n" +
														`in the order specified, but not necessarily contiguous.` + "\n" +
														`If not provided or empty, checks will be performed as "connect only" and be marked as successful when TCP connection is successfully established.`,
												},
												"send": schema.StringAttribute{
													Optional:    true,
													Description: `Base64 encoded content of the message which will be sent during the health check to the target`,
												},
											},
											MarkdownDescription: `TcpHealthCheck defines configuration for specifying bytes to send and` + "\n" +
												`expected response during the health check`,
										},
										"timeout": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `Maximum time to wait for a health check response.` + "\n" +
												`If not specified then the default value is 15s`,
										},
										"unhealthy_threshold": schema.Int32Attribute{
											Optional: true,
											MarkdownDescription: `Number of consecutive unhealthy checks before considering a host` + "\n" +
												`unhealthy.` + "\n" +
												`If not specified then the default value is 5`,
										},
									},
									MarkdownDescription: `Default is a configuration specific to the group of destinations referenced in` + "\n" +
										`'targetRef'`,
								},
								"target_ref": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Optional:    true,
											Description: `Kind of the referenced resource. Not Null; must be one of ["Mesh", "MeshSubset", "MeshGateway", "MeshService", "MeshExternalService", "MeshMultiZoneService", "MeshServiceSubset", "MeshHTTPRoute", "Dataplane"]`,
											Validators: []validator.String{
												speakeasy_stringvalidators.NotNull(),
												stringvalidator.OneOf(
													"Mesh",
													"MeshSubset",
													"MeshGateway",
													"MeshService",
													"MeshExternalService",
													"MeshMultiZoneService",
													"MeshServiceSubset",
													"MeshHTTPRoute",
													"Dataplane",
												),
											},
										},
										"labels": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Labels are used to select group of MeshServices that match labels. Either Labels or` + "\n" +
												`Name and Namespace can be used.`,
										},
										"mesh": schema.StringAttribute{
											Optional:    true,
											Description: `Mesh is reserved for future use to identify cross mesh resources.`,
										},
										"name": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `Name of the referenced resource. Can only be used with kinds: ` + "`" + `MeshService` + "`" + `,` + "\n" +
												`` + "`" + `MeshServiceSubset` + "`" + ` and ` + "`" + `MeshGatewayRoute` + "`" + ``,
										},
										"namespace": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `Namespace specifies the namespace of target resource. If empty only resources in policy namespace` + "\n" +
												`will be targeted.`,
										},
										"proxy_types": schema.ListAttribute{
											Optional: true,
											PlanModifiers: []planmodifier.List{
												custom_listplanmodifier.SupressZeroNullModifier(),
											},
											ElementType: types.StringType,
											MarkdownDescription: `ProxyTypes specifies the data plane types that are subject to the policy. When not specified,` + "\n" +
												`all data plane types are targeted by the policy.`,
										},
										"section_name": schema.StringAttribute{
											Optional: true,
											MarkdownDescription: `SectionName is used to target specific section of resource.` + "\n" +
												`For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.`,
										},
										"tags": schema.MapAttribute{
											Optional:    true,
											ElementType: types.StringType,
											MarkdownDescription: `Tags used to select a subset of proxies by tags. Can only be used with kinds` + "\n" +
												`` + "`" + `MeshSubset` + "`" + ` and ` + "`" + `MeshServiceSubset` + "`" + ``,
										},
									},
									MarkdownDescription: `TargetRef is a reference to the resource that represents a group of` + "\n" +
										`destinations.` + "\n" +
										`Not Null`,
									Validators: []validator.Object{
										speakeasy_objectvalidators.NotNull(),
									},
								},
							},
						},
						Description: `To list makes a match between the consumed services and corresponding configurations`,
					},
				},
				Description: `Spec is the specification of the Kuma MeshHealthCheck resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "MeshHealthCheck"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"MeshHealthCheck",
					),
				},
			},
			"warnings": schema.ListAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					custom_listplanmodifier.SupressZeroNullModifier(),
				},
				ElementType: types.StringType,
				MarkdownDescription: `warnings is a list of warning messages to return to the requesting Kuma API clients.` + "\n" +
					`Warning messages describe a problem the client making the API request should correct or be aware of.`,
			},
		},
	}
}

func (r *MeshHealthCheckResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KongMesh)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.KongMesh, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *MeshHealthCheckResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshHealthCheckResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsCreateMeshHealthCheckRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.MeshHealthCheck.CreateMeshHealthCheck(ctx, *request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MeshHealthCheckCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedMeshHealthCheckCreateOrUpdateSuccessResponse(ctx, res.MeshHealthCheckCreateOrUpdateSuccessResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetMeshHealthCheckRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.MeshHealthCheck.GetMeshHealthCheck(ctx, *request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.MeshHealthCheckItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedMeshHealthCheckItem(ctx, res1.MeshHealthCheckItem)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshHealthCheckResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshHealthCheckResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request, requestDiags := data.ToOperationsGetMeshHealthCheckRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.MeshHealthCheck.GetMeshHealthCheck(ctx, *request)
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
	if !(res.MeshHealthCheckItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedMeshHealthCheckItem(ctx, res.MeshHealthCheckItem)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshHealthCheckResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshHealthCheckResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	request, requestDiags := data.ToOperationsUpdateMeshHealthCheckRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.MeshHealthCheck.UpdateMeshHealthCheck(ctx, *request)
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
	if !(res.MeshHealthCheckCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedMeshHealthCheckCreateOrUpdateSuccessResponse(ctx, res.MeshHealthCheckCreateOrUpdateSuccessResponse)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	request1, request1Diags := data.ToOperationsGetMeshHealthCheckRequest(ctx)
	resp.Diagnostics.Append(request1Diags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res1, err := r.client.MeshHealthCheck.GetMeshHealthCheck(ctx, *request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if !(res1.MeshHealthCheckItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	resp.Diagnostics.Append(data.RefreshFromSharedMeshHealthCheckItem(ctx, res1.MeshHealthCheckItem)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(refreshPlan(ctx, plan, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshHealthCheckResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshHealthCheckResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request, requestDiags := data.ToOperationsDeleteMeshHealthCheckRequest(ctx)
	resp.Diagnostics.Append(requestDiags...)

	if resp.Diagnostics.HasError() {
		return
	}
	res, err := r.client.MeshHealthCheck.DeleteMeshHealthCheck(ctx, *request)
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

}

func (r *MeshHealthCheckResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	dec := json.NewDecoder(bytes.NewReader([]byte(req.ID)))
	dec.DisallowUnknownFields()
	var data struct {
		Mesh string `json:"mesh"`
		Name string `json:"name"`
	}

	if err := dec.Decode(&data); err != nil {
		resp.Diagnostics.AddError("Invalid ID", `The import ID is not valid. It is expected to be a JSON object string with the format: '{ "mesh": "",  "name": ""}': `+err.Error())
		return
	}

	if len(data.Mesh) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field mesh is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("mesh"), data.Mesh)...)
	if len(data.Name) == 0 {
		resp.Diagnostics.AddError("Missing required field", `The field name is required but was not found in the json encoded ID. It's expected to be a value alike '""`)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), data.Name)...)

}
