package main

import "math"

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
	Status RespStatus      `json:"status"`
	Data   []CryptoElement `json:"data"`
}

type RespStatus struct {
	Timestamp string `json:"timestamp"`
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

func roundIter(statuses []CryptoElement) {
	for i := 0; i < len(statuses); i++ {
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
