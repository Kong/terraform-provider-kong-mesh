// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type DeleteMeshTraceRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshTrace
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *DeleteMeshTraceRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *DeleteMeshTraceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type DeleteMeshTraceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshTraceDeleteSuccessResponse *shared.MeshTraceDeleteSuccessResponse
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteMeshTraceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMeshTraceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMeshTraceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMeshTraceResponse) GetMeshTraceDeleteSuccessResponse() *shared.MeshTraceDeleteSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshTraceDeleteSuccessResponse
}

func (o *DeleteMeshTraceResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
