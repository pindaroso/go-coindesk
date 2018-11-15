# Go Coindesk API

![CoinDesk](https://img.shields.io/badge/docs-GoDoc-ff69b4.svg?style=flat-square)
[![Travis Branch](https://img.shields.io/travis/pindaroso/go-coindesk.svg?style=flat-square)](https://travis-ci.org/pindaroso.go-coindesk)
[![OpenCode](https://img.shields.io/badge/Open-Code-ff6a00.svg?style=flat-square)](https://opencode18.github.io)

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

    fmt.Println(exchangeRates.Data.BaseCurrency)

    for _, e := range exchangeRates.Data.Rates {
        fmt.Println(e)
    }
}
```

### Local Depth Cache

See `examples/depth.go`.

### TODO

- [ ] Add CI/CD
- [ ] Improve test coverage
- [ ] Add remaining routes e.g. `/v1/currency/{symbol}/score/graph` and `/v1/currency/{symbol}/ticker`
