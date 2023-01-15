package goanda

import (
	"net/url"
	"time"
)

// Supporting OANDA docs - http://developer.oanda.com/rest-live-v20/transaction-ep/

type TransactionPages struct {
	Count             int       `json:"count"`
	From              time.Time `json:"from"`
	LastTransactionID string    `json:"lastTransactionID"`
	PageSize          int       `json:"pageSize"`
	Pages             []string  `json:"pages"`
	To                time.Time `json:"to"`
}

type Transaction struct {
	LastTransactionID string `json:"lastTransactionID"`
	Transaction       struct {
		AccountBalance string    `json:"accountBalance"`
		AccountID      string    `json:"accountID"`
		BatchID        string    `json:"batchID"`
		Financing      string    `json:"financing"`
		ID             string    `json:"id"`
		Instrument     string    `json:"instrument"`
		OrderID        string    `json:"orderID"`
		Pl             string    `json:"pl"`
		Price          string    `json:"price"`
		Reason         string    `json:"reason"`
		Time           time.Time `json:"time"`
		TradeOpened    struct {
			TradeID string `json:"tradeID"`
			Units   string `json:"units"`
		} `json:"tradeOpened"`
		Type   string `json:"type"`
		Units  string `json:"units"`
		UserID int    `json:"userID"`
	} `json:"transaction"`
}

type Transactions struct {
	LastTransactionID string `json:"lastTransactionID"`
	Transactions      []struct {
		AccountBalance string    `json:"accountBalance"`
		AccountID      string    `json:"accountID"`
		BatchID        string    `json:"batchID"`
		Financing      string    `json:"financing"`
		ID             string    `json:"id"`
		Instrument     string    `json:"instrument"`
		OrderID        string    `json:"orderID"`
		Pl             string    `json:"pl"`
		Price          string    `json:"price"`
		Reason         string    `json:"reason"`
		Time           time.Time `json:"time"`
		TradeOpened    struct {
			TradeID string `json:"tradeID"`
			Units   string `json:"units"`
		} `json:"tradeOpened"`
		Type   string `json:"type"`
		Units  string `json:"units"`
		UserID int    `json:"userID"`
	} `json:"transactions"`
}

// https://golang.org/pkg/time/#Time.AddDate
// https://play.golang.org/p/Dw7D4JJ7EC
func (c *OandaConnection) GetTransactions(from time.Time, to time.Time) (TransactionPages, error, error) {
	toTime := to.Format(time.RFC3339)
	fromTime := from.Format(time.RFC3339)

	endpoint := "v3/accounts/" + c.AccountID + "/transactions?to=" + url.QueryEscape(toTime) + "&from=" + url.QueryEscape(fromTime)

	response, err1, err2 := c.Request(endpoint)
	data := TransactionPages{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}

func (c *OandaConnection) GetTransaction(ticket string) (Transaction, error, error) {

	endpoint := "v3/accounts/" + c.AccountID + "/transactions/" + ticket

	response, err1, err2 := c.Request(endpoint)
	data := Transaction{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}

func (c *OandaConnection) GetTransactionsSinceId(id string) (Transactions, error, error) {

	endpoint := "v3/accounts/" + c.AccountID + "/transactions/sinceid?id=" + id

	response, err1, err2 := c.Request(endpoint)
	data := Transactions{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}
