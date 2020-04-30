package aftership

import (
	"github.com/aftership/aftership-sdk-go/v2/checkpoint"
	"github.com/aftership/aftership-sdk-go/v2/conf"
	"github.com/aftership/aftership-sdk-go/v2/courier"
	"github.com/aftership/aftership-sdk-go/v2/notification"
	"github.com/aftership/aftership-sdk-go/v2/request"
	"github.com/aftership/aftership-sdk-go/v2/tracking"
)

// AfterShip is the client for all AfterShip API calls
type AfterShip struct {
	Courier        courier.Endpoint      // The endpoint to get a list of supported couriers.
	Tracking       tracking.Endpoint     // The endpoint to create trackings, update trackings, and get tracking results.
	LastCheckpoint checkpoint.Endpoint   // The endpoint to get tracking information of the last checkpoint of a tracking.
	Notification   notification.Endpoint // The endpoint to get, add or remove contacts (sms or email) to be notified when the status of a tracking has changed.
}

// NewAfterShip returns the AfterShip client
func NewAfterShip(conf conf.AfterShipConf) *AfterShip {
	if conf.Endpoint == "" {
		conf.Endpoint = "https://api.aftership.com/v4"
	}

	if conf.UserAagentPrefix == "" {
		conf.UserAagentPrefix = "aftership-sdk-go"
	}

	req := request.NewRequest(conf)
	return &AfterShip{
		Courier:        courier.NewEnpoint(req),
		Tracking:       tracking.NewEnpoint(req),
		LastCheckpoint: checkpoint.NewEnpoint(req),
		Notification:   notification.NewEnpoint(req),
	}
}
