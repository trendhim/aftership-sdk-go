package courier

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

func TestGetCouriers(t *testing.T) {
	exp := []Courier{
		Courier{
			Slug: "ups",
			Name: "UPS",
		},
	}
	mockhttp("GET", "https://api.aftership.com/v4/couriers", Envelope{
		response.Meta{200, "", ""},
		List{
			Total:    1,
			Couriers: exp,
		},
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetCouriers()
	assert.Equal(t, exp, res)
}

func TestGetAllCouriers(t *testing.T) {
	exp := []Courier{
		Courier{
			Slug: "ups",
			Name: "ups",
		},
		Courier{
			Slug: "fedex",
			Name: "FeDex",
		},
	}
	mockhttp("GET", "https://api.aftership.com/v4/couriers/all", Envelope{
		response.Meta{200, "", ""},
		List{
			Total:    1,
			Couriers: exp,
		},
	}, map[string]string{
		"X-RateLimit-Reset":     "1458463600",
		"X-RateLimit-Limit":     "",
		"X-RateLimit-Remaining": "",
	})

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetAllCouriers()
	assert.Equal(t, exp, res)
}

func TestDetectCouriers(t *testing.T) {
	exp := []Courier{
		Courier{
			Slug: "ups",
			Name: "ups",
		},
	}
	mockhttp("GET", "https://api.aftership.com/v4/couriers/all", Envelope{
		response.Meta{200, "", ""},
		List{
			Total:    1,
			Couriers: exp,
		},
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	fmt.Println(endpoint.DetectCouriers(DetectParam{
		"906587618687",
		"DA15BU",
		"20131231",
		"1234567890",
		"",
		"",
		[]string{"dhl", "ups", "fedex"},
	}))
	assert.Equal(t, "", "")
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
