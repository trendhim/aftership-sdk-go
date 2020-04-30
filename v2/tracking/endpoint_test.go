package tracking

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/aftership/aftership-sdk-go/v2/conf"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	m.Run()
	httpmock.DeactivateAndReset()
}

func TestCreateTracking(t *testing.T) {
	data := NewTrackingRequest{
		Tracking: NewTracking{
			Slug:           []string{"dhl"},
			TrackingNumber: "tracking_number",
			Title:          "Title Name",
			Smses: []string{
				"+18555072509",
				"+18555072501",
			},
		},
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "dhl",
			Title: "Title Name",
			Smses: []string{
				"+18555072509",
				"+18555072501",
			},
		},
	}

	mockhttp("POST", "https://api.aftership.com/v4/trackings", SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.CreateTracking(data)
	assert.Equal(t, exp, res)
}

func TestDeleteTracking(t *testing.T) {
	p := SingleTrackingParam{
		"",
		"ups",
		"1Z9999999999999998",
		nil,
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "Title Name",
		},
	}

	mockhttp("DELETE", fmt.Sprintf("https://api.aftership.com/v4/trackings/%s/%s", p.Slug, p.TrackingNumber), SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.DeleteTracking(p)
	assert.Equal(t, exp, res)
}

func TestGetTrackings(t *testing.T) {
	p := MultiTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	exp := MultiTrackingsData{
		Page:  1,
		Count: 2,
		Limit: 10,
		Trackings: []Tracking{
			Tracking{
				Slug:  "dhl",
				Title: "Title Name",
			},
			Tracking{
				Slug:  "ups",
				Title: "Title Name",
			},
		},
	}

	mockhttp("GET", fmt.Sprintf("https://api.aftership.com/v4/trackings?page=%d&limit=%d", p.Page, p.Limit), MultiTrackingsEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetTrackings(p)
	assert.Equal(t, exp, res)
}

func TestGetTracking(t *testing.T) {
	p := SingleTrackingParam{
		"",
		"ups",
		"1Z9999999999999998",
		nil,
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "Title Name",
		},
	}

	mockhttp("GET", fmt.Sprintf("https://api.aftership.com/v4/trackings/%s/%s", p.Slug, p.TrackingNumber), SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetTracking(p, GetTrackingParams{})
	assert.Equal(t, exp, res)
}

func TestUpdateTracking(t *testing.T) {
	p := SingleTrackingParam{
		"",
		"ups",
		"1Z9999999999999998",
		nil,
	}

	data := UpdateTrackingRequest{
		UpdateTracking{
			Title: "New Title",
		},
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "New Title",
		},
	}

	mockhttp("PUT", fmt.Sprintf("https://api.aftership.com/v4/trackings/%s/%s", p.Slug, p.TrackingNumber), SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.UpdateTracking(p, data)
	assert.Equal(t, exp, res)
}

func TestReTrack(t *testing.T) {
	p := SingleTrackingParam{
		"",
		"ups",
		"1Z9999999999999998",
		nil,
	}

	exp := SingleTrackingData{
		Tracking: Tracking{
			Slug:  "ups",
			Title: "Title Name",
		},
	}

	mockhttp("POST", fmt.Sprintf("https://api.aftership.com/v4/trackings/%s/%s/retrack", p.Slug, p.TrackingNumber), SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.ReTrack(p)
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
