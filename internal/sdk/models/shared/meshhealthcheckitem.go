// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"time"
)

// MeshHealthCheckItemType - the type of the resource
type MeshHealthCheckItemType string

const (
	MeshHealthCheckItemTypeMeshHealthCheck MeshHealthCheckItemType = "MeshHealthCheck"
)

func (e MeshHealthCheckItemType) ToPointer() *MeshHealthCheckItemType {
	return &e
}
func (e *MeshHealthCheckItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshHealthCheck":
		*e = MeshHealthCheckItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshHealthCheckItemType: %v", v)
	}
}

// MeshHealthCheckItemKind - Kind of the referenced resource
type MeshHealthCheckItemKind string

const (
	MeshHealthCheckItemKindMesh                 MeshHealthCheckItemKind = "Mesh"
	MeshHealthCheckItemKindMeshSubset           MeshHealthCheckItemKind = "MeshSubset"
	MeshHealthCheckItemKindMeshGateway          MeshHealthCheckItemKind = "MeshGateway"
	MeshHealthCheckItemKindMeshService          MeshHealthCheckItemKind = "MeshService"
	MeshHealthCheckItemKindMeshExternalService  MeshHealthCheckItemKind = "MeshExternalService"
	MeshHealthCheckItemKindMeshMultiZoneService MeshHealthCheckItemKind = "MeshMultiZoneService"
	MeshHealthCheckItemKindMeshServiceSubset    MeshHealthCheckItemKind = "MeshServiceSubset"
	MeshHealthCheckItemKindMeshHTTPRoute        MeshHealthCheckItemKind = "MeshHTTPRoute"
	MeshHealthCheckItemKindDataplane            MeshHealthCheckItemKind = "Dataplane"
)

func (e MeshHealthCheckItemKind) ToPointer() *MeshHealthCheckItemKind {
	return &e
}
func (e *MeshHealthCheckItemKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshHealthCheckItemKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshHealthCheckItemKind: %v", v)
	}
}

type MeshHealthCheckItemProxyTypes string

const (
	MeshHealthCheckItemProxyTypesSidecar MeshHealthCheckItemProxyTypes = "Sidecar"
	MeshHealthCheckItemProxyTypesGateway MeshHealthCheckItemProxyTypes = "Gateway"
)

func (e MeshHealthCheckItemProxyTypes) ToPointer() *MeshHealthCheckItemProxyTypes {
	return &e
}
func (e *MeshHealthCheckItemProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshHealthCheckItemProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshHealthCheckItemProxyTypes: %v", v)
	}
}

// MeshHealthCheckItemTargetRef - TargetRef is a reference to the resource the policy takes an effect on.
// The resource could be either a real store object or virtual resource
// defined inplace.
type MeshHealthCheckItemTargetRef struct {
	// Kind of the referenced resource
	Kind MeshHealthCheckItemKind `json:"kind"`
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
	ProxyTypes []MeshHealthCheckItemProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshHealthCheckItemTargetRef) GetKind() MeshHealthCheckItemKind {
	if o == nil {
		return MeshHealthCheckItemKind("")
	}
	return o.Kind
}

func (o *MeshHealthCheckItemTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshHealthCheckItemTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshHealthCheckItemTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshHealthCheckItemTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshHealthCheckItemTargetRef) GetProxyTypes() []MeshHealthCheckItemProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshHealthCheckItemTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshHealthCheckItemTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// Grpc - GrpcHealthCheck defines gRPC configuration which will instruct the service
// the health check will be made for is a gRPC service.
type Grpc struct {
	// The value of the :authority header in the gRPC health check request,
	// by default name of the cluster this health check is associated with
	Authority *string `json:"authority,omitempty"`
	// If true the GrpcHealthCheck is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Service name parameter which will be sent to gRPC service
	ServiceName *string `json:"serviceName,omitempty"`
}

func (o *Grpc) GetAuthority() *string {
	if o == nil {
		return nil
	}
	return o.Authority
}

func (o *Grpc) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *Grpc) GetServiceName() *string {
	if o == nil {
		return nil
	}
	return o.ServiceName
}

type HealthyPanicThresholdType string

const (
	HealthyPanicThresholdTypeInteger HealthyPanicThresholdType = "integer"
	HealthyPanicThresholdTypeStr     HealthyPanicThresholdType = "str"
)

