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

func TestBuildUrl(t *testing.T) {
	// slug and tracking number
	p := SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
	}

	url, err := buildLastCheckpointURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/last_checkpoint/xq-express/LS404494276CN", url)

	// slug and tracking number, has optional parameters
	url, err = buildLastCheckpointURL(p, "slug", "en")
	assert.Nil(t, err)
	assert.Equal(t, "/last_checkpoint/xq-express/LS404494276CN?fields=slug&lang=en", url)

	// id
	p = SingleTrackingParam{
		ID: "1234567890",
	}

	url, err = buildLastCheckpointURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/last_checkpoint/1234567890", url)

	// id, has optional parameters
	url, err = buildLastCheckpointURL(p, "slug", "en")
	assert.Nil(t, err)
	assert.Equal(t, "/last_checkpoint/1234567890?fields=slug&lang=en", url)

	// should get error when no id, slug and tracking number
	p = SingleTrackingParam{
		"",
		"",
		"",
	}
	_, err = buildLastCheckpointURL(p, "", "")
	assert.NotNil(t, err)

	// Encode slug and tracking number
	p = SingleTrackingParam{
		"",
		"usps",
		"ABCD/1234",
	}

	url, err = buildLastCheckpointURL(p, "", "")
	assert.Nil(t, err)
	assert.Equal(t, "/last_checkpoint/usps/ABCD%2F1234", url)
}

func TestGetLastCheckpoint(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

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
	mockhttp("GET", fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), 200, LastCheckpointEnvelope{
		response.Meta{200, "", ""},
		exp,
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{
		APIKey: "YOUR_API_KEY",
	}, nil)
	endpoint := NewEnpoint(req)
	res, err := endpoint.GetLastCheckpoint(p, "", "")
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SingleTrackingParam{
		"",
		"xq-express",
		"LS404494276CN",
	}

	mockhttp("GET", fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), 401, LastCheckpointEnvelope{
		response.Meta{401, "Invalid API key.", "Unauthorized"},
		LastCheckpoint{},
	}, nil)

	req := request.NewRequest(&conf.AfterShipConf{}, nil)
	endpoint := NewEnpoint(req)
	_, err := endpoint.GetLastCheckpoint(p, "", "")
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
