package apiV4

// TODO provide methods on each type to convert string times to date
// TODO provide constructors with handling for special values like dates and mandatory fields

// NewTracking provides parameters for new Tracking API request
type NewTracking struct {
	TrackingNumber             string            `json:"tracking_number"`
	Slug                       []string          `json:"slug,omitempty"`
	TrackingPostalCode         string            `json:"tracking_postal_code,omitempty"`
	TrackingShipDate           string            `json:"tracking_ship_date,omitempty"`
	TrackingAccountNumber      string            `json:"tracking_account_number,omitempty"`
	TrackingKey                string            `json:"tracking_key,omitempty"`
	TrackingDestinationCountry string            `json:"tracking_destination_country,omitempty"`
	Android                    []string          `json:"android,omitempty"`
	Ios                        []string          `json:"ios,omitempty"`
	Emails                     []string          `json:"emails,omitempty"`
	Smses                      []string          `json:"smses,omitempty"`
	Title                      string            `json:"title,omitempty"`
	CustomerName               string            `json:"customer_name,omitempty"`
	DestinationCountryIso3     string            `json:"destination_country_iso3,omitempty"`
	OrderId                    string            `json:"order_id,omitempty"`
	OrderIdPath                string            `json:"order_id_path,omitempty"`
	CustomFields               map[string]string `json:"custom_fields,omitempty"`
}

// CheckPoint represents a CheckPoint returned by the Aftership API
type CheckPoint struct {
	Slug           string   `json:"slug,omitempty"`
	CreatedAt      string   `json:"created_at,omitempty"`
	CheckPointTime string   `json:"checkpoint_time,omitempty"`
	City           string   `json:"city,omitempty"`
	Coordinates    []string `json:"coordinates,omitempty"`
	CountryIso3    string   `json:"country_iso3,omitempty"`
	CountryName    string   `json:"country_name,omitempty"`
	Message        string   `json:"message,omitempty"`
	State          string   `json:"state,omitempty"`
	Tag            string   `json:"tag,omitempty"`
	Zip            string   `json:"zip,omitempty"`
}

// Tracking represents a Tracking returned by the Aftership API
type Tracking struct {
	NewTracking
	// TODO write a function on this type to convert slug field with a switch
	Slug                 interface{}  `json:"slug"`
	Id                   string       `json:"id"`
	CreatedAt            string       `json:"created_at"`
	UpdatedAt            string       `json:"updated_at"`
	Active               bool         `json:"active"`
	ExpectedDelivery     string       `json:"expected_delivery"`
	Note                 string       `json:"note"`
	OriginCountryIso3    string       `json:"origin_country_iso3"`
	ShipmentPackageCount int          `json:"shipment_package_count"`
	ShipmentType         string       `json:"shipment_type"`
	SignedBy             string       `json:"signed_by"`
	Source               string       `json:"source"`
	Tag                  string       `json:"tag"`
	TrackCount           int          `json:"tracked_count"`
	UniqueToken          string       `json:"unique_token"`
	CheckPoints          []CheckPoint `json:"checkpoints"`
}

// TrackingId identifies a Tracking to be deleted
// its mandatory to provide either Id or Slug and TrackingNumber both
type TrackingId struct {
	Id             string
	Slug           string
	TrackingNumber string
}

// TrackingUpdate represents an update to Tracking details
type TrackingUpdate struct {
	Emails                 []string          `json:"emails,omitempty"`
	Smses                  []string          `json:"smses,omitempty"`
	Title                  string            `json:"title,omitempty"`
	CustomerName           string            `json:"customer_name,omitempty"`
	DestinationCountryIso3 string            `json:"destination_country_iso3,omitempty"`
	OrderId                string            `json:"order_id,omitempty"`
	OrderIdPath            string            `json:"order_id_path,omitempty"`
	CustomFields           map[string]string `json:"custom_fields,omitempty"`
}

