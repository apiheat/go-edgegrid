package edgegrid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// An ErrorResponse reports one or more errors caused by an API request.
//
// error
type ErrorResponse struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Status      int    `json:"status"`
	Detail      string `json:"detail"`
	Instance    string `json:"instance"`
	Method      string `json:"method"`
	ServerIP    string `json:"serverIp"`
	ClientIP    string `json:"clientIp"`
	RequestID   string `json:"requestId"`
	RequestTime string `json:"requestTime"`
}

// An ErrorResponse Error() function implementation
//
// error
func (e *ErrorResponse) Error() string {
	return ShowJSONMessage(e)
}

// ShowJSONMessage returns string JSON message
//
// error
func ShowJSONMessage(errType interface{}) string {
	b, err := json.Marshal(errType)
	if err != nil {
		return ""
	}

	var prettyJSON bytes.Buffer
	errprettyJSON := json.Indent(&prettyJSON, []byte(string(b)), "", "    ")
	if errprettyJSON != nil {
		fmt.Println(errprettyJSON)
	}
	return string(prettyJSON.Bytes())
}
