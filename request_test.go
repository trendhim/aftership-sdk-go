package aftership

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

type mockData struct {
	Meta Meta   `json:"meta"`
	Data string `json:"data"`
}

func TestMakeRequest(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := mockData{
		Data: "test",
	}

	// GET with status 200
	mockHTTP("GET", "/test", 200, exp, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})

	var result string
	err := req.makeRequest(context.Background(), "GET", "/test", nil, nil, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp.Data, result)

	// POST with status 201
	mockHTTP("POST", "/test", 201, exp, nil)

	err = req.makeRequest(context.Background(), "POST", "/test", nil, nil, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp.Data, result)

	// PUT with status 200
	mockHTTP("POST", "/test", 200, exp, nil)

	err = req.makeRequest(context.Background(), "POST", "/test", nil, exp, &result)
	assert.Nil(t, err)
	assert.Equal(t, exp.Data, result)
}

func TestNewRequestError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})

	var result string
	// Invalid data
	err := req.makeRequest(context.Background(), "GET", "/test", nil, make(chan int), &result)
	assert.NotNil(t, err)
	//assert.Equal(t, "JsonError", err.Type)

	// Bad method
	err = req.makeRequest(context.Background(), "bad method", "/test", nil, nil, &result)
	assert.NotNil(t, err)
	//assert.Equal(t, "RequestError", err.Type)

	// 500 err
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockHTTP("GET", "/test", 500, nil, nil)

	err = req.makeRequest(context.Background(), "GET", "/test", nil, nil, &result)
	assert.NotNil(t, err)
	// assert.Equal(t, "InternalError", err.Type)

	// String response
	mockStringResponse("GET", "/test")

	err = req.makeRequest(context.Background(), "GET", "/test", nil, nil, &result)
	assert.NotNil(t, err)
	// assert.Equal(t, "RequestError", err.Type)

	// APIError response
	mockErrorResponse("GET", "/test")

	err = req.makeRequest(context.Background(), "GET", "/test", nil, nil, &result)
	assert.NotNil(t, err)
	// assert.Equal(t, "RequestError", err.Type)

	// Invalid URI
	req = newRequestHelper(Config{
		APIKey:  "YOUR_API_KEY",
		BaseURL: "/this/field/is/illegal/and/should/error/",
	})

	err = req.makeRequest(context.Background(), "GET", "/test", nil, nil, &result)
	assert.NotNil(t, err)
	// assert.Equal(t, "RequestError", err.Type)
}

func TestRateLimit(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := mockData{
		Meta: Meta{
			Code:    429,
			Type:    "TooManyRequests",
			Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
		},
		Data: "test",
	}
	mockHTTP("GET", "/test", 429, exp, map[string]string{
		"x-ratelimit-reset":     "1458463600",
		"x-ratelimit-limit":     "10",
		"x-ratelimit-remaining": "9",
	})

	// rateLimit := &RateLimit{}
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})

	var result mockData
	err := req.makeRequest(context.Background(), "GET", "/test", nil, nil, &result)
	assert.NotNil(t, err)
	// assert.Equal(t, "TooManyRequests", err.Type)
	assert.Equal(t, int64(1458463600), req.getRateLimit().Reset)
	assert.Equal(t, 10, req.getRateLimit().Limit)
	assert.Equal(t, 9, req.getRateLimit().Remaining)
}

func mockHTTP(method string, url string, status int, resp interface{}, headers map[string]string) {
	httpmock.RegisterResponder(method, url,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(status, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			for key, value := range headers {
				resp.Header.Set(key, value)
			}
			return resp, nil
		},
	)
}

func mockStringResponse(method string, url string) {
	httpmock.RegisterResponder(method, url,
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, "test")
			return resp, nil
		},
	)
}

func mockErrorResponse(method string, url string) {
	httpmock.RegisterResponder(method, url, httpmock.NewErrorResponder(errors.New("APIError response")))
}
