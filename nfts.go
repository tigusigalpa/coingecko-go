package coingecko

import "encoding/json"

type NFTsAPI struct {
	client *Client
}

type NFTsListOptions struct {
	Order           *string
	AssetPlatformID *string
	PerPage         *int
	Page            *int
}

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

func (n *NFTsAPI) Collection(id string) (map[string]interface{}, error) {
	data, err := n.client.get("/nfts/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (n *NFTsAPI) CollectionByContract(assetPlatformID string, contractAddress string) (map[string]interface{}, error) {
	data, err := n.client.get("/nfts/"+assetPlatformID+"/contract/"+contractAddress, nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type NFTsMarketsOptions struct {
	AssetPlatformID *string
	Order           *string
	PerPage         *int
	Page            *int
}

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

func (n *NFTsAPI) MarketChart(id string, days int) (map[string]interface{}, error) {
	params := map[string]string{
		"days": intToString(days),
	}

	data, err := n.client.get("/nfts/"+id+"/market_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (n *NFTsAPI) MarketChartByContract(assetPlatformID string, contractAddress string, days int) (map[string]interface{}, error) {
	params := map[string]string{
		"days": intToString(days),
	}

	data, err := n.client.get("/nfts/"+assetPlatformID+"/contract/"+contractAddress+"/market_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (n *NFTsAPI) Tickers(id string) (map[string]interface{}, error) {
	data, err := n.client.get("/nfts/"+id+"/tickers", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
