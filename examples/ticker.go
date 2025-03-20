package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func Ticker() {

	// spot ticker
	tickers, err := client.NewTradingDayTickerService().Symbol("BTCUSDT").Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ticker := range tickers {
		marshal, _ := json.Marshal(ticker)
		fmt.Println(string(marshal))
	}

	// futures ticker
	futuresTickers, err2 := futuresClient.NewListBookTickersService().Symbol("BTCUSDT").Do(context.Background())
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	for _, futuresTicker := range futuresTickers {
		marshal, _ := json.Marshal(futuresTicker)
		fmt.Println(string(marshal))
	}

}
