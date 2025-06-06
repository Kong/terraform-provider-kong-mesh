// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type GetMeshCircuitBreakerRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshCircuitBreaker
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshCircuitBreakerRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshCircuitBreakerRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshCircuitBreakerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshCircuitBreakerItem *shared.MeshCircuitBreakerItem
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetMeshCircuitBreakerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshCircuitBreakerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshCircuitBreakerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshCircuitBreakerResponse) GetMeshCircuitBreakerItem() *shared.MeshCircuitBreakerItem {
	if o == nil {
		return nil
	}
	return o.MeshCircuitBreakerItem
}

func (o *GetMeshCircuitBreakerResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
