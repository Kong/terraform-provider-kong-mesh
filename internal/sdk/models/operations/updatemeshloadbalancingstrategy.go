// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type UpdateMeshLoadBalancingStrategyRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshLoadBalancingStrategy
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshLoadBalancingStrategyItem shared.MeshLoadBalancingStrategyItemInput `request:"mediaType=application/json"`
}

func (o *UpdateMeshLoadBalancingStrategyRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *UpdateMeshLoadBalancingStrategyRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateMeshLoadBalancingStrategyRequest) GetMeshLoadBalancingStrategyItem() shared.MeshLoadBalancingStrategyItemInput {
	if o == nil {
		return shared.MeshLoadBalancingStrategyItemInput{}
	}
	return o.MeshLoadBalancingStrategyItem
}

type UpdateMeshLoadBalancingStrategyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	MeshLoadBalancingStrategyCreateOrUpdateSuccessResponse *shared.MeshLoadBalancingStrategyCreateOrUpdateSuccessResponse
}

func (o *UpdateMeshLoadBalancingStrategyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMeshLoadBalancingStrategyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMeshLoadBalancingStrategyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateMeshLoadBalancingStrategyResponse) GetMeshLoadBalancingStrategyCreateOrUpdateSuccessResponse() *shared.MeshLoadBalancingStrategyCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshLoadBalancingStrategyCreateOrUpdateSuccessResponse
}
