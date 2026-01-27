package coingecko

import "encoding/json"

type WebSocketAPI struct {
	client *Client
}

const (
	WebSocketURLPro = "wss://ws-pro.coingecko.com/v1"
)

func (w *WebSocketAPI) GetWebSocketURL() string {
	return WebSocketURLPro
}

type WebSocketSubscription struct {
	Type   string                 `json:"type"`
	Params map[string]interface{} `json:"params,omitempty"`
}

func (w *WebSocketAPI) CreateSimplePriceSubscription(coinIDs []string, vsCurrencies []string) (string, error) {
	subscription := WebSocketSubscription{
		Type: "subscribe",
		Params: map[string]interface{}{
			"channel": "cg_simple_price",
			"coin_ids": coinIDs,
			"vs_currencies": vsCurrencies,
		},
	}

	data, err := json.Marshal(subscription)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (w *WebSocketAPI) CreateOnchainSimpleTokenPriceSubscription(network string, addresses []string, vsCurrencies []string) (string, error) {
	subscription := WebSocketSubscription{
		Type: "subscribe",
		Params: map[string]interface{}{
			"channel": "cg_onchain_simple_token_price",
			"network": network,
			"addresses": addresses,
			"vs_currencies": vsCurrencies,
		},
	}

	data, err := json.Marshal(subscription)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (w *WebSocketAPI) CreateOnchainTradeSubscription(network string, poolAddress string) (string, error) {
	subscription := WebSocketSubscription{
		Type: "subscribe",
		Params: map[string]interface{}{
			"channel": "cg_onchain_trade",
			"network": network,
			"pool_address": poolAddress,
		},
	}

	data, err := json.Marshal(subscription)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (w *WebSocketAPI) CreateOnchainOHLCVSubscription(network string, poolAddress string, timeframe string) (string, error) {
	subscription := WebSocketSubscription{
		Type: "subscribe",
		Params: map[string]interface{}{
			"channel": "cg_onchain_ohlcv",
			"network": network,
			"pool_address": poolAddress,
			"timeframe": timeframe,
		},
	}

	data, err := json.Marshal(subscription)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (w *WebSocketAPI) CreateUnsubscribeMessage(channel string) (string, error) {
	subscription := WebSocketSubscription{
		Type: "unsubscribe",
		Params: map[string]interface{}{
			"channel": channel,
		},
	}

	data, err := json.Marshal(subscription)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (w *WebSocketAPI) CreatePingMessage() (string, error) {
	ping := map[string]string{
		"type": "ping",
	}

	data, err := json.Marshal(ping)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
