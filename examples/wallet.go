package main

import (
	"context"
	"fmt"

	"github.com/exchange-now/go-binance-th"
)

func WalletBalance() {
	apiKey := ""
	secret := ""
	client := binance.NewClient(apiKey, secret)

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
