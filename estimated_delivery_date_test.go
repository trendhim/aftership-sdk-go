package aftership

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_BatchPredictEstimatedDeliveryDate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/estimated-delivery-date/predict-batch", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		w.Write([]byte(`{
          "meta": {
            "code": 200
          },
          "data": {
            "estimated_delivery_dates": [
              {
                "slug": "fedex",
                "service_type_name": "FEDEX HOME DELIVERY",
                "origin_address": {
                  "country": "USA",
                  "state": "WA",
                  "postal_code": "98108",
                  "raw_location": "Seattle, Washington, 98108, USA, United States",
                  "city": null
                },
                "destination_address": {
                  "country": "USA",
                  "state": "CA",
                  "postal_code": "92019",
                  "raw_location": "El Cajon, California, 92019, USA, United States",
                  "city": null
                },
                "weight": {
                  "unit": "kg",
                  "value": 11
                },
                "package_count": 1,
                "pickup_time": "2021-07-01 15:00:00",
                "estimated_pickup": null,
                "estimated_delivery_date": "2021-07-05",
                "estimated_delivery_date_min": "2021-07-04",
                "estimated_delivery_date_max": "2021-07-05"
              }
            ]
          }
        }`))
	})

	dates := []EstimatedDeliveryDate{
		{
			Slug:            "fedex",
			ServiceTypeName: "FEDEX HOME DELIVERY",
			OriginAddress: &Address{
				Country:     "USA",
				State:       "WA",
				PostalCode:  "98108",
				RawLocation: "Seattle, Washington, 98108, USA, United States",
			},
			DestinationAddress: &Address{
				Country:     "USA",
				State:       "CA",
				PostalCode:  "92019",
				RawLocation: "El Cajon, California, 92019, USA, United States",
			},
			Weight: &Weight{
				Unit:  "kg",
				Value: 11,
			},
			PackageCount: 1,
			PickupTime:   "2021-07-01 15:00:00",
		},
	}

	marshal, _ := json.Marshal(dates)
	fmt.Println(string(marshal))

	res, err := client.BatchPredictEstimatedDeliveryDate(context.Background(), dates)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(res.Dates))
}
