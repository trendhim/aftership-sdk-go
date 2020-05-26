package aftership

import (
	"encoding/json"
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
