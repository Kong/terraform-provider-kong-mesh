// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type CreateMeshFaultInjectionRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshFaultInjection
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshFaultInjectionItem shared.MeshFaultInjectionItemInput `request:"mediaType=application/json"`
}

func (o *CreateMeshFaultInjectionRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *CreateMeshFaultInjectionRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeshFaultInjectionRequest) GetMeshFaultInjectionItem() shared.MeshFaultInjectionItemInput {
	if o == nil {
		return shared.MeshFaultInjectionItemInput{}
	}
	return o.MeshFaultInjectionItem
}

type CreateMeshFaultInjectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	MeshFaultInjectionCreateOrUpdateSuccessResponse *shared.MeshFaultInjectionCreateOrUpdateSuccessResponse
}

func (o *CreateMeshFaultInjectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMeshFaultInjectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMeshFaultInjectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMeshFaultInjectionResponse) GetMeshFaultInjectionCreateOrUpdateSuccessResponse() *shared.MeshFaultInjectionCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshFaultInjectionCreateOrUpdateSuccessResponse
}
