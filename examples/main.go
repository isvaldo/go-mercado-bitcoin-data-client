package main

import (
	"fmt"
	"github.com/isvaldo/go-mercado-bitcoin-data-client"
)

func main() {
	cli := marketbit.New("https://www.mercadobitcoin.net/api")
	tickerItem, err := cli.GetTicker("btc")
	if err != nil {
		panic(err)
	}
	fmt.Println(tickerItem.Low)
}
