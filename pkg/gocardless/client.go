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
		baseURL: config.AppConfig.GocardlessBaseUrl,
	}
}

type TokenResponse struct {
	Access         string `json:"access"`
	AccessExpires  int    `json:"access_expires"`
	Refresh        string `json:"refresh"`
	RefreshExpires int    `json:"refresh_expires"`
}

func (c *Client) GetAccessToken() (*TokenResponse, error) {
	url := fmt.Sprintf("%s/token/new/", c.baseURL)
	payload := map[string]string{
		"secret_key": c.secretKey,
		"secret_id":  c.secretToken,
	}

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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service error: %d", resp.StatusCode)
	}

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
