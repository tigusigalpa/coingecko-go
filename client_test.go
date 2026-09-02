package coingecko

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()

	if client == nil {
		t.Fatal("Expected client to be created")
	}

	if client.baseURL != BaseURLFree {
		t.Errorf("Expected baseURL to be %s, got %s", BaseURLFree, client.baseURL)
	}

	if client.httpClient == nil {
		t.Fatal("Expected httpClient to be initialized")
	}
}

func TestClientWithAPIKey(t *testing.T) {
	apiKey := "test-api-key"
	client := NewClient(WithAPIKey(apiKey))

	if client.apiKey != apiKey {
		t.Errorf("Expected apiKey to be %s, got %s", apiKey, client.apiKey)
	}
}

func TestClientWithProAPI(t *testing.T) {
	client := NewClient(WithProAPI(true))

	if !client.isPro {
		t.Error("Expected isPro to be true")
	}

	if client.baseURL != BaseURLPro {
		t.Errorf("Expected baseURL to be %s, got %s", BaseURLPro, client.baseURL)
	}
}

func TestClientCanReturnToDemoAPI(t *testing.T) {
	client := NewClient(WithProAPI(true), WithProAPI(false))

	if client.isPro {
		t.Error("Expected isPro to be false")
	}
	if client.baseURL != BaseURLFree {
		t.Errorf("Expected baseURL to be %s, got %s", BaseURLFree, client.baseURL)
	}
}

func TestClientWithTimeout(t *testing.T) {
	timeout := 60 * time.Second
	client := NewClient(WithTimeout(timeout))

	if client.httpClient.Timeout != timeout {
		t.Errorf("Expected timeout to be %v, got %v", timeout, client.httpClient.Timeout)
	}
}

func TestNilClientOptionsAreSafe(t *testing.T) {
	client := NewClient(nil, WithHTTPClient(nil), WithTimeout(time.Second))

	if client.httpClient == nil {
		t.Fatal("Expected a default HTTP client")
	}
	if client.httpClient.Timeout != time.Second {
		t.Errorf("Expected timeout to be %v, got %v", time.Second, client.httpClient.Timeout)
	}
}

func TestRequestAuthenticationHeader(t *testing.T) {
	tests := []struct {
		name           string
		options        []ClientOption
		expectedHeader string
	}{
		{
			name:           "demo API",
			options:        []ClientOption{WithAPIKey("demo-key")},
			expectedHeader: "x-cg-demo-api-key",
		},
		{
			name:           "pro API",
			options:        []ClientOption{WithAPIKey("pro-key"), WithProAPI(true)},
			expectedHeader: "x-cg-pro-api-key",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if got := r.Header.Get(test.expectedHeader); got == "" {
					t.Fatalf("Expected %s header", test.expectedHeader)
				}
				w.WriteHeader(http.StatusNoContent)
			}))
			defer server.Close()

			client := NewClient(test.options...)
			client.baseURL = server.URL
			if _, err := client.get("/ping", nil); err != nil {
				t.Fatalf("Request failed: %v", err)
			}
		})
	}
}

func TestRequestEscapesPathSegmentsAndQueryParameters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.URL.EscapedPath(), "/coins/bitcoin%2Ftest"; got != want {
			t.Errorf("Expected escaped path %q, got %q", want, got)
		}
		if got, want := r.URL.Query().Get("query"), "bitcoin & ethereum"; got != want {
			t.Errorf("Expected query %q, got %q", want, got)
		}
		_, _ = w.Write([]byte("{}"))
	}))
	defer server.Close()

	client := NewClient()
	client.baseURL = server.URL
	if _, err := client.get("/coins/"+escapePathSegment("bitcoin/test"), map[string]string{"query": "bitcoin & ethereum"}); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
}

func TestRequestReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "rate limited", http.StatusTooManyRequests)
	}))
	defer server.Close()

	client := NewClient()
	client.baseURL = server.URL
	_, err := client.get("/ping", nil)
	if err == nil {
		t.Fatal("Expected an error")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("Expected APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Expected status %d, got %d", http.StatusTooManyRequests, apiErr.StatusCode)
	}
}

func TestJoinStrings(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"bitcoin", "ethereum"}, "bitcoin,ethereum"},
		{[]string{"usd"}, "usd"},
		{[]string{}, ""},
	}

	for _, test := range tests {
		result := joinStrings(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

func TestBoolToString(t *testing.T) {
	if boolToString(true) != "true" {
		t.Error("Expected 'true'")
	}

	if boolToString(false) != "false" {
		t.Error("Expected 'false'")
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{100, "100"},
		{0, "0"},
		{-1, "-1"},
	}

	for _, test := range tests {
		result := intToString(test.input)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}
