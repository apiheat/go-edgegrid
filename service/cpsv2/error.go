package cpsv2

import "fmt"

// CpsErrorv2
type CpsErrorv2 struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

// Errors type
type Errors struct {
	Error     string `json:"error"`
	FieldName string `json:"fieldName"`
}

// Error returns the string representation of the error.
// See ErrorWithExtra for formatting.
// Satisfies the error interface.
func (cps CpsErrorv2) Error() string {
	return fmt.Sprintf("%s\n\t%s(%s)", cps.Title, cps.Detail, cps.Instance)
}
