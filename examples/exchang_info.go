package main

import (
	"context"
	"fmt"

	"example/data"
)

func ExchangeInfo() {

	showPermissionSets := false
	exchangeInfo, err := client.NewExchangeInfoService().
		//Symbol("BTCUSDT").
		ShowPermissionSets(&showPermissionSets).
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	var symbols []string
	symbolList := exchangeInfo.Symbols
	for _, symbol := range symbolList {
		// 只需要以 "USDT" 结尾的交易对
		if symbol.Symbol[len(symbol.Symbol)-4:] == "USDT" {
			symbols = append(symbols, symbol.Symbol)
		}
	}
	// 写入到文件
	filePath := "data/symbols/symbols.txt"
	err = data.WriteToFile(filePath, symbols)

	//account, _ := json.Marshal(exchangeInfo)
	//fmt.Println(string(account))
}
