# CoinGecko Go API Client - Cryptocurrency Data Integration Library

![CoinGecko Golang SDK](https://github.com/user-attachments/assets/2f5f44bb-df92-4206-8233-05a94397d29a)

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-tigusigalpa%2Fcoingecko--go-blue)](https://github.com/tigusigalpa/coingecko-go)

**The most comprehensive Golang library for CoinGecko API integration** - Access real-time cryptocurrency prices, market
data, historical charts, DeFi analytics, NFT data, and more with a type-safe, production-ready Go client.

**📦 Package:** [pkg.go.dev/github.com/tigusigalpa/coingecko-go](https://pkg.go.dev/github.com/tigusigalpa/coingecko-go)

Perfect for building crypto trading bots, portfolio trackers, market analysis tools, DeFi applications, and blockchain
data platforms.

[Features](#-features) • [Installation](#-installation) • [Quick Start](#-quick-start) • [Getting Started](#-getting-started-guide) • [API Documentation](#-api-documentation) • [Examples](#-real-world-examples)

</div>

---

## 🌟 Why Choose CoinGecko Go Client?

This **Go cryptocurrency API library** provides enterprise-grade access to CoinGecko's comprehensive crypto market data.
Whether you're building a **Bitcoin price tracker**, **Ethereum market analyzer**, **DeFi portfolio manager**, or *
*crypto trading bot**, this library offers everything you need:

- **14+ API Endpoint Groups** covering coins, exchanges, DeFi, NFTs, derivatives, and more
- **Real-time WebSocket Support** for live cryptocurrency price feeds
- **Historical Data Access** for backtesting trading strategies and market analysis
- **Smart Contract Integration** for ERC-20 tokens and multi-chain assets
- **Production-Ready** with comprehensive error handling and rate limiting support
- **Type-Safe** Go implementation with zero runtime surprises
- **Well-Tested** with unit tests and real-world examples

## 🚀 Features

### Core Capabilities

- ✅ **Complete API Coverage** - All CoinGecko API v3 endpoints implemented (Simple, Coins, Contracts, Exchanges, DeFi,
  NFTs, and more)
- 🔐 **Pro API Support** - Full support for CoinGecko Pro API with authentication and higher rate limits
- 💎 **Modern Go** - Built with Go 1.21+ features, generics support, and idiomatic Go patterns
- 🛡️ **Type Safety** - Strongly typed parameters and return values prevent runtime errors
- 📦 **Minimal Dependencies** - Only standard library (plus gorilla/websocket for WebSocket support)
- 🔄 **Flexible Configuration** - Customizable HTTP client, timeouts, retry logic, and middleware
- 🎨 **Clean API Design** - Intuitive, chainable methods for excellent developer experience
- ⚡ **High Performance** - Efficient HTTP client with connection pooling and concurrent request support
- 📚 **Comprehensive Documentation** - Detailed examples, API reference, and integration guides

### Cryptocurrency Data Access

- 💰 **Price Data** - Real-time and historical prices for 10,000+ cryptocurrencies
- 📊 **Market Data** - Market cap, volume, supply, price changes, and market dominance
- 📈 **Charts & OHLC** - Historical price charts, candlestick data, and technical indicators
- 🔍 **Search & Discovery** - Search coins, trending cryptocurrencies, and new listings
- 🏦 **Exchange Data** - Trading volumes, tickers, and exchange rankings
- 💹 **Derivatives** - Futures, perpetuals, and derivatives market data
- 🎨 **NFT Analytics** - NFT collection data, floor prices, and trading volumes
- 🏛️ **Treasury Data** - Corporate and institutional crypto holdings
- 🌐 **Global Metrics** - Total market cap, DeFi TVL, and market statistics

## 📋 Requirements

- **Go 1.21 or higher** (supports Go 1.21, 1.22, and later)
- **Operating Systems**: Linux, macOS, Windows
- **Optional**: CoinGecko API key for Pro tier access

## 📦 Installation

Install the CoinGecko Golang client using Go modules:

```bash
go get github.com/tigusigalpa/coingecko-go
```

Import in your Go application:

```go
import coingecko "github.com/tigusigalpa/coingecko-go"
```

## 🎯 Quick Start

Get up and running with the CoinGecko Golang client in under 60 seconds:

### Basic Usage - Free API

```go
package main

import (
    "fmt"
    "log"
    
    coingecko "github.com/tigusigalpa/coingecko-go"
)

func main() {
    // Create a new CoinGecko client (no API key needed for free tier)
    cg := coingecko.New()
    
    // Get Bitcoin price in USD
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

### Pro API with Authentication

Unlock higher rate limits and advanced features with CoinGecko Pro:

```go
cg := coingecko.New(
    coingecko.WithAPIKey("your-api-key-here"),
    coingecko.WithProAPI(true),
)
```

### Advanced Configuration

Customize the HTTP client for production environments:

```go
import (
    "net/http"
    "time"
)

// Configure custom HTTP client with timeout and retry logic
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
    coingecko.WithTimeout(60 * time.Second),
)
```

---

## 📘 Getting Started Guide

### Common Use Cases

#### 1. Get Current Cryptocurrency Prices

Fetch real-time prices for single or multiple cryptocurrencies:

```go
// Single coin, single currency
prices, err := cg.Simple().Price(
    []string{"bitcoin"},
    []string{"usd"},
    nil,
)

// Multiple coins with market data
prices, err := cg.Simple().Price(
    []string{"bitcoin", "ethereum", "cardano", "solana"},
    []string{"usd", "eur", "gbp"},
    &coingecko.SimplePriceOptions{
        IncludeMarketCap:   true,
        Include24hrVol:     true,
        Include24hrChange:  true,
    },
)
```

#### 2. Get Market Data and Rankings

Access comprehensive market data for crypto assets:

```go
perPage := 10
page := 1

markets, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
    Order:   stringPtr("market_cap_desc"),
    PerPage: &perPage,
    Page:    &page,
})
```

#### 3. Get Historical Price Data

Retrieve historical data for backtesting and analysis:

```go
// Market chart with price, volume, and market cap
chartData, err := cg.Coins().MarketChart(
    "bitcoin",
    "usd",
    "30", // days: 1, 7, 14, 30, 90, 180, 365, "max"
    nil,
)

