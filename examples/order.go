package main

import (
	"context"
	"fmt"

	"github.com/exchange-now/go-binance-th"
	"github.com/exchange-now/go-binance-th/futures"
)

func SpotOrder() {
	binance.UseTestnet = true
	apiKey := ""
	secret := ""
	client := binance.NewClient(apiKey, secret)

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
	apiKey := ""
	secret := ""
	client := binance.NewFuturesClient(apiKey, secret)

	symbol := "LTCUSDT"
	side := futures.SideTypeSell
	orderType := futures.OrderTypeMarket
	quantity := "0.1"

	res, err := client.NewCreateOrderService().Symbol(symbol).Side(side).
		Type(orderType).Quantity(quantity).PositionSide(futures.PositionSideTypeLong).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

}
