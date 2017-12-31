package marketbit

import (
	"fmt"
	"testing"
)

func TestTickerMustBeReturnMockedValues(t *testing.T) {

	cli := New("https://www.mercadobitcoin.net/api")

	tickerItem, err := cli.GetTicker("btc")
	if err != nil {
		panic(err)
	}

	fmt.Println(tickerItem.Buy)
	fmt.Println(tickerItem.Last)
	fmt.Println(tickerItem.Sell)
	fmt.Println(tickerItem.Vol)
}

func TestOrderBookMustBeReturnMockedValues(t *testing.T) {

	cli := New("https://www.mercadobitcoin.net/api")

	orderItem, err := cli.GetOrderBook("btc")
	if err != nil {
		panic(err)
	}

	for _, askItem := range orderItem.Asks {
		fmt.Println(askItem.Amount())
		fmt.Println(askItem.Price())
	}
	for _, bidsItem := range orderItem.Bids {
		fmt.Println(bidsItem.Amount())
		fmt.Println(bidsItem.Price())
	}
}

func TestTradesMustBeReturnMockedValues(t *testing.T) {

	cli := New("https://www.mercadobitcoin.net/api")

	tradesResponse, err := cli.GetLastTrades("btc")
	if err != nil {
		panic(err)
	}

	for _, tradeItens := range tradesResponse.Trades {
		fmt.Println(tradeItens.Price)
		fmt.Println(tradeItens.Amount)
		fmt.Println(tradeItens.Date)
		fmt.Println(tradeItens.Tid)
		fmt.Println(tradeItens.Type)
	}
}

func TestSummaryMustBeReturnMockedValues(t *testing.T) {

	cli := New("https://www.mercadobitcoin.net/api")

	summary, err := cli.GetSummaryAt("btc", "2018-12-30")
	if err != nil {
		panic(err)
	}

	fmt.Println(summary.Opening)
	fmt.Println(summary.Closing)
}