// OHLC candlestick data for technical analysis
ohlc, err := cg.Coins().OHLC("bitcoin", "usd", "7", nil)
```

#### 4. Search and Trending Cryptocurrencies

Discover trending coins and search the crypto market:

```go
// Search for specific cryptocurrency
results, err := cg.Search().Search("bitcoin")

// Get trending cryptocurrencies
trending, err := cg.Search().Trending()
```

#### 5. Smart Contract Token Data

Access ERC-20 and other token data by contract address:

```go
// Get token data by contract address
tokenData, err := cg.Contract().Coin(
    "ethereum",
    "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984", // Uniswap (UNI) token
)
```

### Helper Functions for Optional Parameters

When working with optional parameters, use pointer helpers:

```go
func stringPtr(s string) *string {
    return &s
}

func intPtr(i int) *int {
    return &i
}

// Usage example
category := "decentralized-finance-defi"
perPage := 50

markets, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
    Category: &category,
    PerPage:  &perPage,
})
```

### Error Handling Best Practices

Always implement proper error handling in production code:

```go
data, err := cg.Coins().List(false)
if err != nil {
    log.Printf("CoinGecko API error: %v", err)
    // Implement retry logic or fallback mechanism
    return
}
```

### Rate Limiting and Caching

Implement caching to respect API rate limits and improve performance:

```go
import (
    "time"
    "sync"
)

type Cache struct {
    data      map[string]interface{}
    timestamp map[string]time.Time
    mutex     sync.RWMutex
    ttl       time.Duration
}

