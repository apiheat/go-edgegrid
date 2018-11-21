package edgegrid

import (
	"encoding/json"
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
	List                                []string  `json:"list"`
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

// NetworkListsOptionsv2 represents struct required to create items for network list
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type NetworkListsOptionsv2 struct {
	Name        string   `json:"name,omitempty"`
	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	List        []string `json:"list,omitempty"`
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

// NetworkListActivationOptsv2 represents object used for activating network list in Akamai
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type NetworkListActivationOptsv2 struct {
	Comments               string   `json:"comments"`
	NotificationRecipients []string `json:"notificationRecipients"`
	Fast                   bool     `json:"fast"`
}

// NetworkListActivationStatusv2 represents object used for status of network list activation
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type NetworkListActivationStatusv2 struct {
	ActivationID       int    `json:"activationId"`
	ActivationComments string `json:"activationComments"`
	ActivationStatus   string `json:"activationStatus"`
	SyncPoint          int    `json:"syncPoint"`
	UniqueID           string `json:"uniqueId"`
	Fast               bool   `json:"fast"`
}

// An AkamaiNetworkListErrorv2 Error() function implementation
//
// error
func (e *AkamaiNetworkListErrorv2) Error() string {
	return ShowJSONMessage(e)
}

// An NetworkListDeleteResponse represents response from deleting a list
//
// error
type NetworkListDeleteResponse struct {
	Status    int    `json:"status"`
	UniqueID  string `json:"uniqueId"`
	SyncPoint int    `json:"syncPoint"`
}

// An NetworkListSubscription represents object used for (un)subscribing for notifications
//
// error
type NetworkListSubscription struct {
	Recipients []string `json:"recipients"`
	UniqueIds  []string `json:"uniqueIds"`
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

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return &netListsv2.NetworkLists, APIClientResponse, netListError
	}

	return &netListsv2.NetworkLists, APIClientResponse, nil

}

// CreateNetworkList Create a new network list
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *NetworkListServicev2) CreateNetworkList(opts NetworkListsOptionsv2) (*AkamaiNetworkListv2, *ClientResponse, error) {

	apiURI := NetworkListPathV2

	var k *AkamaiNetworkListv2
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodPost, apiURI, opts, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return k, APIClientResponse, netListError
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

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return k, APIClientResponse, netListError
	}

	return k, APIClientResponse, APIclientError

}

// AppendListNetworkList Appends list of items
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *NetworkListServicev2) AppendListNetworkList(ListID string, opts NetworkListsOptionsv2) (*AkamaiNetworkListv2, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s/append",
		NetworkListPathV2,
		ListID)

	var k *AkamaiNetworkListv2
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodPost, apiURI, opts, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return k, APIClientResponse, netListError
	}

	return k, APIClientResponse, APIclientError
}

// RemoveNetworkListElement Remove network list element
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) RemoveNetworkListElement(ListID, element string) (*AkamaiNetworkListv2, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s/elements?element=%s",
		NetworkListPathV2,
		ListID,
		element)

	var k *AkamaiNetworkListv2
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodDelete, apiURI, nil, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return k, APIClientResponse, netListError
	}

	return k, APIClientResponse, APIclientError
}

// ActivateNetworkList Activates network list on specified network ( PRODUCTION or STAGING )
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) ActivateNetworkList(ListID string, targetEnv AkamaiEnvironment, opts NetworkListActivationOptsv2) (*NetworkListActivationStatusv2, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s/environments/%s/activate",
		NetworkListPathV2,
		ListID,
		targetEnv)

	var k *NetworkListActivationStatusv2
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodPost, apiURI, opts, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return k, APIClientResponse, netListError
	}

	return k, APIClientResponse, APIclientError
}

// GetNetworkListActStatus Gets activation network list status on specified network ( PRODUCTION or STAGING )
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) GetNetworkListActStatus(ListID string, targetEnv AkamaiEnvironment) (*NetworkListActivationStatusv2, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s/environments/%s/status",
		NetworkListPathV2,
		ListID,
		targetEnv)

	var k *NetworkListActivationStatusv2
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return k, APIClientResponse, netListError
	}

	return k, APIClientResponse, APIclientError
}

// DeleteNetworkList Remove network list element
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) DeleteNetworkList(ListID string) (*NetworkListDeleteResponse, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/%s",
		NetworkListPathV2,
		ListID)

	var k *NetworkListDeleteResponse
	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodDelete, apiURI, nil, &k)
	if APIclientError != nil {
		return nil, APIClientResponse, APIclientError
	}

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return k, APIClientResponse, netListError
	}

	return k, APIClientResponse, APIclientError
}

// NetworkListNotification Manage network list subscription
//
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) NetworkListNotification(action AkamaiSubscription, sub NetworkListSubscription) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("/network-list/v2/notifications/%s", action)

	APIClientResponse, APIclientError := nls.client.NewRequest(http.MethodPost, apiURI, sub, nil)
	if APIclientError != nil {
		return APIClientResponse, APIclientError
	}

	if APIClientResponse.Response.StatusCode >= http.StatusBadRequest {
		netListError := &AkamaiNetworkListErrorv2{}
		err := json.Unmarshal([]byte(APIClientResponse.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return APIClientResponse, netListError
	}

	return APIClientResponse, APIclientError
}
