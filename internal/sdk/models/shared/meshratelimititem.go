// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"time"
)

// MeshRateLimitItemType - the type of the resource
type MeshRateLimitItemType string

const (
	MeshRateLimitItemTypeMeshRateLimit MeshRateLimitItemType = "MeshRateLimit"
)

func (e MeshRateLimitItemType) ToPointer() *MeshRateLimitItemType {
	return &e
}
func (e *MeshRateLimitItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshRateLimit":
		*e = MeshRateLimitItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRateLimitItemType: %v", v)
	}
}

type MeshRateLimitItemAdd struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *MeshRateLimitItemAdd) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItemAdd) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type MeshRateLimitItemSet struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *MeshRateLimitItemSet) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItemSet) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// MeshRateLimitItemHeaders - The Headers to be added to the HTTP response on a rate limit event
type MeshRateLimitItemHeaders struct {
	Add []MeshRateLimitItemAdd `json:"add,omitempty"`
	Set []MeshRateLimitItemSet `json:"set,omitempty"`
}

func (o *MeshRateLimitItemHeaders) GetAdd() []MeshRateLimitItemAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *MeshRateLimitItemHeaders) GetSet() []MeshRateLimitItemSet {
	if o == nil {
		return nil
	}
	return o.Set
}

// MeshRateLimitItemSpecFromOnRateLimit - Describes the actions to take on a rate limit event
type MeshRateLimitItemSpecFromOnRateLimit struct {
	// The Headers to be added to the HTTP response on a rate limit event
	Headers *MeshRateLimitItemHeaders `json:"headers,omitempty"`
	// The HTTP status code to be set on a rate limit event
	Status *int `json:"status,omitempty"`
}

func (o *MeshRateLimitItemSpecFromOnRateLimit) GetHeaders() *MeshRateLimitItemHeaders {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *MeshRateLimitItemSpecFromOnRateLimit) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// MeshRateLimitItemSpecFromRequestRate - Defines how many requests are allowed per interval.
type MeshRateLimitItemSpecFromRequestRate struct {
	// The interval the number of units is accounted for.
	Interval string `json:"interval"`
	// Number of units per interval (depending on usage it can be a number of requests,
	// or a number of connections).
	Num int `json:"num"`
}

func (o *MeshRateLimitItemSpecFromRequestRate) GetInterval() string {
	if o == nil {
		return ""
	}
	return o.Interval
}

func (o *MeshRateLimitItemSpecFromRequestRate) GetNum() int {
	if o == nil {
		return 0
	}
	return o.Num
}

// MeshRateLimitItemHTTP - LocalHTTP defines configuration of local HTTP rate limiting
// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter
type MeshRateLimitItemHTTP struct {
	// Define if rate limiting should be disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// Describes the actions to take on a rate limit event
	OnRateLimit *MeshRateLimitItemSpecFromOnRateLimit `json:"onRateLimit,omitempty"`
	// Defines how many requests are allowed per interval.
	RequestRate *MeshRateLimitItemSpecFromRequestRate `json:"requestRate,omitempty"`
}

func (o *MeshRateLimitItemHTTP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *MeshRateLimitItemHTTP) GetOnRateLimit() *MeshRateLimitItemSpecFromOnRateLimit {
	if o == nil {
		return nil
	}
	return o.OnRateLimit
}

func (o *MeshRateLimitItemHTTP) GetRequestRate() *MeshRateLimitItemSpecFromRequestRate {
	if o == nil {
		return nil
	}
	return o.RequestRate
}

// ConnectionRate - Defines how many connections are allowed per interval.
type ConnectionRate struct {
	// The interval the number of units is accounted for.
	Interval string `json:"interval"`
	// Number of units per interval (depending on usage it can be a number of requests,
	// or a number of connections).
	Num int `json:"num"`
}

func (o *ConnectionRate) GetInterval() string {
	if o == nil {
		return ""
	}
	return o.Interval
}

func (o *ConnectionRate) GetNum() int {
	if o == nil {
		return 0
	}
	return o.Num
}

// MeshRateLimitItemTCP - LocalTCP defines confguration of local TCP rate limiting
// https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter
type MeshRateLimitItemTCP struct {
	// Defines how many connections are allowed per interval.
	ConnectionRate *ConnectionRate `json:"connectionRate,omitempty"`
	// Define if rate limiting should be disabled.
	// Default: false
	Disabled *bool `json:"disabled,omitempty"`
}

