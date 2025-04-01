// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"time"
)

// HostnameGeneratorItemType - the type of the resource
type HostnameGeneratorItemType string

const (
	HostnameGeneratorItemTypeHostnameGenerator HostnameGeneratorItemType = "HostnameGenerator"
)

func (e HostnameGeneratorItemType) ToPointer() *HostnameGeneratorItemType {
	return &e
}
func (e *HostnameGeneratorItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HostnameGenerator":
		*e = HostnameGeneratorItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HostnameGeneratorItemType: %v", v)
	}
}

type MeshExternalService struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

func (o *MeshExternalService) GetMatchLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.MatchLabels
}

type MeshMultiZoneService struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

func (o *MeshMultiZoneService) GetMatchLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.MatchLabels
}

type MeshService struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

func (o *MeshService) GetMatchLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.MatchLabels
}

type Selector struct {
	MeshExternalService  *MeshExternalService  `json:"meshExternalService,omitempty"`
	MeshMultiZoneService *MeshMultiZoneService `json:"meshMultiZoneService,omitempty"`
	MeshService          *MeshService          `json:"meshService,omitempty"`
}

func (o *Selector) GetMeshExternalService() *MeshExternalService {
	if o == nil {
		return nil
	}
	return o.MeshExternalService
}

func (o *Selector) GetMeshMultiZoneService() *MeshMultiZoneService {
	if o == nil {
		return nil
	}
	return o.MeshMultiZoneService
}

func (o *Selector) GetMeshService() *MeshService {
	if o == nil {
		return nil
	}
	return o.MeshService
}

// HostnameGeneratorItemSpec - Spec is the specification of the Kuma HostnameGenerator resource.
type HostnameGeneratorItemSpec struct {
	Selector *Selector `json:"selector,omitempty"`
	Template *string   `json:"template,omitempty"`
}

func (o *HostnameGeneratorItemSpec) GetSelector() *Selector {
	if o == nil {
		return nil
	}
	return o.Selector
}

func (o *HostnameGeneratorItemSpec) GetTemplate() *string {
	if o == nil {
		return nil
	}
	return o.Template
}

// HostnameGeneratorItem - Successful response
type HostnameGeneratorItem struct {
	// the type of the resource
	Type HostnameGeneratorItemType `json:"type"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma HostnameGenerator resource.
	Spec HostnameGeneratorItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (h HostnameGeneratorItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HostnameGeneratorItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HostnameGeneratorItem) GetType() HostnameGeneratorItemType {
	if o == nil {
		return HostnameGeneratorItemType("")
	}
	return o.Type
}

func (o *HostnameGeneratorItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *HostnameGeneratorItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *HostnameGeneratorItem) GetSpec() HostnameGeneratorItemSpec {
	if o == nil {
		return HostnameGeneratorItemSpec{}
	}
	return o.Spec
}

func (o *HostnameGeneratorItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *HostnameGeneratorItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

type HostnameGeneratorItemInput struct {
	// the type of the resource
	Type HostnameGeneratorItemType `json:"type"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma HostnameGenerator resource.
	Spec HostnameGeneratorItemSpec `json:"spec"`
}

func (o *HostnameGeneratorItemInput) GetType() HostnameGeneratorItemType {
	if o == nil {
		return HostnameGeneratorItemType("")
	}
	return o.Type
}

func (o *HostnameGeneratorItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *HostnameGeneratorItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *HostnameGeneratorItemInput) GetSpec() HostnameGeneratorItemSpec {
	if o == nil {
		return HostnameGeneratorItemSpec{}
	}
	return o.Spec
}
