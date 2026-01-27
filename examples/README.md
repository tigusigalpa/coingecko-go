# CoinGecko Go Client - Examples

This directory contains example applications demonstrating various use cases of the CoinGecko Go client.

## Running Examples

### Basic Example

Demonstrates basic API usage including price lookups and server status.

```bash
cd examples/basic
go run main.go
```

### Portfolio Tracker

Shows how to track a cryptocurrency portfolio and calculate its total value.

```bash
cd examples/portfolio
go run main.go
```

### Market Analysis

Demonstrates market analysis features including top coins, gainers/losers, and global market data.

```bash
cd examples/market_analysis
go run main.go
```

### DeFi Tracker

Shows how to track DeFi tokens by contract address and analyze DeFi categories.

```bash
cd examples/defi_tracker
go run main.go
```

## Using with Pro API

To use these examples with a Pro API key, set the environment variable:

```bash
export COINGECKO_API_KEY="your-api-key-here"
```

Then modify the example code to use:

```go
apiKey := os.Getenv("COINGECKO_API_KEY")
cg := coingecko.New(
    coingecko.WithAPIKey(apiKey),
    coingecko.WithProAPI(true),
)
```

## Building All Examples

From the root directory:

```bash
make examples
```

## Notes

- Examples use the free API by default (no API key required)
- Rate limits apply - add delays between requests if needed
- Some endpoints may require Pro API access
