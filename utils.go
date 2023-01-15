package goanda

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

func checkApiErr(body []byte, route string) {
	bodyString := string(body[:])
	if strings.Contains(bodyString, "errorMessage") {
		panic("\nOANDA API Error: " + bodyString + "\nOn route: " + route)
	}
}

func marshalJson(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	//checkErr(err)
	
	return bytes, err
}

func unmarshalJson(body []byte, data interface{}) (error) {
	jsonErr := json.Unmarshal(body, &data)
  return jsonErr
	//checkErr(jsonErr)
}

func createUrl(host string, endpoint string) string {
	var buffer bytes.Buffer
	// Generate the auth header
	buffer.WriteString(host)
	buffer.WriteString(endpoint)

	url := buffer.String()
	return url
}

func makeRequest(c *OandaConnection, endpoint string, client http.Client, req *http.Request) ([]byte, error, error) {
	req.Header.Set("User-Agent", c.Headers.agent)
	req.Header.Set("Authorization", c.Headers.auth)
	req.Header.Set("Content-Type", c.Headers.contentType)

	res, getErr := client.Do(req)
	//checkErr(getErr)
	body, readErr := ioutil.ReadAll(res.Body)
	//checkErr(readErr)
	checkApiErr(body, endpoint)

	return body, getErr, readErr
}
