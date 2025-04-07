// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type GetMeshGlobalRateLimitRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshGlobalRateLimit
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetMeshGlobalRateLimitRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *GetMeshGlobalRateLimitRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetMeshGlobalRateLimitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	MeshGlobalRateLimitItem *shared.MeshGlobalRateLimitItem
	// Not Found
	NotFoundError *shared.NotFoundError
}

func (o *GetMeshGlobalRateLimitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshGlobalRateLimitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshGlobalRateLimitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshGlobalRateLimitResponse) GetMeshGlobalRateLimitItem() *shared.MeshGlobalRateLimitItem {
	if o == nil {
		return nil
	}
	return o.MeshGlobalRateLimitItem
}

func (o *GetMeshGlobalRateLimitResponse) GetNotFoundError() *shared.NotFoundError {
	if o == nil {
		return nil
	}
	return o.NotFoundError
}
