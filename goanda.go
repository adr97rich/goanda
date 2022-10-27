package goanda

import (
	"bytes"
	"net/http"
	"time"
)

type Headers struct {
	contentType    string
	agent          string
	DatetimeFormat string
	auth           string
}

type Connection interface {
	Request(endpoint string) []byte
	Send(endpoint string, data []byte) []byte
	Update(endpoint string, data []byte) []byte
	GetOrderDetails(instrument string, units string) OrderDetails
	GetAccountSummary() AccountSummary
	CreateOrder(body OrderPayload) OrderResponse
}

type OandaConnection struct {
	Hostname        string
	Port            int
	Ssl             bool
	Token           string
	AccountID       string
	DatetimeFormat  string
	Headers         *Headers
}

const OANDA_AGENT string = "v20-golang/0.0.1"

func NewConnection(accountID string, token string, live bool) *OandaConnection {
	hostname := ""
	// should we use the live API?
	if live {
		hostname = "https://api-fxtrade.oanda.com/"
	} else {
		hostname = "https://api-fxpractice.oanda.com/"
	}

	var buffer bytes.Buffer
	// Generate the auth header
	buffer.WriteString("Bearer ")
	buffer.WriteString(token)

	authHeader := buffer.String()
	// Create headers for oanda to be used in requests
	headers := &Headers{
		contentType:    "application/json",
		agent:          OANDA_AGENT,
		DatetimeFormat: "RFC3339",
		auth:           authHeader,
	}
	// Create the connection object
	connection := &OandaConnection{
		Hostname:  hostname,
		Port:      443,
		Ssl:       true,
		Token:     token,
		Headers:   headers,
		AccountID: accountID,
	}

	return connection
}

// TODO: include params as a second option
func (c *OandaConnection) Request(endpoint string) []byte {
	client := http.Client {
		Timeout: time.Second * 5, // 5 sec timeout
	}

	url := createUrl(c.Hostname, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkErr(err)
  
  body := makeRequest(c, endpoint, client, req)
	
	return body
}

func (c *OandaConnection) Send(endpoint string, data []byte) []byte {
	client := http.Client{
		Timeout: time.Second * 5, // 5 sec timeout
	}

	url := createUrl(c.Hostname, endpoint)

	// New request object
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
  checkErr(err)

  body := makeRequest(c, endpoint, client, req)

	return body
}

func (c *OandaConnection) Update(endpoint string, data []byte) []byte {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	url := createUrl(c.Hostname, endpoint)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	checkErr(err)

	body := makeRequest(c, endpoint, client, req)

	return body
}
