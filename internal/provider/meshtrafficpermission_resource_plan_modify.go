package provider

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	sdkerrors "github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/errors"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"net/http"
	"strconv"

	// ... other relevant imports ...
)

// Ensure the resource implements the ModifyPlan interface.
var _ resource.ResourceWithModifyPlan = &MeshTrafficPermissionResource{}

func (r *MeshTrafficPermissionResource) ModifyPlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	// Skip if resource already exists in state
	if !req.State.Raw.IsNull() {
		return
	}

	var plannedResource MeshTrafficPermissionResourceModel
	diags := req.Plan.Get(ctx, &plannedResource)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Skip if required fields for reading the resource are not yet known
	if plannedResource.Name.IsUnknown() || plannedResource.Mesh.IsUnknown() {
		return
	}

	mtp, err := r.client.MeshTrafficPermission.GetMeshTrafficPermission(ctx, operations.GetMeshTrafficPermissionRequest{
		Name: plannedResource.Name.ValueString(),
		Mesh: plannedResource.Mesh.ValueString(),
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

	if mtp.StatusCode != http.StatusNotFound {
		showAlreadyExistsError(resp, plannedResource)
	}
}

func showAlreadyExistsError(resp *resource.ModifyPlanResponse, plannedResource MeshTrafficPermissionResourceModel) {
	resp.Diagnostics.AddError(
		"MeshTrafficPermission already exists",
		"A resource with the name "+plannedResource.Name.String()+" already exists in the mesh "+plannedResource.Mesh.String()+" - to be managed via Terraform it needs to be imported first",
	)
}
