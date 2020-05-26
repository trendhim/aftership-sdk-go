package aftership

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetLastCheckpoint(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}
	exp := LastCheckpoint{
		ID:             "5b74f4958776db0e00b6f5ed",
		TrackingNumber: "111111111111",
		Checkpoint: Checkpoint{
			Slug: "slug",
		},
	}
	mockHTTP("GET", fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), 200, Response{
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
	endpoint := newCheckpointsEndpoint(req)
	res, err := endpoint.GetLastCheckpoint(context.Background(), p, GetCheckpointParams{})
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetLastCheckpointWithOptionalParams(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	op := GetCheckpointParams{
		Fields: "slug",
		Lang:   "en",
	}

	exp := LastCheckpoint{
		ID:             "5b74f4958776db0e00b6f5ed",
		TrackingNumber: "111111111111",
		Checkpoint: Checkpoint{
			Slug: "slug",
		},
	}
	mockHTTP("GET", fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), 200, Response{
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
	endpoint := newCheckpointsEndpoint(req)
	res, err := endpoint.GetLastCheckpoint(context.Background(), p, op)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestError(t *testing.T) {
	req := newRequestHelper(Config{})
	endpoint := newCheckpointsEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.GetLastCheckpoint(context.Background(), p, GetCheckpointParams{})
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("GET", fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: LastCheckpoint{},
	}, nil)

	_, err = endpoint.GetLastCheckpoint(context.Background(), p, GetCheckpointParams{})
	assert.NotNil(t, err)
	assert.Equal(t, &APIError{
		Code:    401,
		Type:    "Unauthorized",
		Message: "Invalid API key.",
		Path:    "/last_checkpoint/xq-express/LS404494276CN",
	}, err)
}
