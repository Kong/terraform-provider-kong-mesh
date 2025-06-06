// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

type InspectResourcesRequest struct {
	// The mesh the policy is part of
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// The type of the policy
	PolicyType string `pathParam:"style=simple,explode=false,name=policyType"`
	// The type of the policy
	PolicyName string `pathParam:"style=simple,explode=false,name=policyName"`
	// The max number of items to return
	Size *int64 `queryParam:"style=form,explode=true,name=size"`
	// The offset of result
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
	// A sub string to filter resources by name
	Name *string `queryParam:"style=form,explode=true,name=name"`
}

func (o *InspectResourcesRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *InspectResourcesRequest) GetPolicyType() string {
	if o == nil {
		return ""
	}
	return o.PolicyType
}

func (o *InspectResourcesRequest) GetPolicyName() string {
	if o == nil {
		return ""
	}
	return o.PolicyName
}

func (o *InspectResourcesRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *InspectResourcesRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *InspectResourcesRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type InspectResourcesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response containing dataplanes that match a policy.
	InspectDataplanesForPolicy *shared.InspectDataplanesForPolicy
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Internal
	BaseError *shared.BaseError
}

func (o *InspectResourcesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InspectResourcesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InspectResourcesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *InspectResourcesResponse) GetInspectDataplanesForPolicy() *shared.InspectDataplanesForPolicy {
	if o == nil {
		return nil
	}
	return o.InspectDataplanesForPolicy
}

func (o *InspectResourcesResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *InspectResourcesResponse) GetBaseError() *shared.BaseError {
	if o == nil {
		return nil
	}
	return o.BaseError
}
