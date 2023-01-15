package goanda

import (
	"net/url"
	"strings"
	"time"
)

// Supporting OANDA docs - http://developer.oanda.com/rest-live-v20/pricing-ep/

type Pricings struct {
	Prices []struct {
		Asks []struct {
			Liquidity int    `json:"liquidity"`
			Price     string `json:"price"`
		} `json:"asks"`
		Bids []struct {
			Liquidity int    `json:"liquidity"`
			Price     string `json:"price"`
		} `json:"bids"`
		CloseoutAsk                string `json:"closeoutAsk"`
		CloseoutBid                string `json:"closeoutBid"`
		Instrument                 string `json:"instrument"`
		QuoteHomeConversionFactors struct {
			NegativeUnits string `json:"negativeUnits"`
			PositiveUnits string `json:"positiveUnits"`
		} `json:"quoteHomeConversionFactors"`
		Status         string    `json:"status"`
		Time           time.Time `json:"time"`
		UnitsAvailable struct {
			Default struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"default"`
			OpenOnly struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"openOnly"`
			ReduceFirst struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"reduceFirst"`
			ReduceOnly struct {
				Long  string `json:"long"`
				Short string `json:"short"`
			} `json:"reduceOnly"`
		} `json:"unitsAvailable"`
	} `json:"prices"`
}

func (c *OandaConnection) GetPricingForInstruments(instruments []string) (Pricings, error, error) {
	instrumentString := strings.Join(instruments, ",")
	endpoint := "v3/accounts/" + c.AccountID + "/pricing?instruments=" + url.QueryEscape(instrumentString)

	response, err1, err2 := c.Request(endpoint)
	data := Pricings{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}
