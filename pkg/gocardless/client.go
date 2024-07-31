package gocardless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iacopoghilardi/mynance-service-api/internal/config"
)

type Client struct {
	secretKey   string
	secretToken string
	httpClient  *http.Client
	baseURL     string
}

func NewClient(secretKey, secretToken string) *Client {
	return &Client{
		secretKey:   secretKey,
		secretToken: secretToken,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: "https://bankaccountdata.gocardless.com/api/v2",
	}
}

type TokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (c *Client) GetAccessToken() (*TokenResponse, error) {
	url := fmt.Sprintf("%s/token/new/", c.baseURL)
	payload := map[string]string{
		"secret_id":  c.secretKey,
		"secret_key": c.secretToken,
	}

	fmt.Printf("SECRET: %v", c.secretKey)
	fmt.Printf("SECRET: %v", c.secretToken)

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("STATUS CODE: %v", resp.Status)
	fmt.Printf("STATUS CODE: %v", resp.Body)

	// if resp.StatusCode != http.StatusOK {
	// 	return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	// }

	var tokenResponse TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &tokenResponse, nil
}

func NewGoCardlessClient() *Client {
	configs := config.AppConfig
	return NewClient(configs.GocardlessSecret, configs.GocardlessToken)
}