// HealthyPanicThreshold - Allows to configure panic threshold for Envoy cluster. If not specified,
// the default is 50%. To disable panic mode, set to 0%.
// Either int or decimal represented as string.
// Deprecated: the setting has been moved to MeshCircuitBreaker policy,
// please use MeshCircuitBreaker policy instead.
type HealthyPanicThreshold struct {
	Integer *int64  `queryParam:"inline"`
	Str     *string `queryParam:"inline"`

	Type HealthyPanicThresholdType
}

func CreateHealthyPanicThresholdInteger(integer int64) HealthyPanicThreshold {
	typ := HealthyPanicThresholdTypeInteger

	return HealthyPanicThreshold{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateHealthyPanicThresholdStr(str string) HealthyPanicThreshold {
	typ := HealthyPanicThresholdTypeStr

	return HealthyPanicThreshold{
		Str:  &str,
		Type: typ,
	}
}

func (u *HealthyPanicThreshold) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = HealthyPanicThresholdTypeInteger
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = HealthyPanicThresholdTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for HealthyPanicThreshold", string(data))
}

func (u HealthyPanicThreshold) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type HealthyPanicThreshold: all fields are null")
}

type Add struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *Add) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Add) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type Set struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (o *Set) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Set) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// RequestHeadersToAdd - The list of HTTP headers which should be added to each health check
// request
type RequestHeadersToAdd struct {
	Add []Add `json:"add,omitempty"`
	Set []Set `json:"set,omitempty"`
}

func (o *RequestHeadersToAdd) GetAdd() []Add {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *RequestHeadersToAdd) GetSet() []Set {
	if o == nil {
		return nil
	}
	return o.Set
}

// MeshHealthCheckItemHTTP - HttpHealthCheck defines HTTP configuration which will instruct the service
// the health check will be made for is an HTTP service.
type MeshHealthCheckItemHTTP struct {
	// If true the HttpHealthCheck is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// List of HTTP response statuses which are considered healthy
	ExpectedStatuses []int64 `json:"expectedStatuses,omitempty"`
	// The HTTP path which will be requested during the health check
	// (ie. /health)
	// If not specified then the default value is "/"
	Path *string `json:"path,omitempty"`
	// The list of HTTP headers which should be added to each health check
	// request
	RequestHeadersToAdd *RequestHeadersToAdd `json:"requestHeadersToAdd,omitempty"`
}

func (o *MeshHealthCheckItemHTTP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *MeshHealthCheckItemHTTP) GetExpectedStatuses() []int64 {
	if o == nil {
		return nil
	}
	return o.ExpectedStatuses
}

func (o *MeshHealthCheckItemHTTP) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *MeshHealthCheckItemHTTP) GetRequestHeadersToAdd() *RequestHeadersToAdd {
	if o == nil {
		return nil
	}
	return o.RequestHeadersToAdd
}

// TCP - TcpHealthCheck defines configuration for specifying bytes to send and
// expected response during the health check
type TCP struct {
	// If true the TcpHealthCheck is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// List of Base64 encoded blocks of strings expected as a response. When checking the response,
	// "fuzzy" matching is performed such that each block must be found, and
	// in the order specified, but not necessarily contiguous.
	// If not provided or empty, checks will be performed as "connect only" and be marked as successful when TCP connection is successfully established.
	Receive []string `json:"receive,omitempty"`
	// Base64 encoded content of the message which will be sent during the health check to the target
	Send *string `json:"send,omitempty"`
}

func (o *TCP) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *TCP) GetReceive() []string {
	if o == nil {
		return nil
	}
	return o.Receive
}

func (o *TCP) GetSend() *string {
	if o == nil {
		return nil
	}
	return o.Send
}

