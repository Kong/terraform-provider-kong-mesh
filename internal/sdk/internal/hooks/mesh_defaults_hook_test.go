package hooks

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestCheckMatchingFeatures(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		matchingFn     func(*http.Request) bool
		expectedOutput bool
	}

	testCases := []testCase{
		{
			name:           "onprem does not have features",
			input:          "/",
			matchingFn:     onPremMatchFeatures,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			matched := tc.matchingFn(&http.Request{
				Method: http.MethodPost,
				URL: &url.URL{
					Path: tc.input,
				},
			})

			require.Equal(t, tc.expectedOutput, matched)
		})
	}
}

func TestCheckMatchingPolicies(t *testing.T) {
	type testCase struct {
		name           string
		input          string
		matchingFn     func(*http.Request) bool
		expectedOutput bool
	}

	testCases := []testCase{
		{
			name:           "on prem matches",
			input:          "/meshes/default",
			matchingFn:     onPremMatchPolicies,
			expectedOutput: true,
		},
		{
			name:           "on prem does not match",
			input:          "/meshtrafficpermissions/default",
			matchingFn:     onPremMatchPolicies,
			expectedOutput: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			matched := tc.matchingFn(&http.Request{
				Method: http.MethodPut,
				URL: &url.URL{
					Path: tc.input,
				},
			})

			require.Equal(t, tc.expectedOutput, matched)
		})
	}
}
