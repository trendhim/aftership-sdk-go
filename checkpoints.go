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
	GetLastCheckpoint(ctx context.Context, identifier TrackingIdentifier, params GetCheckpointParams) (LastCheckpoint, error)
}

// checkpointsEndpointImpl is the implementation of checkpoint endpoint
type checkpointsEndpointImpl struct {
	helper requestHelper
}

// newCheckpointsEndpoint creates a instance of checkpoint endpoint
func newCheckpointsEndpoint(helper requestHelper) CheckpointsEndpoint {
	return &checkpointsEndpointImpl{
		helper: helper,
	}
}

// GetLastCheckpoint returns the tracking information of the last checkpoint of a single tracking.
func (impl *checkpointsEndpointImpl) GetLastCheckpoint(ctx context.Context, identifier TrackingIdentifier, params GetCheckpointParams) (LastCheckpoint, error) {
	uriPath := fmt.Sprintf("/last_checkpoint%s", identifier.URIPath())
	var lastCheckpoint LastCheckpoint
	err := impl.helper.makeRequest(ctx, http.MethodGet, uriPath, params, nil, &lastCheckpoint)
	return lastCheckpoint, err
}
