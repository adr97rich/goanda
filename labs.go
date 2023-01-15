package goanda

type Spreads struct {

  Max   [][]float64 `json:"max"`
  Avg   [][]float64 `json:"avg"`
  Min   [][]float64 `json:"min"`

}

func (c *OandaConnection) GetOrderBookData(instrument string, period string) (interface{}, error, error) {
	var json_data interface{}
  endpoint := "labs/v1/orderbook_data?instrument=" + instrument + "&period=" + period
	data, err1, err2 := c.Request(endpoint)
	_ = unmarshalJson(data, &json_data)

	return json_data, err1, err2
}

func (c *OandaConnection) GetSpreads(instrument string, period string) (Spreads, error, error) {
  endpoint := "labs/v1/spreads?instrument=" + instrument + "&period=" + period
	data, err1, err2 := c.Request(endpoint)
  spreads := Spreads{}
	_ = unmarshalJson(data, &spreads)

	return spreads, err1, err2
}
