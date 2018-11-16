package coindesk

import (
	"fmt"
)

// Get Exchange Rates
func (c *Coindesk) GetExchangeRates() (exchangeRates ExchangeRates, err error) {
	reqUrl := fmt.Sprintf("exchangerates")
	_, err = c.client.do("GET", reqUrl, "", &exchangeRates)
	return
}
