package binance

import (
	"context"
	"errors"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func connect(_ context.Context, d *plugin.QueryData) (*binance.Client, error) {
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

	api := binance.NewClient(apiKey, apiSecret)
	return api, nil
}
