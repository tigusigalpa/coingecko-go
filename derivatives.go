package coingecko

import "encoding/json"

type DerivativesAPI struct {
	client *Client
}

func (d *DerivativesAPI) Tickers(includeTickers *string) ([]map[string]interface{}, error) {
	params := map[string]string{}
	
	if includeTickers != nil {
		params["include_tickers"] = *includeTickers
	}

	data, err := d.client.get("/derivatives", params)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type DerivativesExchangesOptions struct {
	Order   *string
	PerPage *int
	Page    *int
}

func (d *DerivativesAPI) Exchanges(options *DerivativesExchangesOptions) ([]map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		if options.Order != nil {
			params["order"] = *options.Order
		}
		if options.PerPage != nil {
			params["per_page"] = intToString(*options.PerPage)
		}
		if options.Page != nil {
			params["page"] = intToString(*options.Page)
		}
	}

	data, err := d.client.get("/derivatives/exchanges", params)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *DerivativesAPI) Exchange(id string, includeTickers *string) (map[string]interface{}, error) {
	params := map[string]string{}
	
	if includeTickers != nil {
		params["include_tickers"] = *includeTickers
	}

	data, err := d.client.get("/derivatives/exchanges/"+id, params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (d *DerivativesAPI) ExchangesList() ([]map[string]interface{}, error) {
	data, err := d.client.get("/derivatives/exchanges/list", nil)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