func (o *MeshRateLimitItemTCP) GetConnectionRate() *ConnectionRate {
	if o == nil {
		return nil
	}
	return o.ConnectionRate
}

func (o *MeshRateLimitItemTCP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

// Local - LocalConf defines local http or/and tcp rate limit configuration
type Local struct {
	// LocalHTTP defines configuration of local HTTP rate limiting
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter
	HTTP *MeshRateLimitItemHTTP `json:"http,omitempty"`
	// LocalTCP defines confguration of local TCP rate limiting
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter
	TCP *MeshRateLimitItemTCP `json:"tcp,omitempty"`
}

func (o *Local) GetHTTP() *MeshRateLimitItemHTTP {
	if o == nil {
		return nil
	}
	return o.HTTP
}

func (o *Local) GetTCP() *MeshRateLimitItemTCP {
	if o == nil {
		return nil
	}
	return o.TCP
}

// MeshRateLimitItemDefault - Default is a configuration specific to the group of clients referenced in
// 'targetRef'
type MeshRateLimitItemDefault struct {
	// LocalConf defines local http or/and tcp rate limit configuration
	Local *Local `json:"local,omitempty"`
}

func (o *MeshRateLimitItemDefault) GetLocal() *Local {
	if o == nil {
		return nil
	}
	return o.Local
}

// MeshRateLimitItemSpecKind - Kind of the referenced resource
type MeshRateLimitItemSpecKind string

const (
	MeshRateLimitItemSpecKindMesh                 MeshRateLimitItemSpecKind = "Mesh"
	MeshRateLimitItemSpecKindMeshSubset           MeshRateLimitItemSpecKind = "MeshSubset"
	MeshRateLimitItemSpecKindMeshGateway          MeshRateLimitItemSpecKind = "MeshGateway"
	MeshRateLimitItemSpecKindMeshService          MeshRateLimitItemSpecKind = "MeshService"
	MeshRateLimitItemSpecKindMeshExternalService  MeshRateLimitItemSpecKind = "MeshExternalService"
	MeshRateLimitItemSpecKindMeshMultiZoneService MeshRateLimitItemSpecKind = "MeshMultiZoneService"
	MeshRateLimitItemSpecKindMeshServiceSubset    MeshRateLimitItemSpecKind = "MeshServiceSubset"
	MeshRateLimitItemSpecKindMeshHTTPRoute        MeshRateLimitItemSpecKind = "MeshHTTPRoute"
	MeshRateLimitItemSpecKindDataplane            MeshRateLimitItemSpecKind = "Dataplane"
)

func (e MeshRateLimitItemSpecKind) ToPointer() *MeshRateLimitItemSpecKind {
	return &e
}
func (e *MeshRateLimitItemSpecKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshRateLimitItemSpecKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRateLimitItemSpecKind: %v", v)
	}
}

type MeshRateLimitItemSpecProxyTypes string

const (
	MeshRateLimitItemSpecProxyTypesSidecar MeshRateLimitItemSpecProxyTypes = "Sidecar"
	MeshRateLimitItemSpecProxyTypesGateway MeshRateLimitItemSpecProxyTypes = "Gateway"
)

func (e MeshRateLimitItemSpecProxyTypes) ToPointer() *MeshRateLimitItemSpecProxyTypes {
	return &e
}
func (e *MeshRateLimitItemSpecProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshRateLimitItemSpecProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRateLimitItemSpecProxyTypes: %v", v)
	}
}

// MeshRateLimitItemSpecTargetRef - TargetRef is a reference to the resource that represents a group of
// clients.
type MeshRateLimitItemSpecTargetRef struct {
	// Kind of the referenced resource
	Kind MeshRateLimitItemSpecKind `json:"kind"`
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
	ProxyTypes []MeshRateLimitItemSpecProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshRateLimitItemSpecTargetRef) GetKind() MeshRateLimitItemSpecKind {
	if o == nil {
		return MeshRateLimitItemSpecKind("")
	}
	return o.Kind
}

func (o *MeshRateLimitItemSpecTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRateLimitItemSpecTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRateLimitItemSpecTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshRateLimitItemSpecTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshRateLimitItemSpecTargetRef) GetProxyTypes() []MeshRateLimitItemSpecProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshRateLimitItemSpecTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshRateLimitItemSpecTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshRateLimitItemFrom struct {
	// Default is a configuration specific to the group of clients referenced in
	// 'targetRef'
	Default *MeshRateLimitItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource that represents a group of
	// clients.
	TargetRef MeshRateLimitItemSpecTargetRef `json:"targetRef"`
}

func (o *MeshRateLimitItemFrom) GetDefault() *MeshRateLimitItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshRateLimitItemFrom) GetTargetRef() MeshRateLimitItemSpecTargetRef {
	if o == nil {
		return MeshRateLimitItemSpecTargetRef{}
	}
	return o.TargetRef
}

type MeshRateLimitItemSpecAdd struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *MeshRateLimitItemSpecAdd) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItemSpecAdd) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type MeshRateLimitItemSpecSet struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *MeshRateLimitItemSpecSet) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItemSpecSet) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// MeshRateLimitItemSpecHeaders - The Headers to be added to the HTTP response on a rate limit event
type MeshRateLimitItemSpecHeaders struct {
	Add []MeshRateLimitItemSpecAdd `json:"add,omitempty"`
	Set []MeshRateLimitItemSpecSet `json:"set,omitempty"`
}

func (o *MeshRateLimitItemSpecHeaders) GetAdd() []MeshRateLimitItemSpecAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *MeshRateLimitItemSpecHeaders) GetSet() []MeshRateLimitItemSpecSet {
	if o == nil {
		return nil
	}
	return o.Set
}

// MeshRateLimitItemOnRateLimit - Describes the actions to take on a rate limit event
type MeshRateLimitItemOnRateLimit struct {
	// The Headers to be added to the HTTP response on a rate limit event
	Headers *MeshRateLimitItemSpecHeaders `json:"headers,omitempty"`
	// The HTTP status code to be set on a rate limit event
	Status *int `json:"status,omitempty"`
}

func (o *MeshRateLimitItemOnRateLimit) GetHeaders() *MeshRateLimitItemSpecHeaders {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *MeshRateLimitItemOnRateLimit) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// MeshRateLimitItemRequestRate - Defines how many requests are allowed per interval.
type MeshRateLimitItemRequestRate struct {
	// The interval the number of units is accounted for.
	Interval string `json:"interval"`
	// Number of units per interval (depending on usage it can be a number of requests,
	// or a number of connections).
	Num int `json:"num"`
}

func (o *MeshRateLimitItemRequestRate) GetInterval() string {
	if o == nil {
		return ""
	}
	return o.Interval
}

func (o *MeshRateLimitItemRequestRate) GetNum() int {
	if o == nil {
		return 0
	}
	return o.Num
}

// MeshRateLimitItemSpecHTTP - LocalHTTP defines configuration of local HTTP rate limiting
// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter
type MeshRateLimitItemSpecHTTP struct {
	// Define if rate limiting should be disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// Describes the actions to take on a rate limit event
	OnRateLimit *MeshRateLimitItemOnRateLimit `json:"onRateLimit,omitempty"`
	// Defines how many requests are allowed per interval.
	RequestRate *MeshRateLimitItemRequestRate `json:"requestRate,omitempty"`
}

func (o *MeshRateLimitItemSpecHTTP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *MeshRateLimitItemSpecHTTP) GetOnRateLimit() *MeshRateLimitItemOnRateLimit {
	if o == nil {
		return nil
	}
	return o.OnRateLimit
}

func (o *MeshRateLimitItemSpecHTTP) GetRequestRate() *MeshRateLimitItemRequestRate {
	if o == nil {
		return nil
	}
	return o.RequestRate
}

// MeshRateLimitItemConnectionRate - Defines how many connections are allowed per interval.
type MeshRateLimitItemConnectionRate struct {
	// The interval the number of units is accounted for.
	Interval string `json:"interval"`
	// Number of units per interval (depending on usage it can be a number of requests,
	// or a number of connections).
	Num int `json:"num"`
}

func (o *MeshRateLimitItemConnectionRate) GetInterval() string {
	if o == nil {
		return ""
	}
	return o.Interval
}

func (o *MeshRateLimitItemConnectionRate) GetNum() int {
	if o == nil {
		return 0
	}
	return o.Num
}

