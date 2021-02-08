package billingv2

import "fmt"

// BillingErrorv2 API error response
type BillingErrorv2 struct {
	Details []struct {
		Field   string `json:"field"`
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			Field string `json:"field"`
			Value string `json:"value"`
		} `json:"data"`
	} `json:"details"`
	Code       string      `json:"code"`
	Title      string      `json:"title"`
	IncidentID string      `json:"incidentId"`
	Resolution interface{} `json:"resolution"`
	HelpURL    interface{} `json:"helpUrl"`
}

// Errors type
type Errors struct {
	Error     string `json:"error"`
	FieldName string `json:"fieldName"`
}

// Error returns the string representation of the error.
// See ErrorWithExtra for formatting.
// Satisfies the error interface.
func (b BillingErrorv2) Error() string {
	msg := fmt.Sprintf("%s\n\t%s", b.Title, b.Code)

	return msg
}
