package contractsv1

import (
	"fmt"
	"strings"
)

// ContractsErrorv1 type
type ContractsErrorv1 struct {
	Status string   `json:"status"`
	Title  string   `json:"title"`
	Type   string   `json:"type"`
	Errors []Errors `json:"errors"`
}

// Errors type
type Errors struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// Error returns the string representation of the error.
// See ErrorWithExtra for formatting.
// Satisfies the error interface.
func (c ContractsErrorv1) Error() string {
	var details []string
	if len(c.Errors) > 0 {
		for _, detail := range c.Errors {
			details = append(details, detail.Detail)
		}

		return fmt.Sprintf("%s\n\tResponse status %s\n\t Details: %s", c.Title, c.Status, strings.Join(details, ","))
	}

	return fmt.Sprintf("%s\n\tResponse status %s", c.Title, c.Status)
}
