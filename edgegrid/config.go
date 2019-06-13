package edgegrid

// Config represents options that are passed during client initialization
type Config struct {
	// Defines account switch key used to manage sub-accounts with partner API keys
	AccountSwitchKey string

	//BasePath is used to provide information about what is the base path we would be using for our calls.
	BasePath string

	// Credentials holds the current credentials configuration
	Credentials *Credentials

	// LocalTesting determines if the host we would be using is local - so we can run tests
	LocalTesting bool

	// Defines log level output i.e. debug/error/warning/info
	LogVerbosity string

	// Used for adding the User Agent header for the requests we make towards APIs
	UserAgent string
}

// NewConfig returns a new Config pointer that can be chained with builder
// methods to set multiple configuration values inline without using pointers.
//
//   // Create config with account switch key defined which can be used by the
//   // service clients.
//   cfg := edgegrid.NewConfig().WithAccountSwitchKey("MS-123BV")
//
func NewConfig() *Config {
	return &Config{}
}

// WithAccountSwitchKey sets account switch key used across calls
// a Config pointer.
func (c *Config) WithAccountSwitchKey(ask string) *Config {
	c.AccountSwitchKey = ask
	return c
}

// WithLogVerbosity sets a config log verbosity and returns
// a Config pointer.
func (c *Config) WithLogVerbosity(logVerbosity string) *Config {
	c.LogVerbosity = logVerbosity
	return c
}

// WithCredentials sets a config Credentials value returning a Config pointer
// for chaining.
func (c *Config) WithCredentials(creds *Credentials) *Config {
	c.Credentials = creds
	return c
}
