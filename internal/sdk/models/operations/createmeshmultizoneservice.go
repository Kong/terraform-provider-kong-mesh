// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type CreateMeshMultiZoneServiceRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshMultiZoneService
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshMultiZoneServiceItem shared.MeshMultiZoneServiceItemInput `request:"mediaType=application/json"`
}

func (o *CreateMeshMultiZoneServiceRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *CreateMeshMultiZoneServiceRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeshMultiZoneServiceRequest) GetMeshMultiZoneServiceItem() shared.MeshMultiZoneServiceItemInput {
	if o == nil {
		return shared.MeshMultiZoneServiceItemInput{}
	}
	return o.MeshMultiZoneServiceItem
}

type CreateMeshMultiZoneServiceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	MeshMultiZoneServiceCreateOrUpdateSuccessResponse *shared.MeshMultiZoneServiceCreateOrUpdateSuccessResponse
}

func (o *CreateMeshMultiZoneServiceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMeshMultiZoneServiceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMeshMultiZoneServiceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMeshMultiZoneServiceResponse) GetMeshMultiZoneServiceCreateOrUpdateSuccessResponse() *shared.MeshMultiZoneServiceCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshMultiZoneServiceCreateOrUpdateSuccessResponse
}
