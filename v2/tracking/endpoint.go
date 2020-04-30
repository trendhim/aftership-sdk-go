package tracking

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/google/go-querystring/query"
)

// Endpoint provides the interface for all trackings API calls
type Endpoint interface {
	// CreateTracking Creates a tracking.
	CreateTracking(newTracking NewTrackingRequest) (SingleTrackingData, *error.AfterShipError)

	// DeleteTracking Deletes a tracking.
	DeleteTracking(param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError)

	// GetTrackings Gets tracking results of multiple trackings.
	GetTrackings(params MultiTrackingsParams) (MultiTrackingsData, *error.AfterShipError)

	// GetTracking Gets tracking results of a single tracking.
	// fields : List of fields to include in the http. Use comma for multiple values.
	// Fields to include: tracking_postal_code,tracking_ship_date,tracking_account_number,
	// tracking_key,tracking_destination_country, title,order_id,tag,checkpoints,
	// checkpoint_time, message, country_name
	GetTracking(param SingleTrackingParam, optionalParams GetTrackingParams) (SingleTrackingData, *error.AfterShipError)

	// UpdateTracking Updates a tracking.
	UpdateTracking(param SingleTrackingParam, update UpdateTrackingRequest) (SingleTrackingData, *error.AfterShipError)

	// ReTrack an expired tracking once. Max. 3 times per tracking.
	ReTrack(param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError)
}

// EndpointImpl is the implementaion of tracking endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEnpoint creates a instance of tracking endpoint
func NewEnpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// CreateTracking creates a new tracking
func (impl *EndpointImpl) CreateTracking(newTracking NewTrackingRequest) (SingleTrackingData, *error.AfterShipError) {
	var envelope SingleTrackingEnvelope
	err := impl.request.MakeRequest("POST", "/trackings", newTracking, &envelope)
	if err != nil {
		return SingleTrackingData{}, err
	}

	return envelope.Data, nil
}

// DeleteTracking Deletes a tracking.
func (impl *EndpointImpl) DeleteTracking(param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError) {
	url, err := BuildTrackingURL(param, "trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("DELETE", url, nil, &envelope)
	if err != nil {
		return SingleTrackingData{}, err
	}
	return envelope.Data, nil
}

// GetTrackings Gets tracking results of multiple trackings.
func (impl *EndpointImpl) GetTrackings(params MultiTrackingsParams) (MultiTrackingsData, *error.AfterShipError) {
	url, err := BuildURLWithQueryString("/trackings", params)
	if err != nil {
		return MultiTrackingsData{}, err
	}

	var envelope MultiTrackingsEnvelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	if err != nil {
		return MultiTrackingsData{}, err
	}
	return envelope.Data, nil
}

// GetTracking Gets tracking results of a single tracking.
func (impl *EndpointImpl) GetTracking(param SingleTrackingParam, optionalParams GetTrackingParams) (SingleTrackingData, *error.AfterShipError) {
	url, err := BuildTrackingURL(param, "trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	url, err = BuildURLWithQueryString(url, optionalParams)
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("GET", url, nil, &envelope)
	if err != nil {
		return SingleTrackingData{}, err
	}
	return envelope.Data, nil
}

// UpdateTracking Updates a tracking.
func (impl *EndpointImpl) UpdateTracking(param SingleTrackingParam, update UpdateTrackingRequest) (SingleTrackingData, *error.AfterShipError) {
	url, err := BuildTrackingURL(param, "trackings", "")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("PUT", url, update, &envelope)
	if err != nil {
		return SingleTrackingData{}, err
	}
	return envelope.Data, nil
}

// ReTrack an expired tracking once. Max. 3 times per tracking.
func (impl *EndpointImpl) ReTrack(param SingleTrackingParam) (SingleTrackingData, *error.AfterShipError) {
	url, err := BuildTrackingURL(param, "trackings", "retrack")
	if err != nil {
		return SingleTrackingData{}, err
	}

	var envelope SingleTrackingEnvelope
	err = impl.request.MakeRequest("POST", url, nil, &envelope)
	if err != nil {
		return SingleTrackingData{}, err
	}
	return envelope.Data, nil
}

// BuildTrackingURL returns the tracking URL
func BuildTrackingURL(param SingleTrackingParam, path string, subPath string) (string, *error.AfterShipError) {
	if path == "" {
		path = "trackings"
	}

	var trackingURL string
	if param.ID != "" {
		trackingURL = fmt.Sprintf("/%s/%s", path, url.QueryEscape(param.ID))
	} else if param.Slug != "" && param.TrackingNumber != "" {
		trackingURL = fmt.Sprintf("/%s/%s/%s", path, url.QueryEscape(param.Slug), url.QueryEscape(param.TrackingNumber))
	} else {
		return "", error.MakeSdkError(error.ErrorTypeHandlerError, "You must specify the id or slug and tracking number", param)
	}

	if subPath != "" {
		trackingURL += fmt.Sprintf("/%s", subPath)
	}

	if param.OptionalParams != nil {
		url, err := BuildURLWithQueryString(trackingURL, param.OptionalParams)
		if err != nil {
			return "", err
		}
		return url, nil
	}

	return trackingURL, nil
}

// BuildURLWithQueryString returns the url with query string
func BuildURLWithQueryString(uri string, params interface{}) (string, *error.AfterShipError) {
	queryStringObj, err := query.Values(params)
	if err != nil {
		return "", error.MakeSdkError(error.ErrorTypeHandlerError, err.Error(), params)
	}

	queryString := queryStringObj.Encode()
	if queryString != "" {
		conn := "?"
		if strings.Contains(uri, "?") {
			conn = "&"
		}

		uri += conn + queryString
	}

	return uri, nil
}
