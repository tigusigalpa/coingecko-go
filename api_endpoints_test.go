package coingecko

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func TestEndpointGroups(t *testing.T) {
	client := New(WithHTTPClient(&http.Client{Transport: roundTripFunc(mockAPIResponse)}))
	client.client.baseURL = "https://api.example.test"

	text := "value"
	page := 1
	limit := 10

	assertNoError(t, func() error {
		_, err := client.Simple().Price([]string{"bitcoin"}, []string{"usd"}, &SimplePriceOptions{true, true, true, true})
		return err
	}())
	assertNoError(t, func() error {
		_, err := client.Simple().TokenPrice("ethereum", []string{"0xabc"}, []string{"usd"}, &TokenPriceOptions{true, true, true, true})
		return err
	}())
	assertNoError(t, func() error { _, err := client.Simple().SupportedVsCurrencies(); return err }())

	assertNoError(t, func() error { _, err := client.Coins().List(true); return err }())
	assertNoError(t, func() error { _, err := client.Coins().TopGainersLosers("usd", &text, &limit); return err }())
	assertNoError(t, func() error { _, err := client.Coins().RecentlyAdded(); return err }())
	assertNoError(t, func() error {
		_, err := client.Coins().Markets("usd", &MarketsOptions{[]string{"bitcoin"}, &text, &text, &limit, &page, true, &text, &text, &text})
		return err
	}())
	assertNoError(t, func() error {
		_, err := client.Coins().Coin("bitcoin", &CoinOptions{true, true, true, true, true, true})
		return err
	}())
	assertNoError(t, func() error { _, err := client.Coins().Coin("bitcoin", nil); return err }())
	assertNoError(t, func() error {
		_, err := client.Coins().Tickers("bitcoin", &TickersOptions{[]string{"binance"}, true, &page, &text, &text})
		return err
	}())
	assertNoError(t, func() error { _, err := client.Coins().History("bitcoin", "01-01-2024", true); return err }())
	assertNoError(t, func() error {
		_, err := client.Coins().MarketChart("bitcoin", "usd", "30", &MarketChartOptions{&text, &text})
		return err
	}())
	assertNoError(t, func() error { _, err := client.Coins().MarketChartRange("bitcoin", "usd", 1, 2, &text); return err }())
	assertNoError(t, func() error { _, err := client.Coins().OHLC("bitcoin", "usd", "7", &text); return err }())
	assertNoError(t, func() error { _, err := client.Coins().CirculatingSupplyChart("bitcoin", "30", &text); return err }())
	assertNoError(t, func() error { _, err := client.Coins().CirculatingSupplyChartRange("bitcoin", 1, 2); return err }())
	assertNoError(t, func() error { _, err := client.Coins().TotalSupplyChart("bitcoin", "30", &text); return err }())
	assertNoError(t, func() error { _, err := client.Coins().TotalSupplyChartRange("bitcoin", 1, 2); return err }())

	assertNoError(t, func() error { _, err := client.Contract().Coin("ethereum", "0xabc"); return err }())
	assertNoError(t, func() error { _, err := client.Contract().MarketChart("ethereum", "0xabc", "usd", "30"); return err }())
	assertNoError(t, func() error {
		_, err := client.Contract().MarketChartRange("ethereum", "0xabc", "usd", 1, 2)
		return err
	}())
	assertNoError(t, func() error { _, err := client.AssetPlatforms().List(&text); return err }())
	assertNoError(t, func() error { _, err := client.AssetPlatforms().TokenLists("ethereum", &page); return err }())
	assertNoError(t, func() error { _, err := client.Categories().List(); return err }())
	assertNoError(t, func() error { _, err := client.Categories().ListWithMarketData(&text); return err }())

	assertNoError(t, func() error { _, err := client.Derivatives().Tickers(&text); return err }())
	assertNoError(t, func() error {
		_, err := client.Derivatives().Exchanges(&DerivativesExchangesOptions{&text, &limit, &page})
		return err
	}())
	assertNoError(t, func() error { _, err := client.Derivatives().Exchange("binance", &text); return err }())
	assertNoError(t, func() error { _, err := client.Derivatives().ExchangesList(); return err }())
	assertNoError(t, func() error { _, err := client.Entities().List(); return err }())
	assertNoError(t, func() error { _, err := client.Entities().TreasuryByCoinID("bitcoin"); return err }())
	assertNoError(t, func() error { _, err := client.Entities().TreasuryByEntityID("entity", "bitcoin"); return err }())
	assertNoError(t, func() error { _, err := client.Entities().TreasuryChart("entity", "bitcoin", 30); return err }())
	assertNoError(t, func() error {
		_, err := client.Entities().TransactionHistory("entity", "bitcoin", &TransactionHistoryOptions{&page, &limit})
		return err
	}())

	assertNoError(t, func() error { _, err := client.ExchangeRates().Rates(); return err }())
	assertNoError(t, func() error { _, err := client.Exchanges().List(&ExchangesListOptions{&limit, &page}); return err }())
	assertNoError(t, func() error { _, err := client.Exchanges().ListIDMap(); return err }())
	assertNoError(t, func() error { _, err := client.Exchanges().Exchange("binance"); return err }())
	assertNoError(t, func() error {
		_, err := client.Exchanges().Tickers("binance", &ExchangeTickersOptions{[]string{"bitcoin"}, true, &page, &text, &text})
		return err
	}())
	assertNoError(t, func() error { _, err := client.Exchanges().VolumeChart("binance", 30); return err }())
	assertNoError(t, func() error { _, err := client.Exchanges().VolumeChartRange("binance", 1, 2); return err }())
	assertNoError(t, func() error { _, err := client.Global().Crypto(); return err }())
	assertNoError(t, func() error { _, err := client.Global().DeFi(); return err }())
	assertNoError(t, func() error { _, err := client.Global().MarketCapChart(30, "usd"); return err }())

	assertNoError(t, func() error { _, err := client.NFTs().List(&NFTsListOptions{&text, &text, &limit, &page}); return err }())
	assertNoError(t, func() error { _, err := client.NFTs().Collection("collection"); return err }())
	assertNoError(t, func() error { _, err := client.NFTs().CollectionByContract("ethereum", "0xabc"); return err }())
	assertNoError(t, func() error {
		_, err := client.NFTs().Markets(&NFTsMarketsOptions{&text, &text, &limit, &page})
		return err
	}())
	assertNoError(t, func() error { _, err := client.NFTs().MarketChart("collection", 30); return err }())
	assertNoError(t, func() error { _, err := client.NFTs().MarketChartByContract("ethereum", "0xabc", 30); return err }())
	assertNoError(t, func() error { _, err := client.NFTs().Tickers("collection"); return err }())
	assertNoError(t, func() error { _, err := client.Ping().Ping(); return err }())
	assertNoError(t, func() error { _, err := client.Ping().APIUsage(); return err }())
	assertNoError(t, func() error { _, err := client.Search().Search("bitcoin"); return err }())
	assertNoError(t, func() error { _, err := client.Search().Trending(); return err }())
}

