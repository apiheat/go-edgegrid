package edgegrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// AkamaiEnvironmentVar represents Akamai's env variables used
//
// client
type AkamaiEnvironmentVar string

// AkamaiEnvironmentVar const represents Akamai's env variables to be used.
//
// client
const (
	EnvVarEdgercPath        AkamaiEnvironmentVar = "AKAMAI_EDGERC_CONFIG"
	EnvVarEdgercSection     AkamaiEnvironmentVar = "AKAMAI_EDGERC_SECTION"
	EnvVarDebugLevelSection AkamaiEnvironmentVar = "AKAMAI_EDGERC_DEBUGLEVEL"
)

// AkamaiEnvironment represents Akamai's target environment type.
//
// client
type AkamaiEnvironment string

// AkamaiEnvironment const represents Akamai's target environment to be used in calls.
//
// client
const (
	Production AkamaiEnvironment = "production"
	Staging    AkamaiEnvironment = "staging"
)

const (
	userAgent = "go-edgegrid"
)

// Client represents Akamai's API client for communicating with service
//
// client
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// This base URL comes from edgerc config.
	baseURL *url.URL

	// edgerc credentials
	credentials *EdgercCredentials

	// Services used for talking to different parts of the Akamai API.
	Auth         *AuthService
	Debug        *DebugService
	NetworkLists *NetworkListService
	PropertyAPI  *PropertyAPIService
	ReportingAPI *ReportingAPIService
}

// ClientOptions represents options we can pass during client creation
//
// client
type ClientOptions struct {
	ConfigPath    string
	ConfigSection string
	DebugLevel    string
}

// ClientResponse represents response from our API call
//
// client
type ClientResponse struct {
	Body     string
	Response *http.Response
}

var (
	apiPaths = map[string]string{
		"network_list": "/network-list/v1/network_lists",
		"papi_v1":      "/papi/v1",
		"reporting_v1": "/reporting-api/v1/reports",
	}
)

// NewClient returns a new edgegrid.Client for API. If a nil httpClient is
// provided, http.DefaultClient will be used.
//
// client
func NewClient(httpClient *http.Client, conf *ClientOptions) (*Client, error) {
	var (
		path, section, debuglvl string
	)

	// If we do not pass config we will try to to use env variables
	if conf != nil {
		path = conf.ConfigPath
		section = conf.ConfigSection
		debuglvl = conf.DebugLevel
	} else {
		path = os.Getenv(string(EnvVarEdgercPath))
		section = os.Getenv(string(EnvVarEdgercSection))
		debuglvl, _ = os.LookupEnv(string(EnvVarDebugLevelSection))
	}

	switch debuglvl {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	log.WithFields(log.Fields{
		"path":     path,
		"section":  section,
		"debuglvl": debuglvl,
	}).Info("Create new edge client")

	return nil, newClient(httpClient, path, section)
}

// newClient *private* function to initiaite client
//
// client
func newClient(httpClient *http.Client, edgercPath, edgercSection string) *Client {
	var errInitEdgerc error

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	log.Debug("[newClient]::Create new client object")
	c := &Client{client: httpClient}

	log.Debug("[newClient]::Create credentials")
	c.credentials, errInitEdgerc = InitEdgerc(edgercPath, edgercSection)

	if errInitEdgerc != nil {
		fmt.Printf("Error loading file? %s", errInitEdgerc)
		return nil
	}

	// Set base URL for making all API requests
	log.Debug("[newClient]::SetBaseURL")
	c.SetBaseURL(c.credentials.host, false)

	// Create all the public services.
	log.Debug("[newClient]::Create service Auth")
	c.Auth = &AuthService{client: c}

	log.Debug("[newClient]::Create service NetworkLists")
	c.NetworkLists = &NetworkListService{client: c}

	log.Debug("[newClient]::Create service PropertyAPI")
	c.PropertyAPI = &PropertyAPIService{client: c}

	log.Debug("[newClient]::Create service ReportingAPI")
	c.ReportingAPI = &ReportingAPIService{client: c}

	log.Debug("[newClient]::Create service Debug")
	c.Debug = &DebugService{client: c}

	return c
}

// newRequest creates an HTTP request that can be sent to Akamai APIs. A relative URL can be provided in path, which will be resolved to the
// Host specified in Config. If body is specified, it will be sent as the request body.
//
// client
func (cl *Client) NewRequest(method, path string, vreq, vresp interface{}) (*ClientResponse, error) {

	targetURL, _ := prepareURL(cl.baseURL, path)

	log.WithFields(log.Fields{
		"method": method,
		"base":   cl.baseURL,
		"path":   path,
	}).Info("Create new request")

	log.Debug("[NewRequest]::Create http request")
	req, err := http.NewRequest(method, targetURL.String(), nil)
	if err != nil {
		return nil, nil
	}

	if method == "POST" || method == "PUT" {
		log.Info("Prepare request body object")
		log.Debug("[NewRequest]::Method is POST/PUT")
		log.Debug("[NewRequest]::Marshal request object")
		bodyBytes, err := json.Marshal(vreq)
		if err != nil {
			return nil, err
		}
		bodyReader := bytes.NewReader(bodyBytes)

		log.Debug("[NewRequest]::Body object added to request")
		req.Body = ioutil.NopCloser(bodyReader)
		req.ContentLength = int64(bodyReader.Len())

		log.Debug("[NewRequest]::Body object is:" + string(bodyBytes))
		log.Debug("[NewRequest]::Set header Content-Type to 'application/json' ")
		req.Header.Set("Content-Type", "application/json")

	}

	authorizationHeader := AuthString(cl.credentials, req, []string{})
	log.Debug("[NewRequest]::Set header Authorization")
	req.Header.Add("Authorization", authorizationHeader)

	log.Info("Execute http request")
	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Debug("[NewRequest]::Process response")
	clientResp := &ClientResponse{}

	err = CheckResponse(resp)
	if err != nil {
		clientResp.Response = resp
		clientResp.Body = ""

		return clientResp, err
	}

	byt, _ := ioutil.ReadAll(resp.Body)

	clientResp.Response = resp
	clientResp.Body = string(byt)

	if vresp != nil {
		if err = json.Unmarshal([]byte(byt), &vresp); err != nil {
			return clientResp, err
		}
	}

	log.Debug("[NewRequest]::Return response")

	return clientResp, err
}

// SetBaseURL sets the base URL for API requests to a custom endpoint.
//
// client
func (cl *Client) SetBaseURL(urlStr string, passThrough bool) error {

	log.WithFields(log.Fields{
		"urlStr":      urlStr,
		"passThrough": passThrough,
	}).Info("Set BaseURL for client")

	var err error

	if passThrough == true {
		cl.baseURL, err = url.Parse(urlStr)
	} else {
		if strings.HasPrefix(urlStr, "https://") {
			cl.baseURL, err = url.Parse(urlStr)
		} else {
			cl.baseURL, err = url.Parse("https://" + urlStr)
		}
	}

	log.WithFields(log.Fields{
		"baseURL": cl.baseURL.String(),
	}).Debug("[SetBaseURL]::Base URL set")

	return err
}

// prepareURL returns URL which is used to make API call
//
// client
func prepareURL(url *url.URL, path string) (*url.URL, error) {

	rel, err := url.Parse(strings.TrimPrefix(path, "/"))
	if err != nil {
		return nil, err
	}

	u := url.ResolveReference(rel)

	return u, nil
}
