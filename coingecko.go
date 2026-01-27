package coingecko

type CoinGecko struct {
	client *Client
}

func New(options ...ClientOption) *CoinGecko {
	return &CoinGecko{
		client: NewClient(options...),
	}
}

func (cg *CoinGecko) Simple() *SimpleAPI {
	return &SimpleAPI{client: cg.client}
}

func (cg *CoinGecko) Coins() *CoinsAPI {
	return &CoinsAPI{client: cg.client}
}

func (cg *CoinGecko) Contract() *ContractAPI {
	return &ContractAPI{client: cg.client}
}

func (cg *CoinGecko) AssetPlatforms() *AssetPlatformsAPI {
	return &AssetPlatformsAPI{client: cg.client}
}

func (cg *CoinGecko) Categories() *CategoriesAPI {
	return &CategoriesAPI{client: cg.client}
}

func (cg *CoinGecko) Exchanges() *ExchangesAPI {
	return &ExchangesAPI{client: cg.client}
}

func (cg *CoinGecko) Derivatives() *DerivativesAPI {
	return &DerivativesAPI{client: cg.client}
}

func (cg *CoinGecko) Entities() *EntitiesAPI {
	return &EntitiesAPI{client: cg.client}
}

func (cg *CoinGecko) NFTs() *NFTsAPI {
	return &NFTsAPI{client: cg.client}
}

func (cg *CoinGecko) ExchangeRates() *ExchangeRatesAPI {
	return &ExchangeRatesAPI{client: cg.client}
}

func (cg *CoinGecko) Search() *SearchAPI {
	return &SearchAPI{client: cg.client}
}

func (cg *CoinGecko) Global() *GlobalAPI {
	return &GlobalAPI{client: cg.client}
}

func (cg *CoinGecko) WebSocket() *WebSocketAPI {
	return &WebSocketAPI{client: cg.client}
}

func (cg *CoinGecko) Ping() *PingAPI {
	return &PingAPI{client: cg.client}
}
