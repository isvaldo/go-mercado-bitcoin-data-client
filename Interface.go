package marketbit

import "net/http"

type Interface interface {
	GetTicker(coin string) (*TickerItem, error)
}

//NewWithClient return a instance from MarketTicker, with optional client configuration
func NewWithClient(baseURL string, client *http.Client) Interface {
	return NewDataMarket(baseURL, client)
}

//NewWithClient Normal call
func New(baseURL string) Interface {
	return NewDataMarket(baseURL, http.DefaultClient)
}

//NewMedia return a instance from MarketTicker, with optional client configuration
func NewDataMarket(baseURL string, client *http.Client) Interface {
	return &HTTP{
		BaseURL:    baseURL,
		HTTPClient: client,
	}
}
