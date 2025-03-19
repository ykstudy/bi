package main

import (
	binance "bicore"
	"context"
	"encoding/json"
	"fmt"
)

func ExchangeInfo() {
	client := binance.NewClient(ApiKey, Secret)
	client.Debug = true

	res, err := client.NewExchangeInfoService().
		Symbol("BTCUSDT").
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	account, _ := json.Marshal(res)
	fmt.Println(string(account))
}
