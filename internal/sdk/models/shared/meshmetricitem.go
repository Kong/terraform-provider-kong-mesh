// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"time"
)

// MeshMetricItemType - the type of the resource
type MeshMetricItemType string

const (
	MeshMetricItemTypeMeshMetric MeshMetricItemType = "MeshMetric"
)

func (e MeshMetricItemType) ToPointer() *MeshMetricItemType {
	return &e
}
func (e *MeshMetricItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshMetric":
		*e = MeshMetricItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshMetricItemType: %v", v)
	}
}

type Applications struct {
	// Address on which an application listens.
	Address *string `json:"address,omitempty"`
	// Name of the application to scrape
	Name *string `json:"name,omitempty"`
	// Path on which an application expose HTTP endpoint with metrics.
	Path *string `default:"/metrics" json:"path"`
	// Port on which an application expose HTTP endpoint with metrics.
	Port int `json:"port"`
}

func (a Applications) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Applications) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Applications) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Applications) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Applications) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *Applications) GetPort() int {
	if o == nil {
		return 0
	}
	return o.Port
}

// OpenTelemetry backend configuration
type OpenTelemetry struct {
	// Endpoint for OpenTelemetry collector
	Endpoint string `json:"endpoint"`
	// RefreshInterval defines how frequent metrics should be pushed to collector
	RefreshInterval *string `json:"refreshInterval,omitempty"`
}

func (o *OpenTelemetry) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *OpenTelemetry) GetRefreshInterval() *string {
	if o == nil {
		return nil
	}
	return o.RefreshInterval
}

// MeshMetricItemMode - Configuration of TLS for Prometheus listener.
type MeshMetricItemMode string

const (
	MeshMetricItemModeDisabled          MeshMetricItemMode = "Disabled"
	MeshMetricItemModeProvidedTLS       MeshMetricItemMode = "ProvidedTLS"
	MeshMetricItemModeActiveMtlsBackend MeshMetricItemMode = "ActiveMTLSBackend"
)

func (e MeshMetricItemMode) ToPointer() *MeshMetricItemMode {
	return &e
}
func (e *MeshMetricItemMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Disabled":
		fallthrough
	case "ProvidedTLS":
		fallthrough
	case "ActiveMTLSBackend":
		*e = MeshMetricItemMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshMetricItemMode: %v", v)
	}
}

// MeshMetricItemTLS - Configuration of TLS for prometheus listener.
type MeshMetricItemTLS struct {
	// Configuration of TLS for Prometheus listener.
	Mode *MeshMetricItemMode `default:"Disabled" json:"mode"`
}

func (m MeshMetricItemTLS) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshMetricItemTLS) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshMetricItemTLS) GetMode() *MeshMetricItemMode {
	if o == nil {
		return nil
	}
	return o.Mode
}

// Prometheus backend configuration.
type Prometheus struct {
	// ClientId of the Prometheus backend. Needed when using MADS for DP discovery.
	ClientID *string `json:"clientId,omitempty"`
	// Path on which a dataplane should expose HTTP endpoint with Prometheus metrics.
	Path *string `default:"/metrics" json:"path"`
	// Port on which a dataplane should expose HTTP endpoint with Prometheus metrics.
	Port *int `default:"5670" json:"port"`
	// Configuration of TLS for prometheus listener.
	TLS *MeshMetricItemTLS `json:"tls,omitempty"`
}

func (p Prometheus) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Prometheus) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Prometheus) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *Prometheus) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *Prometheus) GetPort() *int {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *Prometheus) GetTLS() *MeshMetricItemTLS {
	if o == nil {
		return nil
	}
	return o.TLS
}

// MeshMetricItemSpecType - Type of the backend that will be used to collect metrics. At the moment only Prometheus backend is available.
type MeshMetricItemSpecType string

const (
	MeshMetricItemSpecTypePrometheus    MeshMetricItemSpecType = "Prometheus"
	MeshMetricItemSpecTypeOpenTelemetry MeshMetricItemSpecType = "OpenTelemetry"
)

func (e MeshMetricItemSpecType) ToPointer() *MeshMetricItemSpecType {
	return &e
}
func (e *MeshMetricItemSpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Prometheus":
		fallthrough
	case "OpenTelemetry":
		*e = MeshMetricItemSpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshMetricItemSpecType: %v", v)
	}
}

