package coingecko

import "encoding/json"

// GlobalAPI provides global cryptocurrency market data.
type GlobalAPI struct {
	client *Client
}

// Crypto returns global cryptocurrency market data.
func (g *GlobalAPI) Crypto() (map[string]interface{}, error) {
	data, err := g.client.get("/global", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// DeFi returns global decentralized-finance market data.
func (g *GlobalAPI) DeFi() (map[string]interface{}, error) {
	data, err := g.client.get("/global/decentralized_finance_defi", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// MarketCapChart returns global market-cap history for a currency.
func (g *GlobalAPI) MarketCapChart(days int, vsCurrency string) (map[string]interface{}, error) {
	params := map[string]string{
		"days":        intToString(days),
		"vs_currency": vsCurrency,
	}

	data, err := g.client.get("/global/market_cap_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
