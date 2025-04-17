// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/hooks"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/errors"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/models/shared"
	"github.com/kong/terraform-provider-kong-mesh/internal/sdk/retry"
	"net/http"
	"time"
)

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client HTTPClient

	ServerURL         string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	return c.ServerURL, map[string]string{}
}

// KongMesh - Kong Mesh: This is a BETA Mesh specification. Endpoints in this specification may change with zero notice
type KongMesh struct {
	System                    *System
	GlobalInsight             *GlobalInsight
	Inspect                   *Inspect
	MeshAccessLog             *MeshAccessLog
	MeshCircuitBreaker        *MeshCircuitBreaker
	MeshFaultInjection        *MeshFaultInjection
	MeshHealthCheck           *MeshHealthCheck
	MeshHTTPRoute             *MeshHTTPRoute
	MeshLoadBalancingStrategy *MeshLoadBalancingStrategy
	MeshMetric                *MeshMetric
	MeshPassthrough           *MeshPassthrough
	MeshProxyPatch            *MeshProxyPatch
	MeshRateLimit             *MeshRateLimit
	MeshRetry                 *MeshRetry
	MeshTCPRoute              *MeshTCPRoute
	MeshTimeout               *MeshTimeout
	MeshTLS                   *MeshTLS
	MeshTrace                 *MeshTrace
	MeshTrafficPermission     *MeshTrafficPermission
	Mesh                      *Mesh
	MeshGateway               *MeshGateway
	HostnameGenerator         *HostnameGenerator
	MeshExternalService       *MeshExternalService
	MeshMultiZoneService      *MeshMultiZoneService
	MeshService               *MeshService
	Tenants                   *Tenants
	MeshGlobalRateLimit       *MeshGlobalRateLimit
	MeshOPA                   *MeshOPA

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*KongMesh)

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *KongMesh) {
		sdk.sdkConfiguration.Client = client
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *KongMesh) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *KongMesh) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided serverURL and options
func New(serverURL string, opts ...SDKOption) *KongMesh {
	sdk := &KongMesh{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "2.0.0",
			SDKVersion:        "0.1.1",
			GenVersion:        "2.546.3",
			UserAgent:         "speakeasy-sdk/terraform 0.1.1 2.546.3 2.0.0 github.com/kong/terraform-provider-kong-mesh/internal/sdk",
			ServerURL:         serverURL,
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL := serverURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.System = newSystem(sdk.sdkConfiguration)

	sdk.GlobalInsight = newGlobalInsight(sdk.sdkConfiguration)

	sdk.Inspect = newInspect(sdk.sdkConfiguration)

	sdk.MeshAccessLog = newMeshAccessLog(sdk.sdkConfiguration)

	sdk.MeshCircuitBreaker = newMeshCircuitBreaker(sdk.sdkConfiguration)

	sdk.MeshFaultInjection = newMeshFaultInjection(sdk.sdkConfiguration)

	sdk.MeshHealthCheck = newMeshHealthCheck(sdk.sdkConfiguration)

	sdk.MeshHTTPRoute = newMeshHTTPRoute(sdk.sdkConfiguration)

	sdk.MeshLoadBalancingStrategy = newMeshLoadBalancingStrategy(sdk.sdkConfiguration)

	sdk.MeshMetric = newMeshMetric(sdk.sdkConfiguration)

	sdk.MeshPassthrough = newMeshPassthrough(sdk.sdkConfiguration)

	sdk.MeshProxyPatch = newMeshProxyPatch(sdk.sdkConfiguration)

	sdk.MeshRateLimit = newMeshRateLimit(sdk.sdkConfiguration)

	sdk.MeshRetry = newMeshRetry(sdk.sdkConfiguration)

	sdk.MeshTCPRoute = newMeshTCPRoute(sdk.sdkConfiguration)

	sdk.MeshTimeout = newMeshTimeout(sdk.sdkConfiguration)

	sdk.MeshTLS = newMeshTLS(sdk.sdkConfiguration)

	sdk.MeshTrace = newMeshTrace(sdk.sdkConfiguration)

	sdk.MeshTrafficPermission = newMeshTrafficPermission(sdk.sdkConfiguration)

	sdk.Mesh = newMesh(sdk.sdkConfiguration)

	sdk.MeshGateway = newMeshGateway(sdk.sdkConfiguration)

	sdk.HostnameGenerator = newHostnameGenerator(sdk.sdkConfiguration)

	sdk.MeshExternalService = newMeshExternalService(sdk.sdkConfiguration)

	sdk.MeshMultiZoneService = newMeshMultiZoneService(sdk.sdkConfiguration)

	sdk.MeshService = newMeshService(sdk.sdkConfiguration)

	sdk.Tenants = newTenants(sdk.sdkConfiguration)

	sdk.MeshGlobalRateLimit = newMeshGlobalRateLimit(sdk.sdkConfiguration)

	sdk.MeshOPA = newMeshOPA(sdk.sdkConfiguration)

	return sdk
}

// GetDataplanesXdsConfig - Get a proxy XDS config on a CP, this endpoint is only available on zone CPs.
// Returns the [xds](https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol) configuration of the proxy.
func (s *KongMesh) GetDataplanesXdsConfig(ctx context.Context, request operations.GetDataplanesXdsConfigRequest, opts ...operations.Option) (*operations.GetDataplanesXdsConfigResponse, error) {
	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionTimeout,
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}

	var baseURL string
	if o.ServerURL == nil {
		baseURL = utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	} else {
		baseURL = *o.ServerURL
	}
	opURL, err := utils.GenerateURL(ctx, baseURL, "/meshes/{mesh}/dataplanes/{name}/_config", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	hookCtx := hooks.HookContext{
		BaseURL:        baseURL,
		Context:        ctx,
		OperationID:    "get-dataplanes-xds-config",
		OAuth2Scopes:   []string{},
		SecuritySource: nil,
	}

	timeout := o.Timeout
	if timeout == nil {
		timeout = s.sdkConfiguration.Timeout
	}

	if timeout != nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, *timeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/problem+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	for k, v := range o.SetHeaders {
		req.Header.Set(k, v)
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.GetDataplanesXdsConfigResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out shared.DataplaneXDSConfig
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.DataplaneXDSConfig = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out shared.BadRequestError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.BadRequestError = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/problem+json`):
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}

			var out shared.BaseError
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.BaseError = &out
		default:
			rawBody, err := utils.ConsumeRawBody(httpRes)
			if err != nil {
				return nil, err
			}
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	default:
		rawBody, err := utils.ConsumeRawBody(httpRes)
		if err != nil {
			return nil, err
		}
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
