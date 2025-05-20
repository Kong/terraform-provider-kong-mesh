package hooks

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// OverrideProtectionHook guards against overriding a resource that was not imported to the state.
type OverrideProtectionHook struct {
	Enabled bool
}

// BeforeRequest modifies the request before sending it.
func (e OverrideProtectionHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if !e.Enabled {
		return req, nil
	}

	if req.Method != http.MethodPut {
		return req, nil
	}

	if !strings.HasPrefix(req.URL.Path, "/meshes/") {
		return req, nil
	}

	// Split path and ensure it has enough parts
	// Example: /meshes/{mesh}/meshtrafficpermissions/{name}
	segments := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	if len(segments) < 4 {
		return req, nil
	}

	mesh := segments[1]
	resourceType := segments[2]
	name := segments[3]

	// Construct GET URL to check for existence
	getURL := *req.URL
	getURL.Path = fmt.Sprintf("/meshes/%s/%s/%s", mesh, resourceType, name)
	getReq, err := http.NewRequest(http.MethodGet, getURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create GET request: %w", err)
	}

	resp, err := http.DefaultClient.Do(getReq)
	if err != nil {
		return nil, fmt.Errorf("override protection GET request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil, errors.New("resource already exists, use terraform import to import it to the state")
	}

	return req, nil
}
