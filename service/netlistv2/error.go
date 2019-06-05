package netlistv2

import "fmt"

// NetworkListErrorv2 represents the error returned from Akamai
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#errors
type NetworkListErrorv2 struct {
	Detail      string `json:"detail"`
	Instance    string `json:"instance"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	FieldErrors struct {
		Entry []struct {
			Key   string   `json:"key"`
			Value []string `json:"value"`
		} `json:"entry"`
	} `json:"fieldErrors"`
}

// Error returns the string representation of the error.
//
// See ErrorWithExtra for formatting.
//
// Satisfies the error interface.
func (b NetworkListErrorv2) Error() string {
	msg := fmt.Sprintf("%s\n\t%s", b.Title, b.Detail)

	return msg
}
