// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type DeleteMeshCircuitBreakerRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshCircuitBreaker
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *DeleteMeshCircuitBreakerRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *DeleteMeshCircuitBreakerRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type DeleteMeshCircuitBreakerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshCircuitBreakerDeleteSuccessResponse *shared.MeshCircuitBreakerDeleteSuccessResponse
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *DeleteMeshCircuitBreakerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteMeshCircuitBreakerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteMeshCircuitBreakerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteMeshCircuitBreakerResponse) GetMeshCircuitBreakerDeleteSuccessResponse() *shared.MeshCircuitBreakerDeleteSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshCircuitBreakerDeleteSuccessResponse
}

func (o *DeleteMeshCircuitBreakerResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
