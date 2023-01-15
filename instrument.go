package goanda

// Supporting OANDA docs - http://developer.oanda.com/rest-live-v20/instrument-ep/

import (
	"time"
)

type Candle struct {
	Open  string `json:"o"`
	Close string `json:"c"`
	Low   string `json:"l"`
	High  string `json:"h"`
}

type Candles struct {
	Complete bool      `json:"complete"`
	Volume   int       `json:"volume"`
	Time     time.Time `json:"time"`
	Mid      Candle    `json:"mid"`
}

type BidAskCandles struct {
	Candles []struct {
		Ask struct {
			C string `json:"c"`
			H string `json:"h"`
			L string `json:"l"`
			O string `json:"o"`
		} `json:"ask"`
		Bid struct {
			C string `json:"c"`
			H string `json:"h"`
			L string `json:"l"`
			O string `json:"o"`
		} `json:"bid"`
		Complete bool      `json:"complete"`
		Time     time.Time `json:"time"`
		Volume   int       `json:"volume"`
	} `json:"candles"`
}

type InstrumentHistory struct {
	Instrument  string    `json:"instrument"`
	Granularity string    `json:"granularity"`
	Candles     []Candles `json:"candles"`
}

type Bucket struct {
	Price             string `json:"price"`
	LongCountPercent  string `json:"longCountPercent"`
	ShortCountPercent string `json:"shortCountPercent"`
}

type BrokerBook struct {
	Instrument  string    `json:"instrument"`
	Time        time.Time `json:"time"`
	Price       string    `json:"price"`
	BucketWidth string    `json:"bucketWidth"`
	Buckets     []Bucket  `json:"buckets"`
}

type InstrumentPricing struct {
	Time   time.Time `json:"time"`
	Prices []struct {
		Type string    `json:"type"`
		Time time.Time `json:"time"`
		Bids []struct {
			Price     float64 `json:"price,string"`
			Liquidity int     `json:"liquidity"`
		} `json:"bids"`
		Asks []struct {
			Price     float64 `json:"price,string"`
			Liquidity int     `json:"liquidity"`
		} `json:"asks"`
		CloseoutBid    float64 `json:"closeoutBid,string"`
		CloseoutAsk    float64 `json:"closeoutAsk,string"`
		Status         string  `json:"status"`
		Tradeable      bool    `json:"tradeable"`
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
		QuoteHomeConversionFactors struct {
			PositiveUnits string `json:"positiveUnits"`
			NegativeUnits string `json:"negativeUnits"`
		} `json:"quoteHomeConversionFactors"`
		Instrument string `json:"instrument"`
	} `json:"prices"`
}

func (c *OandaConnection) GetCandles(instrument string, count string, granularity string) (InstrumentHistory, error, error) {
	endpoint := "v3/instruments/" + instrument + "/candles?count=" + count + "&granularity=" + granularity
	candles, err1, err2 := c.Request(endpoint)
	data := InstrumentHistory{}
  _ = unmarshalJson(candles, &data)

	return data, err1, err2
}

func (c *OandaConnection) GetBidAskCandles(instrument string, count string, granularity string) (BidAskCandles, error, error) {
	endpoint := "v3/instruments/" + instrument + "/candles?count=" + count + "&granularity=" + granularity + "&price=BA"
	candles, err1, err2 := c.Request(endpoint)
	data := BidAskCandles{}
  _ = unmarshalJson(candles, &data)

	return data, err1, err2
}

func (c *OandaConnection) OrderBook(instrument string) (BrokerBook, error, error) {
	var json_data interface{}
	endpoint := "v3/instruments/" + instrument + "/orderBook"
	data, err1, err2 := c.Request(endpoint)
	orderbook := BrokerBook{}
  _ = unmarshalJson(data, &json_data)
	json_data = json_data.(map[string]interface{})["orderBook"]
	bytes, _ := marshalJson(json_data)
	_ = unmarshalJson(bytes, &orderbook)

	return orderbook, err1, err2
}

func (c *OandaConnection) PositionBook(instrument string) (BrokerBook, error, error) {
	var json_data interface{}
	endpoint := "v3/instruments/" + instrument + "/positionBook"
	data, err1, err2 := c.Request(endpoint)
	positionbook := BrokerBook{}
	_ = unmarshalJson(data, &json_data)
	json_data = json_data.(map[string]interface{})["positionBook"]
	bytes, _ := marshalJson(json_data)
	_ = unmarshalJson(bytes, &positionbook)

	return positionbook, err1, err2
}

func (c *OandaConnection) GetInstrumentPrice(instrument string) (InstrumentPricing, error, error) {
	endpoint := "v3/accounts/" + c.AccountID + "/pricing?instruments=" + instrument
	pricing, err1, err2 := c.Request(endpoint)
	data := InstrumentPricing{}
	_ = unmarshalJson(pricing, &data)

	return data, err1, err2
}
