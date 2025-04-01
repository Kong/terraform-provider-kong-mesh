// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"time"
)

// MeshServiceItemType - the type of the resource
type MeshServiceItemType string

const (
	MeshServiceItemTypeMeshService MeshServiceItemType = "MeshService"
)

func (e MeshServiceItemType) ToPointer() *MeshServiceItemType {
	return &e
}
func (e *MeshServiceItemType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MeshService":
		*e = MeshServiceItemType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshServiceItemType: %v", v)
	}
}

type MeshServiceItemSpecType string

const (
	MeshServiceItemSpecTypeServiceTag MeshServiceItemSpecType = "ServiceTag"
)

func (e MeshServiceItemSpecType) ToPointer() *MeshServiceItemSpecType {
	return &e
}
func (e *MeshServiceItemSpecType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ServiceTag":
		*e = MeshServiceItemSpecType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshServiceItemSpecType: %v", v)
	}
}

type Identities struct {
	Type  MeshServiceItemSpecType `json:"type"`
	Value string                  `json:"value"`
}

func (o *Identities) GetType() MeshServiceItemSpecType {
	if o == nil {
		return MeshServiceItemSpecType("")
	}
	return o.Type
}

func (o *Identities) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type TargetPortType string

const (
	TargetPortTypeInteger TargetPortType = "integer"
	TargetPortTypeStr     TargetPortType = "str"
)

type TargetPort struct {
	Integer *int64  `queryParam:"inline"`
	Str     *string `queryParam:"inline"`

	Type TargetPortType
}

func CreateTargetPortInteger(integer int64) TargetPort {
	typ := TargetPortTypeInteger

	return TargetPort{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateTargetPortStr(str string) TargetPort {
	typ := TargetPortTypeStr

	return TargetPort{
		Str:  &str,
		Type: typ,
	}
}

func (u *TargetPort) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = TargetPortTypeInteger
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = TargetPortTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TargetPort", string(data))
}

func (u TargetPort) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type TargetPort: all fields are null")
}

type MeshServiceItemPorts struct {
	// Protocol identifies a protocol supported by a service.
	AppProtocol *string     `default:"tcp" json:"appProtocol"`
	Name        *string     `json:"name,omitempty"`
	Port        int         `json:"port"`
	TargetPort  *TargetPort `json:"targetPort,omitempty"`
}

func (m MeshServiceItemPorts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshServiceItemPorts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshServiceItemPorts) GetAppProtocol() *string {
	if o == nil {
		return nil
	}
	return o.AppProtocol
}

func (o *MeshServiceItemPorts) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MeshServiceItemPorts) GetPort() int {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *MeshServiceItemPorts) GetTargetPort() *TargetPort {
	if o == nil {
		return nil
	}
	return o.TargetPort
}

type DataplaneRef struct {
	Name *string `json:"name,omitempty"`
}

func (o *DataplaneRef) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type MeshServiceItemSelector struct {
	DataplaneRef  *DataplaneRef     `json:"dataplaneRef,omitempty"`
	DataplaneTags map[string]string `json:"dataplaneTags,omitempty"`
}

func (o *MeshServiceItemSelector) GetDataplaneRef() *DataplaneRef {
	if o == nil {
		return nil
	}
	return o.DataplaneRef
}

func (o *MeshServiceItemSelector) GetDataplaneTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.DataplaneTags
}

// State of MeshService. Available if there is at least one healthy endpoint. Otherwise, Unavailable.
// It's used for cross zone communication to check if we should send traffic to it, when MeshService is aggregated into MeshMultiZoneService.
type State string

const (
	StateAvailable   State = "Available"
	StateUnavailable State = "Unavailable"
)

func (e State) ToPointer() *State {
	return &e
}
func (e *State) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Available":
		fallthrough
	case "Unavailable":
		*e = State(v)
		return nil
	default:
		return fmt.Errorf("invalid value for State: %v", v)
	}
}

// MeshServiceItemSpec - Spec is the specification of the Kuma MeshService resource.
type MeshServiceItemSpec struct {
	Identities []Identities             `json:"identities,omitempty"`
	Ports      []MeshServiceItemPorts   `json:"ports,omitempty"`
	Selector   *MeshServiceItemSelector `json:"selector,omitempty"`
	// State of MeshService. Available if there is at least one healthy endpoint. Otherwise, Unavailable.
	// It's used for cross zone communication to check if we should send traffic to it, when MeshService is aggregated into MeshMultiZoneService.
	State *State `default:"Unavailable" json:"state"`
}

func (m MeshServiceItemSpec) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshServiceItemSpec) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshServiceItemSpec) GetIdentities() []Identities {
	if o == nil {
		return nil
	}
	return o.Identities
}

func (o *MeshServiceItemSpec) GetPorts() []MeshServiceItemPorts {
	if o == nil {
		return nil
	}
	return o.Ports
}

func (o *MeshServiceItemSpec) GetSelector() *MeshServiceItemSelector {
	if o == nil {
		return nil
	}
	return o.Selector
}

