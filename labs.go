package goanda

func (c *OandaConnection) GetOrderBookData(instrument string, period string) interface{} {
	var json_data interface{}
  endpoint := "labs/v1/orderbook_data?instrument=" + instrument + "&period=" + period
	data := c.Request(endpoint)
	unmarshalJson(data, &json_data)

	return json_data
}
