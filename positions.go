package goanda

import "encoding/json"

// Supporting OANDA docs - http://developer.oanda.com/rest-live-v20/position-ep/

type OpenPositions struct {
	LastTransactionID string `json:"lastTransactionID"`
	Positions         []struct {
		Instrument string `json:"instrument"`
		Long       struct {
			AveragePrice string   `json:"averagePrice"`
			Pl           string   `json:"pl"`
			ResettablePL string   `json:"resettablePL"`
			TradeIDs     []string `json:"tradeIDs"`
			Units        string   `json:"units"`
			UnrealizedPL string   `json:"unrealizedPL"`
		} `json:"long"`
		Pl           string `json:"pl"`
		ResettablePL string `json:"resettablePL"`
		Short        struct {
			AveragePrice string   `json:"averagePrice"`
			Pl           string   `json:"pl"`
			ResettablePL string   `json:"resettablePL"`
			TradeIDs     []string `json:"tradeIDs"`
			Units        string   `json:"units"`
			UnrealizedPL string   `json:"unrealizedPL"`
		} `json:"short"`
		UnrealizedPL string `json:"unrealizedPL"`
	} `json:"positions"`
}

type ClosePositionPayload struct {
	LongUnits  string `json:"longUnits,omitempty"`
	ShortUnits string `json:"shortUnits,omitempty"`
}

func (c *OandaConnection) GetOpenPositions() (OpenPositions, error, error) {
	endpoint := "v3/accounts/" + c.AccountID + "/openPositions"

	response, err1, err2 := c.Request(endpoint)
	data := OpenPositions{}
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}

func (c *OandaConnection) ClosePosition(instrument string, body ClosePositionPayload) (ModifiedTrade, error, error) {
	endpoint := "v3/accounts/" + c.AccountID + "/positions/" + instrument + "/close"
	jsonBody, _ := json.Marshal(body)
	response, err1, err2 := c.Update(endpoint, jsonBody)
	data := make(map[string]string)
	_ = unmarshalJson(response, &data)

	return data, err1, err2
}
