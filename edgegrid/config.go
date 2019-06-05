package edgegrid

// Config represents options we can pass during client creation
type Config struct {
	// Defines log level output i.e. debug/error/warning/info
	LogLevel string

	// Defines account switch key used to manage sub-accounts with partner API keys
	AccountSwitchKey string

	// Used for adding the User Agent header for the requests we make towards APIs
	UserAgent string
}
