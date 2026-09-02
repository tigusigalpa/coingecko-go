package coingecko

import "encoding/json"

// AssetPlatformsAPI provides access to blockchain platform metadata.
type AssetPlatformsAPI struct {
	client *Client
}

// AssetPlatform describes a blockchain platform supported by CoinGecko.
type AssetPlatform struct {
	ID              string  `json:"id"`
	ChainIdentifier *int    `json:"chain_identifier,omitempty"`
	Name            string  `json:"name"`
	Shortname       string  `json:"shortname"`
	NativeCoinID    string  `json:"native_coin_id,omitempty"`
	ImageURL        *string `json:"image,omitempty"`
}

// List returns supported asset platforms, optionally filtered by type.
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

// TokenLists returns token-list data for an asset platform.
func (a *AssetPlatformsAPI) TokenLists(assetPlatformID string, page *int) (map[string]interface{}, error) {
	params := map[string]string{}

	if page != nil {
		params["page"] = intToString(*page)
	}

	data, err := a.client.get("/token_lists/"+escapePathSegment(assetPlatformID)+"/all.json", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
