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

<details>
<summary>Iniciando client</summary>

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
<summary>Resumo 24 horas</summary>

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

### Donate

https://isvaldo.github.io/donate-bitcoin/