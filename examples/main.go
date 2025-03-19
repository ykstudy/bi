package main

import (
	binance "bicore"
	"context"
	"fmt"
)

const (
	ApiKey   = "xxx"
	Secret   = "xxx"
	ProxyUrl = "xxx"
)

var client *binance.Client

func init() {
	client = binance.NewProxiedClient(ApiKey, Secret, ProxyUrl)

	timeOffset, err := client.NewSetServerTimeService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	client.TimeOffset = timeOffset
}

func main() {
	//Ticker()
	//Ohlcv()
	//SpotOrder()
	//FuturesOrder()
	//WalletBalance()
	//Account()
	//ExchangeInfo()
}
