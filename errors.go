package aftership

import (
	"encoding/json"
)

// Error messages
const (
	errEmptyAPIKey                 = "invalid credentials: API Key must not be empty"
	errEmptyAPISecret              = "invalid credentials: API Secret must not be empty"
	errMissingTrackingNumber       = "tracking number is empty and must be provided"
	errMissingTrackingID           = "tracking id is empty and must be provided"
	errMissingSlugOrTrackingNumber = "slug or tracking number is empty, both of them must be provided"
	errExceedRateLimit             = "You have exceeded the API call rate limit. The default limit is 10 requests per second."
	errMarshallingJSON             = "Invalid JSON data."
)

// System error code
const (
	codeRateLimiting = iota + 4900
	codeJSONError
	codeBadRequest
	codeBadParam
	codeSignatureError
	codeRequestFailed
	codeEmptyBody
	codeRequestTimeout
)

// APIError is the error in AfterShip API calls
type APIError struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// TooManyRequestsError is the too many requests error in AfterShip API calls
type TooManyRequestsError struct {
	APIError
	RateLimit *RateLimit `json:"rate_limit"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *TooManyRequestsError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}
