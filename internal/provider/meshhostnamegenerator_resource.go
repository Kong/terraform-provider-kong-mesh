// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	custom_listplanmodifier "github.com/kong/terraform-provider-kong-mesh/internal/planmodifiers/listplanmodifier"
	speakeasy_stringplanmodifier "github.com/kong/terraform-provider-kong-mesh/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/kong/terraform-provider-kong-mesh/internal/provider/types"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &MeshHostnameGeneratorResource{}
var _ resource.ResourceWithImportState = &MeshHostnameGeneratorResource{}

func NewMeshHostnameGeneratorResource() resource.Resource {
	return &MeshHostnameGeneratorResource{}
}

// MeshHostnameGeneratorResource defines the resource implementation.
type MeshHostnameGeneratorResource struct {
	client *sdk.KongMesh
}

// MeshHostnameGeneratorResourceModel describes the resource data model.
type MeshHostnameGeneratorResourceModel struct {
	CreationTime     types.String                      `tfsdk:"creation_time"`
	Labels           map[string]types.String           `tfsdk:"labels"`
	ModificationTime types.String                      `tfsdk:"modification_time"`
	Name             types.String                      `tfsdk:"name"`
	Spec             tfTypes.HostnameGeneratorItemSpec `tfsdk:"spec"`
	Type             types.String                      `tfsdk:"type"`
	Warnings         []types.String                    `tfsdk:"warnings"`
}

func (r *MeshHostnameGeneratorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mesh_hostname_generator"
}

func (r *MeshHostnameGeneratorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "MeshHostnameGenerator Resource",
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
				Optional:    true,
				ElementType: types.StringType,
				Description: `The labels to help identity resources`,
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
				Required:    true,
				Description: `name of the HostnameGenerator`,
			},
			"spec": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"selector": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"mesh_external_service": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
							"mesh_multi_zone_service": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
							"mesh_service": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"match_labels": schema.MapAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
					},
					"template": schema.StringAttribute{
						Optional: true,
					},
				},
				Description: `Spec is the specification of the Kuma HostnameGenerator resource.`,
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `the type of the resource. must be "HostnameGenerator"`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"HostnameGenerator",
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

func (r *MeshHostnameGeneratorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *MeshHostnameGeneratorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *MeshHostnameGeneratorResourceModel
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

	var name string
	name = data.Name.ValueString()

	hostnameGeneratorItem := *data.ToSharedHostnameGeneratorItemInput()
	request := operations.CreateHostnameGeneratorRequest{
		Name:                  name,
		HostnameGeneratorItem: hostnameGeneratorItem,
	}
	res, err := r.client.HostnameGenerator.CreateHostnameGenerator(ctx, request)
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
	if !(res.HostnameGeneratorCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedHostnameGeneratorCreateOrUpdateSuccessResponse(res.HostnameGeneratorCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetHostnameGeneratorRequest{
		Name: name1,
	}
	res1, err := r.client.HostnameGenerator.GetHostnameGenerator(ctx, request1)
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
	if !(res1.HostnameGeneratorItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedHostnameGeneratorItem(res1.HostnameGeneratorItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshHostnameGeneratorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *MeshHostnameGeneratorResourceModel
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

	var name string
	name = data.Name.ValueString()

	request := operations.GetHostnameGeneratorRequest{
		Name: name,
	}
	res, err := r.client.HostnameGenerator.GetHostnameGenerator(ctx, request)
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
	if !(res.HostnameGeneratorItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedHostnameGeneratorItem(res.HostnameGeneratorItem)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshHostnameGeneratorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *MeshHostnameGeneratorResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var name string
	name = data.Name.ValueString()

	hostnameGeneratorItem := *data.ToSharedHostnameGeneratorItemInput()
	request := operations.UpdateHostnameGeneratorRequest{
		Name:                  name,
		HostnameGeneratorItem: hostnameGeneratorItem,
	}
	res, err := r.client.HostnameGenerator.UpdateHostnameGenerator(ctx, request)
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
	if !(res.HostnameGeneratorCreateOrUpdateSuccessResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedHostnameGeneratorCreateOrUpdateSuccessResponse(res.HostnameGeneratorCreateOrUpdateSuccessResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	var name1 string
	name1 = data.Name.ValueString()

	request1 := operations.GetHostnameGeneratorRequest{
		Name: name1,
	}
	res1, err := r.client.HostnameGenerator.GetHostnameGenerator(ctx, request1)
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
	if !(res1.HostnameGeneratorItem != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedHostnameGeneratorItem(res1.HostnameGeneratorItem)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MeshHostnameGeneratorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *MeshHostnameGeneratorResourceModel
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

	var name string
	name = data.Name.ValueString()

	request := operations.DeleteHostnameGeneratorRequest{
		Name: name,
	}
	res, err := r.client.HostnameGenerator.DeleteHostnameGenerator(ctx, request)
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

func (r *MeshHostnameGeneratorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("name"), req.ID)...)
}
