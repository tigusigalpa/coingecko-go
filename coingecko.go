package coingecko

// CoinGecko provides access to the CoinGecko API endpoint groups.
type CoinGecko struct {
	client *Client
}

// New creates a CoinGecko API client with the supplied options.
func New(options ...ClientOption) *CoinGecko {
	return &CoinGecko{
		client: NewClient(options...),
	}
}

// Simple returns the Simple API endpoint group.
func (cg *CoinGecko) Simple() *SimpleAPI {
	return &SimpleAPI{client: cg.client}
}

// Coins returns the Coins API endpoint group.
func (cg *CoinGecko) Coins() *CoinsAPI {
	return &CoinsAPI{client: cg.client}
}

// Contract returns the contract-token API endpoint group.
func (cg *CoinGecko) Contract() *ContractAPI {
	return &ContractAPI{client: cg.client}
}

// AssetPlatforms returns the asset-platform API endpoint group.
func (cg *CoinGecko) AssetPlatforms() *AssetPlatformsAPI {
	return &AssetPlatformsAPI{client: cg.client}
}

// Categories returns the categories API endpoint group.
func (cg *CoinGecko) Categories() *CategoriesAPI {
	return &CategoriesAPI{client: cg.client}
}

// Exchanges returns the exchanges API endpoint group.
func (cg *CoinGecko) Exchanges() *ExchangesAPI {
	return &ExchangesAPI{client: cg.client}
}

// Derivatives returns the derivatives API endpoint group.
func (cg *CoinGecko) Derivatives() *DerivativesAPI {
	return &DerivativesAPI{client: cg.client}
}

// Entities returns the treasury-entities API endpoint group.
func (cg *CoinGecko) Entities() *EntitiesAPI {
	return &EntitiesAPI{client: cg.client}
}

// NFTs returns the NFT API endpoint group.
func (cg *CoinGecko) NFTs() *NFTsAPI {
	return &NFTsAPI{client: cg.client}
}

// ExchangeRates returns the exchange-rates API endpoint group.
func (cg *CoinGecko) ExchangeRates() *ExchangeRatesAPI {
	return &ExchangeRatesAPI{client: cg.client}
}

// Search returns the search API endpoint group.
func (cg *CoinGecko) Search() *SearchAPI {
	return &SearchAPI{client: cg.client}
}

// Global returns the global market-data API endpoint group.
func (cg *CoinGecko) Global() *GlobalAPI {
	return &GlobalAPI{client: cg.client}
}

// WebSocket returns the WebSocket API helper.
func (cg *CoinGecko) WebSocket() *WebSocketAPI {
	return &WebSocketAPI{client: cg.client}
}

// Ping returns the Ping API endpoint group.
func (cg *CoinGecko) Ping() *PingAPI {
	return &PingAPI{client: cg.client}
}
