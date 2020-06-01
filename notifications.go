package aftership

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Notification is the model describing an AfterShip notification
type Notification struct {
	Emails []string `json:"emails"`
	SMSes  []string `json:"smses"`
}

// notificationWrapper is the notification wrapper.
type notificationWrapper struct {
	Notification Notification `json:"notification"`
}

// GetNotification gets contact information for the users to notify when the tracking changes.
// Please note that only customer receivers will be returned.
// Any email, sms or webhook that belongs to the Store will not be returned.
func (client *Client) GetNotification(ctx context.Context, identifier TrackingIdentifier) (Notification, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Notification{}, errors.Wrap(err, "error getting notification")
	}

	uriPath = fmt.Sprintf("/notifications%s", uriPath)
	var wrapper notificationWrapper
	err = client.makeRequest(ctx, http.MethodGet, uriPath, nil, nil, &wrapper)
	return wrapper.Notification, err
}

// AddNotification adds notifications to a single tracking.
func (client *Client) AddNotification(ctx context.Context, identifier TrackingIdentifier, notification Notification) (Notification, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Notification{}, errors.Wrap(err, "error adding notification")
	}

	uriPath = fmt.Sprintf("/notifications%s/add", uriPath)
	var wrapper notificationWrapper
	err = client.makeRequest(ctx, http.MethodPost, uriPath, nil,
		&notificationWrapper{Notification: notification}, &wrapper)
	return wrapper.Notification, err
}

// RemoveNotification removes notifications from a single tracking.
func (client *Client) RemoveNotification(ctx context.Context, identifier TrackingIdentifier, notification Notification) (Notification, error) {
	uriPath, err := identifier.URIPath()
	if err != nil {
		return Notification{}, errors.Wrap(err, "error removing notification")
	}

	uriPath = fmt.Sprintf("/notifications%s/remove", uriPath)
	var wrapper notificationWrapper
	err = client.makeRequest(ctx, http.MethodPost, uriPath, nil,
		&notificationWrapper{Notification: notification}, &wrapper)
	return wrapper.Notification, err
}
