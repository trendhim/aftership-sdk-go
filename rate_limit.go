package aftership

import "time"

// RateLimit is the X-RateLimit value in API response headers
type RateLimit struct {
	Reset     int64 `json:"reset"`     // The unix timestamp when the rate limit will be reset.
	Limit     int   `json:"limit"`     // The rate limit ceiling for your account per sec.
	Remaining int   `json:"remaining"` // The number of requests left for the 1 second window.
}

func (rateLimit *RateLimit) isReached() bool {
	return rateLimit.Remaining == 0 && rateLimit.Reset >= time.Now().Unix()
}
