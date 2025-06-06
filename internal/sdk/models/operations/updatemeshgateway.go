// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type UpdateMeshGatewayRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshGateway
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshGatewayItem shared.MeshGatewayItem `request:"mediaType=application/json"`
}

func (o *UpdateMeshGatewayRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *UpdateMeshGatewayRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateMeshGatewayRequest) GetMeshGatewayItem() shared.MeshGatewayItem {
	if o == nil {
		return shared.MeshGatewayItem{}
	}
	return o.MeshGatewayItem
}

type UpdateMeshGatewayResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	MeshGatewayCreateOrUpdateSuccessResponse *shared.MeshGatewayCreateOrUpdateSuccessResponse
}

func (o *UpdateMeshGatewayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMeshGatewayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMeshGatewayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateMeshGatewayResponse) GetMeshGatewayCreateOrUpdateSuccessResponse() *shared.MeshGatewayCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshGatewayCreateOrUpdateSuccessResponse
}
