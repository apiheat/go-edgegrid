package diagnosticv2

import (
	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/apiheat/go-edgegrid/v6/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/diagnostic-tools/v2"
)

// Diagnosticv2 provides the API operation methods for making requests to
// Akamai diagnostic tools. See this package's package overview docs
// for details on the service.
type Diagnosticv2 struct {
	*client.Client
}

// New creates a new instance of the Diagnosticv2 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     // Create a Diagnosticv2 client from just a config.
//     svc := diagnosticv2.New(myConfig))
func New(cfgs *edgegrid.Config) *Diagnosticv2 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Diagnosticv2 {
	svc := &Diagnosticv2{
		Client: client.New(cfg),
	}

	return svc
}
