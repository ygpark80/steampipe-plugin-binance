package binance

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableBinanceStaking(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "binance_staking",
		Description: "Binance Staking, dedicated to increasing user staking income",
		List: &plugin.ListConfig{
			Hydrate: listStaking,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromCamel()},
			{Name: "project_id", Type: proto.ColumnType_STRING, Transform: transform.FromCamel()},
			{Name: "asset", Type: proto.ColumnType_STRING},
			{Name: "up_limit", Type: proto.ColumnType_STRING},
			{Name: "purchased", Type: proto.ColumnType_STRING},
			{Name: "end_time", Type: proto.ColumnType_INT},
			{Name: "issue_start_time", Type: proto.ColumnType_INT},
			{Name: "issue_end_time", Type: proto.ColumnType_INT},
			{Name: "duration", Type: proto.ColumnType_STRING},
			{Name: "expect_redeem_date", Type: proto.ColumnType_INT},
			{Name: "interest_per_unit", Type: proto.ColumnType_STRING},
			{Name: "withWhiteList", Type: proto.ColumnType_BOOL},
			{Name: "display", Type: proto.ColumnType_BOOL},
			{Name: "display_priority", Type: proto.ColumnType_STRING},
			{Name: "status", Type: proto.ColumnType_STRING},
			{Name: "annual_interest_rate", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.AnnualInterestRate")},
			{Name: "daily_interest_rate", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.DailyInterestRate")},
			{Name: "extra_interest_asset", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.ExtraInterestAsset")},
			{Name: "extra_annual_interest_rate", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.ExtraAnnualInterestRate")},
			{Name: "extra_daily_interest_rate", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.ExtraDailyInterestRate")},
			{Name: "min_purchase_amount", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.MinPurchaseAmount")},
			{Name: "max_purchase_amount_per_user", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.MaxPurchaseAmountPerUser")},
			{Name: "chain_process_period", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.ChainProcessPeriod")},
			{Name: "redeem_period", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.RedeemPeriod")},
			{Name: "pay_interest_period", Type: proto.ColumnType_STRING, Transform: transform.FromField("Config.PayInterestPeriod")},
			{Name: "sell_out", Type: proto.ColumnType_BOOL},
			{Name: "create_timestamp", Type: proto.ColumnType_INT},
			{Name: "selected", Type: proto.ColumnType_BOOL},
			{Name: "auto_renew", Type: proto.ColumnType_BOOL},
		},
	}
}

func listStaking(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	stakingRes := Locked()
	for _, t := range stakingRes.Data {
		for _, j := range t.Projects {
			d.StreamListItem(ctx, j)
			logger.Warn("item2", "Id", j.Id, "ProjectId", j.ProjectId, "Display", j.Display)
		}
	}

	return nil, nil
}
