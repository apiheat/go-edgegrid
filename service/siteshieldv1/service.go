package siteshieldv1

import (
	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/apiheat/go-edgegrid/v6/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/siteshield/v1/maps"
)

// Siteshieldv1 provides the API operation methods for making requests to
// Akamai Siteshieldv1. See this package's package overview docs
// for details on the service.
type Siteshieldv1 struct {
	*client.Client
}

// New creates a new instance of the Siteshieldv1 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     // Create a Siteshieldv1 client from just a config.
//     svc := Siteshieldv1.New(myConfig))
func New(cfgs *edgegrid.Config) *Siteshieldv1 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Siteshieldv1 {
	svc := &Siteshieldv1{
		Client: client.New(cfg),
	}

	return svc
}
