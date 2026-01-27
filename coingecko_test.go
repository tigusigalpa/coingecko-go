package coingecko

import "testing"

func TestNew(t *testing.T) {
	cg := New()

	if cg == nil {
		t.Fatal("Expected CoinGecko instance to be created")
	}

	if cg.client == nil {
		t.Fatal("Expected client to be initialized")
	}
}

func TestNewWithOptions(t *testing.T) {
	apiKey := "test-key"
	cg := New(
		WithAPIKey(apiKey),
		WithProAPI(true),
	)

	if cg.client.apiKey != apiKey {
		t.Errorf("Expected apiKey to be %s, got %s", apiKey, cg.client.apiKey)
	}

	if !cg.client.isPro {
		t.Error("Expected isPro to be true")
	}
}

func TestAPIAccessors(t *testing.T) {
	cg := New()

	if cg.Simple() == nil {
		t.Error("Expected Simple API to be accessible")
	}

	if cg.Coins() == nil {
		t.Error("Expected Coins API to be accessible")
	}

	if cg.Contract() == nil {
		t.Error("Expected Contract API to be accessible")
	}

	if cg.AssetPlatforms() == nil {
		t.Error("Expected AssetPlatforms API to be accessible")
	}

	if cg.Categories() == nil {
		t.Error("Expected Categories API to be accessible")
	}

	if cg.Exchanges() == nil {
		t.Error("Expected Exchanges API to be accessible")
	}

	if cg.Derivatives() == nil {
		t.Error("Expected Derivatives API to be accessible")
	}

	if cg.Entities() == nil {
		t.Error("Expected Entities API to be accessible")
	}

	if cg.NFTs() == nil {
		t.Error("Expected NFTs API to be accessible")
	}

	if cg.ExchangeRates() == nil {
		t.Error("Expected ExchangeRates API to be accessible")
	}

	if cg.Search() == nil {
		t.Error("Expected Search API to be accessible")
	}

	if cg.Global() == nil {
		t.Error("Expected Global API to be accessible")
	}

	if cg.WebSocket() == nil {
		t.Error("Expected WebSocket API to be accessible")
	}

	if cg.Ping() == nil {
		t.Error("Expected Ping API to be accessible")
	}
}
