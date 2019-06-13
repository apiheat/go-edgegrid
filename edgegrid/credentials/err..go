package credentials

// Error represents an error caused during credentials retrieval
type Error struct {
	ErrorMessage string `json:"error_message"`
	ErrorType    string `json:"error_type"`
}

// Error implements the error interface.
func (e Error) Error() string {
	return e.ErrorMessage
}