// DeletedTracking is a deleted tracking object returned by API
type DeletedTracking struct {
	Id                    string `json:"id,omitempty"`
	Slug                  string `json:"slug,omitempty"`
	TrackingNumber        string `json:"tracking_number"`
	TrackingPostalCode    string `json:"tracking_postal_code,omitempty"`
	TrackingShipDate      string `json:"tracking_ship_date,omitempty"`
	TrackingAccountNumber string `json:"tracking_account_number,omitempty"`
	TrackingKey           string `json:"tracking_key,omitempty"`
}

// GetTrackingsParams represents the set of params for get Trackings API
type GetTrackingsParams struct {
	// Page to show. (Default: 1)
	Page int `url:"page,omitempty" json:"page,omitempty"`

	// Number of trackings each page contain. (Default: 100, Max: 200)
	Limit int `url:"limit,omitempty" json:"limit,omitempty"`

	// Search the content of the tracking record fields:tracking_number,
	// title, order_id, customer_name, custom_fields, order_id, emails, smses
	Keyword string `url:"keyword,omitempty" json:"keyword,omitempty"`

	// Unique courier code Use comma for multiple values. (Example: dhl,ups,usps)
	Slug string `url:"slug,omitempty" json:"slug,omitempty"`

	// "Total delivery time in days.
	// - Difference of 1st checkpoint time and delivered time for delivered shipments
	// - Difference of 1st checkpoint time and current time for non-delivered shipments
	// Value as 0 for pending shipments or delivered shipment with only one checkpoint."
	DeliveryTime int `url:"delivery_time,omitempty" json:"delivery_time,omitempty"`

	// Origin country of trackings. Use ISO Alpha-3 (three letters). Use comma for multiple values. (Example: USA,HKG)
	Origin string `url:"origin,omitempty" json:"origin,omitempty"`

	// Destination country of trackings. Use ISO Alpha-3 (three letters). Use comma for multiple values. (Example: USA,HKG)
	Destination string `url:"destination,omitempty" json:"destination,omitempty"`

	// Current status of tracking. Values include Pending, InfoReceived, InTransit, OutForDelivery, AttemptFail,
	// Delivered, Exception, Expired(See status definition)
	Tag string `url:"tag,omitempty" json:"tag,omitempty"`

	// "Start date and time of trackings created. AfterShip only stores data of 90 days.
	// (Defaults: 30 days ago, Example: 2013-03-15T16:41:56+08:00)"
	CreatedAtMin string `url:"created_at_min,omitempty" json:"created_at_min,omitempty"`

	// "End date and time of trackings created.
	// (Defaults: now, Example: 2013-04-15T16:41:56+08:00)"
	CreatedAtMax string `url:"created_at_max,omitempty" json:"created_at_max,omitempty"`

	// "List of fields to include in the response. Use comma for multiple values.
	// Fields to include: title, order_id, tag, checkpoints, checkpoint_time, message, country_name
	// Defaults: none, Example: title,order_id"
	Fields string `url:"fields,omitempty" json:"fields,omitempty"`

	// "Default: '' / Example: 'en'
	// Support Chinese to English translation for china-ems and china-post only"
	Lang string `url:"lang,omitempty" json:"lang,omitempty"`

	// For GetTrackingsExport pass the cursor returned in the previous response for retrieve next page.
	// When the browsing reaches the end of the index, the returned cursor will be an empty string.
	Cursor string `url:"cursor,omitempty" json:"cursor,omitempty"`
}

type TrackingsData struct {
	GetTrackingsParams
	Origin      []string   `json:"origin"`
	Destination []string   `json:"destination"`
	Cursor      string     `json:"cursor"`
	Trackings   []Tracking `json:"trackings"`
}

// TrackingResponseData is a model for data part of the tracking API responses
type TrackingResponseData struct {
	Tracking Tracking `json:"tracking"`
}

// TrackingEnvelope is the message envelope for the tracking API responses
type TrackingEnvelope struct {
	Meta ResponseMeta         `json:"meta"`
	Data TrackingResponseData `json:"data"`
}

