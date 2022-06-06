package binance

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// https://binance-docs.github.io/apidocs/spot/en/

const BASE_URL = "https://api.binance.com"

type Client struct {
	APIKey    string
	SecretKey string
}

func NewClient(apiKey, secretKey string) *Client {
	c := &Client{
		APIKey:    apiKey,
		SecretKey: secretKey,
	}
	return c
}

// public

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

func (c *Client) Locked() earnLockedStakingResponse {
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

	var response earnLockedStakingResponse
	json.Unmarshal(body, &response)

	// fmt.Print(string(body))

	return response
}

func (c *Client) Defi() {
	// url := "https://www.binance.com/bapi/earn/v2/friendly/defi-pos/groupByAssetAndPartnerNameList?pageSize=15&pageIndex=1&status=ALL"
}

// private

type BswapUnclaimedRewardsRequest struct {
	Type uint
}

type BswapUnclaimedRewardsResponse struct {
	TotalUnclaimedRewards map[string]string            `json:"totalUnclaimedRewards"`
	Details               map[string]map[string]string `json:"details"`
}

type BswapLiquidityResponse struct {
	PoolId     uint              `json:"poolId"`
	PoolName   string            `json:"poolNmae"`
	UpdateTime uint              `json:"updateTime"`
	Liquidity  map[string]string `json:"liquidity"`
	Share      struct {
		ShareAmount     string            `json:"shareAmount"`
		SharePercentage string            `json:"sharePercentage"`
		Asset           map[string]string `json:"asset"`
	} `json:"share"`
}

func (c *Client) BswapLiquidity() []BswapLiquidityResponse {
	method := "GET"
	endpoint := "/sapi/v1/bswap/liquidity"

	query := url.Values{}
	data := c.callSigned(method, endpoint, query)

	var response []BswapLiquidityResponse
	json.Unmarshal(data, &response)

	return response
}

func (c *Client) BswapUnclaimedRewards(request BswapUnclaimedRewardsRequest) BswapUnclaimedRewardsResponse {
	method := "GET"
	endpoint := "/sapi/v1/bswap/unclaimedRewards"

	query := url.Values{}
	query.Set("type", fmt.Sprintf("%v", request.Type))

	data := c.callSigned(method, endpoint, query)

	var response BswapUnclaimedRewardsResponse
	json.Unmarshal(data, &response)

	return response
}

func (c *Client) callSigned(method string, endpoint string, query url.Values) []byte {
	query.Set("timestamp", fmt.Sprintf("%v", time.Now().UnixNano()/int64(time.Millisecond)))
	form := url.Values{}
	bodyString := form.Encode()

	query.Set("timestamp", fmt.Sprintf("%v", time.Now().UnixNano()/int64(time.Millisecond)))
	queryString := query.Encode()
	raw := fmt.Sprintf("%s%s", queryString, bodyString)
	mac := hmac.New(sha256.New, []byte(c.SecretKey))
	_, _ = mac.Write([]byte(raw))
	signature := fmt.Sprintf("%x", (mac.Sum(nil)))

	query.Set("signature", signature)

	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("X-MBX-APIKEY", c.APIKey)

	fullURL := fmt.Sprintf("%s%s", BASE_URL, endpoint)
	queryString = query.Encode()
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	fmt.Println("fullURL=", fullURL)

	client := http.DefaultClient
	req, _ := http.NewRequest(method, fullURL, &bytes.Buffer{})
	req.Header = header
	res, _ := client.Do(req)
	data, _ := ioutil.ReadAll(res.Body)

	// str1 := string(data[:])
	// fmt.Println("String =", str1)

	return data
}
