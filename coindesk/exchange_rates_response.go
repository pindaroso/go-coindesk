package coindesk

type ExchangeRatesResp struct {
    Data        ExchangeRates   `json:"data"`
    Error       bool            `json:"error"`
    Message     string          `json:"message"`
    Status      int             `json:"status"`
    Timestamp   string          `json:"timestamp"`
}

// Result from: GET /api/v1/exchangerates
type ExchangeRates struct {
    BaseCurrency    Currency    `json:"baseCurrency"`
    Rates           []Rate      `json:"rates"`
}

type Currency struct {
    Iso     string  `json:",iso"`
    Name    string  `json:",name"`
    Slug    string  `json:",slug"`
}

type Rate struct {
    Iso     string  `json:",iso"`
    Name    string  `json:",name"`
    Slug    string  `json:",slug"`
    Rate    float64 `json:",rate"`
}
