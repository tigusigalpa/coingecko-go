package coingecko

import "encoding/json"

// PingAPI provides API status and usage endpoints.
type PingAPI struct {
	client *Client
}

// PingResponse is the response returned by Ping.
type PingResponse struct {
	GeckoSays string `json:"gecko_says"`
}

// Ping checks the availability of the CoinGecko API.
func (p *PingAPI) Ping() (*PingResponse, error) {
	data, err := p.client.get("/ping", nil)
	if err != nil {
		return nil, err
	}

	var result PingResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// APIUsage returns usage information for the configured API key.
func (p *PingAPI) APIUsage() (map[string]interface{}, error) {
	data, err := p.client.get("/key", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
