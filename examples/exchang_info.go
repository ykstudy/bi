package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExchangeInfo() {

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
