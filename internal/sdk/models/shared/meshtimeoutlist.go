// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// MeshTimeoutList - List
type MeshTimeoutList struct {
	Items []MeshTimeoutItem `json:"items,omitempty"`
	// The total number of entities
	Total *float64 `json:"total,omitempty"`
	// URL to the next page
	Next *string `json:"next,omitempty"`
}

func (o *MeshTimeoutList) GetItems() []MeshTimeoutItem {
	if o == nil {
		return nil
	}
	return o.Items
}

func (o *MeshTimeoutList) GetTotal() *float64 {
	if o == nil {
		return nil
	}
	return o.Total
}

func (o *MeshTimeoutList) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}
