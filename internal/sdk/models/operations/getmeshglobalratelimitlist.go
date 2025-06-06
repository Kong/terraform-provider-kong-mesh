// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

// GetMeshGlobalRateLimitListQueryParamFilter - filter by labels when multiple filters are present, they are ANDed
type GetMeshGlobalRateLimitListQueryParamFilter struct {
	Key   *string `queryParam:"name=key"`
	Value *string `queryParam:"name=value"`
}

func (o *GetMeshGlobalRateLimitListQueryParamFilter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMeshGlobalRateLimitListQueryParamFilter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetMeshGlobalRateLimitListRequest struct {
	// offset in the list of entities
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// the number of items per page
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// filter by labels when multiple filters are present, they are ANDed
	Filter *GetMeshGlobalRateLimitListQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
}

func (g GetMeshGlobalRateLimitListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMeshGlobalRateLimitListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMeshGlobalRateLimitListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetMeshGlobalRateLimitListRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetMeshGlobalRateLimitListRequest) GetFilter() *GetMeshGlobalRateLimitListQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetMeshGlobalRateLimitListRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

type GetMeshGlobalRateLimitListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List
	MeshGlobalRateLimitList *shared.MeshGlobalRateLimitList
}

func (o *GetMeshGlobalRateLimitListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshGlobalRateLimitListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshGlobalRateLimitListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshGlobalRateLimitListResponse) GetMeshGlobalRateLimitList() *shared.MeshGlobalRateLimitList {
	if o == nil {
		return nil
	}
	return o.MeshGlobalRateLimitList
}
