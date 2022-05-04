package binance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Locked() earnLockedStakingResponse {
	url := "https://www.binance.com/bapi/earn/v2/friendly/pos/union?pageSize=200&pageIndex=1&status=ALL"
	res, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
	}

	var stakingRes earnLockedStakingResponse
	json.Unmarshal(body, &stakingRes)

	// fmt.Print(string(body))

	return stakingRes
}

func Defi() {
	// url := "https://www.binance.com/bapi/earn/v2/friendly/defi-pos/groupByAssetAndPartnerNameList?pageSize=15&pageIndex=1&status=ALL"
}

type earnLockedStakingResponse struct {
	Code string `json:"code"`
	Data []struct {
		AnnualInterestRate string     `json:"annualInterestRate"`
		Asset              string     `json:"asset"`
		MinPurchaseAmount  string     `json:"minPurchaseAmount"`
		RedeemPeriod       string     `json:"redeemPeriod"`
		Products           []struct{} `json:"products"`
		Projects           []struct {
			Id               string `json:"id"`
			ProjectId        string `json:"projectId"`
			Asset            string `json:"asset"`
			UpLimit          string `json:"upLimit"`
			Purchased        string `json:"purchased"`
			EndTime          string `json:"endTime"`
			IssueStartTime   string `json:"issueStartTime"`
			IssueEndTime     string `json:"issueEndTime"`
			Duration         string `json:"duration"`
			ExpectRedeemDate string `json:"expectRedeemDate"`
			InterestPerUnit  string `json:"interestPerUnit"`
			WithWhiteList    bool   `json:"withWhiteList"`
			Display          bool   `json:"display"`
			DisplayPriority  string `json:"displayPriority"`
			Status           string `json:"status"`
			Config           struct {
				Id                       string `json:"id"`
				AnnualInterestRate       string `json:"annualInterestRate"`
				DailyInterestRate        string `json:"dailyInterestRate"`
				ExtraInterestAsset       string `json:"extraInterestAsset"`
				ExtraAnnualInterestRate  string `json:"extraAnnualInterestRate"`
				ExtraDailyInterestRate   string `json:"extraDailyInterestRate"`
				MinPurchaseAmount        string `json:"minPurchaseAmount"`
				MaxPurchaseAmountPerUser string `json:"maxPurchaseAmountPerUser"`
				ChainProcessPeriod       string `json:"chainProcessPeriod"`
				RedeemPeriod             string `json:"redeemPeriod"`
				PayInterestPeriod        string `json:"payInterestPeriod"`
			} `json:"config"`
			SellOut         bool   `json:"sellOut"`
			CreateTimestamp string `json:"createTimestamp"`
			Selected        *bool  `json:"selected"`
			AutoRenew       bool   `json:"autoRenew"`
		} `json:"projects"`
	} `json:"data"`
	Total   int  `json:"total"`
	Success bool `json:"success"`
}
