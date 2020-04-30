package notification

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/conf"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/aftership/aftership-sdk-go/v2/tracking"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	m.Run()
	httpmock.DeactivateAndReset()
}

func TestAddNotification(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := tracking.SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	exp := Data{
		Notification{
			[]string{"vimukthi@aftership.net"},
			[]string{"+85254469627"},
		},
	}
	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/add", p.Slug, p.TrackingNumber), 200, Envelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.AddNotification(p, exp)
	assert.Equal(t, exp, res)
}

func TestAddNotificationError(t *testing.T) {
	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)

	// empty id, slug and tracking_number
	p := tracking.SingleTrackingParam{
		"",
		"",
		"",
		nil,
	}

	_, err := endpoint.AddNotification(p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = tracking.SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/add", p.Slug, p.TrackingNumber), 401, Envelope{
		response.Meta{401, "Invalid API key.", "Unauthorized"},
		Data{},
	}, nil)

	_, err = endpoint.AddNotification(p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestRemoveNotification(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := tracking.SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	exp := Data{
		Notification{
			[]string{"vimukthi@aftership.net"},
			[]string{"+85254469627"},
		},
	}
	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/remove", p.Slug, p.TrackingNumber), 200, Envelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.RemoveNotification(p, exp)
	assert.Equal(t, exp, res)
}

func TestRemoveNotificationError(t *testing.T) {
	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)

	// empty id, slug and tracking_number
	p := tracking.SingleTrackingParam{
		"",
		"",
		"",
		nil,
	}

	_, err := endpoint.RemoveNotification(p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = tracking.SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	mockhttp("POST", fmt.Sprintf("/notifications/%s/%s/remove", p.Slug, p.TrackingNumber), 401, Envelope{
		response.Meta{401, "Invalid API key.", "Unauthorized"},
		Data{},
	}, nil)

	_, err = endpoint.RemoveNotification(p, Data{})
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func TestGetNotificationSetting(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := tracking.SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	exp := Data{
		Notification{
			[]string{"vimukthi@aftership.net"},
			[]string{"+85254469627"},
		},
	}
	mockhttp("GET", fmt.Sprintf("/notifications/%s/%s", p.Slug, p.TrackingNumber), 200, Envelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetNotification(p)
	assert.Equal(t, exp, res)
}

func TestGetotificationError(t *testing.T) {
	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)

	// empty id, slug and tracking_number
	p := tracking.SingleTrackingParam{
		"",
		"",
		"",
		nil,
	}

	_, err := endpoint.GetNotification(p)
	assert.NotNil(t, err)
	assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = tracking.SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	mockhttp("GET", fmt.Sprintf("/notifications/%s/%s", p.Slug, p.TrackingNumber), 401, Envelope{
		response.Meta{401, "Invalid API key.", "Unauthorized"},
		Data{},
	}, nil)

	_, err = endpoint.GetNotification(p)
	assert.NotNil(t, err)
	assert.Equal(t, "Unauthorized", err.Type)
}

func mockhttp(method string, url string, status int, resp interface{}, headers map[string]string) {
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
