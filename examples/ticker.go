package main

import (
	"context"
	"fmt"

	"bicore"
)

func Ticker() {
	client := binance.NewClient(ApiKey, Secret)

	// spot ticker
	ticker, err := client.NewTradingDayTickerService().Symbol("BTCUSDT").Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ticker := range ticker {
		fmt.Println(ticker)
	}

	// futures ticker
	futuresClient := binance.NewFuturesClient(ApiKey, Secret)
	futuresTicker, err2 := futuresClient.NewListBookTickersService().Symbol("BTCUSDT").Do(context.Background())
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	for _, ticker := range futuresTicker {
		fmt.Println(ticker)
	}

}
