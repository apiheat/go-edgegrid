package edgegrid

import (
	"fmt"
	"time"
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
	NetworkLists []AkamaiNetworkListv2      `json:"networkLists"`
	Links        []AkamaiNetworkListLinksv2 `json:"links"`
	Create       AkamaiNetworkListLinks     `json:"create"`
}

// AkamaiNetworkListv2 represents the network list structure
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type AkamaiNetworkListv2 struct {
	NetworkListType                     string                   `json:"networkListType"`
	AccessControlGroup                  string                   `json:"accessControlGroup"`
	Name                                string                   `json:"name"`
	ElementCount                        int                      `json:"elementCount"`
	SyncPoint                           int                      `json:"syncPoint"`
	Type                                string                   `json:"type"`
	UniqueID                            string                   `json:"uniqueId"`
	CreateDate                          time.Time                `json:"createDate"`
	CreatedBy                           string                   `json:"createdBy"`
	ExpeditedProductionActivationStatus string                   `json:"expeditedProductionActivationStatus"`
	ExpeditedStagingActivationStatus    string                   `json:"expeditedStagingActivationStatus"`
	ProductionActivationStatus          string                   `json:"productionActivationStatus"`
	StagingActivationStatus             string                   `json:"stagingActivationStatus"`
	UpdateDate                          time.Time                `json:"updateDate"`
	UpdatedBy                           string                   `json:"updatedBy"`
	Links                               []AkamaiNetworkListLinks `json:"links"`
}

