package billingv2

import (
	"github.com/apiheat/go-edgegrid/v6/edgegrid"
	"github.com/apiheat/go-edgegrid/v6/edgegrid/client"
)

const (
	// Represents base path used for Akamai calls towards APIs.
	basePath = "/billing-center-api/v2"
)

// Billingv2 provides the API operation methods for making requests to
// Akamai Billingv2. See this package's package overview docs
// for details on the service.
type Billingv2 struct {
	*client.Client
}

// New creates a new instance of the Billingv2 client with a config.
// If additional configuration is needed for the client instance use the optional
// edgegrid.Config parameter to add your extra config.
//
// Example:
//     Create a Billingv2 client from just a config.
//     svc := Billingv2.New(myConfig))
func New(cfgs *edgegrid.Config) *Billingv2 {
	return newClient(cfgs)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg *edgegrid.Config) *Billingv2 {
	svc := &Billingv2{
		Client: client.New(cfg),
	}

	return svc
}