func NewCache(ttl time.Duration) *Cache {
    return &Cache{
        data:      make(map[string]interface{}),
        timestamp: make(map[string]time.Time),
        ttl:       ttl,
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()
    
    data, exists := c.data[key]
    if !exists {
        return nil, false
    }
    
    if time.Since(c.timestamp[key]) > c.ttl {
        return nil, false
    }
    
    return data, true
}

func (c *Cache) Set(key string, value interface{}) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    
    c.data[key] = value
    c.timestamp[key] = time.Now()
}
```

---

## 📖 API Documentation

### Complete API Endpoint Reference

The CoinGecko Go client provides access to all major cryptocurrency data endpoints, organized into 14 logical API groups
for easy navigation:

- **[Simple API](#simple-api)** - Quick price lookups and supported currencies
- **[Coins API](#coins-api)** - Comprehensive coin data, markets, and historical charts
- **[Contract API](#contract-api)** - Token data by contract address
- **[Asset Platforms API](#asset-platforms-api)** - Asset platforms and token lists
- **[Categories API](#categories-api)** - Coin categories with market data
- **[Exchanges API](#exchanges-api)** - Exchange data, tickers, and volume charts
- **[Derivatives API](#derivatives-api)** - Derivatives tickers and exchanges
- **[Entities API](#entities-api)** - Crypto treasury holdings and transactions
- **[NFTs API](#nfts-api)** - NFT collections and market data
- **[Exchange Rates API](#exchange-rates-api)** - BTC-to-currency exchange rates
- **[Search API](#search-api)** - Search queries and trending data
- **[Global API](#global-api)** - Global crypto and DeFi market data
- **[WebSocket API](#websocket-api)** - Real-time WebSocket subscriptions
- **[Ping API](#ping-api)** - Server status and API usage

---

### Simple API

Get cryptocurrency prices and basic market data quickly.

#### Get Coin Prices

```go
// Single coin, single currency
prices, err := cg.Simple().Price(
    []string{"bitcoin"},
    []string{"usd"},
    nil,
)

// Multiple coins with additional data
prices, err := cg.Simple().Price(
    []string{"bitcoin", "ethereum", "cardano"},
    []string{"usd", "eur", "gbp"},
    &coingecko.SimplePriceOptions{
        IncludeMarketCap:   true,
        Include24hrVol:     true,
        Include24hrChange:  true,
    },
)
```

#### Get Token Prices by Contract Address

```go
tokenPrices, err := cg.Simple().TokenPrice(
    "ethereum",
    []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"}, // UNI token
    []string{"usd"},
    &coingecko.TokenPriceOptions{
        IncludeMarketCap: true,
        Include24hrVol:   true,
    },
)
```

#### Get Supported Currencies

```go
currencies, err := cg.Simple().SupportedVsCurrencies()
// Returns: ["usd", "eur", "gbp", "jpy", ...]
```

---

### Coins API

Access detailed coin information, market data, and historical charts.

#### List All Coins

```go
// Get list of all coins with IDs
coinsList, err := cg.Coins().List(false)

// Include platform information
coinsWithPlatforms, err := cg.Coins().List(true)
```

#### Top Gainers & Losers

```go
duration := "24h"
topCoins := 100

topMovers, err := cg.Coins().TopGainersLosers("usd", &duration, &topCoins)
```

#### Recently Added Coins

```go
newCoins, err := cg.Coins().RecentlyAdded()
```

#### Coins Market Data

```go
// Get top 100 coins by market cap
perPage := 100
page := 1
priceChange := "1h,24h,7d"

markets, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
    Order:                 stringPtr("market_cap_desc"),
    PerPage:               &perPage,
    Page:                  &page,
    Sparkline:             true,
    PriceChangePercentage: &priceChange,
})

// Filter by specific coins
markets, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
    IDs: []string{"bitcoin", "ethereum", "cardano"},
})

// Filter by category
category := "decentralized-finance-defi"
markets, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
    Category: &category,
    PerPage:  &perPage,
})
```

#### Detailed Coin Data

```go
bitcoin, err := cg.Coins().Coin("bitcoin", &coingecko.CoinOptions{
    Localization:  true,
    Tickers:       true,
    MarketData:    true,
    CommunityData: true,
    DeveloperData: true,
    Sparkline:     true,
})
```

#### Market Chart Data

```go
// Get price, market cap, and volume chart data
chartData, err := cg.Coins().MarketChart(
    "bitcoin",
    "usd",
    "30", // days: 1, 7, 14, 30, 90, 180, 365, "max"
    &coingecko.MarketChartOptions{
        Interval: stringPtr("daily"),
    },
)

// Get chart data for specific time range
rangeData, err := cg.Coins().MarketChartRange(
    "ethereum",
    "usd",
    1609459200, // Unix timestamp
    1640995200, // Unix timestamp
    nil,
)
```

#### OHLC Chart Data

```go
ohlc, err := cg.Coins().OHLC("bitcoin", "usd", "7", nil)
// Returns: [][]float64 - [[timestamp, open, high, low, close], ...]
```

---

### Contract API

Get token data using contract addresses on various blockchain platforms.

```go
// Get comprehensive token data
tokenData, err := cg.Contract().Coin(
    "ethereum",
    "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984", // UNI token
)

