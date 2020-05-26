package aftership

import (
	"context"
	"fmt"
	"net/http"
)

// GetCheckpointParams is the additional parameters in checkpoint query
type GetCheckpointParams struct {
	// List of fields to include in the response. Use comma for multiple values.
	// Fields to include:slug,created_at,checkpoint_time,city,coordinates,country_iso3,
	// country_name,message,state,tag,zip
	// Default: none, Example: city,tag
	Fields string `url:"fields,omitempty" json:"fields,omitempty"`

	// Support Chinese to English translation for china-ems  and  china-post  only (Example: en)
	Lang string `url:"lang,omitempty" json:"lang,omitempty"`
}

// LastCheckpoint is the last checkpoint API response
type LastCheckpoint struct {
	ID             string     `json:"id,omitempty"`
	Slug           string     `json:"slug,omitempty"`
	TrackingNumber string     `json:"tracking_number"`
	Tag            string     `json:"tag"`
	Checkpoint     Checkpoint `json:"checkpoint"`
}

// CheckpointsEndpoint provides the interface for all checkpoint API calls
type CheckpointsEndpoint interface {
	// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
	GetLastCheckpoint(ctx context.Context, id TrackingIdentifier, optionalParams GetCheckpointParams) (LastCheckpoint, error)
}

// CheckpointsEndpointImpl is the implementation of checkpoint endpoint
type CheckpointsEndpointImpl struct {
	request requestHelper
}

// NewCheckpointsEndpoint creates a instance of checkpoint endpoint
func NewCheckpointsEndpoint(req requestHelper) CheckpointsEndpoint {
	return &CheckpointsEndpointImpl{
		request: req,
	}
}

// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
func (impl *CheckpointsEndpointImpl) GetLastCheckpoint(ctx context.Context, id TrackingIdentifier, optionalParams GetCheckpointParams) (LastCheckpoint, error) {
	uriPath := fmt.Sprintf("/last_checkpoint%s", id.URIPath())
	var lastCheckpoint LastCheckpoint
	err := impl.request.makeRequest(ctx, http.MethodGet, uriPath, optionalParams, nil, &lastCheckpoint)
	return lastCheckpoint, err
}
