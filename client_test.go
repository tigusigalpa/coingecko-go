package coingecko

import (
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

func TestClientWithTimeout(t *testing.T) {
	timeout := 60 * time.Second
	client := NewClient(WithTimeout(timeout))

	if client.httpClient.Timeout != timeout {
		t.Errorf("Expected timeout to be %v, got %v", timeout, client.httpClient.Timeout)
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
