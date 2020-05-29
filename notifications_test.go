package aftership

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNotification(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	uri := fmt.Sprintf("/notifications/%s/%s", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"notification": {
							"emails": ["user1@gmail.com","user2@gmail.com"],
							"smses": ["+85291239123", "+85261236123"]
					}
			}
	}`))
	})

	exp := Notification{
		[]string{"user1@gmail.com", "user2@gmail.com"},
		[]string{"+85291239123", "+85261236123"},
	}

	res, err := client.GetNotification(context.Background(), p)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetNotificationError(t *testing.T) {
	// empty slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := client.GetNotification(context.Background(), p)
	assert.NotNil(t, err)
}

func TestAddNotification(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	uri := fmt.Sprintf("/notifications/%s/%s/add", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"notification": {
							"emails": ["user1@gmail.com","user2@gmail.com"],
							"smses": ["+85291239123", "+85261236123"]
					}
			}
	}`))
	})

	req := Notification{
		[]string{"user1@gmail.com", "user2@gmail.com", "invalid EMail @ Gmail. com"},
		[]string{"+85291239123", "+85261236123", "Invalid Mobile Phone Number"},
	}

	exp := Notification{
		[]string{"user1@gmail.com", "user2@gmail.com"},
		[]string{"+85291239123", "+85261236123"},
	}

	res, err := client.AddNotification(context.Background(), p, req)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestAddNotificationError(t *testing.T) {
	// empty slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := client.AddNotification(context.Background(), p, Notification{})
	assert.NotNil(t, err)
}

func TestRemoveNotification(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	uri := fmt.Sprintf("/notifications/%s/%s/remove", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"notification": {
							"emails": [],
							"smses": ["+85261236888"]
					}
			}
	}`))
	})

	req := Notification{
		[]string{"user1@gmail.com", "user2@gmail.com", "invalid EMail @ Gmail. com"},
		[]string{"+85291239123", "Invalid Mobile Phone Number"},
	}

	exp := Notification{
		[]string{},
		[]string{"+85261236888"},
	}

	res, err := client.RemoveNotification(context.Background(), p, req)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestRemoveNotificationError(t *testing.T) {
	// empty slug or tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := client.RemoveNotification(context.Background(), p, Notification{})
	assert.NotNil(t, err)
}
