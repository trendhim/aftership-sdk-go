package aftership

import (
	"context"
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
	assert.Equal(t, "https://api.aftership.com/2023-10", client.Config.BaseURL)
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

func TestGetRateLimit(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("x-ratelimit-reset", "1458463600")
		w.Header().Set("x-ratelimit-limit", "10")
		w.Header().Set("x-ratelimit-remaining", "9")
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte(`{
			"meta": {
					"code": 429,
					"type": "TooManyRequests",
					"message": "You have exceeded the API call rate limit. Default limit is 10 requests per second."
			},
			"data": {}
		}`))
	})

	exp := RateLimit{
		Reset:     int64(1458463600),
		Limit:     10,
		Remaining: 9,
	}

	var result mockData
	err := client.makeRequest(context.Background(), http.MethodGet, "/test", nil, nil, &result)

	assert.NotNil(t, err)
	assert.Equal(t, exp, client.GetRateLimit())
}
