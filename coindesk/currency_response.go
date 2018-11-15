package coindesk

type CurrencyAllResp struct {
	Data      CurrencyAll `json:"data"`
	Error     bool        `json:"error"`
	Message   string      `json:"message"`
	Status    int         `json:"status"`
	Timestamp string      `json:"timestamp"`
}

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
