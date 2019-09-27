package fastpurgev3

import (
	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/apiheat/go-edgegrid/v6/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/ccu/v3"
)

// Fastpurgev3 provides the API operation methods for making requests to
// Akamai Fastpurgev3. See this package's package overview docs
// for details on the service.
type Fastpurgev3 struct {
	*client.Client
}

// New creates a new instance of the Fastpurgev3 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     // Create a Fastpurgev3 client from just a config.
//     svc := Fastpurgev3.New(myConfig))
func New(cfgs *edgegrid.Config) *Fastpurgev3 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Fastpurgev3 {
	svc := &Fastpurgev3{
		Client: client.New(cfg),
	}

	return svc
}
