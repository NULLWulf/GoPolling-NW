package main

import (
	"fmt"
	"github.com/nullwulf/loggly"
	"os"
	//loggly "github.com/nullwulf/loggly"
)

func main() {

	//top10CryptoUrl := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=10"
	tag := "Top10Cryptos"

	client := loggly.New(tag)
	fmt.Println(client)
	fmt.Println(os.Getenv("LOGGLY_TOKEN"))
	fmt.Println(os.Getenv("CMP_TOKEN"))
}
