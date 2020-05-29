package aftership

import (
	"errors"
	"net/http"
)

// Config is the config of AfterShip SDK client
type Config struct {
	// APIKey.
	APIKey string
	// BaseURL is the base URL of AfterShip API. Defaults to 'https://api.aftership.com/v4'
	BaseURL string
	// UserAgentPrefix is the prefix of User-Agent in headers. Defaults to 'aftership-sdk-go'
	UserAgentPrefix string
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

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.aftership.com/v4"
	}

	if cfg.UserAgentPrefix == "" {
		cfg.UserAgentPrefix = "aftership-sdk-go"
	}

	client := &Client{
		Config:    cfg,
		rateLimit: &RateLimit{},
	}

	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
	}

	return client, nil
}
