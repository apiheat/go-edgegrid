package client

import (
	"fmt"

	"github.com/apiheat/go-edgegrid/edgegrid"
	"github.com/apiheat/go-edgegrid/edgegrid/signer"

	"gopkg.in/resty.v1"
)

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	Config  *edgegrid.Config
	Rclient *resty.Client
}

// New will return a pointer to a new initialized service client.
func New(cfg *edgegrid.Config, options ...func(*Client)) *Client {
	svc := &Client{
		Config: cfg,
	}

	// Create instance of resty client
	svc.Rclient = resty.New()
	svc.Rclient.SetDebug(true)

	if svc.Config.LocalTesting {
		svc.Rclient.SetHostURL(svc.Config.TestingURL)

	} else {
		svc.Rclient.SetHostURL(fmt.Sprintf("%s://%s", svc.Config.Scheme, svc.Config.Credentials.Host))
	}

	// Create inistance of auth signer
	authSigner := signer.New(svc.Config.Credentials, svc.Config.Scheme, svc.Config.Credentials.Host)

	// Registering Request Middleware - which will run just before every request
	svc.Rclient.SetPreRequestHook(func(c *resty.Client, req *resty.Request) error {

		// Set authentication header with signed data based on request
		req.SetHeader("Authorization", authSigner.SignRequest(req, []string{}, svc.Config.LocalTesting))

		return nil
	})

	return svc
}
