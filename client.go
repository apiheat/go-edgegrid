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
	libraryVersion = "0.0.1"
	userAgent      = "go-edgegrid/" + libraryVersion
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
	NetworkLists *NetworkListService
}

// ClientOptions represents options we can pass during client creation
//
// client
type ClientOptions struct {
	ConfigPath    string
	ConfigSection string
}

var (
	apiPaths = map[string]string{
		"network_list": "/network-list/v1/network_lists",
	}
)

// NewClient returns a new edgegrid.Client for API. If a nil httpClient is
// provided, http.DefaultClient will be used.
//
// client
func NewClient(httpClient *http.Client, conf *ClientOptions) *Client {
	var (
		path, section string
	)

	// If we do not pass config we will try to to use env variables
	if conf != nil {
		path = conf.ConfigPath
		section = conf.ConfigSection
	} else {
		path = os.Getenv("AKAMAI_EDGERC_PATH")
		section = os.Getenv("AKAMAI_EDGERC_SECTION")
	}

	return newClient(httpClient, path, section)
}

// newClient *private* function to initiaite client
//
// client
func newClient(httpClient *http.Client, edgercPath, edgercSection string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient}
	c.credentials, _ = InitEdgerc(edgercPath, edgercSection)

	// Set base URL for making all API requests
	c.SetBaseURL(c.credentials.host)

	// Create all the public services.
	c.Auth = &AuthService{client: c}
	c.NetworkLists = &NetworkListService{client: c}

	return c
}

// newRequest creates an HTTP request that can be sent to Akamai APIs. A relative URL can be provided in path, which will be resolved to the
// Host specified in Config. If body is specified, it will be sent as the request body.
//
// client
func (cl *Client) NewRequest(method, path string, v interface{}) (*http.Response, error) {

	targetURL, _ := prepareURL(cl.baseURL, path)

	fmt.Println(fmt.Sprintf("client.NewRequest() => Target URL: %s ", targetURL))

	req, err := http.NewRequest(method, targetURL.String(), nil)
	if err != nil {
		return nil, nil
	}

	if method == "POST" || method == "PUT" {
		bodyBytes, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		bodyReader := bytes.NewReader(bodyBytes)

		req.Body = ioutil.NopCloser(bodyReader)
		req.ContentLength = int64(bodyReader.Len())

		req.Header.Set("Content-Type", "application/json")

	}

	authorizationHeader := AuthString(cl.credentials, req, []string{})
	req.Header.Add("Authorization", authorizationHeader)

	resp, err := cl.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	byt, _ := ioutil.ReadAll(resp.Body)

	if err = json.Unmarshal([]byte(byt), &v); err != nil {
		return nil, err
	}

	return resp, err
}

// SetBaseURL sets the base URL for API requests to a custom endpoint.
//
// client
func (cl *Client) SetBaseURL(urlStr string) error {

	if strings.HasPrefix(urlStr, "https://") {
		cl.baseURL, _ = url.Parse(urlStr)
	} else {
		cl.baseURL, _ = url.Parse("https://" + urlStr)
	}
	// // Make sure the given URL end with a slash
	// if !strings.HasSuffix(urlStr, "/") {
	// 	urlStr += "/"
	// }

	var err error
	cl.baseURL, err = url.Parse(urlStr)
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
