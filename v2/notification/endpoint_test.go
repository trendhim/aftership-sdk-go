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
	mockhttp("POST", fmt.Sprintf("https://api.aftership.com/v4/notifications/%s/%s/add", p.Slug, p.TrackingNumber), Envelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.AddNotification(p, exp)
	assert.Equal(t, exp, res)
}

func TestRemoveNotification(t *testing.T) {
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
	mockhttp("POST", fmt.Sprintf("https://api.aftership.com/v4/notifications/%s/%s/remove", p.Slug, p.TrackingNumber), Envelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.RemoveNotification(p, exp)
	assert.Equal(t, exp, res)
}

func TestGetNotificationSetting(t *testing.T) {
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
	mockhttp("GET", fmt.Sprintf("https://api.aftership.com/v4/notifications/%s/%s", p.Slug, p.TrackingNumber), Envelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetNotification(p)
	assert.Equal(t, exp, res)
}

func mockhttp(method string, url string, resp interface{}, headers map[string]string) {
	httpmock.RegisterResponder(method, url,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
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