// Get price chart for a token
tokenChart, err := cg.Contract().MarketChart(
    "ethereum",
    "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
    "usd",
    "30",
)
```

---

### Asset Platforms API

```go
// Get all asset platforms
platforms, err := cg.AssetPlatforms().List(nil)

// Filter platforms
filter := "nft"
filtered, err := cg.AssetPlatforms().List(&filter)

// Get all tokens on a specific platform
tokens, err := cg.AssetPlatforms().TokenLists("ethereum", nil)
```

---

### Categories API

```go
// Get list of all coin categories
categories, err := cg.Categories().List()

// Get categories with market data
order := "market_cap_desc"
categoriesData, err := cg.Categories().ListWithMarketData(&order)
```

---

### Exchanges API

```go
// Get all exchanges with data
perPage := 100
page := 1
exchanges, err := cg.Exchanges().List(&coingecko.ExchangesListOptions{
    PerPage: &perPage,
    Page:    &page,
})

// Get detailed data for a specific exchange
binance, err := cg.Exchanges().Exchange("binance")

// Get exchange tickers
tickers, err := cg.Exchanges().Tickers("binance", &coingecko.ExchangeTickersOptions{
    CoinIDs: []string{"bitcoin", "ethereum"},
    Page:    &page,
})

// Get volume chart
volumeChart, err := cg.Exchanges().VolumeChart("binance", 30)
```

---

### Global API

```go
// Get global crypto market data
globalData, err := cg.Global().Crypto()

// Get global DeFi market data
defiData, err := cg.Global().DeFi()

// Get global market cap chart
marketCapChart, err := cg.Global().MarketCapChart(30, "usd")
```

---

### Search API

```go
// Search for coins, exchanges, categories, and NFTs
results, err := cg.Search().Search("bitcoin")

// Get trending searches
trending, err := cg.Search().Trending()
```

---

### Ping API

```go
// Check server status
status, err := cg.Ping().Ping()
fmt.Println(status.GeckoSays) // "(V3) To the Moon!"

// Check API usage (Pro API only)
usage, err := cg.Ping().APIUsage()
```

---

---

## 💡 Real-World Examples

### Building a Cryptocurrency Portfolio Tracker

Track your crypto portfolio value in real-time with automatic price updates:

```go
package main

import (
    "fmt"
    "log"
    
    coingecko "github.com/tigusigalpa/coingecko-go"
)

