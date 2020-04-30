package checkpoint

import (
	"fmt"
	"strings"

	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/request"
)

// Endpoint provides the interface for all checkpoint API calls
type Endpoint interface {
	// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
	GetLastCheckpoint(param SingleTrackingParam, fields string, lang string) (LastCheckpoint, *error.AfterShipError)
}

// EndpointImpl is the implementaion of checkpoint endpoint
type EndpointImpl struct {
	request request.APIRequest
}

// NewEnpoint creates a instance of checkpoint endpoint
func NewEnpoint(req request.APIRequest) Endpoint {
	return &EndpointImpl{
		request: req,
	}
}

// GetLastCheckpoint Return the tracking information of the last checkpoint of a single tracking.
func (impl *EndpointImpl) GetLastCheckpoint(param SingleTrackingParam, fields string, lang string) (LastCheckpoint, *error.AfterShipError) {
	var envelope LastCheckpointEnvelope
	var url string
	if param.ID != "" {
		url = fmt.Sprintf("/last_checkpoint/%s", param.ID)
	} else if param.Slug != "" && param.TrackingNumber != "" {
		url = fmt.Sprintf("/last_checkpoint/%s/%s", param.Slug, param.TrackingNumber)
	}

	if fields != "" || lang != "" {
		extraParams := []string{}

		if fields != "" {
			extraParams = append(extraParams, fmt.Sprintf("fields=%s", fields))
		}

		if lang != "" {
			extraParams = append(extraParams, fmt.Sprintf("lang=%s", lang))
		}

		url += fmt.Sprintf("?%s", strings.Join(extraParams, "&"))
	}

	err := impl.request.MakeRequest("GET", url, nil, &envelope)
	if err != nil {
		return LastCheckpoint{}, err
	}

	return envelope.Data, nil
}
