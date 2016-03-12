package apiV4

// Courier is the model describing an AfterShip courier
type Courier struct {
	Slug                   string   `json:"slug"`
	Name                   string   `json:"name"`
	Phone                  string   `json:"phone"`
	OtherName              string   `json:"other_name"`
	WebUrl                 string   `json:"web_url"`
	RequiredFields         []string `json:"required_fields"`
	DefaultLanguage        string   `json:"default_language"`
	SupportedLanguages     []string `json:"supported_languages"`
	ServiceFromCountryISO3 []string `json:"service_from_country_iso3"`
}

// CourierResponseData is a model for data part of the courier API responses
type CourierResponseData struct {
	Couriers []Courier `json:"couriers"`
}

// CourierEnvelope is the message envelope for the courier API responses
type CourierEnvelope struct {
	Meta ResponseMeta        `json:"meta"`
	Data CourierResponseData `json:"data"`
}

// ResponseCode provides implementation of Response.ResponseCode()
// for CourierEnvelope struct
func (envelope *CourierEnvelope) ResponseCode() ResponseMeta {
	return envelope.Meta
}

// CourierDetectParam contains fields required and optional fields for
// courier detection
type CourierDetectParam struct {

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

	// Slug If not specified, Aftership will automatically detect the courier based on the tracking number format and
	// your selected couriers.
	// Use comma separated to input a list of couriers for auto detect.
	Slug []string `json:"slug,omitempty"`
}

type CourierDetectParamReqBody struct {
	Tracking CourierDetectParam `json:"tracking"`
}

// CourierHandler provides the interface for all courier handling API calls
// in AfterShip APIV4
type CourierHandler interface {

	// GetCouriers returns a list of couriers activated at your AfterShip account.
	GetCouriers() ([]Courier, AfterShipApiError)

	// GetAllCouriers returns a list of all couriers.
	GetAllCouriers() ([]Courier, AfterShipApiError)

	// DetectCouriers returns a list of matched couriers based on tracking number format
	// and selected couriers or a list of couriers.
	DetectCouriers(params CourierDetectParam) ([]Courier, AfterShipApiError)
}
