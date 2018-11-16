package coindesk

// Result from: GET /api/v1/currency/all
type CurrencyAll struct {
	CoindeskMarketCap CoindeskMarketCap          `json:"coindeskMarketCap"`
	Currency          map[string]CurrencyVerbose `json:"currency"`
}

type CoindeskMarketCap struct {
	FactorConfigId int     `json:"FactorConfigId"`
	CreatedAt      string  `json:"createdAt"`
	Data           float64 `json:"data"`
}

type CurrencyVerbose struct {
	CirculatingSupply float64 `json:"circulatingSupply"`
	Iso               string  `json:"iso"`
	Name              string  `json:"name"`
}
