package goanda

// Supporting OANDA docs - http://developer.oanda.com/rest-live-v20/trade-ep/

import (
	"encoding/json"
	"time"
)

type ReceivedTrades struct {
	LastTransactionID string `json:"lastTransactionID"`
	Trades            []struct {
		CurrentUnits string    `json:"currentUnits"`
		Financing    string    `json:"financing"`
		ID           string    `json:"id"`
		InitialUnits string    `json:"initialUnits"`
		Instrument   string    `json:"instrument"`
		OpenTime     time.Time `json:"openTime"`
		Price        string    `json:"price"`
		RealizedPL   string    `json:"realizedPL"`
		State        string    `json:"state"`
		UnrealizedPL string    `json:"unrealizedPL"`
	} `json:"trades"`
}

type ReceivedTrade struct {
	LastTransactionID string `json:"lastTransactionID"`
	Trades            struct {
		CurrentUnits string    `json:"currentUnits"`
		Financing    string    `json:"financing"`
		ID           string    `json:"id"`
		InitialUnits string    `json:"initialUnits"`
		Instrument   string    `json:"instrument"`
		OpenTime     time.Time `json:"openTime"`
		Price        string    `json:"price"`
		RealizedPL   string    `json:"realizedPL"`
		State        string    `json:"state"`
		UnrealizedPL string    `json:"unrealizedPL"`
	} `json:"trade"`
}

type CloseTradePayload struct {
	Units string
}

type ModifiedTrade struct {
	OrderCreateTransaction struct {
		Type         string `json:"type"`
		Instrument   string `json:"instrument"`
		Units        string `json:"units"`
		TimeInForce  string `json:"timeInForce"`
		PositionFill string `json:"positionFill"`
		Reason       string `json:"reason"`
		TradeClose   struct {
			Units   string `json:"units"`
			TradeID string `json:"tradeID"`
		} `json:"tradeClose"`
		ID        string    `json:"id"`
		UserID    int       `json:"userID"`
		AccountID string    `json:"accountID"`
		BatchID   string    `json:"batchID"`
		RequestID string    `json:"requestID"`
		Time      time.Time `json:"time"`
	} `json:"orderCreateTransaction"`
	OrderFillTransaction struct {
		Type           string `json:"type"`
		Instrument     string `json:"instrument"`
		Units          string `json:"units"`
		Price          string `json:"price"`
		FullPrice      string `json:"fullPrice"`
		PL             string `json:"pl"`
		Financing      string `json:"financing"`
		Commission     string `json:"commission"`
		AccountBalance string `json:"accountBalance"`
		TradeOpened    string `json:"tradeOpened"`
		TimeInForce    string `json:"timeInForce"`
		PositionFill   string `json:"positionFill"`
		Reason         string `json:"reason"`
		TradesClosed   []struct {
			TradeID    string `json:"tradeID"`
			Units      string `json:"units"`
			RealizedPL string `json:"realizedPL"`
			Financing  string `json:"financing"`
		} `json:"tradesClosed"`
		TradeReduced struct {
			TradeID    string `json:"tradeID"`
			Units      string `json:"units"`
			RealizedPL string `json:"realizedPL"`
			Financing  string `json:"financing"`
		} `json:"tradeReduced"`
		ID            string    `json:"id"`
		UserID        int       `json:"userID"`
		AccountID     string    `json:"accountID"`
		BatchID       string    `json:"batchID"`
		RequestID     string    `json:"requestID"`
		OrderID       string    `json:"orderId"`
		ClientOrderID string    `json:"clientOrderId"`
		Time          time.Time `json:"time"`
	} `json:"orderFillTransaction"`
	OrderCancelTransaction struct {
		Type      string    `json:"type"`
		OrderID   string    `json:"orderID"`
		Reason    string    `json:"reason"`
		ID        string    `json:"id"`
		UserID    int       `json:"userID"`
		AccountID string    `json:"accountID"`
		BatchID   string    `json:"batchID"`
		RequestID string    `json:"requestID"`
		Time      time.Time `json:"time"`
	} `json:"orderCancelTransaction"`
	RelatedTransactionIDs []string `json:"relatedTransactionIDs"`
	LastTransactionID     string   `json:"lastTransactionID"`
}

func (c *OandaConnection) GetTradesForInstrument(instrument string) (ReceivedTrades, error, error) {
	endpoint := "v3/accounts/" + c.AccountID + "/trades" + "?instrument=" + instrument

	response, err1, err2 := c.Request(endpoint)
	data := ReceivedTrades{}
	_ = unmarshalJson(response, &data)
	return data, err1, err2
}

func (c *OandaConnection) GetOpenTrades() (ReceivedTrades, error, error) {
	endpoint := "v3/accounts/" + c.AccountID + "/openTrades"

	response, err1, err2 := c.Request(endpoint)
	data := ReceivedTrades{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}

func (c *OandaConnection) GetTrade(ticket string) (ReceivedTrade, error, error) {
	endpoint := "v3/accounts/" + c.AccountID + "/trades" + "/" + ticket

	response, err1, err2 := c.Request(endpoint)
	data := ReceivedTrade{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}

// Default is close the whole position using the string "ALL" in body.units
func (c *OandaConnection) ReduceTradeSize(ticket string, body CloseTradePayload) (ModifiedTrade, error, error) {
	endpoint := "v3/accounts/" + c.AccountID + "/trades/" + ticket
	jsonBody, _ := json.Marshal(body)
	//checkErr(err)
	response, err1, err2 := c.Update(endpoint, jsonBody)
	data := ModifiedTrade{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}
