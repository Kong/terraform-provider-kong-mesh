package hooks

import (
	shared_speakeasy "github.com/Kong/shared-speakeasy/hooks/mesh_defaults"
	"net/http"
	"regexp"
)

var onPremMatchFeatures = func(req *http.Request) bool {
	return false
}

var onPremMatchPolicies = func(req *http.Request) bool {
	match, err := regexp.MatchString(`^/meshes/[^/]+$`, req.URL.Path)
	if err != nil {
		return false
	}

	return match && req.Method == http.MethodPut
}

// MeshDefaultsHook is a struct that implements the BeforeRequestHook interface.
type MeshDefaultsHook struct{}

// BeforeRequest modifies the request before sending it.
func (e MeshDefaultsHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	return shared_speakeasy.BeforeRequest(onPremMatchFeatures, onPremMatchPolicies)(req)
}
