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

	fmt.Println(tickerItem.Low)

}
