<p align="center"><img src="docs/img/mkbtc.png" width="360"></p>
<p align="center">
  <a href="https://travis-ci.org/isvaldo/go-mercado-bitcoin-data-client"><img src="https://travis-ci.org/isvaldo/go-mercado-bitcoin-data-client.svg?branch=master"/></a>
  <a href="https://ci.appveyor.com/project/isvaldo/go-mercado-bitcoin-data-client"><img src="https://ci.appveyor.com/api/projects/status/github/isvaldo/go-mercado-bitcoin-data-client?svg=true&branch=master&passingText=Windows%20-%20OK&failingText=Windows%20-%20failed&pendingText=Windows%20-%20pending" alt="Windows Build Status"></a>
  <a href="https://goreportcard.com/report/github.com/isvaldo/go-mercado-bitcoin-data-client"><img src="https://goreportcard.com/badge/github.com/isvaldo/go-mercado-bitcoin-data-client" /></a>
  <a href="https://codeclimate.com/github/isvaldo/go-mercado-bitcoin-data-client/test_coverage"><img src="https://api.codeclimate.com/v1/badges/0003cc5bae42ebbbb8c9/test_coverage" /></a>
  <a href="https://codeclimate.com/github/isvaldo/go-mercado-bitcoin-data-client/maintainability"><img src="https://api.codeclimate.com/v1/badges/0003cc5bae42ebbbb8c9/maintainability" /></a>
  </p>



# go-mercado-bitcoin-data-client

Api de dados do mercado bitcoin oferece uma maneira automatica de obter os dados
do orderbook e negociações, esse client facilita o acesso a essas infomações atraves da api
oficial  <a href="https://www.mercadobitcoin.com.br/api-doc/">Docs</a>


## Install
Para baixar as dependencias do projeto rode o seguinte comando

```sh
dep ensure
``` 

### Examples


#### Client
<details>
<summary>Iniciando Client</summary>

```go
package main

import (
	"fmt"
	"github.com/isvaldo/go-mercado-bitcoin-data-client"
)

func main() {
	cli := marketbit.New("https://www.mercadobitcoin.net/api")
	fmt.Println(cli)
}
```
</details>
<details>

<summary>Customizando http client</summary>

```go
package main

import (
	"fmt"
	"github.com/isvaldo/go-mercado-bitcoin-data-client"
)

func main() {
	cli := marketbit.NewWithClient("https://www.mercadobitcoin.net/api",createHTTPClient(10,10))
	fmt.Println(cli)
}

func createHTTPClient(maxIdleConnections, requestTimeout int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: maxIdleConnections,
		},
		Timeout: time.Duration(requestTimeout) * time.Second,
	}
}
```
</details>


#### Ticker
Retorna informações com o resumo das últimas 24 horas de negociações.


<details>
<summary>Resumo das ultimas 24 horas BTC</summary>

```go
package main

import (
	"fmt"
	"github.com/isvaldo/go-mercado-bitcoin-data-client"
)

func main() {
	cli := marketbit.New("https://www.mercadobitcoin.net/api")
	tickerItem, err := cli.GetTicker("btc") // btc, ltc, bch
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
```

<details>
<summary>Http Response</summary>

```javascript
'ticker': {
'high': 14481.47000000,
'low': 13706.00002000,
'vol': 443.73564488,
'last': 14447.01000000,
'buy': 14447.00100000,
'sell': 14447.01000000,
'date': 1502977646
}
```
</details>

<details>
<summary>Code Response</summary>

```sh
/tmp/___go_build_main_go #gosetup
1514840316
115.06076518
48550.00000000
48550.00000000
48400.00000000
49599.00000000
47500.00000000
```

</details>

</details>

#### OrderBook
Livro de ofertas é composto por duas listas: (1) uma lista com as ofertas de compras ordenadas pelo maior valor; (2) uma lista com as ofertas de venda ordenadas pelo menor valor. O livro mostra até 1000 ofertas de compra e até 1000 ofertas de venda.

Uma oferta é constituída por uma ou mais ordens, sendo assim, a quantidade da oferta é o resultado da soma das quantidades das ordens de mesmo preço unitário. Caso uma oferta represente mais de uma ordem, a prioridade de execução se dá com base na data de criação da ordem, da mais antiga para a mais nova.

<details>
<summary>Resumo do orderbook</summary>


```go
package main

import (
	"fmt"
	"github.com/isvaldo/go-mercado-bitcoin-data-client"
)

func main() {
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

```

</details>

#### Trades
Histórico de negociações realizadas.

<details>
<summary>Resumo do orderbook</summary>


```go
package main

import (
	"fmt"
	"github.com/isvaldo/go-mercado-bitcoin-data-client"
)

func main() {
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

```

</details>



### Donate

https://isvaldo.github.io/donate-bitcoin/