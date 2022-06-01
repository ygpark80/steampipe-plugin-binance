package binance

import (
	"context"
	"strings"

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
			{Name: "type", Type: proto.ColumnType_STRING},
			{Name: "free", Type: proto.ColumnType_DOUBLE},
			{Name: "locked", Type: proto.ColumnType_DOUBLE},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, _ := connect(ctx, d)
	res, _ := client.NewGetAccountService().Do(ctx)

	for _, t := range res.Balances {
		d.StreamListItem(ctx, Balance{
			Asset:  getAsset(t.Asset),
			Type:   getType(t.Asset),
			Free:   t.Free,
			Locked: t.Locked,
		})
	}

	return nil, nil
}

func getAsset(asset string) string {
	if asset == "LDBAKE" {
		return "QTUM"
	}
	if asset == "LDBAKET" {
		return "BAKE"
	}
	if asset == "LDERD" {
		return "EGLD"
	}
	if asset == "LDLEND" {
		return "TKO"
	}
	if asset == "LDSHIB2" {
		return "SHIB"
	}
	if getType(asset) == "Flexible" {
		return asset[2:]
	}
	return asset
}

func getType(asset string) string {
	if strings.HasPrefix(asset, "LD") && asset != "LDO" {
		return "Flexible"
	}
	return "Spot"
}

type Balance struct {
	Asset  string
	Type   string
	Free   string
	Locked string
}
