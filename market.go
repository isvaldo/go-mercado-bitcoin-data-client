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

type Ask []float64
type Bid []float64

func (c Ask) Price() float64 {
	return c[0]
}
func (c Ask) Amount() float64 {
	return c[1]
}
func (c Bid) Price() float64 {
	return c[0]
}
func (c Bid) Amount() float64 {
	return c[1]
}

type OrderBookResponse struct {
	Asks []Ask `json:"asks"`
	Bids []Bid `json:"bids"`
}
