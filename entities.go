package coingecko

import "encoding/json"

// EntitiesAPI provides cryptocurrency treasury-entity data.
type EntitiesAPI struct {
	client *Client
}

// List returns all treasury entities.
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

// TreasuryByCoinID returns treasury data for a coin ID.
func (e *EntitiesAPI) TreasuryByCoinID(coinID string) (map[string]interface{}, error) {
	data, err := e.client.get("/entities/"+escapePathSegment(coinID), nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TreasuryByEntityID returns an entity's treasury data for a coin ID.
func (e *EntitiesAPI) TreasuryByEntityID(entityID string, coinID string) (map[string]interface{}, error) {
	data, err := e.client.get("/entities/"+escapePathSegment(entityID)+"/"+escapePathSegment(coinID), nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TreasuryChart returns historical treasury data for an entity and coin.
func (e *EntitiesAPI) TreasuryChart(entityID string, coinID string, days int) (map[string]interface{}, error) {
	params := map[string]string{
		"days": intToString(days),
	}

	data, err := e.client.get("/entities/"+escapePathSegment(entityID)+"/"+escapePathSegment(coinID)+"/chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TransactionHistoryOptions configures pagination for TransactionHistory.
type TransactionHistoryOptions struct {
	Page    *int
	PerPage *int
}

// TransactionHistory returns an entity's treasury transaction history.
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

	data, err := e.client.get("/entities/"+escapePathSegment(entityID)+"/"+escapePathSegment(coinID)+"/transactions", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
