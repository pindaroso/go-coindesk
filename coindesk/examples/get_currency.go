package main

import (
    "fmt"

    "github.com/go-coindesk/coindesk"
)

func main() {
    client := coindesk.New()

    q := coindesk.CurrencyQuery{StartDate: "2018-11-14", EndDate: "2018-11-15", Interval: "120-min", Convert: "USD"}
    currency, err := client.GetCurrencyAll(q)

    if err != nil {
        panic(err)
    }

    fmt.Println(currency.Data)
}
