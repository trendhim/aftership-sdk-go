package aftership

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the API client being tested
	client *Client

	// server is a test HTTP server used to provide mock API responses
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// Cloudflare client configured to use test server
	client, _ = NewClient(Config{
		APIKey:  "YOUR_API_KEY",
		BaseURL: server.URL,
	})
}

func teardown() {
	server.Close()
}

func TestInvalidAPIKey(t *testing.T) {
	// API Key is not specified
	_, err := NewClient(Config{})
	assert.NotNil(t, err)
	assert.Equal(t, errEmptyAPIKey, err.Error())
}

func TestDefaultConfig(t *testing.T) {
	// API Key is specified
	client, err := NewClient(Config{
		APIKey: "YOUR_API_KEY",
	})

	assert.Nil(t, err)
	assert.Equal(t, "YOUR_API_KEY", client.Config.APIKey)
	assert.Equal(t, "https://api.aftership.com/v4", client.Config.BaseURL)
	assert.Equal(t, "aftership-sdk-go", client.Config.UserAgentPrefix)
}

func TestSpecifiedConfig(t *testing.T) {
	apiKey := "YOUR_API_KEY"
	endpoint := "YOUR_ENDPOINT"
	agent := "YOUR_AGENT"

	client, err := NewClient(Config{
		APIKey:          apiKey,
		BaseURL:         endpoint,
		UserAgentPrefix: agent,
	})

	assert.Nil(t, err)
	assert.Equal(t, apiKey, client.Config.APIKey)
	assert.Equal(t, endpoint, client.Config.BaseURL)
	assert.Equal(t, agent, client.Config.UserAgentPrefix)
}
