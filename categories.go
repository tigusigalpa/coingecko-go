package coingecko

import "encoding/json"

type CategoriesAPI struct {
	client *Client
}

type Category struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

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
