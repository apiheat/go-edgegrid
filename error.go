package edgegrid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// An ErrorResponse reports one or more errors caused by an API request.
//
// error
type ErrorResponse struct {
	Type        string         `json:"type"`
	Title       string         `json:"title"`
	Status      int            `json:"status"`
	Detail      string         `json:"detail"`
	Instance    string         `json:"instance"`
	Method      string         `json:"method"`
	ServerIP    string         `json:"serverIp"`
	ClientIP    string         `json:"clientIp"`
	RequestID   string         `json:"requestId"`
	RequestTime string         `json:"requestTime"`
	Response    *http.Response `json:"-"`
}

// An ErrorResponse Error() function implementation
//
// error
func (e *ErrorResponse) Error() string {
	path, _ := url.QueryUnescape(e.Response.Request.URL.Opaque)
	u := fmt.Sprintf("%s://%s%s", e.Response.Request.URL.Scheme, e.Response.Request.URL.Host, path)
	return fmt.Sprintf("%s %s: %d %s", e.Response.Request.Method, u, e.Response.StatusCode, e.Detail)
}

// CheckResponse checks the API response for errors, and returns them if present.
//
// error
// #TODO: Should be deprecated
func CheckResponse(r *http.Response) error {

	log.Debug("[CheckResponse]::Check response code")
	log.Debug("[CheckResponse]::Response code is:" + strconv.Itoa(r.StatusCode))
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	log.Debug("[CheckResponse]::Create error response object")
	errorResponse := &ErrorResponse{Response: r}

	log.Debug("[CheckResponse]::Read body")
	data, err := ioutil.ReadAll(r.Body)

	log.Debug("[CheckResponse]::Error body is: " + string(data))

	if err == nil && data != nil {
		log.Debug("[CheckResponse]::Process response body")
		errRespData := json.Unmarshal(data, &errorResponse)

		if errRespData == nil {
			log.Debug("[CheckResponse]::Succesfully processed body")
			errorResponse.Status = r.StatusCode
			errorResponse.Title = r.Status
		} else {
			log.Debug("[CheckResponse]::Failover to provide raw data object.response")
		}

		errorResponse.Response = r
	}

	return errorResponse
}

// CheckRespForErrorv2 checks the API response for errors, and returns bool value
//
// error
func CheckRespForErrorv2(r *http.Response) bool {

	log.Debug("[CheckResponse]::Check response code")
	log.Debug("[CheckResponse]::Response code is:" + strconv.Itoa(r.StatusCode))
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		log.Debug(fmt.Sprintf("[CheckResponse]::Response code is %v - matches 200, 201, 202, 204, 304", r.StatusCode))
		return false
	}

	log.Debug("[CheckResponse]::Read body")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Debug(fmt.Sprintf("[CheckResponse]::Error reading body - %s", err))
	}
	log.Debug("[CheckResponse]::Error body ... ")
	log.Debug(fmt.Sprintf("[CheckResponse]:: %s", string(data)))

	return true
}
