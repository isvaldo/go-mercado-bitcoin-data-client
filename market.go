package marketbit

type TickerItem struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
	Date int    `json:"date"`
}

type TickerResponse struct {
	TickerItem `json:"ticker"`
}
