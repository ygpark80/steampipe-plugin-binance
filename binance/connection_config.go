package binance

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type binanceConfig struct {
	ApiKey    *string `hcl:"api_key"`
	ApiSecret *string `hcl:"api_secret"`
}

func ConfigInstance() interface{} {
	return &binanceConfig{}
}

func GetConfig(connection *plugin.Connection) binanceConfig {
	if connection == nil || connection.Config == nil {
		return binanceConfig{}
	}
	config, _ := connection.Config.(binanceConfig)
	return config
}
