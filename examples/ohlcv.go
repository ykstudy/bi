package main

import (
	"context"
	"fmt"

	"bicore"
)

func Ohlcv() {

	// spot ohlcv
	ohlcv, err := client.NewKlinesService().Symbol("BTCUSDT").Interval("1m").Limit(5).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, kline := range ohlcv {
		fmt.Println(kline.OpenTime, kline.Open, kline.High, kline.Low, kline.Close, kline.Volume)
	}

	// futures ohlcv
	futuresClient := binance.NewFuturesProxiedClient(ApiKey, Secret, ProxyUrl)
	futuresOHLCV, err2 := futuresClient.NewKlinesService().Symbol("BTCUSDT").Interval("1m").Limit(5).Do(context.Background())
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	for _, kline := range futuresOHLCV {
		fmt.Println(kline.OpenTime, kline.Open, kline.High, kline.Low, kline.Close, kline.Volume)
	}

}
