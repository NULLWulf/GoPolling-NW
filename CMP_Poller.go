package main

import (
	"encoding/json"
	"fmt"
	"github.com/nullwulf/loggly"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"unsafe"
)

const top10CryptoUrl = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=10"
const tag = "Top10Cryptos"

func main() {

	// CMP = Coin Market Pro API
	// URL Endpoint that queries for top most valuable cryptos in USD
	// Loggly Tag

	callCmpApi()
	ticker := time.NewTicker(1 * time.Hour)
	for _ = range ticker.C {
		callCmpApi()
	}
}

func callCmpApi() {
	// Instantiate Loggly Client
	lgglyClient := loggly.New(tag)

	// Instantiate CMP Client and set headers
	cmpClient := &http.Client{}
	req, err := http.NewRequest("GET", top10CryptoUrl, nil)
	req.Header.Add("X-CMC_PRO_API_KEY", os.Getenv("CMP_TOKEN"))
	req.Header.Add("Content-type", "application/json")

	// If error during cmp client initialize output to loggly
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		return
	}

	// Execute API call to CMP
	resp, err := cmpClient.Do(req)
	// If error during client request output error to loggly
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		return
	}

	// Read response body of request and get body size
	body, err := io.ReadAll(resp.Body)
	sz := unsafe.Sizeof(body)
	sz2string := strconv.FormatInt(int64(sz), 10)

	// If error during call output to loggly
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		return
		// Otherwise Unmarshall to CMP data structure
	} else {
		msg := "Successful call to URL: " + top10CryptoUrl + ".\nResponse Body Size: " + sz2string + " bytes."
		lgglyClient.EchoSend("info", msg)
		res := CmpResponse{}
		err := json.Unmarshal(body, &res)
		// If error during marshalling output to loggly
		if err != nil {
			lgglyClient.EchoSend("error", err.Error())
			return
		}
		// Prints Unmarshalled structure in key:value pair format
		fmt.Printf("%+v", res)

	}

	// Gracefully close the client asdad
	err = resp.Body.Close()
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		return
	}
}
