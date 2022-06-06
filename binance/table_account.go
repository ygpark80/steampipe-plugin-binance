package binance

import (
	"context"
	"strconv"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"golang.org/x/exp/maps"
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
			{Name: "symbol", Type: proto.ColumnType_STRING},
			{Name: "price", Type: proto.ColumnType_DOUBLE},
			{Name: "total", Type: proto.ColumnType_DOUBLE},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	clients, _ := connect(ctx, d)
	account, _ := clients.binance.NewGetAccountService().Do(ctx)
	bswapLiquidity := clients.api.BswapLiquidity()
	bswapUnclaimedRewardsResponse := clients.api.BswapUnclaimedRewards(BswapUnclaimedRewardsRequest{Type: 1})

	symbols := map[string]string{}
	for _, t := range account.Balances {
		free, _ := strconv.ParseFloat(t.Free, 32)
		if free > 0 {
			symbols[getAsset(t.Asset)] = ""
		}
	}
	for _, v := range bswapLiquidity {
		for k, vv := range v.Share.Asset {
			vvv, _ := strconv.ParseFloat(vv, 32)
			if vvv > 0 {
				symbols[k] = ""
			}
		}
	}
	for k := range bswapUnclaimedRewardsResponse.TotalUnclaimedRewards {
		symbols[k] = ""
	}
	delete(symbols, "BUSD")
	delete(symbols, "USDT")

	exchangeInfo, _ := clients.binance.NewExchangeInfoService().Do(ctx)
	for asset := range symbols {
		orders := []string{"USDT", "BUSD", "BTC"}
		for _, order := range orders {
			if symbols[asset] != "" {
				break
			}

			symbol := asset + order
			for _, v := range exchangeInfo.Symbols {
				if v.Symbol == symbol {
					symbols[asset] = symbol
					break
				}
			}
		}

		if symbols[asset] == "" {
			for _, v := range exchangeInfo.Symbols {
				if strings.HasPrefix(v.Symbol, asset) {
					symbols[asset] = v.Symbol
					break
				}
			}
		}
	}

	prices := map[string]string{}
	logger.Warn("listAccount", "len(symbols)=", len(removeEmptyStrings(maps.Values(symbols))))
	tickerPrices := clients.api.TickerPrice(removeEmptyStrings(maps.Values(symbols)))
	for _, v := range tickerPrices {
		prices[v.Symbol] = v.Price
	}

	for _, t := range account.Balances {
		asset := getAsset(t.Asset)

		d.StreamListItem(ctx, Balance{
			Asset:  getAsset(t.Asset),
			Type:   getType(t.Asset),
			Free:   t.Free,
			Locked: t.Locked,
			Symbol: symbols[asset],
			Price:  getPrice(asset, symbols, prices),
			Total:  multiply(t.Free, getPrice(asset, symbols, prices)),
		})
	}

	// TODO: should group by asset
	for _, v := range bswapLiquidity {
		for k, vv := range v.Share.Asset {
			vvv, _ := strconv.ParseFloat(vv, 32)
			if vvv > 0 {
				asset := k
				d.StreamListItem(ctx, Balance{
					Asset:  k,
					Type:   "Liquidity",
					Free:   vv,
					Locked: "0",
					Symbol: symbols[asset],
					Price:  getPrice(asset, symbols, prices),
					Total:  multiply(vv, getPrice(asset, symbols, prices)),
				})
			}
		}
	}

	for k, v := range bswapUnclaimedRewardsResponse.TotalUnclaimedRewards {
		asset := k

		d.StreamListItem(ctx, Balance{
			Asset:  k,
			Type:   "Unclaimed Rewards",
			Free:   v,
			Locked: "0",
			Symbol: symbols[asset],
			Price:  getPrice(asset, symbols, prices),
			Total:  multiply(v, getPrice(asset, symbols, prices)),
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

func getPrice(asset string, symbols map[string]string, prices map[string]string) string {
	if asset == "USDT" || asset == "BUSD" {
		return "1"
	}
	price := prices[symbols[asset]]
	if price == "" {
		return "0"
	}
	return prices[symbols[asset]]
}

func multiply(free string, price string) float64 {
	if free == "" {
		free = "0"
	}
	if price == "" {
		price = "0"
	}
	_free, _ := strconv.ParseFloat(free, 32)
	_price, _ := strconv.ParseFloat(price, 32)

	return _free * _price
}

type Balance struct {
	Asset  string
	Type   string
	Free   string
	Locked string
	Symbol string
	Price  string
	Total  float64
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}
