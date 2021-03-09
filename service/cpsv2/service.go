package cpsv2

import (
	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/apiheat/go-edgegrid/v6/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/cps/v2"
)

// Cpsv2 provides the API operation methods for making requests to
// Akamai Cpsv2. See this package's package overview docs
// for details on the service.
type Cpsv2 struct {
	*client.Client
}

// New creates a new instance of the Cpsv2 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     // Create a Cpsv2 client from just a config.
//     svc := Cpsv2.New(myConfig))
func New(cfgs *edgegrid.Config) *Cpsv2 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Cpsv2 {
	svc := &Cpsv2{
		Client: client.New(cfg),
	}

	return svc
}
