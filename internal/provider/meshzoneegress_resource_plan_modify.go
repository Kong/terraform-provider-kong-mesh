package provider

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	sdkerrors "github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/errors"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
)

var _ resource.ResourceWithModifyPlan = &MeshZoneEgressResource{}

func (r *MeshZoneEgressResource) ModifyPlan(
	ctx context.Context,
	req resource.ModifyPlanRequest,
	resp *resource.ModifyPlanResponse,
) {
	if !req.State.Raw.IsNull() {
		return
	}

	var name types.String
	if diags := req.Plan.GetAttribute(ctx, path.Root("name"), &name); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if name.IsUnknown() {
		return
	}

	request := operations.GetZoneEgressRequest{
		Name: name.ValueString(),
	}
	res, err := r.client.ZoneEgress.GetZoneEgress(ctx, request)

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
			"MeshZoneEgress already exists",
			"A resource with the name "+name.String()+" already exists - to be managed via Terraform it needs to be imported first",
		)
	}
}
