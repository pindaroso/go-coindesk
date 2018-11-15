package ping_test

import (
    "github.com/pindaroso/go-coindesk/coindesk"

    "testing"
)

func TestExchangeRate(t *testing.T) {
    client := coindesk.New()

    if resp, err := client.GetExchangeRates(); err != nil {
        t.Fatal(err)
    } else {
        t.Logf("%+v\n", ping)
    }

}
