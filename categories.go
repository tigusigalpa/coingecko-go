package coingecko

import "encoding/json"

// CategoriesAPI provides access to coin categories.
type CategoriesAPI struct {
	client *Client
}

// Category identifies a CoinGecko coin category.
type Category struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// List returns all coin categories.
func (c *CategoriesAPI) List() ([]Category, error) {
	data, err := c.client.get("/coins/categories/list", nil)
	if err != nil {
		return nil, err
	}

	var result []Category
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListWithMarketData returns categories with market data in the requested order.
func (c *CategoriesAPI) ListWithMarketData(order *string) ([]map[string]interface{}, error) {
	params := map[string]string{}

	if order != nil {
		params["order"] = *order
	}

	data, err := c.client.get("/coins/categories", params)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
