package binance

import (
	"context"
	"errors"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type Clients struct {
	binance *binance.Client
	api     *Client
}

func connect(_ context.Context, d *plugin.QueryData) (*Clients, error) {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_API_SECRET")

	binanceConfig := GetConfig(d.Connection)
	if binanceConfig.ApiKey != nil {
		apiKey = *binanceConfig.ApiKey
	}
	if binanceConfig.ApiSecret != nil {
		apiSecret = *binanceConfig.ApiSecret
	}

	if apiKey == "" || apiSecret == "" {
		return nil, errors.New("'api_key' and 'api_secret' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	binance := binance.NewClient(apiKey, apiSecret)
	api := NewClient(apiKey, apiSecret)

	return &Clients{binance: binance, api: api}, nil
}
