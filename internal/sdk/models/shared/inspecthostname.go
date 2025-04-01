// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// InspectHostname - An supported hostname along with the zones it exists in
type InspectHostname struct {
	// Generated hostname
	Hostname string                `json:"hostname"`
	Zones    []InspectHostnameZone `json:"zones"`
}

func (o *InspectHostname) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *InspectHostname) GetZones() []InspectHostnameZone {
	if o == nil {
		return []InspectHostnameZone{}
	}
	return o.Zones
}
