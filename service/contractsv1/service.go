package contractsv1

import (
	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/apiheat/go-edgegrid/v6/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/contract-api/v1"
)

// Contractsv1 provides the API operation methods for making requests to
// Akamai Contractsv1. See this package's package overview docs
// for details on the service.
type Contractsv1 struct {
	*client.Client
}

// New creates a new instance of the Contractsv1 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     // Create a Contractsv1 client from just a config.
//     svc := Contractsv1.New(myConfig))
func New(cfgs *edgegrid.Config) *Contractsv1 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Contractsv1 {
	svc := &Contractsv1{
		Client: client.New(cfg),
	}

	return svc
}
