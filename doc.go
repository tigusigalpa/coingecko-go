// Package coingecko provides a comprehensive Go client for the CoinGecko API.
//
// CoinGecko is a cryptocurrency data aggregation platform that provides real-time
// pricing, market data, historical charts, and more for thousands of cryptocurrencies.
//
// # Features
//
//   - Complete API coverage for all CoinGecko endpoints
//   - Support for both Free and Pro API tiers
//   - Type-safe API with strongly typed parameters
//   - Configurable HTTP client with custom options
//   - WebSocket support for real-time data
//   - Comprehensive error handling
//
// # Installation
//
//	go get github.com/tigusigalpa/coingecko-go
//
// # Quick Start
//
// Basic usage with the free API:
//
//	package main
//
//	import (
//	    "fmt"
//	    "log"
//	    coingecko "github.com/tigusigalpa/coingecko-go"
//	)
//
//	func main() {
//	    cg := coingecko.New()
//
//	    prices, err := cg.Simple().Price(
//	        []string{"bitcoin"},
//	        []string{"usd"},
//	        nil,
//	    )
//	    if err != nil {
//	        log.Fatal(err)
//	    }
//
//	    fmt.Printf("Bitcoin: $%.2f\n", prices["bitcoin"]["usd"])
//	}
//
// # Using Pro API
//
// To use the Pro API with authentication:
//
//	cg := coingecko.New(
//	    coingecko.WithAPIKey("your-api-key"),
//	    coingecko.WithProAPI(true),
//	)
//
// # Custom HTTP Client
//
// You can provide a custom HTTP client with specific settings:
//
//	import (
//	    "net/http"
//	    "time"
//	)
//
//	httpClient := &http.Client{
//	    Timeout: 60 * time.Second,
//	}
//
//	cg := coingecko.New(
//	    coingecko.WithHTTPClient(httpClient),
//	)
//
// # API Endpoints
//
// The client provides access to all CoinGecko API endpoints organized into logical groups:
//
//   - Simple API: Quick price lookups and supported currencies
//   - Coins API: Comprehensive coin data, markets, and historical charts
//   - Contract API: Token data by contract address
//   - Asset Platforms API: Blockchain platforms and token lists
//   - Categories API: Coin categories with market data
//   - Exchanges API: Exchange data, tickers, and volume charts
//   - Derivatives API: Derivatives market data
//   - Entities API: Crypto treasury holdings
//   - NFTs API: NFT collections and market data
//   - Exchange Rates API: BTC exchange rates
//   - Search API: Search and trending data
//   - Global API: Global market statistics
//   - WebSocket API: Real-time data subscriptions
//   - Ping API: Server status and API usage
//
// # Error Handling
//
// All API methods return an error as the last return value. Always check for errors:
//
//	data, err := cg.Coins().List(false)
//	if err != nil {
//	    log.Printf("API error: %v", err)
//	    return
//	}
//
// # Rate Limiting
//
// Be aware of CoinGecko's rate limits:
//   - Free API: 10-50 calls/minute
//   - Pro API: Higher limits based on your plan
//
// Implement caching to avoid hitting rate limits.
//
// # More Information
//
// For detailed documentation and examples, visit:
// https://github.com/tigusigalpa/coingecko-go
//
// CoinGecko API documentation:
// https://docs.coingecko.com
package coingecko
