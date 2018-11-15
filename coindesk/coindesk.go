package coindesk

const (
    // https://production.api.coindesk.com/v1/exchangeRates
    // https://production.api.coindesk.com/v1/currency/all?graph=true&start_date=2018-11-14&end_date=2018-11-15&interval=120-min&convert=USD
    // https://production.api.coindesk.com/v1/currency/BTC/score/graph?currency=ADA,BCH,BTC,BTG,DASH,EOS,ETC,ETH,IOTA,LSK,LTC,NEO,QTUM,TRX,XEM,XLM,XMR,XRP,ZEC&start_date=2018-10-17&end_date=2018-11-15
    // https://production.api.coindesk.com/v1/currency/XMR/ticker
    BaseUrl = "https://production.api.coindesk.com/v1"
)

type Coindesk struct {
    client *Client
}

func New() *Coindesk {
    client := NewClient()
    return &Coindesk{client}
}
