package binance

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-binance",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo(),
		DefaultConcurrency: &plugin.DefaultConcurrencyConfig{
			TotalMaxConcurrency: 10,
		},
		TableMap: map[string]*plugin.Table{
			"binance_account": tableBinanceAccount(ctx),
			"binance_staking": tableBinanceStaking(ctx),
		},
	}
	return p
}
