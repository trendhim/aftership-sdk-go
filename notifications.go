package aftership

import (
	"context"
	"fmt"
	"net/http"
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

// NotificationsEndpoint provides the interface for all notifications handling API calls
type NotificationsEndpoint interface {
	// AddNotification adds notifications to a single tracking.
	AddNotification(ctx context.Context, id TrackingIdentifier, data Notification) (Notification, error)

	// RemoveNotification removes notifications from a single tracking.
	RemoveNotification(ctx context.Context, id TrackingIdentifier, data Notification) (Notification, error)

	// GetNotification gets contact information for the users to notify when the tracking changes. Please note that only customer receivers will be returned.
	// Any email, sms or webhook that belongs to the Store will not be returned.
	GetNotification(ctx context.Context, id TrackingIdentifier) (Notification, error)
}

// notificationEndpointImpl is the implementation of notification endpoint
type notificationEndpointImpl struct {
	request requestHelper
}

// newNotificationEndpoint creates a instance of notification endpoint
func newNotificationEndpoint(req requestHelper) NotificationsEndpoint {
	return &notificationEndpointImpl{
		request: req,
	}
}

// AddNotification adds notifications to a single tracking.
func (impl *notificationEndpointImpl) AddNotification(ctx context.Context, id TrackingIdentifier, notification Notification) (Notification, error) {
	uriPath := fmt.Sprintf("/notifications%s/add", id.URIPath())
	var wrapper notificationWrapper
	err := impl.request.makeRequest(ctx, http.MethodPost, uriPath, nil,
		&notificationWrapper{Notification: notification}, &wrapper)
	return wrapper.Notification, err
}

// RemoveNotification removes notifications from a single tracking.
func (impl *notificationEndpointImpl) RemoveNotification(ctx context.Context, id TrackingIdentifier, notification Notification) (Notification, error) {
	uriPath := fmt.Sprintf("/notifications%s/remove", id.URIPath())
	var wrapper notificationWrapper
	err := impl.request.makeRequest(ctx, http.MethodPost, uriPath, nil,
		&notificationWrapper{Notification: notification}, &wrapper)
	return wrapper.Notification, err
}

// GetNotification gets contact information for the users to notify when the tracking changes.
// Please note that only customer receivers will be returned.
// Any email, sms or webhook that belongs to the Store will not be returned.
func (impl *notificationEndpointImpl) GetNotification(ctx context.Context, id TrackingIdentifier) (Notification, error) {
	uriPath := fmt.Sprintf("/notifications%s", id.URIPath())
	var wrapper notificationWrapper
	err := impl.request.makeRequest(ctx, http.MethodGet, uriPath, nil, nil, &wrapper)
	return wrapper.Notification, err
}
