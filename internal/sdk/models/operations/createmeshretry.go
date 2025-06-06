// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type CreateMeshRetryRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshRetry
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshRetryItem shared.MeshRetryItemInput `request:"mediaType=application/json"`
}

func (o *CreateMeshRetryRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *CreateMeshRetryRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeshRetryRequest) GetMeshRetryItem() shared.MeshRetryItemInput {
	if o == nil {
		return shared.MeshRetryItemInput{}
	}
	return o.MeshRetryItem
}

type CreateMeshRetryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	MeshRetryCreateOrUpdateSuccessResponse *shared.MeshRetryCreateOrUpdateSuccessResponse
}

func (o *CreateMeshRetryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMeshRetryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMeshRetryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMeshRetryResponse) GetMeshRetryCreateOrUpdateSuccessResponse() *shared.MeshRetryCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshRetryCreateOrUpdateSuccessResponse
}