// MeshRateLimitItemSpecTCP - LocalTCP defines confguration of local TCP rate limiting
// https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter
type MeshRateLimitItemSpecTCP struct {
	// Defines how many connections are allowed per interval.
	ConnectionRate *MeshRateLimitItemConnectionRate `json:"connectionRate,omitempty"`
	// Define if rate limiting should be disabled.
	// Default: false
	Disabled *bool `json:"disabled,omitempty"`
}

func (o *MeshRateLimitItemSpecTCP) GetConnectionRate() *MeshRateLimitItemConnectionRate {
	if o == nil {
		return nil
	}
	return o.ConnectionRate
}

func (o *MeshRateLimitItemSpecTCP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

// MeshRateLimitItemLocal - LocalConf defines local http or/and tcp rate limit configuration
type MeshRateLimitItemLocal struct {
	// LocalHTTP defines configuration of local HTTP rate limiting
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter
	HTTP *MeshRateLimitItemSpecHTTP `json:"http,omitempty"`
	// LocalTCP defines confguration of local TCP rate limiting
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter
	TCP *MeshRateLimitItemSpecTCP `json:"tcp,omitempty"`
}

func (o *MeshRateLimitItemLocal) GetHTTP() *MeshRateLimitItemSpecHTTP {
	if o == nil {
		return nil
	}
	return o.HTTP
}

func (o *MeshRateLimitItemLocal) GetTCP() *MeshRateLimitItemSpecTCP {
	if o == nil {
		return nil
	}
	return o.TCP
}

// MeshRateLimitItemSpecDefault - Default contains configuration of the inbound rate limits
type MeshRateLimitItemSpecDefault struct {
	// LocalConf defines local http or/and tcp rate limit configuration
	Local *MeshRateLimitItemLocal `json:"local,omitempty"`
}

func (o *MeshRateLimitItemSpecDefault) GetLocal() *MeshRateLimitItemLocal {
	if o == nil {
		return nil
	}
	return o.Local
}

type MeshRateLimitItemRules struct {
	// Default contains configuration of the inbound rate limits
	Default *MeshRateLimitItemSpecDefault `json:"default,omitempty"`
}

func (o *MeshRateLimitItemRules) GetDefault() *MeshRateLimitItemSpecDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

// MeshRateLimitItemKind - Kind of the referenced resource
type MeshRateLimitItemKind string

const (
	MeshRateLimitItemKindMesh                 MeshRateLimitItemKind = "Mesh"
	MeshRateLimitItemKindMeshSubset           MeshRateLimitItemKind = "MeshSubset"
	MeshRateLimitItemKindMeshGateway          MeshRateLimitItemKind = "MeshGateway"
	MeshRateLimitItemKindMeshService          MeshRateLimitItemKind = "MeshService"
	MeshRateLimitItemKindMeshExternalService  MeshRateLimitItemKind = "MeshExternalService"
	MeshRateLimitItemKindMeshMultiZoneService MeshRateLimitItemKind = "MeshMultiZoneService"
	MeshRateLimitItemKindMeshServiceSubset    MeshRateLimitItemKind = "MeshServiceSubset"
	MeshRateLimitItemKindMeshHTTPRoute        MeshRateLimitItemKind = "MeshHTTPRoute"
	MeshRateLimitItemKindDataplane            MeshRateLimitItemKind = "Dataplane"
)

func (e MeshRateLimitItemKind) ToPointer() *MeshRateLimitItemKind {
	return &e
}
func (e *MeshRateLimitItemKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshRateLimitItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRateLimitItemKind: %v", v)
	}
}

type MeshRateLimitItemProxyTypes string

const (
	MeshRateLimitItemProxyTypesSidecar MeshRateLimitItemProxyTypes = "Sidecar"
	MeshRateLimitItemProxyTypesGateway MeshRateLimitItemProxyTypes = "Gateway"
)

func (e MeshRateLimitItemProxyTypes) ToPointer() *MeshRateLimitItemProxyTypes {
	return &e
}
func (e *MeshRateLimitItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshRateLimitItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRateLimitItemProxyTypes: %v", v)
	}
}

// MeshRateLimitItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined inplace.
type MeshRateLimitItemTargetRef struct {
	// Kind of the referenced resource
	Kind MeshRateLimitItemKind `json:"kind"`
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
	ProxyTypes []MeshRateLimitItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshRateLimitItemTargetRef) GetKind() MeshRateLimitItemKind {
	if o == nil {
		return MeshRateLimitItemKind("")
	}
	return o.Kind
}

func (o *MeshRateLimitItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRateLimitItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRateLimitItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshRateLimitItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshRateLimitItemTargetRef) GetProxyTypes() []MeshRateLimitItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshRateLimitItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshRateLimitItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshRateLimitItemSpecToAdd struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *MeshRateLimitItemSpecToAdd) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItemSpecToAdd) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type MeshRateLimitItemSpecToSet struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *MeshRateLimitItemSpecToSet) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItemSpecToSet) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// MeshRateLimitItemSpecToHeaders - The Headers to be added to the HTTP response on a rate limit event
type MeshRateLimitItemSpecToHeaders struct {
	Add []MeshRateLimitItemSpecToAdd `json:"add,omitempty"`
	Set []MeshRateLimitItemSpecToSet `json:"set,omitempty"`
}

func (o *MeshRateLimitItemSpecToHeaders) GetAdd() []MeshRateLimitItemSpecToAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *MeshRateLimitItemSpecToHeaders) GetSet() []MeshRateLimitItemSpecToSet {
	if o == nil {
		return nil
	}
	return o.Set
}

// MeshRateLimitItemSpecOnRateLimit - Describes the actions to take on a rate limit event
type MeshRateLimitItemSpecOnRateLimit struct {
	// The Headers to be added to the HTTP response on a rate limit event
	Headers *MeshRateLimitItemSpecToHeaders `json:"headers,omitempty"`
	// The HTTP status code to be set on a rate limit event
	Status *int `json:"status,omitempty"`
}

func (o *MeshRateLimitItemSpecOnRateLimit) GetHeaders() *MeshRateLimitItemSpecToHeaders {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *MeshRateLimitItemSpecOnRateLimit) GetStatus() *int {
	if o == nil {
		return nil
	}
	return o.Status
}

// MeshRateLimitItemSpecRequestRate - Defines how many requests are allowed per interval.
type MeshRateLimitItemSpecRequestRate struct {
	// The interval the number of units is accounted for.
	Interval string `json:"interval"`
	// Number of units per interval (depending on usage it can be a number of requests,
	// or a number of connections).
	Num int `json:"num"`
}

func (o *MeshRateLimitItemSpecRequestRate) GetInterval() string {
	if o == nil {
		return ""
	}
	return o.Interval
}

func (o *MeshRateLimitItemSpecRequestRate) GetNum() int {
	if o == nil {
		return 0
	}
	return o.Num
}

// MeshRateLimitItemSpecToHTTP - LocalHTTP defines configuration of local HTTP rate limiting
// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter
type MeshRateLimitItemSpecToHTTP struct {
	// Define if rate limiting should be disabled.
	Disabled *bool `json:"disabled,omitempty"`
	// Describes the actions to take on a rate limit event
	OnRateLimit *MeshRateLimitItemSpecOnRateLimit `json:"onRateLimit,omitempty"`
	// Defines how many requests are allowed per interval.
	RequestRate *MeshRateLimitItemSpecRequestRate `json:"requestRate,omitempty"`
}

func (o *MeshRateLimitItemSpecToHTTP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *MeshRateLimitItemSpecToHTTP) GetOnRateLimit() *MeshRateLimitItemSpecOnRateLimit {
	if o == nil {
		return nil
	}
	return o.OnRateLimit
}

func (o *MeshRateLimitItemSpecToHTTP) GetRequestRate() *MeshRateLimitItemSpecRequestRate {
	if o == nil {
		return nil
	}
	return o.RequestRate
}

// MeshRateLimitItemSpecConnectionRate - Defines how many connections are allowed per interval.
type MeshRateLimitItemSpecConnectionRate struct {
	// The interval the number of units is accounted for.
	Interval string `json:"interval"`
	// Number of units per interval (depending on usage it can be a number of requests,
	// or a number of connections).
	Num int `json:"num"`
}

func (o *MeshRateLimitItemSpecConnectionRate) GetInterval() string {
	if o == nil {
		return ""
	}
	return o.Interval
}

