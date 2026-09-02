package coingecko

import (
	"encoding/json"
	"fmt"
)

// ExchangesAPI provides cryptocurrency exchange data.
type ExchangesAPI struct {
	client *Client
}

// ExchangesListOptions configures pagination for List.
type ExchangesListOptions struct {
	PerPage *int
	Page    *int
}

// List returns cryptocurrency exchanges.
func (e *ExchangesAPI) List(options *ExchangesListOptions) ([]map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		if options.PerPage != nil {
			params["per_page"] = intToString(*options.PerPage)
		}
		if options.Page != nil {
			params["page"] = intToString(*options.Page)
		}
	}

	data, err := e.client.get("/exchanges", params)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListIDMap returns exchange identifiers.
func (e *ExchangesAPI) ListIDMap() ([]map[string]interface{}, error) {
	data, err := e.client.get("/exchanges/list", nil)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Exchange returns data for an exchange ID.
func (e *ExchangesAPI) Exchange(id string) (map[string]interface{}, error) {
	data, err := e.client.get("/exchanges/"+escapePathSegment(id), nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ExchangeTickersOptions configures filtering for Tickers.
type ExchangeTickersOptions struct {
	CoinIDs             []string
	IncludeExchangeLogo bool
	Page                *int
	Depth               *string
	Order               *string
}

// Tickers returns ticker data for an exchange ID.
func (e *ExchangesAPI) Tickers(id string, options *ExchangeTickersOptions) (map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		if len(options.CoinIDs) > 0 {
			params["coin_ids"] = joinStrings(options.CoinIDs)
		}
		if options.IncludeExchangeLogo {
			params["include_exchange_logo"] = "true"
		}
		if options.Page != nil {
			params["page"] = intToString(*options.Page)
		}
		if options.Depth != nil {
			params["depth"] = *options.Depth
		}
		if options.Order != nil {
			params["order"] = *options.Order
		}
	}

	data, err := e.client.get("/exchanges/"+escapePathSegment(id)+"/tickers", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// VolumeChart returns an exchange's volume history for a period.
func (e *ExchangesAPI) VolumeChart(id string, days int) ([][]interface{}, error) {
	params := map[string]string{
		"days": intToString(days),
	}

	data, err := e.client.get("/exchanges/"+escapePathSegment(id)+"/volume_chart", params)
	if err != nil {
		return nil, err
	}

	var result [][]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// VolumeChartRange returns an exchange's volume history between Unix timestamps.
func (e *ExchangesAPI) VolumeChartRange(id string, from int64, to int64) ([][]interface{}, error) {
	params := map[string]string{
		"from": fmt.Sprintf("%d", from),
		"to":   fmt.Sprintf("%d", to),
	}

	data, err := e.client.get("/exchanges/"+escapePathSegment(id)+"/volume_chart/range", params)
	if err != nil {
		return nil, err
	}

	var result [][]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