// MeshHealthCheckItemDefault - Default is a configuration specific to the group of destinations referenced in
// 'targetRef'
type MeshHealthCheckItemDefault struct {
	// If set to true, health check failure events will always be logged. If set
	// to false, only the initial health check failure event will be logged. The
	// default value is false.
	AlwaysLogHealthCheckFailures *bool `json:"alwaysLogHealthCheckFailures,omitempty"`
	// Specifies the path to the file where Envoy can log health check events.
	// If empty, no event log will be written.
	EventLogPath *string `json:"eventLogPath,omitempty"`
	// If set to true, Envoy will not consider any hosts when the cluster is in
	// 'panic mode'. Instead, the cluster will fail all requests as if all hosts
	// are unhealthy. This can help avoid potentially overwhelming a failing
	// service.
	FailTrafficOnPanic *bool `json:"failTrafficOnPanic,omitempty"`
	// GrpcHealthCheck defines gRPC configuration which will instruct the service
	// the health check will be made for is a gRPC service.
	Grpc *Grpc `json:"grpc,omitempty"`
	// Allows to configure panic threshold for Envoy cluster. If not specified,
	// the default is 50%. To disable panic mode, set to 0%.
	// Either int or decimal represented as string.
	// Deprecated: the setting has been moved to MeshCircuitBreaker policy,
	// please use MeshCircuitBreaker policy instead.
	HealthyPanicThreshold *HealthyPanicThreshold `json:"healthyPanicThreshold,omitempty"`
	// Number of consecutive healthy checks before considering a host healthy.
	// If not specified then the default value is 1
	HealthyThreshold *int `json:"healthyThreshold,omitempty"`
	// HttpHealthCheck defines HTTP configuration which will instruct the service
	// the health check will be made for is an HTTP service.
	HTTP *MeshHealthCheckItemHTTP `json:"http,omitempty"`
	// If specified, Envoy will start health checking after a random time in
	// ms between 0 and initialJitter. This only applies to the first health
	// check.
	InitialJitter *string `json:"initialJitter,omitempty"`
	// Interval between consecutive health checks.
	// If not specified then the default value is 1m
	Interval *string `json:"interval,omitempty"`
	// If specified, during every interval Envoy will add IntervalJitter to the
	// wait time.
	IntervalJitter *string `json:"intervalJitter,omitempty"`
	// If specified, during every interval Envoy will add IntervalJitter *
	// IntervalJitterPercent / 100 to the wait time. If IntervalJitter and
	// IntervalJitterPercent are both set, both of them will be used to
	// increase the wait time.
	IntervalJitterPercent *int `json:"intervalJitterPercent,omitempty"`
	// The "no traffic interval" is a special health check interval that is used
	// when a cluster has never had traffic routed to it. This lower interval
	// allows cluster information to be kept up to date, without sending a
	// potentially large amount of active health checking traffic for no reason.
	// Once a cluster has been used for traffic routing, Envoy will shift back
	// to using the standard health check interval that is defined. Note that
	// this interval takes precedence over any other. The default value for "no
	// traffic interval" is 60 seconds.
	NoTrafficInterval *string `json:"noTrafficInterval,omitempty"`
	// Reuse health check connection between health checks. Default is true.
	ReuseConnection *bool `json:"reuseConnection,omitempty"`
	// TcpHealthCheck defines configuration for specifying bytes to send and
	// expected response during the health check
	TCP *TCP `json:"tcp,omitempty"`
	// Maximum time to wait for a health check response.
	// If not specified then the default value is 15s
	Timeout *string `json:"timeout,omitempty"`
	// Number of consecutive unhealthy checks before considering a host
	// unhealthy.
	// If not specified then the default value is 5
	UnhealthyThreshold *int `json:"unhealthyThreshold,omitempty"`
}

func (o *MeshHealthCheckItemDefault) GetAlwaysLogHealthCheckFailures() *bool {
	if o == nil {
		return nil
	}
	return o.AlwaysLogHealthCheckFailures
}

func (o *MeshHealthCheckItemDefault) GetEventLogPath() *string {
	if o == nil {
		return nil
	}
	return o.EventLogPath
}

func (o *MeshHealthCheckItemDefault) GetFailTrafficOnPanic() *bool {
	if o == nil {
		return nil
	}
	return o.FailTrafficOnPanic
}

func (o *MeshHealthCheckItemDefault) GetGrpc() *Grpc {
	if o == nil {
		return nil
	}
	return o.Grpc
}

func (o *MeshHealthCheckItemDefault) GetHealthyPanicThreshold() *HealthyPanicThreshold {
	if o == nil {
		return nil
	}
	return o.HealthyPanicThreshold
}

