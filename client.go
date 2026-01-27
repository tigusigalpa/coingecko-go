package coingecko

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BaseURLFree = "https://api.coingecko.com/api/v3"
	BaseURLPro  = "https://pro-api.coingecko.com/api/v3"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	isPro      bool
}

type ClientOption func(*Client)

func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithProAPI(isPro bool) ClientOption {
	return func(c *Client) {
		c.isPro = isPro
		if isPro {
			c.baseURL = BaseURLPro
		}
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: BaseURLFree,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func (c *Client) request(method, endpoint string, params map[string]string) ([]byte, error) {
	fullURL := c.baseURL + endpoint

	if len(params) > 0 && method == http.MethodGet {
		queryParams := url.Values{}
		for key, value := range params {
			queryParams.Add(key, value)
		}
		fullURL += "?" + queryParams.Encode()
	}

	req, err := http.NewRequest(method, fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	if c.apiKey != "" {
		req.Header.Set("x-cg-pro-api-key", c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (c *Client) get(endpoint string, params map[string]string) ([]byte, error) {
	return c.request(http.MethodGet, endpoint, params)
}

func (c *Client) post(endpoint string, params map[string]string) ([]byte, error) {
	return c.request(http.MethodPost, endpoint, params)
}

func (c *Client) unmarshalResponse(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return nil
}

func joinStrings(items []string) string {
	return strings.Join(items, ",")
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func intToString(i int) string {
	return fmt.Sprintf("%d", i)
}