// AkamaiNetworkListLinks represents the network list `links` structure
//
// Akamai API docs: https://developer.akamai.com/api/luna/network-list
type AkamaiNetworkListLinksv2 struct {
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

// // CreateNetworkListOptions represents the available options for network list creation
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type CreateNetworkListOptions struct {
// 	Name        string   `json:"name,omitempty"`
// 	Type        string   `json:"type,omitempty"`
// 	AcgID       string   `json:"acgId,omitempty"`
// 	ContractID  string   `json:"contractId,omitempty"`
// 	GroupID     int64    `json:"groupId,omitempty"`
// 	Description string   `json:"description,omitempty"`
// 	List        []string `json:"list"`
// }

// // NetworkListResponse represents the response from network list creation
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type NetworkListResponse struct {
// 	Status           int                      `json:"status,omitempty"`
// 	UniqueID         string                   `json:"unique-id,omitempty"`
// 	Message          string                   `json:"message,omitempty"`
// 	Links            []AkamaiNetworkListLinks `json:"links"`
// 	SyncPoint        int                      `json:"sync-point,omitempty"`
// 	ActivationStatus string                   `json:"activation-status,omitempty"`
// }

// // ActivateNetworkListOptions represents options for network list activation
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type ActivateNetworkListOptions struct {
// 	SiebelTicketID         string   `json:"siebel-ticket-id"`
// 	NotificationRecipients []string `json:"notification-recipients"`
// 	Comments               string   `json:"comments"`
// }

// // ActivateNetworkListStatus represents status of network list activation
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type ActivateNetworkListStatus struct {
// 	Status             int           `json:"status"`
// 	UniqueID           string        `json:"unique-id"`
// 	Links              []interface{} `json:"links"`
// 	SyncPoint          int           `json:"sync-point"`
// 	ActivationStatus   string        `json:"activation-status"`
// 	ActivationComments string        `json:"activation-comments"`
// }

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
	var netListv2Err *AkamaiNetworkListErrorv2
	clientResp, clientErr := nls.client.NewRequestWithCustomError("GET", apiURI, nil, &netListsv2, &netListv2Err)
	if clientErr != nil {
		return nil, clientResp, clientErr
	}

	return &netListsv2.NetworkLists, clientResp, clientErr

}

// // GetNetworkList Gets single Akamai network list
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#getanetworklist
// func (nls *NetworkListServicev2) GetNetworkList(ListID string, opts ListNetworkListsOptions) (*AkamaiNetworkList, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/%s?listType=%s&extended=%t&includeDeprecated=%t&includeElements=%t",
// 		NetworkListPathV1,
// 		ListID,
// 		opts.TypeOflist,
// 		opts.Extended,
// 		opts.IncludeDeprecated,
// 		opts.IncludeElements)

// 	var k *AkamaiNetworkList
// 	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return k, resp, err

// }

// // CreateNetworkList Create a new Network List
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#createanetworklist
// func (nls *NetworkListServicev2) CreateNetworkList(opts CreateNetworkListOptions) (*NetworkListResponse, *ClientResponse, error) {

// 	apiURI := NetworkListPathV1

// 	var k *NetworkListResponse
// 	resp, err := nls.client.NewRequest("POST", apiURI, opts, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return k, resp, err
// }

// // ModifyNetworkList Modify an existing Network List
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#modifyanetworklist
// func (nls *NetworkListServicev2) ModifyNetworkList(ListID string, opts AkamaiNetworkList) (*NetworkListResponse, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/%s",
// 		NetworkListPathV1,
// 		ListID)

// 	var k *NetworkListResponse
// 	resp, err := nls.client.NewRequest("PUT", apiURI, opts, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return k, resp, err
// }

// // AddNetworkListItems Appends a set of IP addresses or geo locations to a list.
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#addtoanetworklist
// func (nls *NetworkListServicev2) AddNetworkListItems(ListID string, opts CreateNetworkListOptions) (*NetworkListResponse, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/%s",
// 		NetworkListPathV1,
// 		ListID)

// 	var k *NetworkListResponse
// 	resp, err := nls.client.NewRequest("POST", apiURI, opts, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return k, resp, err

// }

// // AddNetworkListElement Adds the specified element to the list.
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#addanelement
// func (nls *NetworkListServicev2) AddNetworkListElement(ListID, ListElement string) (*NetworkListResponse, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/%s/element?element=%s",
// 		NetworkListPathV1,
// 		ListID,
// 		ListElement,
// 	)

// 	var k *NetworkListResponse
// 	resp, err := nls.client.NewRequest("PUT", apiURI, nil, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	// networkListResponse := new(NetworkListResponse)
// 	// byt, _ := ioutil.ReadAll(resp.Response.Body)

// 	// if err = json.Unmarshal([]byte(byt), &networkListResponse); err != nil {
// 	// 	return nil, resp, err
// 	// }

// 	return k, resp, err

// }

// // RemoveNetworkListItem Deletes the specified element from the list.
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#removeanelement
// func (nls *NetworkListServicev2) RemoveNetworkListItem(ListID, ListItem string) (*NetworkListResponse, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/%s/element?element=%s",
// 		NetworkListPathV1,
// 		ListID,
// 		ListItem)

// 	var k *NetworkListResponse
// 	resp, err := nls.client.NewRequest("DELETE", apiURI, nil, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return k, resp, err

// }

// // SearchNetworkListItem Retrieves a list of all Network Lists having elements containing the search terms. Each Network List’s status is also provided.
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#searchnetworklists
// func (nls *NetworkListServicev2) SearchNetworkListItem(ListItem string, opts ListNetworkListsOptions) (*[]AkamaiNetworkList, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/search?expression=%s&listType=%s&extended=%t&includeDeprecated=%t",
// 		NetworkListPathV1,
// 		ListItem,
// 		opts.TypeOflist,
// 		opts.Extended,
// 		opts.IncludeDeprecated)

// 	var k *AkamaiNetworkLists
// 	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return &k.NetworkLists, resp, err

// }

// // ActivateNetworkList Activates selected network list in specific env with options specified
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#activateanetworklist
// func (nls *NetworkListServicev2) ActivateNetworkList(ListID string, targetEnvironment AkamaiEnvironment, opts ActivateNetworkListOptions) (*NetworkListResponse, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/%s/activate?env=%s",
// 		NetworkListPathV1,
// 		ListID,
// 		targetEnvironment)

// 	var k *NetworkListResponse
// 	resp, err := nls.client.NewRequest("POST", apiURI, opts, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return k, resp, err
// }

// // GetNetworkListActivationStatus Gets activation status of selected network list in specific env
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#getactivationstatus
// func (nls *NetworkListServicev2) GetNetworkListActivationStatus(ListID string, targetEnvironment AkamaiEnvironment) (*ActivateNetworkListStatus, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s/%s/status?env=%s",
// 		NetworkListPathV1,
// 		ListID,
// 		targetEnvironment)

// 	var k *ActivateNetworkListStatus
// 	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return k, resp, err
// }
