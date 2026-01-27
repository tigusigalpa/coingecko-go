package coingecko

import "encoding/json"

type EntitiesAPI struct {
	client *Client
}

func (e *EntitiesAPI) List() ([]map[string]interface{}, error) {
	data, err := e.client.get("/entities/list", nil)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *EntitiesAPI) TreasuryByCoinID(coinID string) (map[string]interface{}, error) {
	data, err := e.client.get("/entities/"+coinID, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *EntitiesAPI) TreasuryByEntityID(entityID string, coinID string) (map[string]interface{}, error) {
	data, err := e.client.get("/entities/"+entityID+"/"+coinID, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *EntitiesAPI) TreasuryChart(entityID string, coinID string, days int) (map[string]interface{}, error) {
	params := map[string]string{
		"days": intToString(days),
	}

	data, err := e.client.get("/entities/"+entityID+"/"+coinID+"/chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type TransactionHistoryOptions struct {
	Page    *int
	PerPage *int
}

func (e *EntitiesAPI) TransactionHistory(entityID string, coinID string, options *TransactionHistoryOptions) (map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		if options.Page != nil {
			params["page"] = intToString(*options.Page)
		}
		if options.PerPage != nil {
			params["per_page"] = intToString(*options.PerPage)
		}
	}

	data, err := e.client.get("/entities/"+entityID+"/"+coinID+"/transactions", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
