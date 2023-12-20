package aftership

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLastCheckpoint(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	uri := fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"id": "5b74f4958776db0e00b6f5ed",
					"tracking_number": "111111111111",
					"slug": "fedex",
					"tag": "Delivered",
					"subtag": "Delivered_001",
					"subtag_message": "Delivered",
					"checkpoint": {
							"slug": "fedex",
							"created_at": "2018-08-16T03:50:47+00:00",
							"checkpoint_time": "2018-08-01T13:19:47-04:00",
							"city": "Deal",
							"coordinates": [],
							"country_iso3": null,
							"country_name": null,
							"message": "Delivered - Left at front door. Signature Service not requested.",
							"state": "NJ",
							"tag": "Delivered",
							"subtag": "Delivered_001",
							"subtag_message": "Delivered",
							"zip": null,
							"raw_tag": "FPX_L_RPIF"
					}
			}
	}`))
	})

	createdAt, _ := time.Parse(time.RFC3339, "2018-08-16T03:50:47+00:00")
	exp := LastCheckpoint{
		ID:             "5b74f4958776db0e00b6f5ed",
		TrackingNumber: "111111111111",
		Slug:           "fedex",
		Tag:            "Delivered",
		Subtag:         "Delivered_001",
		SubtagMessage:  "Delivered",
		Checkpoint: Checkpoint{
			Slug:           "fedex",
			CreatedAt:      &createdAt,
			CheckpointTime: "2018-08-01T13:19:47-04:00",
			City:           "Deal",
			Coordinates:    []float32{},
			Message:        "Delivered - Left at front door. Signature Service not requested.",
			State:          "NJ",
			Tag:            "Delivered",
			Subtag:         "Delivered_001",
			SubtagMessage:  "Delivered",
			RawTag:         "FPX_L_RPIF",
		},
	}

	res, err := client.GetLastCheckpoint(context.Background(), p, GetCheckpointParams{})
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestGetLastCheckpointWithOptionalParams(t *testing.T) {
	setup()
	defer teardown()

	p := SlugTrackingNumber{
		Slug:           "xq-express",
		TrackingNumber: "LS404494276CN",
	}

	op := GetCheckpointParams{
		Fields: "slug",
		Lang:   "en",
	}

	uri := fmt.Sprintf("/last_checkpoint/%s/%s", p.Slug, p.TrackingNumber)
	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{
			"meta": {
					"code": 200
			},
			"data": {
					"id": "5b74f4958776db0e00b6f5ed",
					"tracking_number": "111111111111",
					"slug": "fedex",
					"tag": "Delivered",
					"subtag": "Delivered_001",
					"subtag_message": "Delivered",
					"checkpoint": {
							"slug": "fedex"
					}
			}
	}`))
	})

	exp := LastCheckpoint{
		ID:             "5b74f4958776db0e00b6f5ed",
		TrackingNumber: "111111111111",
		Slug:           "fedex",
		Tag:            "Delivered",
		Subtag:         "Delivered_001",
		SubtagMessage:  "Delivered",
		Checkpoint: Checkpoint{
			Slug: "fedex",
		},
	}

	res, err := client.GetLastCheckpoint(context.Background(), p, op)
	assert.Equal(t, exp, res)
	assert.Nil(t, err)
}

func TestError(t *testing.T) {
	setup()
	defer teardown()

	// empty id, slug and tracking_number
	p := SlugTrackingNumber{
		Slug:           "",
		TrackingNumber: "",
	}

	_, err := client.GetLastCheckpoint(context.Background(), p, GetCheckpointParams{})
	assert.NotNil(t, err)
}
