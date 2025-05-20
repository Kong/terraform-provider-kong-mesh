package provider

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	sdkerrors "github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/errors"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"net/http"
	"strconv"
)

var _ resource.ResourceWithModifyPlan = &MeshHostnameGeneratorResource{}

func (r *MeshHostnameGeneratorResource) ModifyPlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	if !req.State.Raw.IsNull() {
		return
	}

	var plannedResource MeshHostnameGeneratorResourceModel
	diags := req.Plan.Get(ctx, &plannedResource)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	if plannedResource.Name.IsUnknown() {
		return
	}
	res, err := r.client.Mesh.GetMesh(ctx, operations.GetMeshRequest{
		Name: plannedResource.Name.ValueString(),
	})

	if err != nil {
		var sdkError *sdkerrors.SDKError
		if errors.As(err, &sdkError) {
			if sdkError.StatusCode == http.StatusNotFound {
				return
			} else {
				resp.Diagnostics.AddError(
					"Unexpected error status code",
					"The status code for non existing resource is not 404, got "+strconv.Itoa(sdkError.StatusCode)+" error: "+sdkError.Error(),
				)
				return
			}
		} else {
			resp.Diagnostics.AddError(
				"Couldn't map error to SDKError",
				"Only SDKError is supported for this operation, but got: "+err.Error(),
			)
			return
		}
	}

	if res.StatusCode != http.StatusNotFound {
		resp.Diagnostics.AddError(
			"MeshHostnameGenerator already exists",
			"A resource with the name "+plannedResource.Name.String()+" already exists - to be managed via Terraform it needs to be imported first",
		)
	}
}
