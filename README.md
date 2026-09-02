# CoinGecko Go Client

![CoinGecko Golang SDK](https://i.postimg.cc/rpL6MPf7/coingecko-golang-hero-github.jpg)

[![CI](https://github.com/tigusigalpa/coingecko-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/coingecko-go/actions/workflows/ci.yml)
[![Tests](https://github.com/tigusigalpa/coingecko-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/coingecko-go/actions/workflows/test.yml)
[![Codecov](https://codecov.io/gh/tigusigalpa/coingecko-go/graph/badge.svg)](https://codecov.io/gh/tigusigalpa/coingecko-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/tigusigalpa/coingecko-go.svg)](https://pkg.go.dev/github.com/tigusigalpa/coingecko-go)
[![Go version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)

A small, practical Go client for the [CoinGecko API](https://docs.coingecko.com). Use it to fetch prices, markets, charts, exchanges, NFTs, treasury data, and to connect to CoinGecko's live WebSocket streams—without hand-building URLs or authentication headers.

> The package intentionally stays close to CoinGecko's API. Request options are typed; where API response schemas are broad or evolve frequently, responses are returned as `map[string]interface{}`.

**Documentation:** [pkg.go.dev/github.com/tigusigalpa/coingecko-go](https://pkg.go.dev/github.com/tigusigalpa/coingecko-go)

## Contents

- [Install](#install)
- [Quick start](#quick-start)
- [Authentication](#authentication)
- [Common tasks](#common-tasks)
- [WebSocket streams](#websocket-streams)
- [Endpoint groups](#endpoint-groups)
- [Errors, timeouts, and rate limits](#errors-timeouts-and-rate-limits)
- [Examples and development](#examples-and-development)

## Install

```bash
go get github.com/tigusigalpa/coingecko-go
```

```go
import coingecko "github.com/tigusigalpa/coingecko-go"
```

Requirements:

- Go 1.21 or newer
- A CoinGecko Demo or Pro API key for authenticated access

## Quick start

Fetch the current price of Bitcoin and Ethereum in USD. A nil options value keeps the request minimal.

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
		[]string{"bitcoin", "ethereum"},
		[]string{"usd"},
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("BTC: $%.2f\n", prices["bitcoin"]["usd"])
	fmt.Printf("ETH: $%.2f\n", prices["ethereum"]["usd"])
}
```

## Authentication

The client selects the correct authentication header automatically:

- `WithAPIKey` uses the Demo API endpoint and its Demo key header.
- Adding `WithProAPI(true)` switches to the Pro endpoint and its Pro key header.

Keep keys outside source control. Environment variables are the simplest option for local apps and services.

```go
apiKey := os.Getenv("COINGECKO_API_KEY")

// Demo API
demo := coingecko.New(coingecko.WithAPIKey(apiKey))

// Pro API
pro := coingecko.New(
	coingecko.WithAPIKey(apiKey),
	coingecko.WithProAPI(true),
)
```

For the available plans and their endpoint access, see the [official CoinGecko documentation](https://docs.coingecko.com).

## Common tasks

### Request price extras

```go
prices, err := cg.Simple().Price(
	[]string{"bitcoin", "ethereum"},
	[]string{"usd", "eur"},
	&coingecko.SimplePriceOptions{
		IncludeMarketCap:      true,
		Include24hrVol:        true,
		Include24hrChange:     true,
		IncludeLastUpdatedAt:  true,
	},
)
```

### Browse market leaders

Optional API parameters are pointers so you can omit them precisely.

```go
order := "market_cap_desc"
perPage := 25
page := 1

markets, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
	Order:    &order,
	PerPage:  &perPage,
	Page:     &page,
	Sparkline: true,
})
```

### Look up a token contract

```go
token, err := cg.Contract().Coin(
	"ethereum",
	"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984", // UNI
)
```

### Use a custom HTTP client

Bring your own transport when you need a proxy, custom certificates, retries, or connection-pool settings.

```go
httpClient := &http.Client{
	Timeout: 20 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

cg := coingecko.New(
	coingecko.WithHTTPClient(httpClient),
	coingecko.WithAPIKey(os.Getenv("COINGECKO_API_KEY")),
)
```

`WithTimeout` is a convenient alternative when you only need to set a timeout:

```go
cg := coingecko.New(coingecko.WithTimeout(15 * time.Second))
```

## WebSocket streams

WebSocket access requires a Pro API key and `WithProAPI(true)`. The WebSocket protocol has two steps: subscribe to a channel, then send the channel-specific configuration message.

```go
package main

import (
	"context"
	"log"
	"os"

	coingecko "github.com/tigusigalpa/coingecko-go"
	"github.com/gorilla/websocket"
)

func main() {
	cg := coingecko.New(
		coingecko.WithAPIKey(os.Getenv("COINGECKO_API_KEY")),
		coingecko.WithProAPI(true),
	)

	conn, _, err := cg.WebSocket().Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	subscribe, err := cg.WebSocket().CreateChannelSubscription(coingecko.WebSocketChannelSimplePrice)
	if err != nil {
		log.Fatal(err)
	}
	configure, err := cg.WebSocket().CreateSimplePriceSubscription([]string{"bitcoin"}, []string{"usd"})
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.WriteMessage(websocket.TextMessage, []byte(subscribe)); err != nil {
		log.Fatal(err)
	}
	if err := conn.WriteMessage(websocket.TextMessage, []byte(configure)); err != nil {
		log.Fatal(err)
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("update: %s", message)
	}
}
```

Available channel constants:

- `WebSocketChannelSimplePrice`
- `WebSocketChannelOnchainSimpleTokenPrice`
- `WebSocketChannelOnchainTrade`
- `WebSocketChannelOnchainOHLCV`

For long-running applications, handle disconnects and reconnect with backoff. The server uses normal WebSocket ping/pong control frames; do not send a JSON `ping` message.

## Endpoint groups

The API is organised by domain, so calls are easy to discover from `cg`:

| Group | Typical use |
| --- | --- |
| `Simple()` | Current coin and token prices; supported currencies |
| `Coins()` | Coin lists, markets, metadata, tickers, history, charts, OHLC, supply charts |
| `Contract()` | ERC-20 and other token data by platform and contract address |
| `AssetPlatforms()` | Supported platforms and token lists |
| `Categories()` | Coin categories and category market data |
| `Exchanges()` | Exchange lists, tickers, and volume charts |
| `Derivatives()` | Derivatives tickers and exchanges |
| `Entities()` | Public-company and entity crypto treasury data |
| `NFTs()` | NFT collections, markets, tickers, and charts |
| `ExchangeRates()` | BTC exchange rates |
| `Global()` | Global crypto and DeFi market statistics |
| `Search()` | Search and trending results |
| `Ping()` | Service status and Pro API key usage |
| `WebSocket()` | Live price, onchain token, trade, and OHLCV stream helpers |

The complete method-level reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/tigusigalpa/coingecko-go).

## Errors, timeouts, and rate limits

Every API method returns an error. HTTP responses outside the 2xx range are returned as `*coingecko.APIError`, which exposes the status code and API response body.

```go
prices, err := cg.Simple().Price([]string{"bitcoin"}, []string{"usd"}, nil)
if err != nil {
	var apiErr *coingecko.APIError
	if errors.As(err, &apiErr) {
		log.Printf("CoinGecko returned %d: %s", apiErr.StatusCode, apiErr.Body)
		return
	}
	log.Printf("request failed: %v", err)
	return
}
```

CoinGecko rate limits and endpoint availability vary by plan. Treat 429 responses as a signal to slow down, cache frequently requested data, and use retries with exponential backoff where appropriate. Avoid logging API keys or placing them in client-side applications.

## Examples and development

Runnable examples live in [`examples/`](examples/):

- `basic` — prices and API status
- `portfolio` — portfolio value calculation
- `market_analysis` — market leaders and global data
- `defi_tracker` — token contracts and DeFi categories

```bash
cd examples/basic && go run .
cd examples/portfolio && go run .
cd examples/market_analysis && go run .
cd examples/defi_tracker && go run .
```

To verify the library locally:

```bash
go test ./...
go test -race -covermode=atomic -coverprofile=coverage.out ./...
go vet ./...
```

The repository checks formatting, static analysis, tests, security scanning, and Codecov coverage in CI. Contributions are welcome—please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.

## Security

If you believe you have found a security vulnerability, please email **sovletig@gmail.com** instead of filing a public issue.

## License

Released under the [MIT License](LICENSE).

## Links

- [CoinGecko API documentation](https://docs.coingecko.com)
- [CoinGecko API plans](https://www.coingecko.com/en/api)
- [Changelog](CHANGELOG.md)
- [Issues](https://github.com/tigusigalpa/coingecko-go/issues)
