package gocardless

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iacopoghilardi/mynance-service-api/internal/config"
)

var GoCardlessClient Client

const (
	baseURL = "https://ob.nordigen.com/api/v2"
)

type Client struct {
	secretKey   string
	secretToken string
}

func init() {
	configs := config.AppConfig

	GoCardlessClient = *NewClient(configs.GocardlessSecret, config.AppConfig.GocardlessToken)
}

func NewClient(secretKey string, secretToken string) *Client {
	return &Client{
		secretKey:   secretKey,
		secretToken: secretToken,
	}
}

type TokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (c *Client) GetAccessToken() (*TokenResponse, error) {
	url := fmt.Sprintf("%s/token/new/", baseURL)
	payload := map[string]string{
		"secret_id": c.secretKey,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResponse TokenResponse
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
