package marketbit

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// concrete Implementation of marketbit.Interface
type HTTP struct {
	BaseURL    string
	HTTPClient *http.Client
}

func (c *HTTP) GetTicker(coin string) (*TickerItem, error) {
	response := TickerResponse{TickerItem{}}
	path := fmt.Sprintf("%s/%s/ticker", c.BaseURL, strings.ToUpper(coin))
	if err := c.doRequest("GET", path, nil, &response); err != nil {
		return nil, errors.Wrapf(err, "Error on execute request:[%s]", path)
	}
	return &response.TickerItem, nil
}

func (c *HTTP) GetOrderBook(coin string) (*OrderBookResponse, error) {
	response := OrderBookResponse{}
	path := fmt.Sprintf("%s/%s/orderbook", c.BaseURL, strings.ToUpper(coin))
	if err := c.doRequest("GET", path, nil, &response); err != nil {
		return nil, errors.Wrapf(err, "Error on execute request:[%s]", path)
	}
	return &response, nil
}

func (c *HTTP) GetLastTrades(coin string) (*TradeResponse, error) {
	response := TradeResponse{}
	path := fmt.Sprintf("%s/%s/trades", c.BaseURL, strings.ToUpper(coin))
	if err := c.doRequest("GET", path, nil, &response.Trades); err != nil {
		return nil, errors.Wrapf(err, "Error on execute request:[%s]", path)
	}
	return &response, nil
}

func (c *HTTP) GetTradesRange(coin string, from interface{}, to interface{}) (*TradeResponse, error) {
	fromDate, err := GuessDateType(from)
	if err != nil {
		return nil, errors.Wrap(err, "Error on guess date type")
	}
	toDate, err := GuessDateType(to)
	if err != nil {
		return nil, errors.Wrap(err, "Error on guess date type")
	}

	response := TradeResponse{}
	path := fmt.Sprintf("%s/%s/trades/%d/%d/",
		c.BaseURL,
		strings.ToUpper(coin),
		fromDate.Unix(),
		toDate.Unix())
	if err := c.doRequest("GET", path, nil, &response.Trades); err != nil {
		return nil, errors.Wrapf(err, "Error on execute request:[%s]", path)
	}
	return &response, nil
}

func (c *HTTP) GetSummaryAt(coin string, date interface{}) (*SummaryItem, error) {
	dateDefault, err := GuessDateType(date)
	if err != nil {
		return nil, errors.Wrap(err, "Error on guess date type")
	}
	response := SummaryItem{}
	path := fmt.Sprintf("%s/%s/day-summary/%d/%d/%d",
		c.BaseURL,
		strings.ToUpper(coin),
		dateDefault.Year(), dateDefault.Month(), dateDefault.Day())
	if err := c.doRequest("GET", path, nil, &response); err != nil {
		return nil, errors.Wrapf(err, "Error on execute request:[%s]", path)
	}
	return &response, nil
}

//doRequest is help for execute request and parse the response
func (c *HTTP) doRequest(method, url string, body io.Reader, to interface{}) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if to == nil {
		io.Copy(ioutil.Discard, res.Body)
		return nil
	}
	return json.NewDecoder(res.Body).Decode(to)
}
