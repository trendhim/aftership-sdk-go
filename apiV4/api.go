// apiV4 exposes all API level functions and objects of AfterShip API version 4 for
// go code
package apiV4

const URL = "https://api.aftership.com/v4"
const API_KEY_HEADER_FIELD = "aftership-api-key"
const COURIERS_ENDPOINT = "/couriers"
const COURIERS_ALL_ENDPOINT = "/couriers/all"
const COURIERS_DETECT_ENDPOINT = "/couriers/detect"
const TRACKINGS_ENDPOINT = "/trackings"
const TRACKINGS_EXPORTS_ENDPOINT = "/trackings/exports"
const LAST_CHECKPOINT_ENDPOINT = "/last_checkpoint"
const NOTIFICATIONS = "/notifications"

type Response interface {
	ResponseCode() ResponseMeta
}

// RetryPolicy configures retry policy
type RetryPolicy struct {
	RetryOnError            bool
	ErrorRetryCount         int
	RetryOnHittingRateLimit bool
}

type ResponseMeta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type AfterShipApiError struct {
	ResponseMeta
}
