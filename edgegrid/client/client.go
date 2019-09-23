package client

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/apiheat/go-edgegrid/edgegrid"
	"github.com/apiheat/go-edgegrid/edgegrid/signer"
	"github.com/go-resty/resty/v2"
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

	switch svc.Config.LogVerbosity {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}

	if svc.Config.Credentials == nil {
		log.Fatalln("Cannot create client without credentials!")
	}

	// Create instance of resty client
	svc.Rclient = resty.New()

	//Sets headers and customize the user agent
	svc.Rclient.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   svc.Config.UserAgent,
	})

	svc.Rclient.SetDebug(svc.Config.RequestDebug)

	if svc.Config.LocalTesting {
		svc.Rclient.SetHostURL(svc.Config.TestingURL)

	} else {
		svc.Rclient.SetHostURL(fmt.Sprintf("%s://%s", svc.Config.Scheme, svc.Config.Credentials.Host))
	}

	// Create inistance of auth signer
	authSigner := signer.New(svc.Config.Credentials, svc.Config.Scheme, svc.Config.Credentials.Host)

	// Registering Request Middleware - which will run just before every request is prepared
	svc.Rclient.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {

		if svc.Config.AccountSwitchKey != "" {
			r.SetQueryParam("accountSwitchKey", svc.Config.AccountSwitchKey)
		}

		return nil
	})

	// Registering Request Middleware - which will run just before every request but after
	// preparation of the request.
	svc.Rclient.SetPreRequestHook(func(c *resty.Client, req *http.Request) error {

		req.Header.Set("Authorization", authSigner.SignRequest(req, []string{}))

		// Set authentication header with signed data based on request
		return nil
	})

	return svc
}
