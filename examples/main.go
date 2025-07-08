package main

import (
	binance "bicore"
	"bicore/futures"
	"context"
	"example/config"
	"fmt"
)

var client *binance.Client
var futuresClient *futures.Client

func init() {

	conf := config.NewConfig()
	client = binance.NewProxiedClient(conf.B.ApiKey, conf.B.SecretKey, conf.B.ProxyUrl)

	futuresClient = futures.NewProxiedClient(conf.B.ApiKey, conf.B.SecretKey, conf.B.ProxyUrl)

	timeOffset, err := client.NewSetServerTimeService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	client.TimeOffset = timeOffset
	futuresClient.TimeOffset = timeOffset
}

func main() {
	//Ticker()
	//Ohlcv()
	//SpotOrder()
	//FuturesOrder()
	//WalletBalance()
	//Account()
	ExchangeInfo()
	//Kline()
}
