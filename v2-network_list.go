package edgegrid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
)

// NetworkListServicev2 represents exposed services to manage network lists
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type NetworkListServicev2 struct {
	client *Client
}

// AkamaiNetworkListsv2 represents array of network lists
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type AkamaiNetworkListsv2 struct {
	NetworkLists []AkamaiNetworkListv2 `json:"networkLists"`
	Links        struct {
		Create AkamaiNetworkListLinkv2 `json:"create"`
	} `json:"links"`
}

// AkamaiNetworkListv2 represents the network list structure
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type AkamaiNetworkListv2 struct {
	NetworkListType    string `json:"networkListType"`
	AccessControlGroup string `json:"accessControlGroup"`
	Name               string `json:"name"`
	ElementCount       int    `json:"elementCount"`
	Links              struct {
		ActivateInProduction AkamaiNetworkListLinkv2 `json:"activateInProduction"`
		ActivateInStaging    AkamaiNetworkListLinkv2 `json:"activateInStaging"`
		AppendItems          AkamaiNetworkListLinkv2 `json:"appendItems"`
		Retrieve             AkamaiNetworkListLinkv2 `json:"retrieve"`
		StatusInProduction   AkamaiNetworkListLinkv2 `json:"statusInProduction"`
		StatusInStaging      AkamaiNetworkListLinkv2 `json:"statusInStaging"`
		Update               AkamaiNetworkListLinkv2 `json:"update"`
	} `json:"links"`
	SyncPoint                           int       `json:"syncPoint"`
	Type                                string    `json:"type"`
	UniqueID                            string    `json:"uniqueId"`
	CreateDate                          time.Time `json:"createDate"`
	CreatedBy                           string    `json:"createdBy"`
	ExpeditedProductionActivationStatus string    `json:"expeditedProductionActivationStatus"`
	ExpeditedStagingActivationStatus    string    `json:"expeditedStagingActivationStatus"`
	ProductionActivationStatus          string    `json:"productionActivationStatus"`
	StagingActivationStatus             string    `json:"stagingActivationStatus"`
	UpdateDate                          time.Time `json:"updateDate"`
	UpdatedBy                           string    `json:"updatedBy"`
}

// AkamaiNetworkListLinks represents the network list `links` structure
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type AkamaiNetworkListLinkv2 struct {
	Href   string `json:"href"`
	Method string `json:"method"`
}

// ListNetworkListsOptions represents the available options for listing network lists
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type ListNetworkListsOptionsv2 struct {
	TypeOflist      string
	Extended        bool
	IncludeElements bool
	Search          string
}

// AkamaiNetworkListErrorv2 represents the error returned from Akamai
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#errors
type AkamaiNetworkListErrorv2 struct {
	Detail      string `json:"detail"`
	Instance    string `json:"instance"`
	Status      int    `json:"status"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	FieldErrors struct {
		Entry []struct {
			Key   string   `json:"key"`
			Value []string `json:"value"`
		} `json:"entry"`
	} `json:"fieldErrors"`
}

// An AkamaiNetworkListErrorv2 Error() function implementation
//
// error
func (e *AkamaiNetworkListErrorv2) Error() string {
	b, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}

	var prettyJSON bytes.Buffer
	errprettyJSON := json.Indent(&prettyJSON, []byte(string(b)), "", "    ")
	if errprettyJSON != nil {
		fmt.Println(errprettyJSON)
	}
	return string(prettyJSON.Bytes())
}

// ListNetworkLists List all configured Network Lists for the authenticated user.
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlists
func (nls *NetworkListServicev2) ListNetworkLists(opts ListNetworkListsOptionsv2) (*[]AkamaiNetworkListv2, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s?listType=%s&extended=%t&search=%s&includeElements=%t",
		NetworkListPathV2,
		opts.TypeOflist,
		opts.Extended,
		opts.Search,
		opts.IncludeElements)

	var netListsv2 *AkamaiNetworkListsv2

	log.Debug("[NetworkListServicev2]::Execute request")
	clientResp, clientErr := nls.client.NewRequest("GET", apiURI, nil, &netListsv2)

	if clientErr != nil {
		log.Debug("[NetworkListServicev2]::Client request error")
		log.Debug(fmt.Sprintf("[NetworkListServicev2]:: %s", clientErr))

		return nil, clientResp, clientErr
	}

	/*
		This is MVP for next iteration to be placed in function
	*/
	data, _ := ioutil.ReadAll(clientResp.Response.Body)
	switch clientResp.Response.StatusCode {
	case 200, 201:
		return &netListsv2.NetworkLists, clientResp, nil
	case 400, 403, 404, 409, 500:
		var netListsv2err *AkamaiNetworkListErrorv2
		json.Unmarshal(data, &netListsv2err)
		return nil, clientResp, netListsv2err
	}

	var errorResponse *ErrorResponse
	_ = json.Unmarshal(data, &errorResponse)
	return nil, clientResp, errorResponse

}
