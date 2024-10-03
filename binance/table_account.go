package binance

import (
	"context"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBinanceAccount(_ context.Context) *plugin.Table {
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
	clients, _ := connect(ctx, d)
	res, _ := clients.binance.NewGetAccountService().Do(ctx)

	for _, t := range res.Balances {
		d.StreamListItem(ctx, Balance{
			Asset:  getAsset(t.Asset),
			Type:   getType(t.Asset),
			Free:   t.Free,
			Locked: t.Locked,
		})
	}

	// TODO: should group by asset
	bswapLiquidity := clients.api.BswapLiquidity()
	for _, v := range bswapLiquidity {
		for k, vv := range v.Share.Asset {
			vvv, _ := strconv.ParseFloat(vv, 32)
			if vvv > 0 {
				d.StreamListItem(ctx, Balance{
					Asset:  k,
					Type:   "Liquidity",
					Free:   vv,
					Locked: "0",
				})
			}
		}
	}

	bswapUnclaimedRewardsResponse := clients.api.BswapUnclaimedRewards(BswapUnclaimedRewardsRequest{Type: 1})
	for k, v := range bswapUnclaimedRewardsResponse.TotalUnclaimedRewards {
		d.StreamListItem(ctx, Balance{
			Asset:  k,
			Type:   "Unclaimed Rewards",
			Free:   v,
			Locked: "0",
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