func (o *MeshHealthCheckItemDefault) GetHealthyThreshold() *int {
	if o == nil {
		return nil
	}
	return o.HealthyThreshold
}

func (o *MeshHealthCheckItemDefault) GetHTTP() *MeshHealthCheckItemHTTP {
	if o == nil {
		return nil
	}
	return o.HTTP
}

func (o *MeshHealthCheckItemDefault) GetInitialJitter() *string {
	if o == nil {
		return nil
	}
	return o.InitialJitter
}

func (o *MeshHealthCheckItemDefault) GetInterval() *string {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *MeshHealthCheckItemDefault) GetIntervalJitter() *string {
	if o == nil {
		return nil
	}
	return o.IntervalJitter
}

func (o *MeshHealthCheckItemDefault) GetIntervalJitterPercent() *int {
	if o == nil {
		return nil
	}
	return o.IntervalJitterPercent
}

func (o *MeshHealthCheckItemDefault) GetNoTrafficInterval() *string {
	if o == nil {
		return nil
	}
	return o.NoTrafficInterval
}

func (o *MeshHealthCheckItemDefault) GetReuseConnection() *bool {
	if o == nil {
		return nil
	}
	return o.ReuseConnection
}

func (o *MeshHealthCheckItemDefault) GetTCP() *TCP {
	if o == nil {
		return nil
	}
	return o.TCP
}

func (o *MeshHealthCheckItemDefault) GetTimeout() *string {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *MeshHealthCheckItemDefault) GetUnhealthyThreshold() *int {
	if o == nil {
		return nil
	}
	return o.UnhealthyThreshold
}

// MeshHealthCheckItemSpecKind - Kind of the referenced resource
type MeshHealthCheckItemSpecKind string

const (
	MeshHealthCheckItemSpecKindMesh                 MeshHealthCheckItemSpecKind = "Mesh"
	MeshHealthCheckItemSpecKindMeshSubset           MeshHealthCheckItemSpecKind = "MeshSubset"
	MeshHealthCheckItemSpecKindMeshGateway          MeshHealthCheckItemSpecKind = "MeshGateway"
	MeshHealthCheckItemSpecKindMeshService          MeshHealthCheckItemSpecKind = "MeshService"
	MeshHealthCheckItemSpecKindMeshExternalService  MeshHealthCheckItemSpecKind = "MeshExternalService"
	MeshHealthCheckItemSpecKindMeshMultiZoneService MeshHealthCheckItemSpecKind = "MeshMultiZoneService"
	MeshHealthCheckItemSpecKindMeshServiceSubset    MeshHealthCheckItemSpecKind = "MeshServiceSubset"
	MeshHealthCheckItemSpecKindMeshHTTPRoute        MeshHealthCheckItemSpecKind = "MeshHTTPRoute"
	MeshHealthCheckItemSpecKindDataplane            MeshHealthCheckItemSpecKind = "Dataplane"
)

func (e MeshHealthCheckItemSpecKind) ToPointer() *MeshHealthCheckItemSpecKind {
	return &e
}
func (e *MeshHealthCheckItemSpecKind) UnmarshalJSON(data []byte) error {
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
		*e = MeshHealthCheckItemSpecKind(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshHealthCheckItemSpecKind: %v", v)
	}
}

type MeshHealthCheckItemSpecProxyTypes string

const (
	MeshHealthCheckItemSpecProxyTypesSidecar MeshHealthCheckItemSpecProxyTypes = "Sidecar"
	MeshHealthCheckItemSpecProxyTypesGateway MeshHealthCheckItemSpecProxyTypes = "Gateway"
)

func (e MeshHealthCheckItemSpecProxyTypes) ToPointer() *MeshHealthCheckItemSpecProxyTypes {
	return &e
}
func (e *MeshHealthCheckItemSpecProxyTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Sidecar":
		fallthrough
	case "Gateway":
		*e = MeshHealthCheckItemSpecProxyTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshHealthCheckItemSpecProxyTypes: %v", v)
	}
}

