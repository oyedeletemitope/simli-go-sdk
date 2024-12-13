package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL:    "https://api.simli.ai",
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) post(endpoint string, payload interface{}, response interface{}) error {
	url := c.BaseURL + endpoint
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("API request failed with status: " + resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(response)
}