func (o *MeshRateLimitItemSpecConnectionRate) GetNum() int {
	if o == nil {
		return 0
	}
	return o.Num
}

// MeshRateLimitItemSpecToTCP - LocalTCP defines confguration of local TCP rate limiting
// https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter
type MeshRateLimitItemSpecToTCP struct {
	// Defines how many connections are allowed per interval.
	ConnectionRate *MeshRateLimitItemSpecConnectionRate `json:"connectionRate,omitempty"`
	// Define if rate limiting should be disabled.
	// Default: false
	Disabled *bool `json:"disabled,omitempty"`
}

func (o *MeshRateLimitItemSpecToTCP) GetConnectionRate() *MeshRateLimitItemSpecConnectionRate {
	if o == nil {
		return nil
	}
	return o.ConnectionRate
}

func (o *MeshRateLimitItemSpecToTCP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

// MeshRateLimitItemSpecLocal - LocalConf defines local http or/and tcp rate limit configuration
type MeshRateLimitItemSpecLocal struct {
	// LocalHTTP defines configuration of local HTTP rate limiting
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/local_rate_limit_filter
	HTTP *MeshRateLimitItemSpecToHTTP `json:"http,omitempty"`
	// LocalTCP defines confguration of local TCP rate limiting
	// https://www.envoyproxy.io/docs/envoy/latest/configuration/listeners/network_filters/local_rate_limit_filter
	TCP *MeshRateLimitItemSpecToTCP `json:"tcp,omitempty"`
}

func (o *MeshRateLimitItemSpecLocal) GetHTTP() *MeshRateLimitItemSpecToHTTP {
	if o == nil {
		return nil
	}
	return o.HTTP
}

func (o *MeshRateLimitItemSpecLocal) GetTCP() *MeshRateLimitItemSpecToTCP {
	if o == nil {
		return nil
	}
	return o.TCP
}

// MeshRateLimitItemSpecToDefault - Default is a configuration specific to the group of clients referenced in
// 'targetRef'
type MeshRateLimitItemSpecToDefault struct {
	// LocalConf defines local http or/and tcp rate limit configuration
	Local *MeshRateLimitItemSpecLocal `json:"local,omitempty"`
}

func (o *MeshRateLimitItemSpecToDefault) GetLocal() *MeshRateLimitItemSpecLocal {
	if o == nil {
		return nil
	}
	return o.Local
}

// MeshRateLimitItemSpecToKind - Kind of the referenced resource
type MeshRateLimitItemSpecToKind string

const (
	MeshRateLimitItemSpecToKindMesh                 MeshRateLimitItemSpecToKind = "Mesh"
	MeshRateLimitItemSpecToKindMeshSubset           MeshRateLimitItemSpecToKind = "MeshSubset"
	MeshRateLimitItemSpecToKindMeshGateway          MeshRateLimitItemSpecToKind = "MeshGateway"
	MeshRateLimitItemSpecToKindMeshService          MeshRateLimitItemSpecToKind = "MeshService"
	MeshRateLimitItemSpecToKindMeshExternalService  MeshRateLimitItemSpecToKind = "MeshExternalService"
	MeshRateLimitItemSpecToKindMeshMultiZoneService MeshRateLimitItemSpecToKind = "MeshMultiZoneService"
	MeshRateLimitItemSpecToKindMeshServiceSubset    MeshRateLimitItemSpecToKind = "MeshServiceSubset"
	MeshRateLimitItemSpecToKindMeshHTTPRoute        MeshRateLimitItemSpecToKind = "MeshHTTPRoute"
	MeshRateLimitItemSpecToKindDataplane            MeshRateLimitItemSpecToKind = "Dataplane"
)

func (e MeshRateLimitItemSpecToKind) ToPointer() *MeshRateLimitItemSpecToKind {
	return &e
}
func (e *MeshRateLimitItemSpecToKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshRateLimitItemSpecToKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRateLimitItemSpecToKind: %v", v)
	}
}

type MeshRateLimitItemSpecToProxyTypes string

const (
	MeshRateLimitItemSpecToProxyTypesSidecar MeshRateLimitItemSpecToProxyTypes = "Sidecar"
	MeshRateLimitItemSpecToProxyTypesGateway MeshRateLimitItemSpecToProxyTypes = "Gateway"
)