// MeshHealthCheckItemSpecTargetRef - TargetRef is a reference to the resource that represents a group of
// destinations.
type MeshHealthCheckItemSpecTargetRef struct {
	// Kind of the referenced resource
	Kind MeshHealthCheckItemSpecKind `json:"kind"`
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
	ProxyTypes []MeshHealthCheckItemSpecProxyTypes `json:"proxyTypes,omitempty"`
	// SectionName is used to target specific section of resource.
	// For example, you can target port from MeshService.ports[] by its name. Only traffic to this port will be affected.
	SectionName *string `json:"sectionName,omitempty"`
	// Tags used to select a subset of proxies by tags. Can only be used with kinds
	// `MeshSubset` and `MeshServiceSubset`
	Tags map[string]string `json:"tags,omitempty"`
}

func (o *MeshHealthCheckItemSpecTargetRef) GetKind() MeshHealthCheckItemSpecKind {
	if o == nil {
		return MeshHealthCheckItemSpecKind("")
	}
	return o.Kind
}

func (o *MeshHealthCheckItemSpecTargetRef) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshHealthCheckItemSpecTargetRef) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshHealthCheckItemSpecTargetRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshHealthCheckItemSpecTargetRef) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *MeshHealthCheckItemSpecTargetRef) GetProxyTypes() []MeshHealthCheckItemSpecProxyTypes {
	if o == nil {
		return nil
	}
	return o.ProxyTypes
}

func (o *MeshHealthCheckItemSpecTargetRef) GetSectionName() *string {
	if o == nil {
		return nil
	}
	return o.SectionName
}

func (o *MeshHealthCheckItemSpecTargetRef) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

type MeshHealthCheckItemTo struct {
	// Default is a configuration specific to the group of destinations referenced in
	// 'targetRef'
	Default *MeshHealthCheckItemDefault `json:"default,omitempty"`
	// TargetRef is a reference to the resource that represents a group of
	// destinations.
	TargetRef MeshHealthCheckItemSpecTargetRef `json:"targetRef"`
}

func (o *MeshHealthCheckItemTo) GetDefault() *MeshHealthCheckItemDefault {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *MeshHealthCheckItemTo) GetTargetRef() MeshHealthCheckItemSpecTargetRef {
	if o == nil {
		return MeshHealthCheckItemSpecTargetRef{}
	}
	return o.TargetRef
}

// MeshHealthCheckItemSpec - Spec is the specification of the Kuma MeshHealthCheck resource.
type MeshHealthCheckItemSpec struct {
	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *MeshHealthCheckItemTargetRef `json:"targetRef,omitempty"`
	// To list makes a match between the consumed services and corresponding configurations
	To []MeshHealthCheckItemTo `json:"to,omitempty"`
}

func (o *MeshHealthCheckItemSpec) GetTargetRef() *MeshHealthCheckItemTargetRef {
	if o == nil {
		return nil
	}
	return o.TargetRef
}

func (o *MeshHealthCheckItemSpec) GetTo() []MeshHealthCheckItemTo {
	if o == nil {
		return nil
	}
	return o.To
}

// MeshHealthCheckItem - Successful response
type MeshHealthCheckItem struct {
	// the type of the resource
	Type MeshHealthCheckItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshHealthCheck resource.
	Spec MeshHealthCheckItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
}

func (m MeshHealthCheckItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshHealthCheckItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshHealthCheckItem) GetType() MeshHealthCheckItemType {
	if o == nil {
		return MeshHealthCheckItemType("")
	}
	return o.Type
}

func (o *MeshHealthCheckItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshHealthCheckItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshHealthCheckItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshHealthCheckItem) GetSpec() MeshHealthCheckItemSpec {
	if o == nil {
		return MeshHealthCheckItemSpec{}
	}
	return o.Spec
}

func (o *MeshHealthCheckItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshHealthCheckItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

type MeshHealthCheckItemInput struct {
	// the type of the resource
	Type MeshHealthCheckItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshHealthCheck resource.
	Spec MeshHealthCheckItemSpec `json:"spec"`
}

func (m MeshHealthCheckItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshHealthCheckItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshHealthCheckItemInput) GetType() MeshHealthCheckItemType {
	if o == nil {
		return MeshHealthCheckItemType("")
	}
	return o.Type
}

func (o *MeshHealthCheckItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshHealthCheckItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshHealthCheckItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshHealthCheckItemInput) GetSpec() MeshHealthCheckItemSpec {
	if o == nil {
		return MeshHealthCheckItemSpec{}
	}
	return o.Spec
}
