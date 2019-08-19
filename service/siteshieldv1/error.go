package siteshieldv1

import "fmt"

// SiteshieldErrorv1 represents the error returned from Akamai
type SiteshieldErrorv1 struct {
	Type        string   `json:"type"`
	Title       string   `json:"title"`
	Status      int64    `json:"status"`
	Detail      string   `json:"detail"`
	Instance    string   `json:"instance"`
	Method      string   `json:"method"`
	RequestTime string   `json:"requestTime"`
	Errors      []Errors `json:"errors"`
}

// Errors type
type Errors struct {
	Error     string `json:"error"`
	FieldName string `json:"fieldName"`
}

// Error returns the string representation of the error.
// See ErrorWithExtra for formatting.
// Satisfies the error interface.
func (b SiteshieldErrorv1) Error() string {
	msg := fmt.Sprintf("%s\n\t%s", b.Title, b.Detail)

	return msg
}
