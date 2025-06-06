// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type PutDataplaneRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the Dataplane
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	DataplaneItem shared.DataplaneItem `request:"mediaType=application/json"`
}

func (o *PutDataplaneRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *PutDataplaneRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *PutDataplaneRequest) GetDataplaneItem() shared.DataplaneItem {
	if o == nil {
		return shared.DataplaneItem{}
	}
	return o.DataplaneItem
}

type PutDataplaneResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Updated
	DataplaneCreateOrUpdateSuccessResponse *shared.DataplaneCreateOrUpdateSuccessResponse
}

func (o *PutDataplaneResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDataplaneResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDataplaneResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PutDataplaneResponse) GetDataplaneCreateOrUpdateSuccessResponse() *shared.DataplaneCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.DataplaneCreateOrUpdateSuccessResponse
}
