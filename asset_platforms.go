package coingecko

import "encoding/json"

type AssetPlatformsAPI struct {
	client *Client
}

type AssetPlatform struct {
	ID                string  `json:"id"`
	ChainIdentifier   *int    `json:"chain_identifier,omitempty"`
	Name              string  `json:"name"`
	Shortname         string  `json:"shortname"`
	NativeCoinID      string  `json:"native_coin_id,omitempty"`
	ImageURL          *string `json:"image,omitempty"`
}

func (a *AssetPlatformsAPI) List(filter *string) ([]AssetPlatform, error) {
	params := map[string]string{}
	
	if filter != nil {
		params["filter"] = *filter
	}

	data, err := a.client.get("/asset_platforms", params)
	if err != nil {
		return nil, err
	}

	var result []AssetPlatform
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AssetPlatformsAPI) TokenLists(assetPlatformID string, page *int) (map[string]interface{}, error) {
	params := map[string]string{}
	
	if page != nil {
		params["page"] = intToString(*page)
	}

	data, err := a.client.get("/token_lists/"+assetPlatformID+"/all.json", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
