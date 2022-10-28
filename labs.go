package goanda

type Spreads struct {

  Max   [][]float64 `json:"max"`
  Avg   [][]float64 `json:"avg"`
  Min   [][]float64 `json:"min"`

}

func (c *OandaConnection) GetOrderBookData(instrument string, period string) interface{} {
	var json_data interface{}
  endpoint := "labs/v1/orderbook_data?instrument=" + instrument + "&period=" + period
	data := c.Request(endpoint)
	unmarshalJson(data, &json_data)

	return json_data
}

func (c *OandaConnection) GetSpreads(instrument string, period string) Spreads {
  endpoint := "labs/v1/spreads?instrument=" + instrument + "&period=" + period
	data := c.Request(endpoint)
  spreads := Spreads{}
	unmarshalJson(data, &spreads)

	return spreads
}
