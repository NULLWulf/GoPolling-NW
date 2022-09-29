package main

// Flattened structure of CMP crypto data
type cryptoStats struct {
	name        string
	symbol      string
	totalSupply int
	cmpId       int
	perChange   USD
}

// CmpResponse Response body of CMP api, only interested in data array of crypto objects
type CmpResponse struct {
	Data []CryptoElem `json:"data"`
}

// CryptoElem data type containing nominal and statistical data
type CryptoElem struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	TotalSupply int    `json:"total_supply"`
	CmcRank     int    `json:"cmc_rank"`
	CryptoQuote Quote  `json:"quote"`
}

// Quote containing data relative data to respective queried currency in this case USD
type Quote struct {
	USDStats USD `json:"USD"`
}

// USD United States dollar relative statistics
type USD struct {
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
