package coindesk

import (
    "errors"
)

// Input for: GET /api/v1/depth
type ExchangeRatesQuery struct {}

func (q *ExchangeRatesQuery) ValidateExchangeRatesQuery() error {
    switch {
    case false:
        return errors.New("Oops! Something went wrong")
    default:
        return nil
    }
}