type MeshMetricItemBackends struct {
	// OpenTelemetry backend configuration
	OpenTelemetry *OpenTelemetry `json:"openTelemetry,omitempty"`
	// Prometheus backend configuration.
	Prometheus *Prometheus `json:"prometheus,omitempty"`
	// Type of the backend that will be used to collect metrics. At the moment only Prometheus backend is available.
	Type MeshMetricItemSpecType `json:"type"`
}

func (o *MeshMetricItemBackends) GetOpenTelemetry() *OpenTelemetry {
	if o == nil {
		return nil
	}
	return o.OpenTelemetry
}

func (o *MeshMetricItemBackends) GetPrometheus() *Prometheus {
	if o == nil {
		return nil
	}
	return o.Prometheus
}

func (o *MeshMetricItemBackends) GetType() MeshMetricItemSpecType {
	if o == nil {
		return MeshMetricItemSpecType("")
	}
	return o.Type
}

// Name of the predefined profile, one of: all, basic, none
type Name string

const (
	NameAll   Name = "All"
	NameBasic Name = "Basic"
	NameNone  Name = "None"
)

func (e Name) ToPointer() *Name {
	return &e
}
func (e *Name) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "All":
		fallthrough
	case "Basic":
		fallthrough
	case "None":
		*e = Name(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Name: %v", v)
	}
}

type AppendProfiles struct {
	// Name of the predefined profile, one of: all, basic, none
	Name Name `json:"name"`
}

func (o *AppendProfiles) GetName() Name {
	if o == nil {
		return Name("")
	}
	return o.Name
}

// MeshMetricItemSpecDefaultType - Type defined the type of selector, one of: prefix, regex, exact
type MeshMetricItemSpecDefaultType string

const (
	MeshMetricItemSpecDefaultTypePrefix   MeshMetricItemSpecDefaultType = "Prefix"
	MeshMetricItemSpecDefaultTypeRegex    MeshMetricItemSpecDefaultType = "Regex"
	MeshMetricItemSpecDefaultTypeExact    MeshMetricItemSpecDefaultType = "Exact"
	MeshMetricItemSpecDefaultTypeContains MeshMetricItemSpecDefaultType = "Contains"
)

func (e MeshMetricItemSpecDefaultType) ToPointer() *MeshMetricItemSpecDefaultType {
	return &e
}
func (e *MeshMetricItemSpecDefaultType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Prefix":
		fallthrough
	case "Regex":
		fallthrough
	case "Exact":
		fallthrough
	case "Contains":
		*e = MeshMetricItemSpecDefaultType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshMetricItemSpecDefaultType: %v", v)
	}
}

type Exclude struct {
	// Match is the value used to match using particular Type
	Match string `json:"match"`
	// Type defined the type of selector, one of: prefix, regex, exact
	Type MeshMetricItemSpecDefaultType `json:"type"`
}

func (o *Exclude) GetMatch() string {
	if o == nil {
		return ""
	}
	return o.Match
}

func (o *Exclude) GetType() MeshMetricItemSpecDefaultType {
	if o == nil {
		return MeshMetricItemSpecDefaultType("")
	}
	return o.Type
}

// MeshMetricItemSpecDefaultSidecarType - Type defined the type of selector, one of: prefix, regex, exact
type MeshMetricItemSpecDefaultSidecarType string

const (
	MeshMetricItemSpecDefaultSidecarTypePrefix   MeshMetricItemSpecDefaultSidecarType = "Prefix"
	MeshMetricItemSpecDefaultSidecarTypeRegex    MeshMetricItemSpecDefaultSidecarType = "Regex"
	MeshMetricItemSpecDefaultSidecarTypeExact    MeshMetricItemSpecDefaultSidecarType = "Exact"
	MeshMetricItemSpecDefaultSidecarTypeContains MeshMetricItemSpecDefaultSidecarType = "Contains"
)

