package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/ygpark80/steampipe-plugin-binance/binance"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: binance.Plugin})
}
