package main

import (
	"fmt"

	"github.com/pindaroso/go-coindesk/coindesk"
)

func main() {
	client := coindesk.New()
	exchangeRates, err := client.GetExchangeRates()

	if err != nil {
		panic(err)
	}

	fmt.Println(exchangeRates.BaseCurrency)

	for _, e := range exchangeRates.Rates {
		fmt.Println(e)
	}
}