func TestWebSocketMessageHelpers(t *testing.T) {
	api := New().WebSocket()
	assertNoError(t, func() error { _, err := api.CreateChannelSubscription(WebSocketChannelSimplePrice); return err }())
	assertNoError(t, func() error {
		_, err := api.CreateSimplePriceSubscription([]string{"bitcoin"}, []string{"usd"})
		return err
	}())
	assertNoError(t, func() error {
		_, err := api.CreateOnchainSimpleTokenPriceSubscription("eth", []string{"0xabc"}, nil)
		return err
	}())
	assertNoError(t, func() error { _, err := api.CreateOnchainTradeSubscription("eth", "0xabc"); return err }())
	assertNoError(t, func() error { _, err := api.CreateOnchainOHLCVSubscription("eth", "0xabc", "1m"); return err }())
	assertNoError(t, func() error { _, err := api.CreateUnsubscribeMessage(WebSocketChannelSimplePrice); return err }())
}

func mockAPIResponse(request *http.Request) (*http.Response, error) {
	payload := "{}"
	path := request.URL.Path
	switch {
	case strings.HasSuffix(path, "/simple/supported_vs_currencies"):
		payload = `[]`
	case strings.HasSuffix(path, "/coins/list"), strings.HasSuffix(path, "/coins/list/new"), strings.HasSuffix(path, "/coins/markets"), strings.HasSuffix(path, "/asset_platforms"), strings.HasSuffix(path, "/coins/categories/list"), strings.HasSuffix(path, "/coins/categories"), strings.HasSuffix(path, "/derivatives"), strings.HasSuffix(path, "/derivatives/exchanges"), strings.HasSuffix(path, "/derivatives/exchanges/list"), strings.HasSuffix(path, "/entities/list"), strings.HasSuffix(path, "/exchanges"), strings.HasSuffix(path, "/exchanges/list"), strings.HasSuffix(path, "/nfts/list"), strings.HasSuffix(path, "/nfts/markets"):
		payload = `[]`
	case strings.HasSuffix(path, "/ohlc"):
		payload = `[[1,2,3,4,5]]`
	case strings.Contains(path, "/volume_chart"):
		payload = `[[1,2]]`
	}

	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(payload)),
		Header:     make(http.Header),
		Request:    request,
	}, nil
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
