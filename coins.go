package coingecko

import (
	"encoding/json"
	"fmt"
)

// CoinsAPI provides coin market data, metadata, and historical charts.
type CoinsAPI struct {
	client *Client
}

// CoinListItem identifies a coin returned by the list endpoint.
type CoinListItem struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms,omitempty"`
}

// List returns all supported coins and, optionally, their platform mappings.
func (c *CoinsAPI) List(includePlatform bool) ([]CoinListItem, error) {
	params := map[string]string{}
	if includePlatform {
		params["include_platform"] = "true"
	}

	data, err := c.client.get("/coins/list", params)
	if err != nil {
		return nil, err
	}

	var result []CoinListItem
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TopGainersLosers returns the top gaining and losing coins for a currency.
func (c *CoinsAPI) TopGainersLosers(vsCurrency string, duration *string, topCoins *int) (map[string]interface{}, error) {
	params := map[string]string{
		"vs_currency": vsCurrency,
	}

	if duration != nil {
		params["duration"] = *duration
	}
	if topCoins != nil {
		params["top_coins"] = intToString(*topCoins)
	}

	data, err := c.client.get("/coins/top_gainers_losers", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// RecentlyAdded returns coins that were recently added to CoinGecko.
func (c *CoinsAPI) RecentlyAdded() ([]map[string]interface{}, error) {
	data, err := c.client.get("/coins/list/new", nil)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// MarketsOptions configures filtering and pagination for Markets.
type MarketsOptions struct {
	IDs                   []string
	Category              *string
	Order                 *string
	PerPage               *int
	Page                  *int
	Sparkline             bool
	PriceChangePercentage *string
	Locale                *string
	Precision             *string
}

// Markets returns market data for coins in the requested currency.
func (c *CoinsAPI) Markets(vsCurrency string, options *MarketsOptions) ([]map[string]interface{}, error) {
	params := map[string]string{
		"vs_currency": vsCurrency,
	}

	if options != nil {
		if len(options.IDs) > 0 {
			params["ids"] = joinStrings(options.IDs)
		}
		if options.Category != nil {
			params["category"] = *options.Category
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
		if options.Sparkline {
			params["sparkline"] = "true"
		}
		if options.PriceChangePercentage != nil {
			params["price_change_percentage"] = *options.PriceChangePercentage
		}
		if options.Locale != nil {
			params["locale"] = *options.Locale
		}
		if options.Precision != nil {
			params["precision"] = *options.Precision
		}
	}

	data, err := c.client.get("/coins/markets", params)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CoinOptions controls optional sections returned by Coin.
type CoinOptions struct {
	Localization  bool
	Tickers       bool
	MarketData    bool
	CommunityData bool
	DeveloperData bool
	Sparkline     bool
}

// Coin returns detailed data for a coin ID.
func (c *CoinsAPI) Coin(id string, options *CoinOptions) (map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		params["localization"] = boolToString(options.Localization)
		params["tickers"] = boolToString(options.Tickers)
		params["market_data"] = boolToString(options.MarketData)
		params["community_data"] = boolToString(options.CommunityData)
		params["developer_data"] = boolToString(options.DeveloperData)
		params["sparkline"] = boolToString(options.Sparkline)
	} else {
		params["localization"] = "true"
		params["tickers"] = "true"
		params["market_data"] = "true"
		params["community_data"] = "true"
		params["developer_data"] = "true"
		params["sparkline"] = "false"
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id), params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TickersOptions configures filtering for Tickers.
type TickersOptions struct {
	ExchangeIDs         []string
	IncludeExchangeLogo bool
	Page                *int
	Order               *string
	Depth               *string
}

// Tickers returns ticker data for a coin ID.
func (c *CoinsAPI) Tickers(id string, options *TickersOptions) (map[string]interface{}, error) {
	params := map[string]string{}

	if options != nil {
		if len(options.ExchangeIDs) > 0 {
			params["exchange_ids"] = joinStrings(options.ExchangeIDs)
		}
		if options.IncludeExchangeLogo {
			params["include_exchange_logo"] = "true"
		}
		if options.Page != nil {
			params["page"] = intToString(*options.Page)
		}
		if options.Order != nil {
			params["order"] = *options.Order
		}
		if options.Depth != nil {
			params["depth"] = *options.Depth
		}
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/tickers", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// History returns historical data for a coin at the supplied date.
func (c *CoinsAPI) History(id string, date string, localization bool) (map[string]interface{}, error) {
	params := map[string]string{
		"date":         date,
		"localization": boolToString(localization),
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/history", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// MarketChartOptions configures a market-chart response.
type MarketChartOptions struct {
	Interval  *string
	Precision *string
}

// MarketChart returns historical market data for a coin over a period.
func (c *CoinsAPI) MarketChart(id string, vsCurrency string, days string, options *MarketChartOptions) (map[string]interface{}, error) {
	params := map[string]string{
		"vs_currency": vsCurrency,
		"days":        days,
	}

	if options != nil {
		if options.Interval != nil {
			params["interval"] = *options.Interval
		}
		if options.Precision != nil {
			params["precision"] = *options.Precision
		}
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/market_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// MarketChartRange returns historical market data between Unix timestamps.
func (c *CoinsAPI) MarketChartRange(id string, vsCurrency string, from int64, to int64, precision *string) (map[string]interface{}, error) {
	params := map[string]string{
		"vs_currency": vsCurrency,
		"from":        fmt.Sprintf("%d", from),
		"to":          fmt.Sprintf("%d", to),
	}

	if precision != nil {
		params["precision"] = *precision
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/market_chart/range", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// OHLC returns open, high, low, and close data for a coin.
func (c *CoinsAPI) OHLC(id string, vsCurrency string, days string, precision *string) ([][]float64, error) {
	params := map[string]string{
		"vs_currency": vsCurrency,
		"days":        days,
	}

	if precision != nil {
		params["precision"] = *precision
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/ohlc", params)
	if err != nil {
		return nil, err
	}

	var result [][]float64
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CirculatingSupplyChart returns a coin's circulating-supply history.
func (c *CoinsAPI) CirculatingSupplyChart(id string, days string, interval *string) (map[string]interface{}, error) {
	params := map[string]string{
		"days": days,
	}

	if interval != nil {
		params["interval"] = *interval
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/circulating_supply_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CirculatingSupplyChartRange returns circulating-supply history between Unix timestamps.
func (c *CoinsAPI) CirculatingSupplyChartRange(id string, from int64, to int64) (map[string]interface{}, error) {
	params := map[string]string{
		"from": fmt.Sprintf("%d", from),
		"to":   fmt.Sprintf("%d", to),
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/circulating_supply_chart/range", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TotalSupplyChart returns a coin's total-supply history.
func (c *CoinsAPI) TotalSupplyChart(id string, days string, interval *string) (map[string]interface{}, error) {
	params := map[string]string{
		"days": days,
	}

	if interval != nil {
		params["interval"] = *interval
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/total_supply_chart", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// TotalSupplyChartRange returns total-supply history between Unix timestamps.
func (c *CoinsAPI) TotalSupplyChartRange(id string, from int64, to int64) (map[string]interface{}, error) {
	params := map[string]string{
		"from": fmt.Sprintf("%d", from),
		"to":   fmt.Sprintf("%d", to),
	}

	data, err := c.client.get("/coins/"+escapePathSegment(id)+"/total_supply_chart/range", params)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}