func main() {
    cg := coingecko.New()
    
    // Portfolio holdings
    holdings := map[string]float64{
        "bitcoin":  0.5,
        "ethereum": 10,
        "cardano":  1000,
    }
    
    // Get coin IDs
    coinIDs := make([]string, 0, len(holdings))
    for coinID := range holdings {
        coinIDs = append(coinIDs, coinID)
    }
    
    // Get prices
    prices, err := cg.Simple().Price(coinIDs, []string{"usd"}, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // Calculate total value
    totalValue := 0.0
    for coinID, amount := range holdings {
        price := prices[coinID]["usd"].(float64)
        totalValue += price * amount
    }
    
    fmt.Printf("Portfolio Value: $%.2f\n", totalValue)
}
```

### Building a Cryptocurrency Price Alert System

Create automated price alerts for Bitcoin, Ethereum, and other cryptocurrencies:

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

func checkAlerts(cg *coingecko.CoinGecko, alerts []Alert) []Alert {
    // Get unique coin IDs
    coinIDs := make([]string, 0, len(alerts))
    seen := make(map[string]bool)
    for _, alert := range alerts {
        if !seen[alert.CoinID] {
            coinIDs = append(coinIDs, alert.CoinID)
            seen[alert.CoinID] = true
        }
    }
    
    // Get prices
    prices, err := cg.Simple().Price(coinIDs, []string{"usd"}, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // Check alerts
    triggered := []Alert{}
    for _, alert := range alerts {
        currentPrice := prices[alert.CoinID]["usd"].(float64)
        
        if alert.Type == "above" && currentPrice >= alert.TargetPrice {
            triggered = append(triggered, alert)
        } else if alert.Type == "below" && currentPrice <= alert.TargetPrice {
            triggered = append(triggered, alert)
        }
    }
    
    return triggered
}

func main() {
    cg := coingecko.New()
    
    alerts := []Alert{
        {CoinID: "bitcoin", TargetPrice: 50000, Type: "above"},
        {CoinID: "ethereum", TargetPrice: 3000, Type: "below"},
    }
    
    triggered := checkAlerts(cg, alerts)
    fmt.Printf("Triggered alerts: %d\n", len(triggered))
}
```

### Crypto Market Analysis Dashboard

Build a comprehensive cryptocurrency market analysis tool:

```go
package main

import (
    "fmt"
    "log"
    
    coingecko "github.com/tigusigalpa/coingecko-go"
)

func main() {
    cg := coingecko.New()
    
    // Get top 10 coins by market cap
    perPage := 10
    page := 1
    priceChange := "1h,24h,7d,30d"
    
    topCoins, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
        Order:                 stringPtr("market_cap_desc"),
        PerPage:               &perPage,
        Page:                  &page,
        Sparkline:             true,
        PriceChangePercentage: &priceChange,
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Top 10 Cryptocurrencies by Market Cap:")
    for i, coin := range topCoins {
        name := coin["name"].(string)
        price := coin["current_price"].(float64)
        marketCap := coin["market_cap"].(float64)
        
        fmt.Printf("%d. %s: $%.2f (Market Cap: $%.0f)\n", 
            i+1, name, price, marketCap)
    }
}

func stringPtr(s string) *string {
    return &s
}
```

### Running the Examples

The package includes 4 production-ready example applications:

```bash
# Basic cryptocurrency price lookup
cd examples/basic && go run main.go

# Portfolio value tracker
cd examples/portfolio && go run main.go

# Market analysis with top gainers/losers
cd examples/market_analysis && go run main.go

# DeFi token tracker with contract data
cd examples/defi_tracker && go run main.go
```

---

## 🔧 Production Best Practices

### Error Handling

```go
prices, err := cg.Simple().Price(
    []string{"bitcoin"},
    []string{"usd"},
    nil,
)
if err != nil {
    log.Printf("CoinGecko API Error: %v", err)
    // Handle error appropriately
    return
}
```

### Rate Limiting and Performance Optimization

Understanding and respecting CoinGecko API rate limits is crucial for production applications:

#### Rate Limits by Tier

- **Free API**: 10-50 calls/minute (suitable for personal projects and testing)
- **Pro API**: Higher limits based on your subscription plan (recommended for production)
- **Enterprise**: Custom rate limits for high-volume applications

#### Implementing Smart Caching

Reduce API calls and improve response times with intelligent caching:

```go
import (
    "time"
    "sync"
)

type CachedPrice struct {
    Price     map[string]map[string]interface{}
    Timestamp time.Time
}

var (
    cache      = make(map[string]*CachedPrice)
    cacheMutex sync.RWMutex
    cacheTTL   = 5 * time.Minute
)

func getCachedPrice(cg *coingecko.CoinGecko, coinID, currency string) (map[string]map[string]interface{}, error) {
    cacheKey := fmt.Sprintf("%s_%s", coinID, currency)
    
    cacheMutex.RLock()
    cached, exists := cache[cacheKey]
    cacheMutex.RUnlock()
    
    if exists && time.Since(cached.Timestamp) < cacheTTL {
        return cached.Price, nil
    }
    
    // Fetch fresh data
    prices, err := cg.Simple().Price([]string{coinID}, []string{currency}, nil)
    if err != nil {
        return nil, err
    }
    
    // Update cache
    cacheMutex.Lock()
    cache[cacheKey] = &CachedPrice{
        Price:     prices,
        Timestamp: time.Now(),
    }
    cacheMutex.Unlock()
    
    return prices, nil
}
```

---

## 🧪 Testing and Quality Assurance

Run the comprehensive test suite to ensure reliability:

```bash
# Run all tests
go test ./...

# Run tests with coverage report
go test -v -race -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests with verbose output
go test -v ./...

# Use Makefile for convenience
make test
```

### Building the Package

```bash
# Build the package
go build ./...

# Format code
go fmt ./...

# Run linter
go vet ./...

# Or use Makefile
make build
make lint
```

## 🤝 Contributing to CoinGecko Go Client

We welcome contributions from the cryptocurrency and Go developer community! Whether you're fixing bugs, adding new
features, improving documentation, or creating examples, your help is appreciated.

### How to Contribute

1. **Fork the repository** on GitHub
2. **Create your feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** with clear, well-documented code
4. **Add tests** for new functionality
5. **Run the test suite**: `make test`
6. **Commit your changes**: `git commit -m 'Add cryptocurrency trading bot example'`
7. **Push to your branch**: `git push origin feature/amazing-feature`
8. **Open a Pull Request** with a clear description of your changes

### Contribution Ideas

- Add more real-world examples (trading bots, arbitrage tools, analytics dashboards)
- Improve error handling and retry logic
- Add support for new CoinGecko API endpoints
- Enhance documentation with tutorials and guides
- Optimize performance and reduce memory usage
- Add integration tests with mock API responses

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.

## 🔒 Security and API Key Management

### Protecting Your API Keys

When using CoinGecko Pro API, always protect your API keys:

```go
// ✅ Good: Use environment variables
apiKey := os.Getenv("COINGECKO_API_KEY")
cg := coingecko.New(
    coingecko.WithAPIKey(apiKey),
    coingecko.WithProAPI(true),
)

// ❌ Bad: Never hardcode API keys
// cg := coingecko.New(coingecko.WithAPIKey("hardcoded-key-here"))
```

### Reporting Security Issues

If you discover any security vulnerabilities or issues related to API key handling, cryptocurrency data integrity, or
other security concerns, please email **sovletig@gmail.com** directly instead of using the public issue tracker. We take
security seriously and will respond promptly.

## 📄 License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

## 🙏 Credits and Acknowledgments

- **Author**: [Igor Sazonov](https://github.com/tigusigalpa) - Go cryptocurrency developer
- **CoinGecko**: [CoinGecko.com](https://www.coingecko.com) - World's largest independent cryptocurrency data aggregator
- **Contributors**: All developers who have contributed to this open-source project
- **Community**: Go developers and cryptocurrency enthusiasts worldwide

## 🔗 Useful Links and Resources

### Project Resources

- **GitHub Repository**: [github.com/tigusigalpa/coingecko-go](https://github.com/tigusigalpa/coingecko-go)
- **Issue Tracker**: [Report bugs and request features](https://github.com/tigusigalpa/coingecko-go/issues)
- **Changelog**: [View version history](CHANGELOG.md)
- **Contributing Guide**: [Learn how to contribute](CONTRIBUTING.md)

### CoinGecko Resources

- **CoinGecko API Documentation**: [docs.coingecko.com](https://docs.coingecko.com) - Complete API reference
- **CoinGecko Website**: [coingecko.com](https://www.coingecko.com) - Cryptocurrency prices and market data
- **CoinGecko Pro**: [coingecko.com/en/api](https://www.coingecko.com/en/api) - Professional API plans
- **API Authentication Guide**: [Setting up your API key](https://docs.coingecko.com/docs/setting-up-your-api-key)

### Go Resources

- **Go Documentation**: [golang.org/doc](https://golang.org/doc)
- **Go Modules**: [Go dependency management](https://go.dev/ref/mod)

## 📊 Use Cases and Applications

This Go library is perfect for building:

- 📈 **Cryptocurrency Trading Bots** - Automated trading with real-time price data
- 💼 **Portfolio Management Apps** - Track and analyze crypto investments
- 📊 **Market Analysis Tools** - Technical analysis and market research platforms
- 🔔 **Price Alert Systems** - Real-time cryptocurrency price notifications
- 📱 **Mobile Apps** - iOS and Android crypto apps with Go backends
- 🌐 **Web Dashboards** - Crypto market dashboards and analytics platforms
- 🤖 **Telegram/Discord Bots** - Crypto price bots for social platforms
- 📉 **Backtesting Systems** - Historical data analysis for trading strategies
- 🏦 **DeFi Applications** - Decentralized finance tools and platforms
- 🎨 **NFT Analytics** - NFT market tracking and analysis tools

## 🏷️ Keywords

Go cryptocurrency API, CoinGecko Go client, Bitcoin price API Go, Ethereum market data Go, crypto trading bot Go, DeFi
API Go, NFT data API, cryptocurrency portfolio tracker, Go blockchain library, crypto market analysis Go, real-time
crypto prices, historical cryptocurrency data, Go crypto SDK

---

<div align="center">

**[⬆ back to top](#coingecko-go-api-client)**

</div>
