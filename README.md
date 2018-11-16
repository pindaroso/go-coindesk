# Go Coindesk API

[![Travis Branch](https://img.shields.io/travis/pindaroso/go-coindesk.svg?style=flat-square)](https://travis-ci.org/pindaroso/go-coindesk.svg?branch=master)
[![GoReport Widget]][GoReport Status]
[![GoDoc](https://godoc.org/github.com/pindaroso/go-coindesk/github?status.svg)](https://godoc.org/github.com/pindaroso/go-coindesk/github)

[GoReport Status]: https://goreportcard.com/report/github.com/pindaroso/go-coindesk?cache
[GoReport Widget]: https://goreportcard.com/badge/github.com/pindaroso/go-coindesk?cache

## Summary

Go client for [Coindesk](https://www.coindesk.com)

## Installation

`go get github.com/pindaroso/go-coindesk/coindesk`

## Documentation

Full API Documentation can be found at https://www.coindesk.com/data/glossary

## Setup

Creating a client:

```go
import (
    "os"

    "go-coindesk/coindesk"
)

client :=  coindesk.New()
```

## Examples

### Get Exchange Rates

```go
package main

import (
    "fmt"

    "go-coindesk/coindesk"
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
```
