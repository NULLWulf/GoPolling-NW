package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Flattened structure of CMP crypto data
type cryptoStats struct {
	name        string
	symbol      string
	totalSupply int
	cmpId       int
	perChange   USDRelativeData
}

// CmpResponse Response body of CMP api, only interested in data array of crypto objects
type CmpResponse struct {
	TimeBlockUTC string
	Partition    string
	Data         []CryptoElement `json:"data"`
}

// CryptoElement data type containing nominal and statistical data
type CryptoElement struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	CmcRank     int    `json:"cmc_rank"`
	CryptoQuote Quote  `json:"quote"`
}

// Quote containing data relative data to respective queried currency in this case USD
type Quote struct {
	USDStats USDRelativeData `json:"USD"`
}

// USDRelativeData United States Dollar Relative Data
type USDRelativeData struct {
	Price             float64 `json:"price"`
	Volume24hr        float64 `json:"volume_24h"`
	VolumeChange24hr  float64 `json:"volume_change_24h"`
	PercentChange1hr  float64 `json:"percent_change_1h"`
	PercentChange24hr float64 `json:"percent_change_24h"`
	PercentChange7d   float64 `json:"percent_change_7d"`
	PercentChange30d  float64 `json:"percent_change_30d"`
	PercentChange60d  float64 `json:"percent_change_60d"`
	PercentChange90d  float64 `json:"percent_change_90d"`
}

// Round to 3 decimal places at most
func r(r float64) float64 {
	return math.Round(r*1000) / 1000
}

// Rounds values to float and also assigns a rank based on position in the structure
// assumes array is sorted beforehand
func roundIter(statuses []CryptoElement) {
	rankCounter := 10
	for i := 0; i < len(statuses); i++ {
		statuses[i].CmcRank = rankCounter
		rankCounter--
		statuses[i].CryptoQuote.USDStats.Price = r(statuses[i].CryptoQuote.USDStats.Price)
		statuses[i].CryptoQuote.USDStats.Volume24hr = r(statuses[i].CryptoQuote.USDStats.Volume24hr)
		statuses[i].CryptoQuote.USDStats.VolumeChange24hr = r(statuses[i].CryptoQuote.USDStats.VolumeChange24hr)
		statuses[i].CryptoQuote.USDStats.PercentChange1hr = r(statuses[i].CryptoQuote.USDStats.PercentChange1hr)
		statuses[i].CryptoQuote.USDStats.PercentChange24hr = r(statuses[i].CryptoQuote.USDStats.PercentChange24hr)
		statuses[i].CryptoQuote.USDStats.PercentChange7d = r(statuses[i].CryptoQuote.USDStats.PercentChange7d)
		statuses[i].CryptoQuote.USDStats.PercentChange30d = r(statuses[i].CryptoQuote.USDStats.PercentChange30d)
		statuses[i].CryptoQuote.USDStats.PercentChange60d = r(statuses[i].CryptoQuote.USDStats.PercentChange60d)
		statuses[i].CryptoQuote.USDStats.PercentChange90d = r(statuses[i].CryptoQuote.USDStats.PercentChange90d)
	}
}

func cryptoStructPrint(cryptoStruct CmpResponse) string {
	var b strings.Builder
	fmt.Fprintf(&b, "----==== Displaying Top 10 Ranked Cryptos per CoinMarket PRO API ====----\n")
	fmt.Fprintf(&b, "----====----====----====----====----====----===----===----===----===----===---- \n")
	fmt.Fprintf(&b, "---=== Actual Time Polled %v  ===---\n", time.Now().UTC().Format(time.RFC3339))
	fmt.Fprintf(&b, "Hourly Time Block: %v\n", cryptoStruct.TimeBlockUTC)
	fmt.Fprintf(&b, "\n")
	for i := 0; i < len(cryptoStruct.Data); i++ {
		p := cryptoStruct.Data[i].CryptoQuote.USDStats
		fmt.Fprintf(&b, "---=== Rank %v : %v ===---\n", cryptoStruct.Data[i].CmcRank, cryptoStruct.Data[i].Name)
		fmt.Fprintf(&b, "Price: $%v\n", p.Price)
		fmt.Fprintf(&b, "Volume 24hr: %v\n", p.Volume24hr)
		fmt.Fprintf(&b, "Volume Change 24hr: %v\n", p.VolumeChange24hr)
		fmt.Fprintf(&b, "--== Relative Movement ==--\n")
		fmt.Fprintf(&b, "1 Hourr: %v%%\n", p.PercentChange1hr)
		fmt.Fprintf(&b, "24 Hour: %v%%\n", p.PercentChange24hr)
		fmt.Fprintf(&b, "7 Day: %v%%\n", p.PercentChange7d)
		fmt.Fprintf(&b, "30 Day: %v%%\n", p.PercentChange30d)
		fmt.Fprintf(&b, "60 Day: %v%%\n", p.PercentChange60d)
		fmt.Fprintf(&b, "90 Day: %v%%\n", p.PercentChange90d)
		fmt.Fprintf(&b, "\n")
	}
	//fmt.Println(b.String())

	return b.String()
}
