package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/nullwulf/loggly"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	top10CryptoUrl = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=10"
)

func main() {

	// Loads Environmental variables into program
	// e.g AWS, Loggly CMP token.
	err := godotenv.Load()
	// If detects an error loading .env file terminates program
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initial call of CMP Api
	callCmpApi()
	// Set to call every hour
	ticker := time.NewTicker(15 * time.Minute)
	for _ = range ticker.C {
		callCmpApi()
	}
}

func callCmpApi() {

	// Attemps to get APP_TAG from environment variables file.
	tag := os.Getenv("APP_TAG")
	if tag == "" {
		tag = "Top10Cryptos"
	}
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
	sz := len(body)
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
		err = json.Unmarshal(body, &res)
		// If error during marshalling output to loggly
		// Month - Day - Year - Hour
		res.TimeBlockUTC = time.Now().UTC().Format("01-02-2006-15")
		res.Partition = "Top10Cryptos"
		sort.Slice(res.Data, func(i, j int) bool {
			return res.Data[i].CryptoQuote.USDStats.Price < res.Data[j].CryptoQuote.USDStats.Price
		})
		roundIter(res.Data)

		if err != nil {
			lgglyClient.EchoSend("error", err.Error())
			return
		}
		// Sets cleaner time to Time object

		lgglyClient.EchoSend("info", cryptoStructPrint(res))
		// Prints Unmarshalled structure in key:value pair format
		dynamodbInsert(&res, lgglyClient)

	}

	// Gracefully close the client
	err = resp.Body.Close()
	if err != nil {
		lgglyClient.EchoSend("error", err.Error())
		return
	}
}
