# CoinGecko Go Client

![CoinGecko Golang SDK](https://github.com/user-attachments/assets/f0095cd2-0e50-44c2-94a6-4af1ca483873)

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-tigusigalpa%2Fcoingecko--go-blue)](https://github.com/tigusigalpa/coingecko-go)

Go client for the [CoinGecko API](https://docs.coingecko.com). Prices, market data, charts, exchanges, NFTs, DeFi — all endpoints, typed responses, minimal dependencies.

**Package:** [pkg.go.dev/github.com/tigusigalpa/coingecko-go](https://pkg.go.dev/github.com/tigusigalpa/coingecko-go)

[Installation](#installation) • [Quick Start](#quick-start) • [API Reference](#api-reference) • [Examples](#examples)

</div>

---

## Features

- All CoinGecko API v3 endpoints: Simple, Coins, Contracts, Exchanges, Derivatives, NFTs, Categories, Global, Search, Ping
- Pro API support (API key auth, higher rate limits)
- WebSocket for live price feeds
- Configurable HTTP client, timeouts, custom transport
- Minimal deps — stdlib + `gorilla/websocket`
- Tested, with runnable examples in `examples/`

## Requirements

- Go 1.21+
- CoinGecko API key (optional, for Pro tier)

## Installation

```bash
go get github.com/tigusigalpa/coingecko-go
```

```go
import coingecko "github.com/tigusigalpa/coingecko-go"
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    coingecko "github.com/tigusigalpa/coingecko-go"
)

func main() {
    cg := coingecko.New()
    
    prices, err := cg.Simple().Price(
        []string{"bitcoin"},
        []string{"usd"},
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Bitcoin: $%.2f\n", prices["bitcoin"]["usd"])
}
```

### Pro API

```go
cg := coingecko.New(
    coingecko.WithAPIKey("your-api-key"),
    coingecko.WithProAPI(true),
)
```

### Custom HTTP Client

```go
httpClient := &http.Client{
    Timeout: 60 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}

cg := coingecko.New(
    coingecko.WithHTTPClient(httpClient),
    coingecko.WithAPIKey("your-api-key"),
    coingecko.WithProAPI(true),
)
```

### Pointer Helpers

Many options use pointers. A couple of helpers make this less verbose:

```go
func stringPtr(s string) *string { return &s }
func intPtr(i int) *int { return &i }
```

---

## API Reference

### Simple

Prices and supported currencies.

```go
// Get prices (multiple coins, multiple currencies, with extras)
prices, err := cg.Simple().Price(
    []string{"bitcoin", "ethereum"},
    []string{"usd", "eur"},
    &coingecko.SimplePriceOptions{
        IncludeMarketCap:  true,
        Include24hrVol:    true,
        Include24hrChange: true,
    },
)

// Token price by contract address
tokenPrices, err := cg.Simple().TokenPrice(
    "ethereum",
    []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"}, // UNI
    []string{"usd"},
    nil,
)

// Supported vs currencies
currencies, err := cg.Simple().SupportedVsCurrencies()
```

### Coins

Market data, charts, OHLC, detailed info.

```go
// List all coins
coins, err := cg.Coins().List(false)

// Market rankings
perPage := 100
page := 1
markets, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
    Order:   stringPtr("market_cap_desc"),
    PerPage: &perPage,
    Page:    &page,
})

// Detailed coin data
bitcoin, err := cg.Coins().Coin("bitcoin", &coingecko.CoinOptions{
    MarketData: true,
    Tickers:    true,
    Sparkline:  true,
})

// Price chart (30 days, daily interval)
chart, err := cg.Coins().MarketChart("bitcoin", "usd", "30",
    &coingecko.MarketChartOptions{Interval: stringPtr("daily")},
)

// Chart for a specific time range (unix timestamps)
rangeData, err := cg.Coins().MarketChartRange("ethereum", "usd", 1609459200, 1640995200, nil)

// OHLC candlestick data
ohlc, err := cg.Coins().OHLC("bitcoin", "usd", "7", nil)

// Top gainers/losers
duration := "24h"
topCoins := 100
movers, err := cg.Coins().TopGainersLosers("usd", &duration, &topCoins)

// Recently added
newCoins, err := cg.Coins().RecentlyAdded()
```

### Contract

Token data by contract address.

```go
tokenData, err := cg.Contract().Coin(
    "ethereum",
    "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
)

tokenChart, err := cg.Contract().MarketChart(
    "ethereum",
    "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
    "usd", "30",
)
```

### Asset Platforms

```go
platforms, err := cg.AssetPlatforms().List(nil)

filter := "nft"
filtered, err := cg.AssetPlatforms().List(&filter)

tokens, err := cg.AssetPlatforms().TokenLists("ethereum", nil)
```

### Categories

```go
categories, err := cg.Categories().List()

order := "market_cap_desc"
withData, err := cg.Categories().ListWithMarketData(&order)
```

### Exchanges

```go
perPage := 100
page := 1
exchanges, err := cg.Exchanges().List(&coingecko.ExchangesListOptions{
    PerPage: &perPage,
    Page:    &page,
})

binance, err := cg.Exchanges().Exchange("binance")

tickers, err := cg.Exchanges().Tickers("binance", &coingecko.ExchangeTickersOptions{
    CoinIDs: []string{"bitcoin", "ethereum"},
    Page:    &page,
})

volumeChart, err := cg.Exchanges().VolumeChart("binance", 30)
```

### Global

```go
globalData, err := cg.Global().Crypto()
defiData, err := cg.Global().DeFi()
marketCapChart, err := cg.Global().MarketCapChart(30, "usd")
```

### Search

```go
results, err := cg.Search().Search("bitcoin")
trending, err := cg.Search().Trending()
```

### Ping

```go
status, err := cg.Ping().Ping()
fmt.Println(status.GeckoSays) // "(V3) To the Moon!"

// Pro API only
usage, err := cg.Ping().APIUsage()
```

---

## Examples

### Portfolio Tracker

```go
package main

import (
    "fmt"
    "log"
    
    coingecko "github.com/tigusigalpa/coingecko-go"
)

func main() {
    cg := coingecko.New()
    
    holdings := map[string]float64{
        "bitcoin":  0.5,
        "ethereum": 10,
        "cardano":  1000,
    }
    
    coinIDs := make([]string, 0, len(holdings))
    for id := range holdings {
        coinIDs = append(coinIDs, id)
    }
    
    prices, err := cg.Simple().Price(coinIDs, []string{"usd"}, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    total := 0.0
    for id, amount := range holdings {
        price := prices[id]["usd"].(float64)
        total += price * amount
    }
    
    fmt.Printf("Portfolio: $%.2f\n", total)
}
```

### Price Alerts

```go
package main

import (
    "fmt"
    "log"
    
    coingecko "github.com/tigusigalpa/coingecko-go"
)

type Alert struct {
    CoinID      string
    TargetPrice float64
    Type        string // "above" or "below"
}

func main() {
    cg := coingecko.New()
    
    alerts := []Alert{
        {CoinID: "bitcoin", TargetPrice: 50000, Type: "above"},
        {CoinID: "ethereum", TargetPrice: 3000, Type: "below"},
    }
    
    ids := make([]string, 0)
    seen := map[string]bool{}
    for _, a := range alerts {
        if !seen[a.CoinID] {
            ids = append(ids, a.CoinID)
            seen[a.CoinID] = true
        }
    }
    
    prices, err := cg.Simple().Price(ids, []string{"usd"}, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, a := range alerts {
        price := prices[a.CoinID]["usd"].(float64)
        if (a.Type == "above" && price >= a.TargetPrice) ||
            (a.Type == "below" && price <= a.TargetPrice) {
            fmt.Printf("ALERT: %s is $%.2f (target: $%.0f %s)\n",
                a.CoinID, price, a.TargetPrice, a.Type)
        }
    }
}
```

### Runnable Examples

There are ready-to-run examples in the `examples/` directory:

```bash
cd examples/basic && go run main.go
cd examples/portfolio && go run main.go
cd examples/market_analysis && go run main.go
cd examples/defi_tracker && go run main.go
```

---

## Rate Limits

Keep in mind the CoinGecko rate limits:

- **Free**: 10–50 calls/min
- **Pro**: higher limits depending on plan

If you're hitting limits, add caching on your side — the library doesn't do that for you.

## Testing

```bash
go test ./...
go test -v -race -coverprofile=coverage.out ./...

# or via Makefile
make test
make build
make lint
```

## Security

Don't hardcode API keys. Use environment variables:

```go
apiKey := os.Getenv("COINGECKO_API_KEY")
cg := coingecko.New(
    coingecko.WithAPIKey(apiKey),
    coingecko.WithProAPI(true),
)
```

If you find a security issue, email **sovletig@gmail.com** instead of opening a public issue.

## Contributing

PRs are welcome. Fork, branch, add tests, run `make test`, open a PR. See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

MIT. See [LICENSE](LICENSE).

## Author

[Igor Sazonov](https://github.com/tigusigalpa) — sovletig@gmail.com

## Links

- [CoinGecko API docs](https://docs.coingecko.com)
- [CoinGecko Pro](https://www.coingecko.com/en/api)
- [Changelog](CHANGELOG.md)
- [Issues](https://github.com/tigusigalpa/coingecko-go/issues)
