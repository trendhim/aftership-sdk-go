package aftership

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/google/uuid"
)

// requestHelper is the API request helper interface
type requestHelper interface {
	// makeRequest makes a AfterShip API calls
	makeRequest(ctx context.Context, method string, uri string, queryParams interface{}, inputData interface{}, resultData interface{}) error
	getRateLimit() RateLimit
}

// requestHelperImpl is the implementation of API Request
type requestHelperImpl struct {
	Client          *http.Client
	RateLimit       *RateLimit
	APIKey          string
	BaseURL         string
	UserAgentPrefix string
}

// newRequestHelper returns the instance of API Request helper
func newRequestHelper(cfg Config) requestHelper {
	rh := &requestHelperImpl{
		APIKey:          cfg.APIKey,
		BaseURL:         cfg.BaseURL,
		UserAgentPrefix: cfg.UserAgentPrefix,
		RateLimit:       &RateLimit{},
	}
	if cfg.HTTPClient == nil {
		rh.Client = &http.Client{}
	} else {
		rh.Client = cfg.HTTPClient
	}

	return rh
}

// makeRequest makes a AfterShip API calls
func (impl *requestHelperImpl) makeRequest(ctx context.Context, method string, path string,
	queryParams interface{}, inputData interface{}, resultData interface{}) error {

	var body io.Reader
	if inputData != nil {
		jsonData, err := json.Marshal(inputData)
		if err != nil {
			return err
		}

		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, impl.BaseURL+path, body)
	if err != nil {
		return err
	}

	// Add headers
	req.Header.Add("aftership-api-key", impl.APIKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("request-id", uuid.New().String())
	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", impl.UserAgentPrefix, VERSION))
	req.Header.Add("aftership-agent", fmt.Sprintf("go-sdk-%s", VERSION))

	if queryParams != nil {
		queryStringObj, err := query.Values(queryParams)
		if err != nil {
			return err
		}
		req.URL.RawQuery = queryStringObj.Encode()
	}

	// Send request
	resp, err := impl.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Rate Limit
	setRateLimit(impl.RateLimit, resp)

	result := &Response{
		Meta: Meta{},
		Data: resultData,
	}
	// Unmarshal response object
	err = json.Unmarshal(contents, result)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		// The 2xx range indicate success
		return nil
	}

	return &APIError{
		Type:    result.Meta.Type,
		Code:    result.Meta.Code,
		Message: result.Meta.Message,
		Path:    path,
	}
}

func setRateLimit(rateLimit *RateLimit, resp *http.Response) {
	if rateLimit != nil && resp != nil && resp.Header != nil {
		// reset timestamp
		if reset := resp.Header.Get("x-ratelimit-reset"); reset != "" {
			if n, err := strconv.ParseInt(reset, 10, 64); err == nil {
				rateLimit.Reset = n
			}
		}

		// limit
		if limit := resp.Header.Get("x-ratelimit-limit"); limit != "" {
			if i, err := strconv.Atoi(limit); err == nil {
				rateLimit.Limit = i
			}
		}

		// remaining
		if remaining := resp.Header.Get("x-ratelimit-remaining"); remaining != "" {
			if i, err := strconv.Atoi(remaining); err == nil {
				rateLimit.Remaining = i
			}
		}
	}
}

func (impl *requestHelperImpl) getRateLimit() RateLimit {
	return *impl.RateLimit
}