// ResponseCode provides implementation of Response.ResponseCode()
// for TrackingEnvelope struct
func (envelope *TrackingEnvelope) ResponseCode() ResponseMeta {
	return envelope.Meta
}

// DeleteTrackingResponseData is a model for data part of the tracking API responses
type DeleteTrackingResponseData struct {
	Tracking DeletedTracking `json:"tracking"`
}

// DeleteTrackingEnvelope is the message envelope for the tracking API responses
type DeleteTrackingEnvelope struct {
	Meta ResponseMeta               `json:"meta"`
	Data DeleteTrackingResponseData `json:"data"`
}

// TrackingsEnvelope is the message envelope for the trackings API responses
type TrackingsEnvelope struct {
	Meta ResponseMeta  `json:"meta"`
	Data TrackingsData `json:"data"`
}

// ResponseCode provides implementation of Response.ResponseCode()
// for TrackingEnvelope struct
func (envelope *TrackingsEnvelope) ResponseCode() ResponseMeta {
	return envelope.Meta
}

// ResponseCode provides implementation of Response.ResponseCode()
// for TrackingEnvelope struct
func (envelope *DeleteTrackingEnvelope) ResponseCode() ResponseMeta {
	return envelope.Meta
}

type NewTrackingReqBody struct {
	Tracking NewTracking `json:"tracking"`
}

type TrackingUpdateReqBody struct {
	Tracking TrackingUpdate `json:"tracking"`
}

// LastCheckPoint is the last checkpoint API response
type LastCheckPoint struct {
	Id             string     `json:"id,omitempty"`
	Slug           string     `json:"slug,omitempty"`
	TrackingNumber string     `json:"tracking_number"`
	Tag            string     `json:"tag"`
	CheckPoint     CheckPoint `json:"checkpoint"`
}

// LastCheckPointEnvelope is the message envelope for the last checkpoint API responses
type LastCheckPointEnvelope struct {
	Meta ResponseMeta   `json:"meta"`
	Data LastCheckPoint `json:"data"`
}

// ResponseCode provides implementation of Response.ResponseCode()
// for LastCheckPointEnvelope struct
func (envelope *LastCheckPointEnvelope) ResponseCode() ResponseMeta {
	return envelope.Meta
}

// TrackingsHandler provides the interface for all trackings handling API calls
// in AfterShip APIV4
type TrackingsHandler interface {
	// CreateTracking Creates a tracking.
	CreateTracking(newTracking NewTracking) (Tracking, AfterShipApiError)

	// DeleteTracking Deletes a tracking.
	DeleteTracking(id TrackingId) (DeletedTracking, AfterShipApiError)

	// GetTrackings Gets tracking results of multiple trackings.
	GetTrackings(params GetTrackingsParams) (TrackingsData, AfterShipApiError)

	// GetTrackingsExport Gets all trackings results (for backup or analytics purpose)
	GetTrackingsExport(params GetTrackingsParams) (TrackingsData, AfterShipApiError)

	// GetTracking Gets tracking results of a single tracking.
	// fields : List of fields to include in the response. Use comma for multiple values.
	// Fields to include: tracking_postal_code,tracking_ship_date,tracking_account_number,
	// tracking_key,tracking_destination_country, title,order_id,tag,checkpoints,
	// checkpoint_time, message, country_name
	GetTracking(id TrackingId, fields string, lang string) (Tracking, AfterShipApiError)

	// UpdateTracking Updates a tracking.
	UpdateTracking(id TrackingId, update TrackingUpdate) (Tracking, AfterShipApiError)

	// ReTrack an expired tracking once. Max. 3 times per tracking.
	ReTrack(id TrackingId) (Tracking, AfterShipApiError)

	// LastCheckPoint Return the tracking information of the last checkpoint of a single tracking.
	GetLastCheckPoint(id TrackingId, fields string, lang string) (LastCheckPoint, AfterShipApiError)
}
