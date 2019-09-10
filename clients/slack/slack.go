package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client ...
type Client struct {
	URL string
}

// NewClient ...
func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

// Payload ...
type Payload struct {
	Text string `json:"text"`
}

// Send ...
func (c *Client) Send(msg string) error {
	httpClient := http.DefaultClient
	payload := &Payload{
		Text: msg,
	}

	json, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code was not 200, got:%d", resp.StatusCode)
	}

	return nil
}
