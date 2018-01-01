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
	fmt.Println(tickerItem.Date)
	fmt.Println(tickerItem.Vol)
	fmt.Println(tickerItem.Sell)
	fmt.Println(tickerItem.Last)
	fmt.Println(tickerItem.Buy)
	fmt.Println(tickerItem.High)
	fmt.Println(tickerItem.Low)
}
