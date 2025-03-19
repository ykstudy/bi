package main

import (
	"context"
	"fmt"

	"bicore"
)

func WalletBalance() {
	client := binance.NewClient(ApiKey, Secret)

	quoteAsset := "USDT"

	res, err := client.NewWalletBalanceService().
		QuoteAsset(quoteAsset).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, w := range res {
		fmt.Printf("%s: %s\n", w.WalletName, w.Balance)
	}

}
