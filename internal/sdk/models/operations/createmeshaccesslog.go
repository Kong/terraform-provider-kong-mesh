// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type CreateMeshAccessLogRequest struct {
	// name of the mesh
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// name of the MeshAccessLog
	Name string `pathParam:"style=simple,explode=false,name=name"`
	// Put request
	MeshAccessLogItem shared.MeshAccessLogItemInput `request:"mediaType=application/json"`
}

func (o *CreateMeshAccessLogRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *CreateMeshAccessLogRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateMeshAccessLogRequest) GetMeshAccessLogItem() shared.MeshAccessLogItemInput {
	if o == nil {
		return shared.MeshAccessLogItemInput{}
	}
	return o.MeshAccessLogItem
}

type CreateMeshAccessLogResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created
	MeshAccessLogCreateOrUpdateSuccessResponse *shared.MeshAccessLogCreateOrUpdateSuccessResponse
}

func (o *CreateMeshAccessLogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateMeshAccessLogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateMeshAccessLogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateMeshAccessLogResponse) GetMeshAccessLogCreateOrUpdateSuccessResponse() *shared.MeshAccessLogCreateOrUpdateSuccessResponse {
	if o == nil {
		return nil
	}
	return o.MeshAccessLogCreateOrUpdateSuccessResponse
}
