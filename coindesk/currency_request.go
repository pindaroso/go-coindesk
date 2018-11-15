package coindesk

import (
	"errors"
)

// Input for: GET /api/v1/currency
type CurrencyQuery struct {
	Convert   string
	Interval  string
	EndDate   string
	StartDate string
}

func (q *CurrencyQuery) ValidateCurrencyQuery() error {
	switch {
	case len(q.Convert) == 0:
		return errors.New("CurrencyQuery must contain a symbol")
	default:
		return nil
	}
}