func (e MeshMetricItemSpecDefaultSidecarType) ToPointer() *MeshMetricItemSpecDefaultSidecarType {
	return &e
}
func (e *MeshMetricItemSpecDefaultSidecarType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Prefix":
		fallthrough
	case "Regex":
		fallthrough
	case "Exact":
		fallthrough
	case "Contains":
		*e = MeshMetricItemSpecDefaultSidecarType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshMetricItemSpecDefaultSidecarType: %v", v)
	}
}

type Include struct {
	// Match is the value used to match using particular Type
	Match string `json:"match"`
	// Type defined the type of selector, one of: prefix, regex, exact
	Type MeshMetricItemSpecDefaultSidecarType `json:"type"`
}

func (o *Include) GetMatch() string {
	if o == nil {
		return ""
	}
	return o.Match
}

func (o *Include) GetType() MeshMetricItemSpecDefaultSidecarType {
	if o == nil {
		return MeshMetricItemSpecDefaultSidecarType("")
	}
	return o.Type
}

// Profiles allows to customize which metrics are published.
type Profiles struct {
	// AppendProfiles allows to combine the metrics from multiple predefined profiles.
	AppendProfiles []AppendProfiles `json:"appendProfiles,omitempty"`
	// Exclude makes it possible to exclude groups of metrics from a resulting profile.
	// Exclude is subordinate to Include.
	Exclude []Exclude `json:"exclude,omitempty"`
	// Include makes it possible to include additional metrics in a selected profiles.
	// Include takes precedence over Exclude.
	Include []Include `json:"include,omitempty"`
}

func (o *Profiles) GetAppendProfiles() []AppendProfiles {
	if o == nil {
		return nil
	}
	return o.AppendProfiles
}

func (o *Profiles) GetExclude() []Exclude {
	if o == nil {
		return nil
	}
	return o.Exclude
}

func (o *Profiles) GetInclude() []Include {
	if o == nil {
		return nil
	}
	return o.Include
}

// Sidecar metrics collection configuration
type Sidecar struct {
	// IncludeUnused if false will scrape only metrics that has been by sidecar (counters incremented
	// at least once, gauges changed at least once, and histograms added to at
	// least once). If true will scrape all metrics (even the ones with zeros).
	// If not specified then the default value is false.
	IncludeUnused *bool `json:"includeUnused,omitempty"`
	// Profiles allows to customize which metrics are published.
	Profiles *Profiles `json:"profiles,omitempty"`
}

func (o *Sidecar) GetIncludeUnused() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeUnused
}

func (o *Sidecar) GetProfiles() *Profiles {
	if o == nil {
		return nil
	}
	return o.Profiles
}

// Default - MeshMetric configuration.
type Default struct {
	// Applications is a list of application that Dataplane Proxy will scrape
	Applications []Applications `json:"applications,omitempty"`
	// Backends list that will be used to collect metrics.
	Backends []MeshMetricItemBackends `json:"backends,omitempty"`
	// Sidecar metrics collection configuration
	Sidecar *Sidecar `json:"sidecar,omitempty"`
}

func (o *Default) GetApplications() []Applications {
	if o == nil {
		return nil
	}
	return o.Applications
}

func (o *Default) GetBackends() []MeshMetricItemBackends {
	if o == nil {
		return nil
	}
	return o.Backends
}

func (o *Default) GetSidecar() *Sidecar {
	if o == nil {
		return nil
	}
	return o.Sidecar
}

// MeshMetricItemKind - Kind of the referenced resource
type MeshMetricItemKind string

const (
	MeshMetricItemKindMesh                 MeshMetricItemKind = "Mesh"
	MeshMetricItemKindMeshSubset           MeshMetricItemKind = "MeshSubset"
	MeshMetricItemKindMeshGateway          MeshMetricItemKind = "MeshGateway"
	MeshMetricItemKindMeshService          MeshMetricItemKind = "MeshService"
	MeshMetricItemKindMeshExternalService  MeshMetricItemKind = "MeshExternalService"
	MeshMetricItemKindMeshMultiZoneService MeshMetricItemKind = "MeshMultiZoneService"
	MeshMetricItemKindMeshServiceSubset    MeshMetricItemKind = "MeshServiceSubset"
	MeshMetricItemKindMeshHTTPRoute        MeshMetricItemKind = "MeshHTTPRoute"
	MeshMetricItemKindDataplane            MeshMetricItemKind = "Dataplane"
)

