package binance

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableBinanceAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "binance_account",
		List: &plugin.ListConfig{
			Hydrate: listAccount,
		},
		Columns: []*plugin.Column{
			{Name: "asset", Type: proto.ColumnType_STRING},
			{Name: "free", Type: proto.ColumnType_STRING},
			{Name: "locked", Type: proto.ColumnType_STRING},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, _ := connect(ctx, d)
	res, _ := client.NewGetAccountService().Do(ctx)

	for _, t := range res.Balances {
		d.StreamListItem(ctx, t)
	}

	return nil, nil
}
