package edgegrid

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// An AkamaiGeneralError reports one or more errors caused by an API request.
type AkamaiGeneralError struct {
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

// An AkamaiGeneralError Error() function implementation
func (e *AkamaiGeneralError) Error() string {
	return ShowJSONMessage(e)
}

// An EdgegridError is used to provide higher level clients with
// error which occured. Later on can be casted to specific type if needed
type EdgegridError struct {
	ResponseCode int    `json:"response_code"`
	ResponseBody string `json:"response_body"`
}

// An EdgegridError Error() function implementation
func (e *EdgegridError) Error() string {
	return ShowJSONMessage(e.ResponseBody)
}

// ShowJSONMessage returns string JSON message
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