func (e MeshMetricItemKind) ToPointer() *MeshMetricItemKind {
	return &e
}
func (e *MeshMetricItemKind) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Mesh":
		fallthrough
	case "MeshSubset":
		fallthrough
	case "MeshGateway":
		fallthrough
	case "MeshService":
		fallthrough
	case "MeshExternalService":
		fallthrough
	case "MeshMultiZoneService":
		fallthrough
	case "MeshServiceSubset":
		fallthrough
	case "MeshHTTPRoute":
		fallthrough
	case "Dataplane":
		*e = MeshMetricItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshMetricItemKind: %v", v)
	}
}

type MeshMetricItemProxyTypes string

const (
	MeshMetricItemProxyTypesSidecar MeshMetricItemProxyTypes = "Sidecar"
	MeshMetricItemProxyTypesGateway MeshMetricItemProxyTypes = "Gateway"
)

func (e MeshMetricItemProxyTypes) ToPointer() *MeshMetricItemProxyTypes {
	return &e
}
func (e *MeshMetricItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshMetricItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshMetricItemProxyTypes: %v", v)
	}
}

// MeshMetricItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined in-place.
type MeshMetricItemTargetRef struct {
	// Kind of the referenced resource
	Kind MeshMetricItemKind `json:"kind"`
	// Labels are used to select group of MeshServices that match labels. Either Labels or
	// Name and Namespace can be used.
	Labels map[string]string `json:"labels,omitempty"`
	// Mesh is reserved for future use to identify cross mesh resources.
	Mesh *string `json:"mesh,omitempty"`
	// Name of the referenced resource. Can only be used with kinds: `MeshService`,
	// `MeshServiceSubset` and `MeshGatewayRoute`
	Name *string `json:"name,omitempty"`
	// Namespace specifies the namespace of target resource. If empty only resources in policy namespace
	// will be targeted.
	Namespace *string `json:"namespace,omitempty"`
	// ProxyTypes specifies the data plane types that are subject to the policy. When not specified,
	// all data plane types are targeted by the policy.
	ProxyTypes []MeshMetricItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshMetricItemTargetRef) GetKind() MeshMetricItemKind {
	if o == nil {
		return MeshMetricItemKind("")
	}
	return o.Kind
}

func (o *MeshMetricItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshMetricItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshMetricItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshMetricItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshMetricItemTargetRef) GetProxyTypes() []MeshMetricItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshMetricItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshMetricItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// MeshMetricItemSpec - Spec is the specification of the Kuma MeshMetric resource.
type MeshMetricItemSpec struct {
	// MeshMetric configuration.
	Default *Default `json:"default,omitempty"`
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined in-place.
	TargetRef *MeshMetricItemTargetRef `json:"targetRef,omitempty"`
}

func (o *MeshMetricItemSpec) GetDefault() *Default {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshMetricItemSpec) GetTargetRef() *MeshMetricItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

// MeshMetricItem - Successful response
type MeshMetricItem struct {
	// the type of the resource
	Type MeshMetricItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshMetric resource.
	Spec MeshMetricItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshMetricItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshMetricItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshMetricItem) GetType() MeshMetricItemType {
	if o == nil {
		return MeshMetricItemType("")
	}
	return o.Type
}

func (o *MeshMetricItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshMetricItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshMetricItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshMetricItem) GetSpec() MeshMetricItemSpec {
	if o == nil {
		return MeshMetricItemSpec{}
	}
	return o.Spec
}

func (o *MeshMetricItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshMetricItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

type MeshMetricItemInput struct {
	// the type of the resource
	Type MeshMetricItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshMetric resource.
	Spec MeshMetricItemSpec `json:"spec"`
}

func (m MeshMetricItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshMetricItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshMetricItemInput) GetType() MeshMetricItemType {
	if o == nil {
		return MeshMetricItemType("")
	}
	return o.Type
}

func (o *MeshMetricItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshMetricItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshMetricItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshMetricItemInput) GetSpec() MeshMetricItemSpec {
	if o == nil {
		return MeshMetricItemSpec{}
	}
	return o.Spec
}
