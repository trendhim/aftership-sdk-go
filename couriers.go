package aftership

import (
	"context"
	"errors"
)

// Courier is the model describing an AfterShip courier
type Courier struct {
	Slug                   string   `json:"slug"`                      // Unique code of courier
	Name                   string   `json:"name"`                      // Name of courier
	Phone                  string   `json:"phone"`                     // Contact phone number of courier
	OtherName              string   `json:"other_name"`                // Other name of courier
	WebURL                 string   `json:"web_url"`                   // Website link of courier
	RequiredFields         []string `json:"required_fields"`           // The extra fields need for tracking, such as `tracking_account_number`, `tracking_postal_code`, `tracking_ship_date`, `tracking_key`, `tracking_destination_country`
	OptionalFields         []string `json:"optional_fields"`           // the extra fields which are optional for tracking. Basically it's the same as required_fields, but the difference is that only some of the tracking numbers require these fields.
	DefaultLanguage        string   `json:"default_language"`          // Default language of tracking results
	SupportedLanguages     []string `json:"supported_languages"`       // Other supported languages
	ServiceFromCountryISO3 []string `json:"service_from_country_iso3"` // Country code (ISO Alpha-3) where the courier provides service
}

// CourierList is the model describing an AfterShip courier list
type CourierList struct {
	Total    int       `json:"total"`    // Total number of couriers supported by AfterShip.
	Couriers []Courier `json:"couriers"` // Array of Courier describes the couriers information.
}

// TrackingCouriers is the model describing an AfterShip couriers detect list
type TrackingCouriers struct {
	Total    int       `json:"total"`    // Total number of matched couriers
	Tracking Tracking  `json:"tracking"` // Tracking describes the tracking information.
	Couriers []Courier `json:"couriers"` // A list of matched couriers based on tracking number format.
}

// CourierDetectionParams contains fields required and optional fields for courier detection
type CourierDetectionParams struct {

	// TrackingNumber of a shipment. Mandatory field.
	TrackingNumber string `json:"tracking_number"`

	// TrackingPostalCode is the postal code of receiver's address.
	// Required by some couriers, such as postnl
	TrackingPostalCode string `json:"tracking_postal_code,omitempty"`

	// TrackingShipDate in YYYYMMDD format. Required by some couriers, such as deutsch-post
	TrackingShipDate string `json:"tracking_ship_date,omitempty"`

	// TrackingAccountNumber of the shipper for a specific courier. Required by some couriers, such as dynamic-logistics
	TrackingAccountNumber string `json:"tracking_account_number,omitempty"`

	// TrackingKey of the shipment for a specific courier. Required by some couriers, such as sic-teliway
	TrackingKey string `json:"tracking_key,omitempty"`

	// TrackingDestinationCountry of the shipment for a specific courier. Required by some couriers, such as postnl-3s
	TrackingDestinationCountry string `json:"tracking_destination_country,omitempty"`

	// Slug If not specified, AfterShip will automatically detect the courier based on the tracking number format and
	// your selected couriers.
	// Use array to input a list of couriers for auto detect.
	Slug []string `json:"slug,omitempty"`
}

var ErrorTrackingNumberRequired = errors.New("tracking number is required")

// CouriersEndpoint provides the interface for all courier API calls
type CouriersEndpoint interface {

	// GetCouriers returns a list of couriers activated at your AfterShip account.
	GetCouriers(ctx context.Context) (CourierList, error)

	// GetAllCouriers returns a list of all couriers.
	GetAllCouriers(ctx context.Context) (CourierList, error)

	// DetectCouriers returns a list of matched couriers based on tracking number format
	// and selected couriers or a list of couriers.
	DetectCouriers(ctx context.Context, params CourierDetectionParams) (TrackingCouriers, error)
}

// couriersEndpointImpl is the implementation of courier endpoint
type couriersEndpointImpl struct {
	request requestHelper
}

// newCouriersEndpoint creates a instance of courier endpoint
func newCouriersEndpoint(req requestHelper) CouriersEndpoint {
	return &couriersEndpointImpl{
		request: req,
	}
}

// GetCouriers returns a list of couriers activated at your AfterShip account.
func (impl *couriersEndpointImpl) GetCouriers(ctx context.Context) (CourierList, error) {
	var courierList CourierList
	err := impl.request.makeRequest(ctx, "GET", "/couriers", nil, nil, &courierList)
	return courierList, err
}

// GetAllCouriers returns a list of all couriers.
func (impl *couriersEndpointImpl) GetAllCouriers(ctx context.Context) (CourierList, error) {
	var courierList CourierList
	err := impl.request.makeRequest(ctx, "GET", "/couriers/all", nil, nil, &courierList)
	return courierList, err
}

// detectCourierRequest is a model for detect courier API request
type detectCourierRequest struct {
	Tracking CourierDetectionParams `json:"tracking"`
}

// DetectCouriers returns a list of matched couriers based on tracking number format
// and selected couriers or a list of couriers.
func (impl *couriersEndpointImpl) DetectCouriers(ctx context.Context, params CourierDetectionParams) (TrackingCouriers, error) {
	if params.TrackingNumber == "" {
		return TrackingCouriers{}, ErrorTrackingNumberRequired
	}

	var trackingCouriers TrackingCouriers
	err := impl.request.makeRequest(ctx, "POST", "/couriers/detect", nil,
		&detectCourierRequest{
			Tracking: params,
		}, &trackingCouriers)
	return trackingCouriers, err
}
