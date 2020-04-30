package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aftership/aftership-sdk-go/v2/conf"
	"github.com/aftership/aftership-sdk-go/v2/error"
	"github.com/aftership/aftership-sdk-go/v2/response"
	"github.com/google/uuid"
)

// APIRequest API request interface
type APIRequest interface {
	// MakeRequest Make a AfterShip API calls
	MakeRequest(method string, uri string, data interface{}, result response.AftershipResponse) *error.AfterShipError
}

// APIRequestImpl is the implementation of
type APIRequestImpl struct {
	Client           *http.Client
	APIKey           string
	Endpoint         string
	UserAagentPrefix string
}

// NewRequest returns the instance of API Request
func NewRequest(conf conf.AfterShipConf) APIRequest {
	return &APIRequestImpl{
		APIKey:           conf.AppKey,
		Endpoint:         conf.Endpoint,
		UserAagentPrefix: conf.UserAagentPrefix,
	}
}

// MakeRequest Make a AfterShip API calls
func (impl *APIRequestImpl) MakeRequest(method string, uri string, data interface{}, result response.AftershipResponse) *error.AfterShipError {
	if impl.Client == nil {
		impl.Client = &http.Client{}
	}

	var body io.Reader
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return error.MakeRequestError("JsonError", err, data)
		}

		body = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, impl.Endpoint+uri, body)
	if err != nil {
		return error.MakeRequestError("RequestError", err, data)
	}

	req.Header.Add("aftership-api-key", impl.APIKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("request-id", uuid.New().String())
	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", impl.UserAagentPrefix, conf.VERSION))
	req.Header.Add("aftership-agent", fmt.Sprintf("go-sdk-%s", conf.VERSION))

	resp, err := impl.Client.Do(req)
	if err != nil {
		return error.MakeRequestError("RequestError", err, data)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(contents, &result)

	if resp.StatusCode != http.StatusOK {
		return error.MakeAPIError(result)
	}

	return nil
}
