package coingecko

import "encoding/json"

type PingAPI struct {
	client *Client
}

type PingResponse struct {
	GeckoSays string `json:"gecko_says"`
}

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
