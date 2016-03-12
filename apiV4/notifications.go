package apiV4

// NotificationSetting a notification setting
type NotificationSetting struct {
	Android []string `json:"android"`
	Emails  []string `json:"emails"`
	Ios     []string `json:"ios"`
	Smses   []string `json:"smses"`
}

type NotificationSettingWrapper struct {
	Notification NotificationSetting `json:"notification"`
}

// NotificationSettingEnvelope is the message envelope for the notification API responses
type NotificationSettingEnvelope struct {
	Meta ResponseMeta               `json:"meta"`
	Data NotificationSettingWrapper `json:"data"`
}

// ResponseCode provides implementation of Response.ResponseCode()
// for NotificationSettingEnvelope struct
func (envelope *NotificationSettingEnvelope) ResponseCode() ResponseMeta {
	return envelope.Meta
}

// NotificationsHandler provides the interface for all notifications handling API calls
// in AfterShip APIV4
type NotificationsHandler interface {
	// AddNotification Adds notifications to a tracking number.
	AddNotification(id TrackingId, notification NotificationSetting) (NotificationSetting, AfterShipApiError)

	// RemoveNotification Removes notifications from a tracking number.
	RemoveNotification(id TrackingId, notification NotificationSetting) (NotificationSetting, AfterShipApiError)

	// GetNotificationSetting Gets notifications value from a tracking number.
	GetNotificationSetting(id TrackingId, fields string) (NotificationSetting, AfterShipApiError)
}
