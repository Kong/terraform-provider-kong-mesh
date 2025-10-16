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

var _ resource.ResourceWithModifyPlan = &MeshTrustResource{}

func (r *MeshTrustResource) ModifyPlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	if !req.State.Raw.IsNull() {
		return
	}

	var plannedResource MeshTrustResourceModel
	diags := req.Plan.Get(ctx, &plannedResource)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plannedResource.Name.IsUnknown() {
		return
	}
	if plannedResource.Mesh.IsUnknown() {
		return
	}
	request := operations.GetMeshTrustRequest{
		Name: plannedResource.Name.ValueString(),
	}
	request.Mesh = plannedResource.Mesh.ValueString()
	res, err := r.client.MeshTrust.GetMeshTrust(ctx, request)

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
			"MeshTrust already exists",
			"A resource with the name "+plannedResource.Name.String()+" already exists in the mesh "+plannedResource.Mesh.String()+" - to be managed via Terraform it needs to be imported first",
		)
	}
}
