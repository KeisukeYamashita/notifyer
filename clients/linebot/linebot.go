package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Client ...
type Client struct {
	URL         string
	AccessToken string
	To          string
}

// NewClient ...
func NewClient(accessToken, to string) *Client {
	return &Client{
		URL:         "https://api.line.me/v2/bot/message/push",
		AccessToken: accessToken,
		To:          to,
	}
}

// Payload ...
type Payload struct {
	Messages []*Message `json:"messages"`
	To       string     `json:"to"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Send ...
func (c *Client) Send(msg string) error {
	httpClient := http.DefaultClient
	payload := &Payload{
		Messages: []*Message{&Message{
			Type: "text",
			Text: strings.TrimSuffix(msg, "\n"),
		}},
		To: c.To,
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code was not 200, got:%d", resp.StatusCode)
	}

	return nil
}
