// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

// GetMeshFaultInjectionListQueryParamFilter - filter by labels when multiple filters are present, they are ANDed
type GetMeshFaultInjectionListQueryParamFilter struct {
	Key   *string `queryParam:"name=key"`
	Value *string `queryParam:"name=value"`
}

func (o *GetMeshFaultInjectionListQueryParamFilter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetMeshFaultInjectionListQueryParamFilter) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type GetMeshFaultInjectionListRequest struct {
	// offset in the list of entities
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// the number of items per page
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// filter by labels when multiple filters are present, they are ANDed
	Filter *GetMeshFaultInjectionListQueryParamFilter `queryParam:"style=form,explode=true,name=filter"`
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
}

func (g GetMeshFaultInjectionListRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetMeshFaultInjectionListRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetMeshFaultInjectionListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetMeshFaultInjectionListRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetMeshFaultInjectionListRequest) GetFilter() *GetMeshFaultInjectionListQueryParamFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetMeshFaultInjectionListRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

type GetMeshFaultInjectionListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// List
	MeshFaultInjectionList *shared.MeshFaultInjectionList
}

func (o *GetMeshFaultInjectionListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMeshFaultInjectionListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMeshFaultInjectionListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMeshFaultInjectionListResponse) GetMeshFaultInjectionList() *shared.MeshFaultInjectionList {
	if o == nil {
		return nil
	}
	return o.MeshFaultInjectionList
}
