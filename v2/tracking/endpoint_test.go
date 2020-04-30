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
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

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

	mockhttp("POST", "/trackings", 200, SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.CreateTracking(data)
	assert.Equal(t, exp, res)
}

func TestDeleteTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

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

	mockhttp("DELETE", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.DeleteTracking(p)
	assert.Equal(t, exp, res)
}

func TestGetTrackings(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

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

	mockhttp("GET", fmt.Sprintf("/trackings?page=%d&limit=%d", p.Page, p.Limit), 200, MultiTrackingsEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetTrackings(p)
	assert.Equal(t, exp, res)
}

func TestGetTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

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

	mockhttp("GET", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetTracking(p, GetTrackingParams{})
	assert.Equal(t, exp, res)
}

func TestUpdateTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

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

	mockhttp("PUT", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.UpdateTracking(p, data)
	assert.Equal(t, exp, res)
}

func TestReTrack(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

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

	mockhttp("POST", fmt.Sprintf("/trackings/%s/%s/retrack", p.Slug, p.TrackingNumber), 200, SingleTrackingEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, _ := endpoint.ReTrack(p)
	assert.Equal(t, exp, res)
}

func TestBuildTrackingUrl(t *testing.T) {
	// slug and tracking number
	p := SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		nil,
	}

	url, err := BuildTrackingURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN", url)

	// slug and tracking number, with subpath
	url, err = BuildTrackingURL(p, "trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN/retrack", url)

	// slug and tracking number, with optional parameters
	p = SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
		&SingleTrackingOptionalParams{
			TrackingPostalCode: "1234",
		},
	}
	url, err = BuildTrackingURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN?tracking_postal_code=1234", url)

	// slug and tracking number, with optional parameters and subpath
	url, err = BuildTrackingURL(p, "trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/xq-express/LS404494276CN/retrack?tracking_postal_code=1234", url)

	// id
	p = SingleTrackingParam{
		ID: "1234567890",
	}

	url, err = BuildTrackingURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890", url)

	// id, with subpath
	url, err = BuildTrackingURL(p, "trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890/retrack", url)

	// id, with optional parameters
	p = SingleTrackingParam{
		ID: "1234567890",
		OptionalParams: &SingleTrackingOptionalParams{
			TrackingPostalCode: "1234",
		},
	}
	url, err = BuildTrackingURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890?tracking_postal_code=1234", url)

	// id, with optional parameters and subpath
	url, err = BuildTrackingURL(p, "trackings", "retrack")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/1234567890/retrack?tracking_postal_code=1234", url)

	// Encode slug and tracking number
	p = SingleTrackingParam{
		"",
		"usps",
		"ABCD/1234",
		nil,
	}

	url, err = BuildTrackingURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/trackings/usps/ABCD%2F1234", url)
}

func TestBuildTrackingUrlError(t *testing.T) {
	// should get error when no id, slug and tracking number
	p := SingleTrackingParam{
		"",
		"",
		"",
		nil,
	}
	_, err := BuildTrackingURL(p, "", "")
	assert.NotNil(t, err)

	// should get error when only slug
	p = SingleTrackingParam{
		"",
		"xq-express",
		"",
		nil,
	}
	_, err = BuildTrackingURL(p, "", "")
	assert.NotNil(t, err)

	// should get error when only tracking_number
	p = SingleTrackingParam{
		"",
		"",
		"LS404494276CN",
		nil,
	}
	_, err = BuildTrackingURL(p, "", "")
	assert.NotNil(t, err)
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
