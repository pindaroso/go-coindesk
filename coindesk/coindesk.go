package coindesk

const (
	BaseUrl = "https://production.api.coindesk.com/v1"
)

type Coindesk struct {
	client *Client
}

func New() *Coindesk {
	client := NewClient()
	return &Coindesk{client}
}
