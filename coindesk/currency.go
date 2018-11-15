package coindesk

import (
	"fmt"
)

func (c *Coindesk) GetCurrencyAll(q CurrencyQuery) (currency CurrencyAllResp, err error) {
	err = q.ValidateCurrencyQuery()
	if err != nil {
		return
	}

	reqUrl := fmt.Sprintf("currency/all?graph=true&start_date=%s&end_date=%s&interval=%s&convert=%s", q.StartDate, q.EndDate, q.Interval, q.Convert)
	_, err = c.client.do("GET", reqUrl, "", &currency)

	return
}
