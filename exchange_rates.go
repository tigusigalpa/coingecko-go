package coingecko

import "encoding/json"

type ExchangeRatesAPI struct {
	client *Client
}

func (e *ExchangeRatesAPI) Rates() (map[string]interface{}, error) {
	data, err := e.client.get("/exchange_rates", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
