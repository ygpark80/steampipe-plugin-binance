package main

import (
	"testing"

	"github.com/ygpark80/steampipe-plugin-binance/binance"
)

func TestApi(t *testing.T) {
	earnLockedStaking := binance.NewClient("", "").Locked()

	if earnLockedStaking.Code != "000000" {
		t.Fatal(`Hello("")`, "")
	}
}
