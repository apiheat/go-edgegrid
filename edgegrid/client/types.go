package client

// ClientResponse represents response from our API call
type ClientResponse struct {
	Body     string
	Response interface{}
}

// ClientOptions represents options we can pass during client creation
type ClientOptions struct {
	ConfigPath       string
	ConfigSection    string
	DebugLevel       string
	AccountSwitchKey string
}
