// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type CreateMeshServiceRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshService
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshServiceItem shared.MeshServiceItemInput `request:"mediaType=application/json"`
}

func (o *CreateMeshServiceRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *CreateMeshServiceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeshServiceRequest) GetMeshServiceItem() shared.MeshServiceItemInput {
	if o == nil {
		return shared.MeshServiceItemInput{}
	}
	return o.MeshServiceItem
}

type CreateMeshServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	MeshServiceCreateOrUpdateSuccessResponse *shared.MeshServiceCreateOrUpdateSuccessResponse
}

func (o *CreateMeshServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMeshServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMeshServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMeshServiceResponse) GetMeshServiceCreateOrUpdateSuccessResponse() *shared.MeshServiceCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshServiceCreateOrUpdateSuccessResponse
}
