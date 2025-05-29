package hooks

import (
	shared_speakeasy "github.com/Kong/shared-speakeasy/hooks/mesh_defaults"
	"net/http"
)

// MeshDefaultsHook is a struct that implements the BeforeRequestHook interface.
type MeshDefaultsHook struct{}

// BeforeRequest modifies the request before sending it.
func (e MeshDefaultsHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	apiPrefix := ""
	cpFeatures := false
	initialPolicies := true
	return shared_speakeasy.BeforeRequest(apiPrefix, cpFeatures, initialPolicies)(req)
}
