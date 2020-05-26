package aftership

import (
	"context"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetCouriers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := CourierList{
		Total: 1,
		Couriers: []Courier{
			{
				Slug: "ups",
				Name: "UPS",
			},
		},
	}

	mockHTTP("GET", "/couriers", 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: exp,
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newCouriersEndpoint(req)
	res, err := endpoint.GetCouriers(context.Background())
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetCouriersError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockHTTP("GET", "/couriers", 429, Response{
		Meta: Meta{
			Code:    429,
			Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
			Type:    "TooManyRequests",
		},
		Data: CourierList{},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newCouriersEndpoint(req)
	_, err := endpoint.GetCouriers(context.Background())
	assert.NotNil(t, err)
	assert.Equal(t, &APIError{
		Code:    429,
		Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
		Type:    "TooManyRequests",
		Path:    "/couriers",
	}, err)
}

func TestGetAllCouriers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := CourierList{
		Total: 1,
		Couriers: []Courier{
			{
				Slug: "ups",
				Name: "ups",
			},
			{
				Slug: "fedex",
				Name: "FeDex",
			},
		},
	}

	mockHTTP("GET", "/couriers/all", 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: exp,
	}, map[string]string{
		"X-RateLimit-Reset":     "1458463600",
		"X-RateLimit-Limit":     "",
		"X-RateLimit-Remaining": "",
	})

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newCouriersEndpoint(req)
	res, _ := endpoint.GetAllCouriers(context.Background())
	assert.Equal(t, exp, res)
}

func TestGetAllCouriersError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockHTTP("GET", "/couriers/all", 429, Response{
		Meta: Meta{
			Code:    429,
			Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
			Type:    "TooManyRequests",
		},
		Data: CourierList{},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newCouriersEndpoint(req)
	_, err := endpoint.GetAllCouriers(context.Background())
	assert.NotNil(t, err)
	assert.Equal(t, &APIError{
		Code:    429,
		Message: "You have exceeded the API call rate limit. Default limit is 10 requests per second.",
		Type:    "TooManyRequests",
		Path:    "/couriers/all",
	}, err)
}

func TestDetectCouriers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	exp := TrackingCouriers{
		Total: 2,
		Couriers: []Courier{
			{
				Slug: "ups",
				Name: "ups",
			},
		},
	}

	mockHTTP("POST", "/couriers/detect", 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: exp,
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newCouriersEndpoint(req)
	result, _ := endpoint.DetectCouriers(context.Background(), CourierDetectionParams{
		"906587618687",
		"DA15BU",
		"20131231",
		"1234567890",
		"",
		"",
		[]string{"dhl", "ups", "fedex"},
	})
	assert.Equal(t, exp, result)
}

func TestInvalidDetectCouriers(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newCouriersEndpoint(req)
	_, err := endpoint.DetectCouriers(context.Background(), CourierDetectionParams{})

	assert.NotNil(t, err)
	assert.Equal(t, ErrorTrackingNumberRequired, err)
}

func TestDetectCouriersError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockHTTP("POST", "/couriers/detect", 401, Response{
		Meta: Meta{
			Code:    402,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: CourierList{},
	}, nil)

	req := newRequestHelper(Config{})
	endpoint := newCouriersEndpoint(req)
	_, err := endpoint.DetectCouriers(context.Background(), CourierDetectionParams{
		TrackingNumber: "906587618687",
	})
	assert.NotNil(t, err)
	assert.Equal(t, &APIError{
		Code:    402,
		Message: "Invalid API key.",
		Type:    "Unauthorized",
		Path:    "/couriers/detect",
	}, err)
}
