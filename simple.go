package coingecko

import "encoding/json"

// SimpleAPI provides simple price and currency endpoints.
type SimpleAPI struct {
	client *Client
}

// SimplePriceOptions configures optional fields returned by Price.
type SimplePriceOptions struct {
	IncludeMarketCap     bool
	Include24hrVol       bool
	Include24hrChange    bool
	IncludeLastUpdatedAt bool
}

// Price returns prices for coin IDs in the requested currencies.
func (s *SimpleAPI) Price(ids []string, vsCurrencies []string, options *SimplePriceOptions) (map[string]map[string]interface{}, error) {
	params := map[string]string{
		"ids":           joinStrings(ids),
		"vs_currencies": joinStrings(vsCurrencies),
	}

	if options != nil {
		if options.IncludeMarketCap {
			params["include_market_cap"] = "true"
		}
		if options.Include24hrVol {
			params["include_24hr_vol"] = "true"
		}
		if options.Include24hrChange {
			params["include_24hr_change"] = "true"
		}
		if options.IncludeLastUpdatedAt {
			params["include_last_updated_at"] = "true"
		}
	}

	data, err := s.client.get("/simple/price", params)
	if err != nil {
		return nil, err
	}

	var result map[string]map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TokenPriceOptions configures optional fields returned by TokenPrice.
type TokenPriceOptions struct {
	IncludeMarketCap     bool
	Include24hrVol       bool
	Include24hrChange    bool
	IncludeLastUpdatedAt bool
}

// TokenPrice returns token prices by platform and contract address.
func (s *SimpleAPI) TokenPrice(assetPlatform string, contractAddresses []string, vsCurrencies []string, options *TokenPriceOptions) (map[string]map[string]interface{}, error) {
	params := map[string]string{
		"contract_addresses": joinStrings(contractAddresses),
		"vs_currencies":      joinStrings(vsCurrencies),
	}

	if options != nil {
		if options.IncludeMarketCap {
			params["include_market_cap"] = "true"
		}
		if options.Include24hrVol {
			params["include_24hr_vol"] = "true"
		}
		if options.Include24hrChange {
			params["include_24hr_change"] = "true"
		}
		if options.IncludeLastUpdatedAt {
			params["include_last_updated_at"] = "true"
		}
	}

	data, err := s.client.get("/simple/token_price/"+escapePathSegment(assetPlatform), params)
	if err != nil {
		return nil, err
	}

	var result map[string]map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SupportedVsCurrencies returns all supported quote currencies.
func (s *SimpleAPI) SupportedVsCurrencies() ([]string, error) {
	data, err := s.client.get("/simple/supported_vs_currencies", nil)
	if err != nil {
		return nil, err
	}

	var result []string
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
