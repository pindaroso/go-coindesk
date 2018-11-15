package main

import (
	"fmt"

	"github.com/go-coindesk/coindesk"
)

func main() {
	client := coindesk.New()
	exchangeRates, err := client.GetExchangeRates()

	if err != nil {
		panic(err)
	}

	fmt.Println(exchangeRates.Data.BaseCurrency)

	for _, e := range exchangeRates.Data.Rates {
		fmt.Println(e)
	}
}
