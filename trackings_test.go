package aftership

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	params := CreateTrackingParams{
		Slug:           []string{"dhl"},
		TrackingNumber: "tracking_number",
		Title:          "Title Name",
		SMSes: []string{
			"+18555072509",
			"+18555072501",
		},
	}

	tracking := Tracking{
		Slug:           "dhl",
		TrackingNumber: "tracking_number",
		Title:          "Title Name",
		SMSes: []string{
			"+18555072509",
			"+18555072501",
		},
	}

	mockHTTP("POST", "/trackings", 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: trackingWrapper{
			Tracking: tracking,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)
	res, _ := endpoint.CreateTracking(context.Background(), params)
	assert.Equal(t, tracking, res)
}

func TestCreateTrackingError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	data := CreateTrackingParams{
		TrackingNumber: "tracking_number",
	}

	mockHTTP("POST", "/trackings", 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: trackingWrapper{},
	}, nil)

	_, err := endpoint.CreateTracking(context.Background(), data)
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}

func TestDeleteTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	exp := Tracking{
		Slug:  "ups",
		Title: "Title Name",
	}

	mockHTTP("DELETE", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: trackingWrapper{
			Tracking: exp,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)
	res, _ := endpoint.DeleteTracking(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestDeleteTrackingError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.DeleteTracking(context.Background(), p)
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("DELETE", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: trackingWrapper{},
	}, nil)

	_, err = endpoint.DeleteTracking(context.Background(), p)
	assert.NotNil(t, err)
	//assert.Equal(t, "Unauthorized", err.Type)
}

func TestGetTrackings(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := GetTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	exp := PagedTrackings{
		Page:  1,
		Count: 2,
		Limit: 10,
		Trackings: []Tracking{
			{
				Slug:  "dhl",
				Title: "Title Name",
			},
			{
				Slug:  "ups",
				Title: "Title Name",
			},
		},
	}

	mockHTTP("GET", fmt.Sprintf("/trackings?limit=%d&page=%d", p.Limit, p.Page), 200, Response{
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
	endpoint := newTrackingsEndpoint(req)
	res, _ := endpoint.GetTrackings(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestGetTrackingsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := GetTrackingsParams{
		Page:  1,
		Limit: 10,
	}

	mockHTTP("GET", fmt.Sprintf("/trackings?limit=%d&page=%d", p.Limit, p.Page), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: PagedTrackings{},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)
	_, err := endpoint.GetTrackings(context.Background(), p)
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}

func TestGetTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	exp := Tracking{
		Slug:  "ups",
		Title: "Title Name",
	}

	mockHTTP("GET", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: trackingWrapper{
			Tracking: exp,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)
	res, _ := endpoint.GetTracking(context.Background(), p, GetTrackingParams{})
	assert.Equal(t, exp, res)
}

func TestGetTrackingError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.GetTracking(context.Background(), p, GetTrackingParams{})
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("GET", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: trackingWrapper{},
	}, nil)

	_, err = endpoint.GetTracking(context.Background(), p, GetTrackingParams{})
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}

func TestUpdateTracking(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	data := UpdateTrackingParams{
		Title: "New Title",
	}

	exp := Tracking{
		Slug:  "ups",
		Title: "New Title",
	}

	mockHTTP("PUT", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: trackingWrapper{
			Tracking: exp,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)
	res, _ := endpoint.UpdateTracking(context.Background(), p, data)
	assert.Equal(t, exp, res)
}

func TestUpdateTrackingError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	data := UpdateTrackingParams{
		Title: "New Title",
	}

	_, err := endpoint.UpdateTracking(context.Background(), p, data)
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("PUT", fmt.Sprintf("/trackings/%s/%s", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: trackingWrapper{},
	}, nil)

	_, err = endpoint.UpdateTracking(context.Background(), p, data)
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}

func TestReTrack(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	exp := Tracking{
		Slug:  "ups",
		Title: "Title Name",
	}

	mockHTTP("POST", fmt.Sprintf("/trackings/%s/%s/retrack", p.Slug, p.TrackingNumber), 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: trackingWrapper{
			Tracking: exp,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)
	res, _ := endpoint.ReTrack(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestReTrackError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.ReTrack(context.Background(), p)
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("POST", fmt.Sprintf("/trackings/%s/%s/retrack", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: trackingWrapper{},
	}, nil)

	_, err = endpoint.ReTrack(context.Background(), p)
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}

func TestMarkAsCompleted(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "ups",
		TrackingNumber: "1Z9999999999999998",
	}

	exp := Tracking{
		Slug:  "ups",
		Title: "Title Name",
	}

	mockHTTP("POST", fmt.Sprintf("/trackings/%s/%s/mark-as-completed", p.Slug, p.TrackingNumber), 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: trackingWrapper{
			Tracking: exp,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)
	res, _ := endpoint.MarkAsCompleted(context.Background(), p, CompletedStatusLost)
	assert.Equal(t, exp, res)
}

func TestMarkAsCompletedError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newTrackingsEndpoint(req)

	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.MarkAsCompleted(context.Background(), p, CompletedStatusLost)
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	mockHTTP("POST", fmt.Sprintf("/trackings/%s/%s/mark-as-completed", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: trackingWrapper{},
	}, nil)

	_, err = endpoint.MarkAsCompleted(context.Background(), p, CompletedStatusDelivered)
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}
