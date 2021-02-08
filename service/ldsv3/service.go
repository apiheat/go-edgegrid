package ldsv3

import (
	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/apiheat/go-edgegrid/v6/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/lds-api/v3"
)

// Ldsv3 provides the API operation methods for making requests to
// Akamai Ldsv3. See this package's package overview docs
// for details on the service.
type Ldsv3 struct {
	*client.Client
}

// New creates a new instance of the Ldsv3 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     // Create a Ldsv3 client from just a config.
//     svc := Ldsv3.New(myConfig))
func New(cfgs *edgegrid.Config) *Ldsv3 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Ldsv3 {
	svc := &Ldsv3{
		Client: client.New(cfg),
	}

	return svc
}
