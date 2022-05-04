package binance

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type binanceConfig struct {
	ApiKey    *string `cty:"api_key"`
	ApiSecret *string `cty:"api_secret"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"api_key":    {Type: schema.TypeString},
	"api_secret": {Type: schema.TypeString},
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
