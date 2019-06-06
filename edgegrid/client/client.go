package client

import (
	"fmt"
	"strconv"
	"time"

	"github.com/apiheat/go-edgegrid/edgegrid"
	"github.com/apiheat/go-edgegrid/edgegrid/edgeauth"

	"gopkg.in/resty.v1"
)

// A Client implements the base client request and response handling
// used by all service clients.
type Client struct {
	Config      *edgegrid.Config
	Credentials *edgeauth.Credentials
	REST        *resty.Client
}

// New will return a pointer to a new initialized service client.
func New(cfg *edgegrid.Config, creds *edgeauth.Credentials, options ...func(*Client)) *Client {
	svc := &Client{
		Config:      cfg,
		Credentials: creds,
		REST:        resty.New(),
	}

	svc.REST.SetDebug(true)

	// You can override all below settings and options at request level if you want to
	//--------------------------------------------------------------------------------
	// Host URL for all request. So you can use relative URL in the request
	// TODO: Scheme option / Endpoint option
	svc.REST.SetHostURL(fmt.Sprintf("https://%s", svc.Credentials.Host))
	svc.REST.SetScheme("https")

	// // OnBeforeRequest registers function that will sign our request for Akamai API
	svc.REST.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		var credentialsAccessToken = svc.Credentials.AccessToken
		var credentialsClientSecret = svc.Credentials.ClientSecret
		var credentialsClientToken = svc.Credentials.ClientToken

		// fmt.Println()
		// fmt.Println(req.URL)
		// fmt.Println(req.Header)
		// fmt.Println()
		// fmt.Println(req.RawRequest)

		uno := edgeauth.GenerateEdgeGridAuthString(credentialsClientToken, credentialsClientSecret, credentialsAccessToken, req)
		dataToSign := edgeauth.MakeDataToSign2(req.Method, "https", c.HostURL, "/network-list/v2/network-lists?extended=true&includeElements=true&search=", uno)
		signingKey := edgeauth.Base64HmacSha256(credentialsClientSecret)
		signed := edgeauth.SignRequest2(dataToSign, signingKey)

		fmt.Println(uno)
		fmt.Println(dataToSign)
		fmt.Println(signingKey)
		fmt.Println(signed)

		finalHeader := edgeauth.GenerateEdgeGridAuthString2(uno, signed)

		fmt.Println(finalHeader)
		req.SetHeader("Authorization", finalHeader)

		return nil // if its success otherwise return error
	})
	return svc
}

// NewRequest returns a new Request pointer for the service API
// operation and parameters.
func (c *Client) NewRequest(params interface{}, data interface{}) string {
	req := c.REST.NewRequest()
	req.Method = "GET"
	req.SetQueryParams(map[string]string{
		"page_no": "1",
		"limit":   "20",
		"sort":    "name",
		"order":   "asc",
		"random":  strconv.FormatInt(time.Now().Unix(), 10),
	}).SetHeader("Accept", "application/json").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		Get("/search_result")

	// // POST Struct, default is JSON content type. No need to set one
	// resp, err := resty.R().
	// 	SetBody(User{Username: "testuser", Password: "testpass"}).
	// 	SetResult(&AuthSuccess{}). // or SetResult(AuthSuccess{}).
	// 	SetError(&AuthError{}).    // or SetError(AuthError{}).
	// 	Post("https://myapp.com/login")

	// resty.R().SetPathParams(map[string]string{
	// 	"userId": "sample@sample.com",
	// 	"subAccountId": "100002",
	//  }).
	//  Get("/v1/users/{userId}/{subAccountId}/details")

	return ""
}
