package coingecko

import (
	"encoding/json"
	"fmt"
)

type ContractAPI struct {
	client *Client
}

func (c *ContractAPI) Coin(assetPlatform string, contractAddress string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/coins/%s/contract/%s", assetPlatform, contractAddress)
	
	data, err := c.client.get(endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ContractAPI) MarketChart(assetPlatform string, contractAddress string, vsCurrency string, days string) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/coins/%s/contract/%s/market_chart", assetPlatform, contractAddress)
	
	params := map[string]string{
		"vs_currency": vsCurrency,
		"days":        days,
	}

	data, err := c.client.get(endpoint, params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ContractAPI) MarketChartRange(assetPlatform string, contractAddress string, vsCurrency string, from int64, to int64) (map[string]interface{}, error) {
	endpoint := fmt.Sprintf("/coins/%s/contract/%s/market_chart/range", assetPlatform, contractAddress)
	
	params := map[string]string{
		"vs_currency": vsCurrency,
		"from":        fmt.Sprintf("%d", from),
		"to":          fmt.Sprintf("%d", to),
	}

	data, err := c.client.get(endpoint, params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
