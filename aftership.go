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
	// The HTTP client to use when sending requests. Defaults to `http.DefaultClient`.
	HTTPClient *http.Client
}

// Client is the client for all AfterShip API calls
type Client struct {
	Config       Config                // The config of Client SDK
	Courier      CouriersEndpoint      // The endpoint to get a list of supported couriers.
	Tracking     TrackingsEndpoint     // The endpoint to create trackings, update trackings, and get tracking results.
	Checkpoint   CheckpointsEndpoint   // The endpoint to get tracking information of the last checkpoint of a tracking.
	Notification NotificationsEndpoint // The endpoint to get, add or remove contacts (sms or email) to be notified when the status of a tracking has changed.
}

// NewClient returns the AfterShip client
func NewClient(cfg Config) (*Client, error) {
	if cfg.APIKey == "" {
		return nil, errors.New("api key is required")
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = "https://api.aftership.com/v4"
	}

	if cfg.UserAgentPrefix == "" {
		cfg.UserAgentPrefix = "aftership-sdk-go"
	}

	req := newRequestHelper(cfg)
	return &Client{
		Config:       cfg,
		Courier:      newCouriersEndpoint(req),
		Tracking:     newTrackingsEndpoint(req),
		Checkpoint:   newCheckpointsEndpoint(req),
		Notification: newNotificationEndpoint(req),
	}, nil
}
