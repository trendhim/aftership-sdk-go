package checkpoint

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

func TestGetLastCheckpoint(t *testing.T) {
	p := SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
	}
	exp := LastCheckpoint{
		ID:             "5b74f4958776db0e00b6f5ed",
		TrackingNumber: "111111111111",
		Checkpoint: Checkpoint{
			Slug: "slug",
		},
	}
	mockhttp("GET", fmt.Sprintf("https://api.aftership.com/v4/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), LastCheckpointEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(conf.AfterShipConf{
		AppKey: "YOUR_API_KEY",
	})
	endpoint := NewEnpoint(req)
	res, _ := endpoint.GetLastCheckpoint(p, "", "")
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
