// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ProvisionMeshControlPlaneRequest struct {
	Features []MeshControlPlaneFeature `json:"features,omitempty"`
}

func (o *ProvisionMeshControlPlaneRequest) GetFeatures() []MeshControlPlaneFeature {
	if o == nil {
		return nil
	}
	return o.Features
}
