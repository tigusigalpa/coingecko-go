package coingecko

import "encoding/json"

// NFTsAPI provides NFT collection and market data.
type NFTsAPI struct {
	client *Client
}

// NFTsListOptions configures filtering and pagination for List.
type NFTsListOptions struct {
	Order           *string
	AssetPlatformID *string
	PerPage         *int
	Page            *int
}

// List returns NFT collections.
func (n *NFTsAPI) List(options *NFTsListOptions) ([]map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		if options.Order != nil {
			params["order"] = *options.Order
		}
		if options.AssetPlatformID != nil {
			params["asset_platform_id"] = *options.AssetPlatformID
		}
		if options.PerPage != nil {
			params["per_page"] = intToString(*options.PerPage)
		}
		if options.Page != nil {
			params["page"] = intToString(*options.Page)
		}
	}

	data, err := n.client.get("/nfts/list", params)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Collection returns an NFT collection by ID.
func (n *NFTsAPI) Collection(id string) (map[string]interface{}, error) {
	data, err := n.client.get("/nfts/"+escapePathSegment(id), nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CollectionByContract returns an NFT collection by platform and contract address.
func (n *NFTsAPI) CollectionByContract(assetPlatformID string, contractAddress string) (map[string]interface{}, error) {
	data, err := n.client.get("/nfts/"+escapePathSegment(assetPlatformID)+"/contract/"+escapePathSegment(contractAddress), nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// NFTsMarketsOptions configures filtering and pagination for Markets.
type NFTsMarketsOptions struct {
	AssetPlatformID *string
	Order           *string
	PerPage         *int
	Page            *int
}

// Markets returns NFT market data.
func (n *NFTsAPI) Markets(options *NFTsMarketsOptions) ([]map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		if options.AssetPlatformID != nil {
			params["asset_platform_id"] = *options.AssetPlatformID
		}
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

	data, err := n.client.get("/nfts/markets", params)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// MarketChart returns market history for an NFT collection.
func (n *NFTsAPI) MarketChart(id string, days int) (map[string]interface{}, error) {
	params := map[string]string{
		"days": intToString(days),
	}

	data, err := n.client.get("/nfts/"+escapePathSegment(id)+"/market_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// MarketChartByContract returns NFT market history by contract address.
func (n *NFTsAPI) MarketChartByContract(assetPlatformID string, contractAddress string, days int) (map[string]interface{}, error) {
	params := map[string]string{
		"days": intToString(days),
	}

	data, err := n.client.get("/nfts/"+escapePathSegment(assetPlatformID)+"/contract/"+escapePathSegment(contractAddress)+"/market_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// Tickers returns ticker data for an NFT collection.
func (n *NFTsAPI) Tickers(id string) (map[string]interface{}, error) {
	data, err := n.client.get("/nfts/"+escapePathSegment(id)+"/tickers", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
