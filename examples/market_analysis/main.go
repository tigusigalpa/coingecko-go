package main

import (
	"fmt"
	"log"

	coingecko "github.com/tigusigalpa/coingecko-go"
)

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func main() {
	cg := coingecko.New()

	fmt.Println("=== Crypto Market Analysis ===\n")

	fmt.Println("=== Top 10 Cryptocurrencies by Market Cap ===")
	topCoins, err := cg.Coins().Markets("usd", &coingecko.MarketsOptions{
		Order:                 stringPtr("market_cap_desc"),
		PerPage:               intPtr(10),
		Page:                  intPtr(1),
		PriceChangePercentage: stringPtr("1h,24h,7d"),
	})
	if err != nil {
		log.Fatal(err)
	}

	for i, coin := range topCoins {
		name := coin["name"].(string)
		symbol := coin["symbol"].(string)
		price := coin["current_price"].(float64)
		marketCap := coin["market_cap"].(float64)
		volume := coin["total_volume"].(float64)

		change24h := 0.0
		if ch, ok := coin["price_change_percentage_24h"]; ok && ch != nil {
			change24h = ch.(float64)
		}

		fmt.Printf("%d. %s (%s)\n", i+1, name, symbol)
		fmt.Printf("   Price: $%.2f | Market Cap: $%.0f | Volume: $%.0f\n", price, marketCap, volume)
		fmt.Printf("   24h Change: %.2f%%\n\n", change24h)
	}

	fmt.Println("=== Top Gainers & Losers (24h) ===")
	duration := "24h"
	topCoinsCount := 100
	movers, err := cg.Coins().TopGainersLosers("usd", &duration, &topCoinsCount)
	if err != nil {
		log.Fatal(err)
	}

	if gainers, ok := movers["top_gainers"].([]interface{}); ok && len(gainers) > 0 {
		fmt.Println("\nTop 5 Gainers:")
		for i := 0; i < 5 && i < len(gainers); i++ {
			gainer := gainers[i].(map[string]interface{})
			fmt.Printf("%d. %s: %.2f%%\n", i+1, gainer["name"], gainer["usd_24h_change"])
		}
	}

	if losers, ok := movers["top_losers"].([]interface{}); ok && len(losers) > 0 {
		fmt.Println("\nTop 5 Losers:")
		for i := 0; i < 5 && i < len(losers); i++ {
			loser := losers[i].(map[string]interface{})
			fmt.Printf("%d. %s: %.2f%%\n", i+1, loser["name"], loser["usd_24h_change"])
		}
	}

	fmt.Println("\n=== Global Market Data ===")
	globalData, err := cg.Global().Crypto()
	if err != nil {
		log.Fatal(err)
	}

	if data, ok := globalData["data"].(map[string]interface{}); ok {
		if totalMarketCap, ok := data["total_market_cap"].(map[string]interface{}); ok {
			fmt.Printf("Total Market Cap: $%.0f\n", totalMarketCap["usd"])
		}
		if totalVolume, ok := data["total_volume"].(map[string]interface{}); ok {
			fmt.Printf("Total 24h Volume: $%.0f\n", totalVolume["usd"])
		}
		if btcDominance, ok := data["market_cap_percentage"].(map[string]interface{}); ok {
			fmt.Printf("Bitcoin Dominance: %.2f%%\n", btcDominance["btc"])
		}
	}
}
