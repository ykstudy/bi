package main

import (
	"context"
	"fmt"

	"bicore"
	"bicore/futures"
)

func SpotOrder() {
	binance.UseTestnet = true

	symbol := "BTCUSDT"
	side := binance.SideTypeSell
	orderType := binance.OrderTypeMarket
	quantity := "0.0001"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}

func FuturesOrder() {
	futures.UseTestnet = true
	futuresClient := binance.NewFuturesProxiedClient(ApiKey, Secret, ProxyUrl)

	symbol := "LTCUSDT"
	side := futures.SideTypeSell
	orderType := futures.OrderTypeMarket
	quantity := "0.1"

	res, err := futuresClient.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).PositionSide(futures.PositionSideTypeLong).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

}
