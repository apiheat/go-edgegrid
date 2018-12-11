package edgegrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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

// Akamai Services Paths
const (
	A2PathV1                 = "/adaptive-acceleration/v1/properties"
	NetworkListPathV1        = "/network-list/v1/network_lists"
	NetworkListPathV2        = "/network-list/v2/network-lists"
	PAPIPathV1               = "/papi/v1"
	ReportingPathV1          = "/reporting-api/v1/reports"
	IdentityManagementPathV2 = "/identity-management/v2"
	SiteshieldPathV1         = "/siteshield/v1/maps"
	FRNPathV1                = "/firewall-rules-manager/v1"
	DTPathV2                 = "/diagnostic-tools/v2"
	BillingPathV2            = "/billing-center-api/v2"
)

// AkamaiEnvironment represents Akamai's target environment type.
//
// client
type AkamaiEnvironment string

const (
	Production AkamaiEnvironment = "production"
	Staging    AkamaiEnvironment = "staging"
)

// AkamaiSubscription represents Akamai's notification actions for subscriptions.
//
// client
type AkamaiSubscription string

const (
	Subscribe   AkamaiSubscription = "subscribe"
	Unsubscribe AkamaiSubscription = "unsubscribe"
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
	Auth               *AuthService
	Debug              *DebugService
	NetworkListsv2     *NetworkListServicev2
	Property           *PropertyService
	Reporting          *ReportingService
	A2                 *AdaptiveAccelerationService
	IdentityManagement *IdentityManagementService
	SiteShield         *SiteShieldService
	FRN                *FirewallRulesNotificationsService
	DT                 *DiagToolsService
	Billing            *BillingService
}

// ClientResponse represents response from our API call
//
// client
type ClientResponse struct {
	Body     string
	Response *http.Response
}

// ClientOptions represents options we can pass during client creation
//
// client
type ClientOptions struct {
	ConfigPath    string
	ConfigSection string
	DebugLevel    string
}

// NewClient returns a new edgegrid.Client for API. If a nil httpClient is
// provided, http.DefaultClient will be used.
//
// client
func NewClient(httpClient *http.Client, conf *ClientOptions) (*Client, error) {
	var (
		path, section, debuglvl string
	)

	// Set up path/section and debug level and override if set on ENV variable level
	path = conf.ConfigPath
	section = conf.ConfigSection
	debuglvl = conf.DebugLevel

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
		log.SetLevel(log.ErrorLevel)
	}

	log.WithFields(log.Fields{
		"path":     path,
		"section":  section,
		"debuglvl": debuglvl,
	}).Info("Create new edge client")

	APIClient, errAPIClient := newClient(httpClient, path, section)

	if errAPIClient != nil {
		log.Debug("[newClient]::Create new client object failed: " + errAPIClient.Error())
		return nil, errAPIClient
	}

	return APIClient, nil
}

// newClient *private* function to initiaite client
//
// client
func newClient(httpClient *http.Client, edgercPath, edgercSection string) (*Client, error) {
	var errInitEdgerc error

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	log.Debug("[newClient]::Create new client object")
	c := &Client{client: httpClient}

	log.Debug("[newClient]::Create credentials")
	c.credentials, errInitEdgerc = InitEdgerc(edgercPath, edgercSection)

	if errInitEdgerc != nil {
		return nil, errInitEdgerc
	}

	// Set base URL for making all API requests
	log.Debug("[newClient]::SetBaseURL")
	c.SetBaseURL(c.credentials.host, false)

	// Create all the public services.
	log.Debug("[newClient]::Create service Auth")
	c.Auth = &AuthService{client: c}

	log.Debug("[newClient]::Create service NetworkListsv2")
	c.NetworkListsv2 = &NetworkListServicev2{client: c}

	log.Debug("[newClient]::Create service Property")
	c.Property = &PropertyService{client: c}

	log.Debug("[newClient]::Create service Reporting")
	c.Reporting = &ReportingService{client: c}

	log.Debug("[newClient]::Create service A2")
	c.A2 = &AdaptiveAccelerationService{client: c}

	log.Debug("[newClient]::Create service IdentityManagement")
	c.IdentityManagement = &IdentityManagementService{client: c}

	log.Debug("[newClient]::Create service SiteShield")
	c.SiteShield = &SiteShieldService{client: c}

	log.Debug("[newClient]::Create service FRN")
	c.FRN = &FirewallRulesNotificationsService{client: c}

	log.Debug("[newClient]::Create service DiagnosticTools")
	c.DT = &DiagToolsService{client: c}

	log.Debug("[newClient]::Create service Billing")
	c.Billing = &BillingService{client: c}

	log.Debug("[newClient]::Create service Debug")
	c.Debug = &DebugService{client: c}

	return c, nil
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
		return nil, err
	}

	if method == http.MethodPost || method == http.MethodPut {
		log.Info("Prepare request body object")
		log.Debug("[NewRequest]::Method is POST/PUT")
		log.Debug("[NewRequest]::Marshal request object")

		reqType := reflect.TypeOf(vreq)
		log.Debug(fmt.Sprintf("[NewRequest]::Object request provided type ( %s ) ", reqType))

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
		log.Debug("[NewRequest]::Error making request")
		log.Debug(fmt.Sprintf("[NewRequest]:: %s", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	log.Debug("[NewRequest]::Processing response")
	clientResp := &ClientResponse{}

	log.Debug("[NewRequest]::Read response body")
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debug("[NewRequest]::Error reading response body")
		log.Debug(fmt.Sprintf("[NewRequest]:: %s", err.Error()))
		return nil, err
	}

	log.Debug("[NewRequest]::Set client object response and body")
	clientResp.Response = resp
	clientResp.Body = string(byt)

	log.Debug("[NewRequest]::Response code is:" + strconv.Itoa(resp.StatusCode))
	log.Debug("[NewRequest]::Body is " + clientResp.Body)

	if vresp != nil {
		respType := reflect.TypeOf(vresp)
		log.Debug(fmt.Sprintf("[NewRequest]::Map response to provided type ( %s ) ", respType))

		if err = json.Unmarshal([]byte(byt), &vresp); err != nil {
			log.Debug("[NewRequest]::Error while unmarshaling response body")
			log.Debug(fmt.Sprintf("[NewRequest]:: %s", err.Error()))
			return nil, err
		}
	}

	log.Debug("[NewRequest]::Return response")

	return clientResp, nil
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
