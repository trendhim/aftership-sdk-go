package aftership

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidAPIKey(t *testing.T) {
	// API Key is not specified
	_, err := NewClient(Config{})
	assert.NotNil(t, err)

	expectedMsg := "api key is required"
	assert.Equal(t, expectedMsg, err.Error())
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
