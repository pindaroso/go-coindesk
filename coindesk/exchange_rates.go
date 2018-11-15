package coindesk

import (
	"fmt"
)

// Get Exchange Rates
func (c *Coindesk) GetExchangeRates() (exchangeRatesResp ExchangeRatesResp, err error) {
	reqUrl := fmt.Sprintf("exchangerates")
	_, err = c.client.do("GET", reqUrl, "", &exchangeRatesResp)
	return
}
