// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"net/http"
)

// ResourceType - The type of resource (only some resources support rules api)
type ResourceType string

const (
	ResourceTypeDataplanes   ResourceType = "dataplanes"
	ResourceTypeMeshgateways ResourceType = "meshgateways"
)

func (e ResourceType) ToPointer() *ResourceType {
	return &e
}
func (e *ResourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dataplanes":
		fallthrough
	case "meshgateways":
		*e = ResourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResourceType: %v", v)
	}
}

type InspectDataplanesRulesRequest struct {
	// The mesh the policy is part of
	Mesh string `pathParam:"style=simple,explode=false,name=mesh"`
	// The type of resource (only some resources support rules api)
	ResourceType ResourceType `pathParam:"style=simple,explode=false,name=resourceType"`
	// The name of the resource
	ResourceName string `pathParam:"style=simple,explode=false,name=resourceName"`
}

func (o *InspectDataplanesRulesRequest) GetMesh() string {
	if o == nil {
		return ""
	}
	return o.Mesh
}

func (o *InspectDataplanesRulesRequest) GetResourceType() ResourceType {
	if o == nil {
		return ResourceType("")
	}
	return o.ResourceType
}

func (o *InspectDataplanesRulesRequest) GetResourceName() string {
	if o == nil {
		return ""
	}
	return o.ResourceName
}

type InspectDataplanesRulesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A response containing policies that match a resource
	InspectRules *shared.InspectRules
	// Bad Request
	BadRequestError *shared.BadRequestError
	// Internal
	InternalError *shared.InternalError
}

func (o *InspectDataplanesRulesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *InspectDataplanesRulesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *InspectDataplanesRulesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *InspectDataplanesRulesResponse) GetInspectRules() *shared.InspectRules {
	if o == nil {
		return nil
	}
	return o.InspectRules
}

func (o *InspectDataplanesRulesResponse) GetBadRequestError() *shared.BadRequestError {
	if o == nil {
		return nil
	}
	return o.BadRequestError
}

func (o *InspectDataplanesRulesResponse) GetInternalError() *shared.InternalError {
	if o == nil {
		return nil
	}
	return o.InternalError
}
