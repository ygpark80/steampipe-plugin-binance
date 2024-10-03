package binance

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBinanceExchange(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "binance_exchange",
		List: &plugin.ListConfig{
			Hydrate: listExchange,
		},
		Columns: []*plugin.Column{
			{Name: "symbol", Type: proto.ColumnType_STRING},
			{Name: "status", Type: proto.ColumnType_STRING},
		},
	}
}

func listExchange(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	clients, _ := connect(ctx, d)
	res, _ := clients.binance.NewExchangeInfoService().Do(ctx)

	for _, t := range res.Symbols {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}
