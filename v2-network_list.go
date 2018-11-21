package edgegrid

import (
	"fmt"
	"net/http"
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
	NetworkListType    string `json:"networkListType,omitempty`
	AccessControlGroup string `json:"accessControlGroup,omitempty"`
	Name               string `json:"name,omitempty"`
	ElementCount       int    `json:"elementCount,omitempty"`
	Links              struct {
		ActivateInProduction AkamaiNetworkListLinkv2 `json:"activateInProduction,omitempty"`
		ActivateInStaging    AkamaiNetworkListLinkv2 `json:"activateInStaging,omitempty"`
		AppendItems          AkamaiNetworkListLinkv2 `json:"appendItems,omitempty"`
		Retrieve             AkamaiNetworkListLinkv2 `json:"retrieve,omitempty"`
		StatusInProduction   AkamaiNetworkListLinkv2 `json:"statusInProduction,omitempty"`
		StatusInStaging      AkamaiNetworkListLinkv2 `json:"statusInStaging,omitempty"`
		Update               AkamaiNetworkListLinkv2 `json:"update,omitempty"`
	} `json:"links"`
	SyncPoint                           int       `json:"syncPoint,omitempty"`
	Type                                string    `json:"type,omitempty"`
	UniqueID                            string    `json:"uniqueId,omitempty"`
	CreateDate                          time.Time `json:"createDate,omitempty"`
	CreatedBy                           string    `json:"createdBy,omitempty"`
	ExpeditedProductionActivationStatus string    `json:"expeditedProductionActivationStatus,omitempty"`
	ExpeditedStagingActivationStatus    string    `json:"expeditedStagingActivationStatus,omitempty"`
	ProductionActivationStatus          string    `json:"productionActivationStatus,omitempty"`
	StagingActivationStatus             string    `json:"stagingActivationStatus,omitempty"`
	UpdateDate                          time.Time `json:"updateDate,omitempty"`
	UpdatedBy                           string    `json:"updatedBy,omitempty"`
}

// AkamaiNetworkListLinks represents the network list `links` structure
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type AkamaiNetworkListLinkv2 struct {
	Href   string `json:"href"`
	Method string `json:"method"`
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

// CreateNetworkListsOptionsv2 represents struct required to create new network list in Akamai
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type CreateNetworkListsOptionsv2 struct {
	Name        string   `json:"name,omitempty"`
	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	List        []string `json:"list"`
}

// ListNetworkListsOptionsv2 represents the available options for listing network lists
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type ListNetworkListsOptionsv2 struct {
	TypeOflist      string
	Extended        bool
	IncludeElements bool
	Search          string
}

// An AkamaiNetworkListErrorv2 Error() function implementation
//
// error
func (e *AkamaiNetworkListErrorv2) Error() string {
	return ShowJSONMessage(e)
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
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodGet, apiURI, nil, &netListsv2)

	// This error indicates we had problems connecting to Akamai endpoint(s)
	if APIclientError != nil {
		log.Debug("[NetworkListServicev2]::Client request error")
		log.Debug(fmt.Sprintf("[NetworkListServicev2]:: %s", APIclientError))

		return nil, APIClientResponse, APIclientError
	}

	return &netListsv2.NetworkLists, APIClientResponse, nil

}

// CreateNetworkList Create a new network list
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *NetworkListServicev2) CreateNetworkList(opts CreateNetworkListsOptionsv2) (*AkamaiNetworkListv2, *ClientResponse, error) {

	apiURI := NetworkListPathV2

	var k *AkamaiNetworkListv2
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodPost, apiURI, opts, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	return k, APIClientResponse, APIclientError
}

// GetNetworkList Create a new network list
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlist
func (nls *NetworkListServicev2) GetNetworkList(ListID string, opts ListNetworkListsOptionsv2) (*AkamaiNetworkListv2, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s?extended=%t&includeElements=%t",
		NetworkListPathV2,
		ListID,
		opts.Extended,
		opts.IncludeElements)

	var k *AkamaiNetworkListv2
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	return k, APIClientResponse, APIclientError

}
