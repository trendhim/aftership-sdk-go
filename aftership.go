package aftership

import (
	"errors"
	"net/http"
)

// Config is the config of AfterShip SDK client
type Config struct {
	// APIKey.
	APIKey string

	// Authentication Type
	AuthenticationType AuthenticationType

	// apiSecret
	// if AuthenticationType is RSA, use rsa private key
	// if AuthenticationType is AES, use aes api secret
	APISecret string

	// BaseURL is the base URL of AfterShip API. Defaults to 'https://api.aftership.com/tracking/2023-10'
	BaseURL string

	// UserAgentPrefix is the prefix of User-Agent in headers. Defaults to 'aftership-sdk-go'
	UserAgentPrefix string

	// HTTPClient is the HTTP client to use when making requests. Defaults to http.DefaultClient.
	HTTPClient *http.Client
}

// Client is the client for all AfterShip API calls
type Client struct {
	// The config of Client SDK
	Config Config
	// The HTTP client to use when sending requests. Defaults to `http.DefaultClient`.
	httpClient *http.Client
	// Rate limit
	rateLimit *RateLimit
}

// NewClient returns the AfterShip client
func NewClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, errors.New(errEmptyAPIKey)
	}

	if cfg.AuthenticationType == AES {
		if cfg.APISecret == "" {
			return nil, errors.New(errEmptyAPISecret)
		}
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.aftership.com/tracking/2023-10"
	}

	if cfg.UserAgentPrefix == "" {
		cfg.UserAgentPrefix = "aftership-sdk-go"
	}

	client := &Client{
		Config:     cfg,
		rateLimit:  &RateLimit{},
		httpClient: http.DefaultClient,
	}

	if cfg.HTTPClient != nil {
		client.httpClient = cfg.HTTPClient
	}

	return client, nil
}

// GetRateLimit returns the X-RateLimit value in API response headers
func (client *Client) GetRateLimit() RateLimit {
	return *client.rateLimit
}
