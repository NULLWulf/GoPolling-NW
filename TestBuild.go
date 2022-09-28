package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//loggly "github.com/nullwulf/loggly"
)

func main() {

	//lgglyKey := os.Getenv("LOGGLY_TOKEN")
	cmpKey := os.Getenv("CMP_TOKEN")

	top10CryptoUrl := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=10"
	//tag := "Top10Cryptos"

	//lgglyClient := loggly.New(tag)

	cmpClient := http.Client{}
	req, err := http.NewRequest("GET", top10CryptoUrl, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("X-CMC_PRO_API_KEY", cmpKey)
	req.Header.Add("Content-type", "application/json")

	resp, err := cmpClient.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}
