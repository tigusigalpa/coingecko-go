package main

import (
	"fmt"
	"log"

	coingecko "github.com/tigusigalpa/coingecko-go"
)

func main() {
	cg := coingecko.New()

	fmt.Println("=== CoinGecko Go Client - Basic Example ===\n")

	prices, err := cg.Simple().Price(
		[]string{"bitcoin", "ethereum"},
		[]string{"usd", "eur"},
		&coingecko.SimplePriceOptions{
			IncludeMarketCap:  true,
			Include24hrVol:    true,
			Include24hrChange: true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current Prices:")
	for coin, data := range prices {
		fmt.Printf("\n%s:\n", coin)
		for currency, value := range data {
			fmt.Printf("  %s: %v\n", currency, value)
		}
	}

	fmt.Println("\n=== Supported Currencies ===")
	currencies, err := cg.Simple().SupportedVsCurrencies()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total supported currencies: %d\n", len(currencies))
	fmt.Printf("First 10: %v\n", currencies[:10])

	fmt.Println("\n=== Server Status ===")
	status, err := cg.Ping().Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server says: %s\n", status.GeckoSays)
}
