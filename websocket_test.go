package coingecko

import (
	"context"
	"encoding/json"
	"testing"
)

func TestWebSocketURL(t *testing.T) {
	if got, want := New().WebSocket().GetWebSocketURL(), "wss://stream.coingecko.com/v1"; got != want {
		t.Errorf("Expected URL %q, got %q", want, got)
	}
}

func TestCreateSimplePriceSubscription(t *testing.T) {
	message, err := New().WebSocket().CreateSimplePriceSubscription([]string{"bitcoin"}, []string{"usd"})
	if err != nil {
		t.Fatalf("CreateSimplePriceSubscription returned an error: %v", err)
	}

	decoded := decodeWebSocketMessage(t, message)
	if got, want := decoded.Command, "message"; got != want {
		t.Errorf("Expected command %q, got %q", want, got)
	}
	assertWebSocketChannel(t, decoded.Identifier, WebSocketChannelSimplePrice)

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(decoded.Data), &data); err != nil {
		t.Fatalf("Unable to decode message data: %v", err)
	}
	if got, want := data["action"], "set_tokens"; got != want {
		t.Errorf("Expected action %q, got %q", want, got)
	}
}

func TestCreateOnchainTradeSubscription(t *testing.T) {
	message, err := New().WebSocket().CreateOnchainTradeSubscription("eth", "0xabc")
	if err != nil {
		t.Fatalf("CreateOnchainTradeSubscription returned an error: %v", err)
	}

	decoded := decodeWebSocketMessage(t, message)
	assertWebSocketChannel(t, decoded.Identifier, WebSocketChannelOnchainTrade)

	var data struct {
		Pools  []string `json:"network_id:pool_addresses"`
		Action string   `json:"action"`
	}
	if err := json.Unmarshal([]byte(decoded.Data), &data); err != nil {
		t.Fatalf("Unable to decode message data: %v", err)
	}
	if len(data.Pools) != 1 || data.Pools[0] != "eth:0xabc" {
		t.Errorf("Expected network-qualified pool address, got %#v", data.Pools)
	}
	if data.Action != "set_pools" {
		t.Errorf("Expected set_pools action, got %q", data.Action)
	}
}

func TestCreateChannelSubscriptionAndUnsubscribe(t *testing.T) {
	api := New().WebSocket()
	for _, test := range []struct {
		name    string
		create  func(string) (string, error)
		command string
	}{
		{"subscribe", api.CreateChannelSubscription, "subscribe"},
		{"unsubscribe", api.CreateUnsubscribeMessage, "unsubscribe"},
	} {
		t.Run(test.name, func(t *testing.T) {
			message, err := test.create(WebSocketChannelSimplePrice)
			if err != nil {
				t.Fatalf("Message creation returned an error: %v", err)
			}
			decoded := decodeWebSocketMessage(t, message)
			if decoded.Command != test.command {
				t.Errorf("Expected command %q, got %q", test.command, decoded.Command)
			}
			assertWebSocketChannel(t, decoded.Identifier, WebSocketChannelSimplePrice)
		})
	}
}

func TestCreatePingMessageReturnsError(t *testing.T) {
	if _, err := New().WebSocket().CreatePingMessage(); err == nil {
		t.Fatal("Expected an error")
	}
}

func TestWebSocketConnectRequiresProAPIKey(t *testing.T) {
	_, _, err := New().WebSocket().Connect(context.Background())
	if err == nil {
		t.Fatal("Expected an error without Pro API configuration")
	}
}

func decodeWebSocketMessage(t *testing.T, message string) webSocketMessage {
	t.Helper()
	var decoded webSocketMessage
	if err := json.Unmarshal([]byte(message), &decoded); err != nil {
		t.Fatalf("Unable to decode WebSocket message: %v", err)
	}
	return decoded
}

func assertWebSocketChannel(t *testing.T, identifier string, expectedChannel string) {
	t.Helper()
	var channel struct {
		Channel string `json:"channel"`
	}
	if err := json.Unmarshal([]byte(identifier), &channel); err != nil {
		t.Fatalf("Unable to decode channel identifier: %v", err)
	}
	if channel.Channel != expectedChannel {
		t.Errorf("Expected channel %q, got %q", expectedChannel, channel.Channel)
	}
}
