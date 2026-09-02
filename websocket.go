package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketAPI creates connections and protocol messages for CoinGecko streams.
type WebSocketAPI struct {
	client *Client
}

// WebSocketURLPro and WebSocketChannel constants identify the CoinGecko WebSocket service and its channels.
const (
	WebSocketURLPro = "wss://stream.coingecko.com/v1"

	WebSocketChannelSimplePrice             = "CGSimplePrice"
	WebSocketChannelOnchainSimpleTokenPrice = "OnchainSimpleTokenPrice"
	WebSocketChannelOnchainTrade            = "OnchainTrade"
	WebSocketChannelOnchainOHLCV            = "OnchainOHLCV"
)

// WebSocketSubscription is retained for backwards compatibility.
// Use the Create* methods to generate messages accepted by the API.
type WebSocketSubscription struct {
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type webSocketMessage struct {
	Command    string `json:"command"`
	Identifier string `json:"identifier"`
	Data       string `json:"data,omitempty"`
}

// GetWebSocketURL returns the Pro WebSocket endpoint URL.
func (w *WebSocketAPI) GetWebSocketURL() string {
	return WebSocketURLPro
}

// Connect opens an authenticated connection to the CoinGecko WebSocket API.
// WebSocket access requires a Pro API key and WithProAPI(true).
func (w *WebSocketAPI) Connect(ctx context.Context) (*websocket.Conn, *http.Response, error) {
	if !w.client.isPro {
		return nil, nil, fmt.Errorf("CoinGecko WebSocket requires WithProAPI(true)")
	}
	if w.client.apiKey == "" {
		return nil, nil, fmt.Errorf("CoinGecko WebSocket requires an API key")
	}

	headers := http.Header{}
	headers.Set("x-cg-pro-api-key", w.client.apiKey)
	return websocket.DefaultDialer.DialContext(ctx, WebSocketURLPro, headers)
}

// CreateChannelSubscription creates the first message required before streaming
// data from a channel.
func (w *WebSocketAPI) CreateChannelSubscription(channel string) (string, error) {
	return createWebSocketMessage("subscribe", channel, nil)
}

// CreateSimplePriceSubscription creates a price-stream configuration message.
func (w *WebSocketAPI) CreateSimplePriceSubscription(coinIDs []string, vsCurrencies []string) (string, error) {
	return createWebSocketMessage("message", WebSocketChannelSimplePrice, map[string]interface{}{
		"coin_id":       coinIDs,
		"vs_currencies": vsCurrencies,
		"action":        "set_tokens",
	})
}

// CreateOnchainSimpleTokenPriceSubscription creates a message for USD token
// prices. The third argument is retained for backwards compatibility; the
// WebSocket API does not support selecting quote currencies for this channel.
func (w *WebSocketAPI) CreateOnchainSimpleTokenPriceSubscription(network string, addresses []string, _ []string) (string, error) {
	return createWebSocketMessage("message", WebSocketChannelOnchainSimpleTokenPrice, map[string]interface{}{
		"network_id:token_addresses": prefixNetwork(network, addresses),
		"action":                     "set_tokens",
	})
}

// CreateOnchainTradeSubscription creates an onchain trade-stream configuration message.
func (w *WebSocketAPI) CreateOnchainTradeSubscription(network string, poolAddress string) (string, error) {
	return createWebSocketMessage("message", WebSocketChannelOnchainTrade, map[string]interface{}{
		"network_id:pool_addresses": []string{network + ":" + poolAddress},
		"action":                    "set_pools",
	})
}

// CreateOnchainOHLCVSubscription creates an onchain OHLCV-stream configuration message.
func (w *WebSocketAPI) CreateOnchainOHLCVSubscription(network string, poolAddress string, timeframe string) (string, error) {
	return createWebSocketMessage("message", WebSocketChannelOnchainOHLCV, map[string]interface{}{
		"network_id:pool_addresses": []string{network + ":" + poolAddress},
		"interval":                  timeframe,
		"token":                     "base",
		"action":                    "set_pools",
	})
}

// CreateUnsubscribeMessage creates a message that removes a channel subscription.
func (w *WebSocketAPI) CreateUnsubscribeMessage(channel string) (string, error) {
	return createWebSocketMessage("unsubscribe", channel, nil)
}

// CreatePingMessage is deprecated. WebSocket ping control frames must be sent
// by the WebSocket client, not as a JSON protocol message.
func (w *WebSocketAPI) CreatePingMessage() (string, error) {
	return "", fmt.Errorf("WebSocket ping must be sent as a control frame")
}

func createWebSocketMessage(command string, channel string, data interface{}) (string, error) {
	identifier, err := json.Marshal(map[string]string{"channel": channel})
	if err != nil {
		return "", fmt.Errorf("marshal WebSocket channel identifier: %w", err)
	}

	message := webSocketMessage{
		Command:    command,
		Identifier: string(identifier),
	}
	if data != nil {
		encodedData, err := json.Marshal(data)
		if err != nil {
			return "", fmt.Errorf("marshal WebSocket message data: %w", err)
		}
		message.Data = string(encodedData)
	}

	encodedMessage, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("marshal WebSocket message: %w", err)
	}
	return string(encodedMessage), nil
}

func prefixNetwork(network string, addresses []string) []string {
	values := make([]string, len(addresses))
	for i, address := range addresses {
		values[i] = network + ":" + address
	}
	return values
}
