package marketbit

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

type TickerItem struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
	Date int    `json:"date"`
}

type TradeItem struct {
	Date   int     `json:"date"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid    int     `json:"tid"`
	Type   string  `json:"type"`
}

type OrderBookResponse struct {
	Asks []Ask `json:"asks"`
	Bids []Bid `json:"bids"`
}

type TickerResponse struct {
	TickerItem `json:"ticker"`
}

type TradeResponse struct {
	Trades []TradeItem
}

type SummaryItem struct {
	Lowest   float64 `json:"lowest"`
	Volume   float64 `json:"volume"`
	Amount   float64 `json:"amount"`
	AvgPrice float64 `json:"avg_price"`
	Opening  float64 `json:"opening"`
	Date     string  `json:"date"`
	Closing  float64 `json:"closing"`
	Highest  float64 `json:"highest"`
	Quantity float64 `json:"quantity"`
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

func GuessDateType(date interface{}) (*time.Time, error) {
	dateDefault := time.Now().AddDate(0, 0, -1)
	var err error
	switch date.(type) {
	case string:
		separator := "-"
		for _, bytes := range date.(string) {
			_, err := strconv.Atoi(string(bytes))
			if err != nil {
				separator = string(bytes)
			}
		}
		dateDefault, err = time.Parse(strings.Replace("2006#01#02", "#", separator, 2), date.(string))
		if err != nil {
			return nil, errors.Wrap(err, "Parse Error, format: 2006-01-02")
		}
	case int, int64:
		dateDefault = time.Unix(date.(int64), 64)
	case time.Time:
		dateDefault = date.(time.Time)
	default:
		return nil, errors.New("Unknown date type")
	}

	return &dateDefault, nil
}
