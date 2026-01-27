package main

import (
	"fmt"
	"log"

	coingecko "github.com/tigusigalpa/coingecko-go"
)

func trackToken(cg *coingecko.CoinGecko, platform, contractAddress string) {
	fmt.Printf("\n=== Tracking Token on %s ===\n", platform)
	fmt.Printf("Contract: %s\n\n", contractAddress)

	tokenData, err := cg.Contract().Coin(platform, contractAddress)
	if err != nil {
		log.Printf("Error fetching token data: %v\n", err)
		return
	}

	if name, ok := tokenData["name"].(string); ok {
		fmt.Printf("Name: %s\n", name)
	}
	if symbol, ok := tokenData["symbol"].(string); ok {
		fmt.Printf("Symbol: %s\n", symbol)
	}

	if marketData, ok := tokenData["market_data"].(map[string]interface{}); ok {
		if currentPrice, ok := marketData["current_price"].(map[string]interface{}); ok {
			if usdPrice, ok := currentPrice["usd"].(float64); ok {
				fmt.Printf("Current Price: $%.6f\n", usdPrice)
			}
		}

		if marketCap, ok := marketData["market_cap"].(map[string]interface{}); ok {
			if usdMarketCap, ok := marketCap["usd"].(float64); ok {
				fmt.Printf("Market Cap: $%.2f\n", usdMarketCap)
			}
		}

		if totalVolume, ok := marketData["total_volume"].(map[string]interface{}); ok {
			if usdVolume, ok := totalVolume["usd"].(float64); ok {
				fmt.Printf("24h Volume: $%.2f\n", usdVolume)
			}
		}

		if priceChange24h, ok := marketData["price_change_percentage_24h"].(float64); ok {
			fmt.Printf("24h Change: %.2f%%\n", priceChange24h)
		}
	}

	fmt.Println("\n=== 7-Day Price Chart ===")
	chartData, err := cg.Contract().MarketChart(platform, contractAddress, "usd", "7")
	if err != nil {
		log.Printf("Error fetching chart data: %v\n", err)
		return
	}

	if prices, ok := chartData["prices"].([]interface{}); ok {
		fmt.Printf("Data points: %d\n", len(prices))
		if len(prices) > 0 {
			firstPrice := prices[0].([]interface{})
			lastPrice := prices[len(prices)-1].([]interface{})
			fmt.Printf("First: $%.6f\n", firstPrice[1].(float64))
			fmt.Printf("Last: $%.6f\n", lastPrice[1].(float64))
		}
	}
}

func main() {
	cg := coingecko.New()

	fmt.Println("=== DeFi Token Tracker ===")

	trackToken(cg, "ethereum", "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984")

	fmt.Println("\n=== DeFi Category Overview ===")
	category := "decentralized-finance-defi"
	perPage := 10
	page := 1

	defiCoins, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
		Category: &category,
		PerPage:  &perPage,
		Page:     &page,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nTop %d DeFi Coins:\n", len(defiCoins))
	for i, coin := range defiCoins {
		name := coin["name"].(string)
		price := coin["current_price"].(float64)
		marketCap := coin["market_cap"].(float64)
		fmt.Printf("%d. %s: $%.2f (Market Cap: $%.0f)\n", i+1, name, price, marketCap)
	}
}
