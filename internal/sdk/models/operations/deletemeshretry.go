// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type DeleteMeshRetryRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshRetry
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *DeleteMeshRetryRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *DeleteMeshRetryRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type DeleteMeshRetryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshRetryDeleteSuccessResponse *shared.MeshRetryDeleteSuccessResponse
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteMeshRetryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMeshRetryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMeshRetryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMeshRetryResponse) GetMeshRetryDeleteSuccessResponse() *shared.MeshRetryDeleteSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshRetryDeleteSuccessResponse
}

func (o *DeleteMeshRetryResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
