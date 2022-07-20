package aftership

import (
	"context"
	"net/http"
)

type Address struct {
	// The country/region of the origin location from where the package is
	// picked up by the carrier to be delivered to the final destination.
	// Use 3 letters of ISO 3166-1 country/region code.
	Country string `json:"country,omitempty"`

	// State, province, or the equivalent location of the origin address.
	// Either `origin_address.state` or `origin_address.postal_code` is required.
	State string `json:"state,omitempty"`

	// City of the origin address.
	PostalCode string `json:"postal_code,omitempty"`

	// Postal code of the origin address.
	// Either `origin_address.state` or `origin_address.postal_code` is required.
	RawLocation string `json:"raw_location,omitempty"`

	// Raw location of the origin address.
	City string `json:"city,omitempty"`
}

type EstimatedPickup struct {
	// The local order time of the package.
	OrderTime string `json:"order_time,omitempty"`

	// Order cut off time. AfterShip will set 18:00:00 as the default value.
	OrderCutoffTime string `json:"order_cutoff_time,omitempty"`

	// Operating days in a week. Number refers to the weekday.
	// E.g., [1,2,3,4,5] means operating days are from Monday to Friday.
	// AfterShip will set [1,2,3,4,5] as the default value.
	BusinessDays []int64 `json:"business_days,omitempty"`

	OrderProcessingTime *OrderProcessingTime `json:"order_processing_time,omitempty"`

	// The local pickup time of the package.
	PickupTime string `json:"pickup_time,omitempty"`
}

type OrderProcessingTime struct {
	// Processing time of an order, from being placed to being picked up.
	// Only support day as value now. AfterShip will set day as the default value.
	Unit string `json:"unit,omitempty"`

	// Processing time of an order, from being placed to being picked up.
	// AfterShip will set 0 as the default value.
	Value int64 `json:"value"`
}

type Weight struct {
	// The weight unit of the package.
	Unit string `json:"unit,omitempty"`

	// The weight of the shipment.
	Value int64 `json:"value,omitempty"`
}

// batchPredictEstimatedDeliveryDateRequest is a model for batch predict courier API request
type batchPredictEstimatedDeliveryDateRequest struct {
	EstimatedDeliveryDates []EstimatedDeliveryDate `json:"estimated_delivery_dates,omitempty"`
}

// EstimatedDeliveryDates is a response object of batch predict
type EstimatedDeliveryDates struct {
	Dates []EstimatedDeliveryDate `json:"estimated_delivery_dates,omitempty"`
}

// BatchPredictEstimatedDeliveryDate Batch predict the estimated delivery dates
func (client *Client) BatchPredictEstimatedDeliveryDate(ctx context.Context, params []EstimatedDeliveryDate) (EstimatedDeliveryDates, error) {
	var dates EstimatedDeliveryDates
	err := client.makeRequest(ctx, http.MethodPost, "/estimated-delivery-date/predict-batch", nil,
		&batchPredictEstimatedDeliveryDateRequest{
			EstimatedDeliveryDates: params,
		}, &dates)
	return dates, err
}
