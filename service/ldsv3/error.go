package ldsv3

import "fmt"

// LsdErrorv3
type LsdErrorv3 struct {
	Details []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Action string `json:"action"`
		} `json:"data"`
	} `json:"details"`
	Code       string `json:"code"`
	Title      string `json:"title"`
	IncidentID string `json:"incidentId"`
}

// Errors type
type Errors struct {
	Error     string `json:"error"`
	FieldName string `json:"fieldName"`
}

// Error returns the string representation of the error.
// See ErrorWithExtra for formatting.
// Satisfies the error interface.
func (lde LsdErrorv3) Error() string {
	msg := fmt.Sprintf("%s\n\t%s", lde.Title, lde.Code)

	return msg
}
