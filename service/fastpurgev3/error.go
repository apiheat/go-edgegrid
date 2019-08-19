package fastpurgev3

import "fmt"

type FastpurgeError struct {
	SupportID   string `json:"supportId"`
	Title       string `json:"title"`
	HTTPStatus  int64  `json:"httpStatus"`
	Detail      string `json:"detail"`
	DescribedBy string `json:"describedBy"`
}

// Error returns the string representation of the error.
// See ErrorWithExtra for formatting.
// Satisfies the error interface.
func (b FastpurgeError) Error() string {
	msg := fmt.Sprintf("%s\n\t%s", b.Title, b.Detail)

	return msg
}
