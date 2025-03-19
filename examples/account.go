package main

import (
	"context"
	"encoding/json"
	"fmt"

	"bicore"
)

func Account() {
	client := binance.NewClient(ApiKey, Secret)

	res, err := client.NewGetAccountService().
		OmitZeroBalances(true).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	account, _ := json.Marshal(res)
	fmt.Println(string(account))
}
