package edgegrid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// NetworkListServicev2 represents exposed services to manage network lists
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type NetworkListServicev2 struct {
	client *Client
}

// NetworkListsv2 represents array of network lists
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type NetworkListsv2 struct {
	NetworkLists []NetworkListv2 `json:"networkLists"`
	Links        struct {
		Create NetworkListLinkv2 `json:"create"`
	} `json:"links"`
}

// NetworkListv2 represents the network list structure
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type NetworkListv2 struct {
	NetworkListType    string `json:"networkListType,omitempty`
	AccessControlGroup string `json:"accessControlGroup,omitempty"`
	Name               string `json:"name,omitempty"`
	ElementCount       int    `json:"elementCount,omitempty"`
	Links              struct {
		ActivateInProduction NetworkListLinkv2 `json:"activateInProduction,omitempty"`
		ActivateInStaging    NetworkListLinkv2 `json:"activateInStaging,omitempty"`
		AppendItems          NetworkListLinkv2 `json:"appendItems,omitempty"`
		Retrieve             NetworkListLinkv2 `json:"retrieve,omitempty"`
		StatusInProduction   NetworkListLinkv2 `json:"statusInProduction,omitempty"`
		StatusInStaging      NetworkListLinkv2 `json:"statusInStaging,omitempty"`
		Update               NetworkListLinkv2 `json:"update,omitempty"`
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
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type NetworkListLinkv2 struct {
	Href   string `json:"href"`
	Method string `json:"method"`
}

// NetworkListErrorv2 represents the error returned from Akamai
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#errors
type NetworkListErrorv2 struct {
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
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type NetworkListsOptionsv2 struct {
	Name        string   `json:"name,omitempty"`
	Type        string   `json:"type,omitempty"`
	Description string   `json:"description,omitempty"`
	List        []string `json:"list,omitempty"`
}

// ListNetworkListsOptionsv2 represents the available options for listing network lists
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type ListNetworkListsOptionsv2 struct {
	TypeOflist      string
	Extended        bool
	IncludeElements bool
	Search          string
}

// NetworkListActivationOptsv2 represents object used for activating network list in Akamai
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type NetworkListActivationOptsv2 struct {
	Comments               string   `json:"comments"`
	NotificationRecipients []string `json:"notificationRecipients"`
	Fast                   bool     `json:"fast"`
}

// NetworkListActivationStatusv2 represents object used for status of network list activation
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
type NetworkListActivationStatusv2 struct {
	ActivationID       int    `json:"activationId"`
	ActivationComments string `json:"activationComments"`
	ActivationStatus   string `json:"activationStatus"`
	SyncPoint          int    `json:"syncPoint"`
	UniqueID           string `json:"uniqueId"`
	Fast               bool   `json:"fast"`
}

// NetworkListErrorv2 Error() function implementation
func (e *NetworkListErrorv2) Error() string {
	return ShowJSONMessage(e)
}

// NetworkListDeleteResponse represents response from deleting a list
type NetworkListDeleteResponse struct {
	Status    int    `json:"status"`
	UniqueID  string `json:"uniqueId"`
	SyncPoint int    `json:"syncPoint"`
}

// NetworkListSubscription represents object used for (un)subscribing for notifications
type NetworkListSubscription struct {
	Recipients []string `json:"recipients"`
	UniqueIds  []string `json:"uniqueIds"`
}

// QStrNetworkList includes query params used across network lists
type QStrNetworkList struct {
	IncludeElements bool   `url:"includeElements,omitempty"`
	Extended        bool   `url:"extended,omitempty"`
	Search          string `url:"search,omitempty"`
	Element         string `url:"element,omitempty"`
}

// ListNetworkLists List all configured Network Lists for the authenticated user.
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlists
func (nls *NetworkListServicev2) ListNetworkLists(opts ListNetworkListsOptionsv2) (*[]NetworkListv2, *ClientResponse, error) {

	qParams := QStrNetworkList{
		Extended:        opts.Extended,
		IncludeElements: opts.IncludeElements,
		Search:          opts.Search,
	}

	path := NetworkListPathV2

	var respStruct *NetworkListsv2

	log.Debug("[NetworkListServicev2]::Execute request")
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	// This error indicates we had problems connecting to Akamai endpoint(s)
	if err != nil {
		log.Debug("[NetworkListServicev2]::Client request error")
		log.Debug(fmt.Sprintf("[NetworkListServicev2]:: %s", err))

		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return &respStruct.NetworkLists, resp, netListError
	}

	return &respStruct.NetworkLists, resp, nil

}

// CreateNetworkList Create a new network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *NetworkListServicev2) CreateNetworkList(opts NetworkListsOptionsv2) (*NetworkListv2, *ClientResponse, error) {

	qParams := QStrNetworkList{}
	path := NetworkListPathV2

	var respStruct *NetworkListv2
	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, opts, nil)
	if err != nil {
		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return respStruct, resp, netListError
	}

	return respStruct, resp, err
}

// GetNetworkList Gets a specific network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlist
func (nls *NetworkListServicev2) GetNetworkList(ListID string, opts ListNetworkListsOptionsv2) (*NetworkListv2, *ClientResponse, error) {

	qParams := QStrNetworkList{
		Extended:        opts.Extended,
		IncludeElements: opts.IncludeElements,
	}

	path := fmt.Sprintf("%s/%s", NetworkListPathV2, ListID)

	var respStruct *NetworkListv2
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return respStruct, resp, netListError
	}

	return respStruct, resp, err

}

