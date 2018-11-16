package coindesk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

type BadRequest struct {
	code int64  `json:"code"`
	msg  string `json:"msg,required"`
}

type GoodRequest struct {
	Data      json.RawMessage `json:"data"`
	Error     bool            `json:"error"`
	Message   string          `json:"message"`
	Status    int             `json:"status"`
	Timestamp string          `json:"timestamp"`
}

func handleError(resp *http.Response) error {
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("Bad response Status %s. Response Body: %s", resp.Status, string(body))
	}
	return nil
}

func parseData(resp *http.Response) (_ *bytes.Buffer, err error) {
	req := &GoodRequest{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(req)
	return bytes.NewBuffer(req.Data), err
}

// Creates a new Coindesk HTTP Client
func NewClient() (c *Client) {
	client := &Client{
		httpClient: &http.Client{},
	}
	return client
}

func (c *Client) do(method, resource, payload string, result interface{}) (resp *http.Response, err error) {
	fullUrl := fmt.Sprintf("%s/%s", BaseUrl, resource)

	req, err := http.NewRequest(method, fullUrl, nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// Check for error
	err = handleError(resp)
	if err != nil {
		return
	}

	data, err := parseData(resp)
	if err != nil {
		return
	}

	// Process response
	if resp != nil {
		decoder := json.NewDecoder(data)
		err = decoder.Decode(result)
	}

	return
}
