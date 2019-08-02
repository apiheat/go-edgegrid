package netlistv2

import (
	"github.com/apiheat/go-edgegrid/edgegrid"
	"github.com/apiheat/go-edgegrid/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/network-list/v2/network-lists"
)

// Netlistv2 provides the API operation methods for making requests to
// Akamai Netlistv2. See this package's package overview docs
// for details on the service.
type Netlistv2 struct {
	*client.Client
}

// New creates a new instance of the Netlistv2 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     // Create a netlistv2 client from just a config.
//     svc := netlistv2.New(myConfig))
func New(cfgs *edgegrid.Config) *Netlistv2 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Netlistv2 {
	svc := &Netlistv2{
		Client: client.New(cfg),
	}

	return svc
}