func (o *MeshServiceItemSpec) GetState() *State {
	if o == nil {
		return nil
	}
	return o.State
}

type MeshServiceItemHostnameGeneratorRef struct {
	CoreName string `json:"coreName"`
}

func (o *MeshServiceItemHostnameGeneratorRef) GetCoreName() string {
	if o == nil {
		return ""
	}
	return o.CoreName
}

type MeshServiceItemAddresses struct {
	Hostname             *string                              `json:"hostname,omitempty"`
	HostnameGeneratorRef *MeshServiceItemHostnameGeneratorRef `json:"hostnameGeneratorRef,omitempty"`
	Origin               *string                              `json:"origin,omitempty"`
}

func (o *MeshServiceItemAddresses) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *MeshServiceItemAddresses) GetHostnameGeneratorRef() *MeshServiceItemHostnameGeneratorRef {
	if o == nil {
		return nil
	}
	return o.HostnameGeneratorRef
}

func (o *MeshServiceItemAddresses) GetOrigin() *string {
	if o == nil {
		return nil
	}
	return o.Origin
}

// DataplaneProxies - Data plane proxies statistics selected by this MeshService.
type DataplaneProxies struct {
	// Number of data plane proxies connected to the zone control plane
	Connected *int64 `json:"connected,omitempty"`
	// Number of data plane proxies with all healthy inbounds selected by this MeshService.
	Healthy *int64 `json:"healthy,omitempty"`
	// Total number of data plane proxies.
	Total *int64 `json:"total,omitempty"`
}

func (o *DataplaneProxies) GetConnected() *int64 {
	if o == nil {
		return nil
	}
	return o.Connected
}

func (o *DataplaneProxies) GetHealthy() *int64 {
	if o == nil {
		return nil
	}
	return o.Healthy
}

func (o *DataplaneProxies) GetTotal() *int64 {
	if o == nil {
		return nil
	}
	return o.Total
}

// MeshServiceItemStatusHostnameGeneratorsStatus - status of the condition, one of True, False, Unknown.
type MeshServiceItemStatusHostnameGeneratorsStatus string

const (
	MeshServiceItemStatusHostnameGeneratorsStatusTrue    MeshServiceItemStatusHostnameGeneratorsStatus = "True"
	MeshServiceItemStatusHostnameGeneratorsStatusFalse   MeshServiceItemStatusHostnameGeneratorsStatus = "False"
	MeshServiceItemStatusHostnameGeneratorsStatusUnknown MeshServiceItemStatusHostnameGeneratorsStatus = "Unknown"
)

func (e MeshServiceItemStatusHostnameGeneratorsStatus) ToPointer() *MeshServiceItemStatusHostnameGeneratorsStatus {
	return &e
}
func (e *MeshServiceItemStatusHostnameGeneratorsStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "True":
		fallthrough
	case "False":
		fallthrough
	case "Unknown":
		*e = MeshServiceItemStatusHostnameGeneratorsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshServiceItemStatusHostnameGeneratorsStatus: %v", v)
	}
}

type MeshServiceItemConditions struct {
	// message is a human readable message indicating details about the transition.
	// This may be an empty string.
	Message string `json:"message"`
	// reason contains a programmatic identifier indicating the reason for the condition's last transition.
	// Producers of specific condition types may define expected values and meanings for this field,
	// and whether the values are considered a guaranteed API.
	// The value should be a CamelCase string.
	// This field may not be empty.
	Reason string `json:"reason"`
	// status of the condition, one of True, False, Unknown.
	Status MeshServiceItemStatusHostnameGeneratorsStatus `json:"status"`
	// type of condition in CamelCase or in foo.example.com/CamelCase.
	Type string `json:"type"`
}

func (o *MeshServiceItemConditions) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *MeshServiceItemConditions) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

func (o *MeshServiceItemConditions) GetStatus() MeshServiceItemStatusHostnameGeneratorsStatus {
	if o == nil {
		return MeshServiceItemStatusHostnameGeneratorsStatus("")
	}
	return o.Status
}

func (o *MeshServiceItemConditions) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

type MeshServiceItemStatusHostnameGeneratorRef struct {
	CoreName string `json:"coreName"`
}

func (o *MeshServiceItemStatusHostnameGeneratorRef) GetCoreName() string {
	if o == nil {
		return ""
	}
	return o.CoreName
}

type MeshServiceItemHostnameGenerators struct {
	// Conditions is an array of hostname generator conditions.
	Conditions           []MeshServiceItemConditions               `json:"conditions,omitempty"`
	HostnameGeneratorRef MeshServiceItemStatusHostnameGeneratorRef `json:"hostnameGeneratorRef"`
}

func (o *MeshServiceItemHostnameGenerators) GetConditions() []MeshServiceItemConditions {
	if o == nil {
		return nil
	}
	return o.Conditions
}

