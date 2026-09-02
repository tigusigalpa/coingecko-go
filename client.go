package coingecko

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Base URLs for the CoinGecko Demo and Pro APIs.
const (
	BaseURLFree = "https://api.coingecko.com/api/v3"
	BaseURLPro  = "https://pro-api.coingecko.com/api/v3"
)

// Client performs HTTP requests to the CoinGecko API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
	isPro      bool
}

// ClientOption configures a Client during construction.
type ClientOption func(*Client)

// APIError is returned when the CoinGecko API responds with a non-success status.
// Body contains the response body supplied by the API, when available.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("CoinGecko API request failed with status %d: %s", e.StatusCode, e.Body)
}

// WithAPIKey configures the API key used to authenticate requests.
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// WithProAPI selects the Pro API when isPro is true, or the Demo API otherwise.
func WithProAPI(isPro bool) ClientOption {
	return func(c *Client) {
		c.isPro = isPro
		c.baseURL = BaseURLFree
		if isPro {
			c.baseURL = BaseURLPro
		}
	}
}

// WithHTTPClient uses httpClient for requests when it is non-nil.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		if httpClient != nil {
			c.httpClient = httpClient
		}
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		if c.httpClient == nil {
			c.httpClient = &http.Client{}
		}
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a Client with the supplied options.
func NewClient(options ...ClientOption) *Client {
	client := &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: BaseURLFree,
	}

	for _, option := range options {
		if option != nil {
			option(client)
		}
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
		if c.isPro {
			req.Header.Set("x-cg-pro-api-key", c.apiKey)
		} else {
			req.Header.Set("x-cg-demo-api-key", c.apiKey)
		}
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
		return nil, &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	return body, nil
}

func (c *Client) get(endpoint string, params map[string]string) ([]byte, error) {
	return c.request(http.MethodGet, endpoint, params)
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

func escapePathSegment(value string) string {
	return url.PathEscape(value)
}
