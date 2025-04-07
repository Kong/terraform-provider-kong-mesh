// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type DeleteMeshMultiZoneServiceRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshMultiZoneService
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *DeleteMeshMultiZoneServiceRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *DeleteMeshMultiZoneServiceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type DeleteMeshMultiZoneServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshMultiZoneServiceDeleteSuccessResponse *shared.MeshMultiZoneServiceDeleteSuccessResponse
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteMeshMultiZoneServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMeshMultiZoneServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMeshMultiZoneServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMeshMultiZoneServiceResponse) GetMeshMultiZoneServiceDeleteSuccessResponse() *shared.MeshMultiZoneServiceDeleteSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshMultiZoneServiceDeleteSuccessResponse
}

func (o *DeleteMeshMultiZoneServiceResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
