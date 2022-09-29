package main

import (
	"fmt"
	"github.com/nullwulf/loggly"
	"io"
	"net/http"
	"os"
	"strconv"
	"unsafe"
)

func main() {

	// CMP = Coin Market Pro API
	// URL Endpoint that queries for top most valuable cryptos in USD
	top10CryptoUrl := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=10"
	// Loggly Tag
	tag := "Top10Cryptos"

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

	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		return
	} else {
		msg := "Successful call to URL: " + top10CryptoUrl + ".\nResponse Body Size: " + sz2string + " bytes."
		lgglyClient.EchoSend("info", msg)
		fmt.Println(string(body))
	}
}