// AppendListNetworkList Adds items to network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *NetworkListServicev2) AppendListNetworkList(ListID string, opts NetworkListsOptionsv2) (*NetworkListv2, *ClientResponse, error) {

	qParams := QStrNetworkList{}
	path := fmt.Sprintf("%s/%s/append", NetworkListPathV2, ListID)

	var respStruct *NetworkListv2
	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, opts, nil)
	if err != nil {
		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return respStruct, resp, netListError
	}

	return respStruct, resp, err
}

// RemoveNetworkListElement Removes network list element
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) RemoveNetworkListElement(ListID, element string) (*NetworkListv2, *ClientResponse, error) {

	qParams := QStrNetworkList{
		Element: element,
	}
	path := fmt.Sprintf("%s/%s/elements", NetworkListPathV2, ListID)

	var respStruct *NetworkListv2
	resp, err := nls.client.makeAPIRequest(http.MethodDelete, path, qParams, &respStruct, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return respStruct, resp, netListError
	}

	return respStruct, resp, err
}

// ActivateNetworkList Activates network list on specified network ( PRODUCTION or STAGING )
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) ActivateNetworkList(ListID string, targetEnv AkamaiEnvironment, opts NetworkListActivationOptsv2) (*NetworkListActivationStatusv2, *ClientResponse, error) {

	qParams := QStrNetworkList{}
	path := fmt.Sprintf("%s/%s/environments/%s/activate", NetworkListPathV2, ListID, targetEnv)

	var respStruct *NetworkListActivationStatusv2
	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, nil, nil)
	if err != nil {
		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return respStruct, resp, netListError
	}

	return respStruct, resp, err
}

// GetNetworkListActStatus Gets activation network list status on specified network ( PRODUCTION or STAGING )
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) GetNetworkListActStatus(ListID string, targetEnv AkamaiEnvironment) (*NetworkListActivationStatusv2, *ClientResponse, error) {

	qParams := QStrNetworkList{}
	path := fmt.Sprintf("%s/%s/environments/%s/status", NetworkListPathV2, ListID, targetEnv)

	var respStruct *NetworkListActivationStatusv2
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return respStruct, resp, netListError
	}

	return respStruct, resp, err
}

// DeleteNetworkList Remove network list element
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) DeleteNetworkList(ListID string) (*NetworkListDeleteResponse, *ClientResponse, error) {

	qParams := QStrNetworkList{}
	path := fmt.Sprintf("%s/%s", NetworkListPathV2, ListID)

	var respStruct *NetworkListDeleteResponse
	resp, err := nls.client.makeAPIRequest(http.MethodDelete, path, qParams, &respStruct, nil, nil)

	if err != nil {
		return nil, resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return respStruct, resp, netListError
	}

	return respStruct, resp, err
}

// NetworkListNotification Manage network list subscription
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *NetworkListServicev2) NetworkListNotification(action AkamaiSubscription, sub NetworkListSubscription) (*ClientResponse, error) {

	qParams := QStrNetworkList{}
	path := fmt.Sprintf("/network-list/v2/notifications/%s", action)

	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, nil, sub, nil)
	if err != nil {
		return resp, err
	}

	if resp.Response.StatusCode >= http.StatusBadRequest {
		netListError := &NetworkListErrorv2{}
		err := json.Unmarshal([]byte(resp.Body), &netListError)
		if err != nil {
			log.Fatal(err)
		}
		return resp, err
	}

	return resp, err
}
