package main

import (
	"fmt"
	"testing"

	"github.com/ygpark80/steampipe-plugin-binance/binance"
)

func TestApi(t *testing.T) {
	earnLockedStaking := binance.NewClient("", "").Locked()

	if earnLockedStaking.Code != "000000" {
		t.Fatal(`Hello("")`, "")
	}
}

func TestBswapLiquidity(t *testing.T) {
	config := GetConfig()

	client := binance.NewClient(config.APIKey, config.SecretKey)
	response := client.BswapLiquidity()

	for k, v := range response {
		fmt.Println(k, "=", fmt.Sprintf("%v", v.Share.Asset))
	}
}

func TestBswapUnclaimedRewards(t *testing.T) {
	config := GetConfig()

	client := binance.NewClient(config.APIKey, config.SecretKey)
	response := client.BswapUnclaimedRewards(binance.BswapUnclaimedRewardsRequest{Type: 1})

	for k, v := range response.TotalUnclaimedRewards {
		fmt.Println(k, "=", fmt.Sprintf("%v", v))
	}

	for k, v := range response.Details {
		for kk, vv := range v {
			fmt.Println(k, kk, "=", fmt.Sprintf("%v", vv))
		}
	}
}
