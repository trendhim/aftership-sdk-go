package aftership

import (
	"context"
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestAddNotification(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	exp := notificationWrapper{
		Notification{
			[]string{"vimukthi@aftership.net"},
			[]string{"+85254469627"},
		},
	}

	mockHTTP("POST", fmt.Sprintf("/notifications/%s/%s/add", p.Slug, p.TrackingNumber), 200, Response{
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
	endpoint := newNotificationEndpoint(req)
	res, _ := endpoint.AddNotification(context.Background(), p, exp.Notification)
	assert.Equal(t, exp.Notification, res)
}

func TestAddNotificationError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newNotificationEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.AddNotification(context.Background(), p, Notification{})
	assert.NotNil(t, err)
	//assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("POST", fmt.Sprintf("/notifications/%s/%s/add", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: notificationWrapper{},
	}, nil)

	_, err = endpoint.AddNotification(context.Background(), p, Notification{})
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}

func TestRemoveNotification(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	exp := Notification{
		[]string{"vimukthi@aftership.net"},
		[]string{"+85254469627"},
	}

	mockHTTP("POST", fmt.Sprintf("/notifications/%s/%s/remove", p.Slug, p.TrackingNumber), 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: notificationWrapper{
			Notification: exp,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newNotificationEndpoint(req)
	res, _ := endpoint.RemoveNotification(context.Background(), p, exp)
	assert.Equal(t, exp, res)
}

func TestRemoveNotificationError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newNotificationEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.RemoveNotification(context.Background(), p, Notification{})
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("POST", fmt.Sprintf("/notifications/%s/%s/remove", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: notificationWrapper{},
	}, nil)

	_, err = endpoint.RemoveNotification(context.Background(), p, Notification{})
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}

func TestGetNotificationSetting(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	exp := Notification{
		[]string{"vimukthi@aftership.net"},
		[]string{"+85254469627"},
	}

	mockHTTP("GET", fmt.Sprintf("/notifications/%s/%s", p.Slug, p.TrackingNumber), 200, Response{
		Meta: Meta{
			Code:    200,
			Message: "",
			Type:    "",
		},
		Data: notificationWrapper{
			Notification: exp,
		},
	}, nil)

	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newNotificationEndpoint(req)
	res, _ := endpoint.GetNotification(context.Background(), p)
	assert.Equal(t, exp, res)
}

func TestGetNotificationError(t *testing.T) {
	req := newRequestHelper(Config{
		APIKey: "YOUR_API_KEY",
	})
	endpoint := newNotificationEndpoint(req)

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := endpoint.GetNotification(context.Background(), p)
	assert.NotNil(t, err)
	// assert.Equal(t, "HandlerError", err.Type)

	// Response with error
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	p = SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	mockHTTP("GET", fmt.Sprintf("/notifications/%s/%s", p.Slug, p.TrackingNumber), 401, Response{
		Meta: Meta{
			Code:    401,
			Message: "Invalid API key.",
			Type:    "Unauthorized",
		},
		Data: notificationWrapper{},
	}, nil)

	_, err = endpoint.GetNotification(context.Background(), p)
	assert.NotNil(t, err)
	// assert.Equal(t, "Unauthorized", err.Type)
}
