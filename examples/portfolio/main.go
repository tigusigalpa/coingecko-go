package main

import (
	"fmt"
	"log"

	coingecko "github.com/tigusigalpa/coingecko-go"
)

// Portfolio contains cryptocurrency holdings keyed by coin ID.
type Portfolio struct {
	Holdings map[string]float64
}

// CalculateValue returns the portfolio's current value in USD.
func (p *Portfolio) CalculateValue(cg *coingecko.CoinGecko) (float64, error) {
	coinIDs := make([]string, 0, len(p.Holdings))
	for coinID := range p.Holdings {
		coinIDs = append(coinIDs, coinID)
	}

	prices, err := cg.Simple().Price(coinIDs, []string{"usd"}, &coingecko.SimplePriceOptions{
		IncludeMarketCap:  true,
		Include24hrChange: true,
	})
	if err != nil {
		return 0, err
	}

	totalValue := 0.0
	fmt.Println("\n=== Portfolio Breakdown ===")
	for coinID, amount := range p.Holdings {
		price := prices[coinID]["usd"].(float64)
		value := price * amount
		totalValue += value

		change24h := 0.0
		if ch, ok := prices[coinID]["usd_24h_change"]; ok {
			change24h = ch.(float64)
		}

		fmt.Printf("%s: %.4f coins × $%.2f = $%.2f (24h: %.2f%%)\n",
			coinID, amount, price, value, change24h)
	}

	return totalValue, nil
}

func main() {
	cg := coingecko.New()

	portfolio := Portfolio{
		Holdings: map[string]float64{
			"bitcoin":  0.5,
			"ethereum": 10,
			"cardano":  1000,
			"solana":   50,
		},
	}

	fmt.Println("=== Crypto Portfolio Tracker ===")

	totalValue, err := portfolio.CalculateValue(cg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n=== Total Portfolio Value: $%.2f ===\n", totalValue)
}
