package edgegrid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
func CheckResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		if err := json.Unmarshal(data, &errorResponse); err != nil {
			errorResponse.Status = r.StatusCode
			errorResponse.Title = r.Status
		}

		errorResponse.Response = r
	}

	return errorResponse
}