func (o *MeshServiceItemHostnameGenerators) GetHostnameGeneratorRef() MeshServiceItemStatusHostnameGeneratorRef {
	if o == nil {
		return MeshServiceItemStatusHostnameGeneratorRef{}
	}
	return o.HostnameGeneratorRef
}

type MeshServiceItemStatusStatus string

const (
	MeshServiceItemStatusStatusReady    MeshServiceItemStatusStatus = "Ready"
	MeshServiceItemStatusStatusNotReady MeshServiceItemStatusStatus = "NotReady"
)

func (e MeshServiceItemStatusStatus) ToPointer() *MeshServiceItemStatusStatus {
	return &e
}
func (e *MeshServiceItemStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Ready":
		fallthrough
	case "NotReady":
		*e = MeshServiceItemStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MeshServiceItemStatusStatus: %v", v)
	}
}

type MeshServiceItemTLS struct {
	Status *MeshServiceItemStatusStatus `json:"status,omitempty"`
}

func (o *MeshServiceItemTLS) GetStatus() *MeshServiceItemStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type MeshServiceItemVips struct {
	IP *string `json:"ip,omitempty"`
}

func (o *MeshServiceItemVips) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

// MeshServiceItemStatus - Status is the current status of the Kuma MeshService resource.
type MeshServiceItemStatus struct {
	Addresses []MeshServiceItemAddresses `json:"addresses,omitempty"`
	// Data plane proxies statistics selected by this MeshService.
	DataplaneProxies   *DataplaneProxies                   `json:"dataplaneProxies,omitempty"`
	HostnameGenerators []MeshServiceItemHostnameGenerators `json:"hostnameGenerators,omitempty"`
	TLS                *MeshServiceItemTLS                 `json:"tls,omitempty"`
	Vips               []MeshServiceItemVips               `json:"vips,omitempty"`
}

func (o *MeshServiceItemStatus) GetAddresses() []MeshServiceItemAddresses {
	if o == nil {
		return nil
	}
	return o.Addresses
}

func (o *MeshServiceItemStatus) GetDataplaneProxies() *DataplaneProxies {
	if o == nil {
		return nil
	}
	return o.DataplaneProxies
}

func (o *MeshServiceItemStatus) GetHostnameGenerators() []MeshServiceItemHostnameGenerators {
	if o == nil {
		return nil
	}
	return o.HostnameGenerators
}

func (o *MeshServiceItemStatus) GetTLS() *MeshServiceItemTLS {
	if o == nil {
		return nil
	}
	return o.TLS
}

func (o *MeshServiceItemStatus) GetVips() []MeshServiceItemVips {
	if o == nil {
		return nil
	}
	return o.Vips
}

// MeshServiceItem - Successful response
type MeshServiceItem struct {
	// the type of the resource
	Type MeshServiceItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshService resource.
	Spec MeshServiceItemSpec `json:"spec"`
	// Time at which the resource was created
	CreationTime *time.Time `json:"creationTime,omitempty"`
	// Time at which the resource was updated
	ModificationTime *time.Time `json:"modificationTime,omitempty"`
	// Status is the current status of the Kuma MeshService resource.
	Status *MeshServiceItemStatus `json:"status,omitempty"`
}

func (m MeshServiceItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshServiceItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshServiceItem) GetType() MeshServiceItemType {
	if o == nil {
		return MeshServiceItemType("")
	}
	return o.Type
}

func (o *MeshServiceItem) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshServiceItem) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshServiceItem) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshServiceItem) GetSpec() MeshServiceItemSpec {
	if o == nil {
		return MeshServiceItemSpec{}
	}
	return o.Spec
}

func (o *MeshServiceItem) GetCreationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreationTime
}

func (o *MeshServiceItem) GetModificationTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModificationTime
}

func (o *MeshServiceItem) GetStatus() *MeshServiceItemStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type MeshServiceItemInput struct {
	// the type of the resource
	Type MeshServiceItemType `json:"type"`
	// Mesh is the name of the Kuma mesh this resource belongs to. It may be omitted for cluster-scoped resources.
	Mesh *string `default:"default" json:"mesh"`
	// Name of the Kuma resource
	Name string `json:"name"`
	// The labels to help identity resources
	Labels map[string]string `json:"labels,omitempty"`
	// Spec is the specification of the Kuma MeshService resource.
	Spec MeshServiceItemSpec `json:"spec"`
}

func (m MeshServiceItemInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MeshServiceItemInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MeshServiceItemInput) GetType() MeshServiceItemType {
	if o == nil {
		return MeshServiceItemType("")
	}
	return o.Type
}

func (o *MeshServiceItemInput) GetMesh() *string {
	if o == nil {
		return nil
	}
	return o.Mesh
}

func (o *MeshServiceItemInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MeshServiceItemInput) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}

func (o *MeshServiceItemInput) GetSpec() MeshServiceItemSpec {
	if o == nil {
		return MeshServiceItemSpec{}
	}
	return o.Spec
}