func (e MeshRateLimitItemSpecToProxyTypes) ToPointer() *MeshRateLimitItemSpecToProxyTypes {
	return &e
}
func (e *MeshRateLimitItemSpecToProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshRateLimitItemSpecToProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshRateLimitItemSpecToProxyTypes: %v", v)
	}
}

// MeshRateLimitItemSpecToTargetRef - TargetRef is a reference to the resource that represents a group of
// clients.
type MeshRateLimitItemSpecToTargetRef struct {
	// Kind of the referenced resource
	Kind MeshRateLimitItemSpecToKind `json:"kind"`
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
	ProxyTypes []MeshRateLimitItemSpecToProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshRateLimitItemSpecToTargetRef) GetKind() MeshRateLimitItemSpecToKind {
	if o == nil {
		return MeshRateLimitItemSpecToKind("")
	}
	return o.Kind
}

func (o *MeshRateLimitItemSpecToTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRateLimitItemSpecToTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRateLimitItemSpecToTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshRateLimitItemSpecToTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshRateLimitItemSpecToTargetRef) GetProxyTypes() []MeshRateLimitItemSpecToProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshRateLimitItemSpecToTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshRateLimitItemSpecToTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshRateLimitItemTo struct {
	// Default is a configuration specific to the group of clients referenced in
	// 'targetRef'
	Default *MeshRateLimitItemSpecToDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource that represents a group of
	// clients.
	TargetRef MeshRateLimitItemSpecToTargetRef `json:"targetRef"`
}

func (o *MeshRateLimitItemTo) GetDefault() *MeshRateLimitItemSpecToDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshRateLimitItemTo) GetTargetRef() MeshRateLimitItemSpecToTargetRef {
	if o == nil {
		return MeshRateLimitItemSpecToTargetRef{}
	}
	return o.TargetRef
}

// MeshRateLimitItemSpec - Spec is the specification of the Kuma MeshRateLimit resource.
type MeshRateLimitItemSpec struct {
	// From list makes a match between clients and corresponding configurations
	From []MeshRateLimitItemFrom `json:"from,omitempty"`
	// Rules defines inbound rate limiting configurations. Currently limited to
	// selecting all inbound traffic, as L7 matching is not yet implemented.
	Rules []MeshRateLimitItemRules `json:"rules,omitempty"`
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *MeshRateLimitItemTargetRef `json:"targetRef,omitempty"`
	// To list makes a match between clients and corresponding configurations
	To []MeshRateLimitItemTo `json:"to,omitempty"`
}

func (o *MeshRateLimitItemSpec) GetFrom() []MeshRateLimitItemFrom {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *MeshRateLimitItemSpec) GetRules() []MeshRateLimitItemRules {
	if o == nil {
		return nil
	}
	return o.Rules
}

func (o *MeshRateLimitItemSpec) GetTargetRef() *MeshRateLimitItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

func (o *MeshRateLimitItemSpec) GetTo() []MeshRateLimitItemTo {
	if o == nil {
		return nil
	}
	return o.To
}

// MeshRateLimitItem - Successful response
type MeshRateLimitItem struct {
	// the type of the resource
	Type MeshRateLimitItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshRateLimit resource.
	Spec MeshRateLimitItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshRateLimitItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshRateLimitItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshRateLimitItem) GetType() MeshRateLimitItemType {
	if o == nil {
		return MeshRateLimitItemType("")
	}
	return o.Type
}

func (o *MeshRateLimitItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRateLimitItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRateLimitItem) GetSpec() MeshRateLimitItemSpec {
	if o == nil {
		return MeshRateLimitItemSpec{}
	}
	return o.Spec
}

func (o *MeshRateLimitItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshRateLimitItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

type MeshRateLimitItemInput struct {
	// the type of the resource
	Type MeshRateLimitItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshRateLimit resource.
	Spec MeshRateLimitItemSpec `json:"spec"`
}

func (m MeshRateLimitItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshRateLimitItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshRateLimitItemInput) GetType() MeshRateLimitItemType {
	if o == nil {
		return MeshRateLimitItemType("")
	}
	return o.Type
}

func (o *MeshRateLimitItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshRateLimitItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshRateLimitItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshRateLimitItemInput) GetSpec() MeshRateLimitItemSpec {
	if o == nil {
		return MeshRateLimitItemSpec{}
	}
	return o.Spec
}
