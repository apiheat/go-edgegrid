// package edgegrid

// // NetworkListService represents exposed services to manage network lists
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type NetworkListService struct {
// 	client *Client
// }

// // AkamaiNetworkLists represents array of network lists
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type AkamaiNetworkLists struct {
// 	NetworkLists []AkamaiNetworkList `json:"network_lists"`
// }

// // AkamaiNetworkList represents the network list structure
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type AkamaiNetworkList struct {
// 	Description                string                   `json:"description,omitempty"`
// 	CreateEpoch                int                      `json:"createEpoch"`
// 	UpdateEpoch                int                      `json:"updateEpoch"`
// 	CreateDate                 int64                    `json:"createDate"`
// 	CreatedBy                  string                   `json:"createdBy"`
// 	UpdatedBy                  string                   `json:"updatedBy"`
// 	UpdateDate                 int64                    `json:"updateDate"`
// 	StagingActivationStatus    string                   `json:"stagingActivationStatus"`
// 	ProductionActivationStatus string                   `json:"productionActivationStatus"`
// 	Name                       string                   `json:"name"`
// 	Type                       string                   `json:"type"`
// 	UniqueID                   string                   `json:"unique-id"`
// 	Account                    string                   `json:"account"`
// 	ReadOnly                   bool                     `json:"readOnly"`
// 	SyncPoint                  int                      `json:"sync-point"`
// 	Links                      []AkamaiNetworkListLinks `json:"links"`
// 	List                       []string                 `json:"list"`
// 	NumEntries                 int                      `json:"numEntries"`
// }

// // AkamaiNetworkListLinks represents the network list `links` structure
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type AkamaiNetworkListLinks struct {
// 	Rel  string `json:"rel"`
// 	Href string `json:"href"`
// }

// // ListNetworkListsOptions represents the available options for listing network lists
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list
// type ListNetworkListsOptions struct {
// 	TypeOflist        string
// 	Extended          bool
// 	IncludeDeprecated bool
// 	IncludeElements   bool
// }

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

// // ListNetworkLists List all configured Network Lists for the authenticated user.
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#listnetworklists
// func (nls *NetworkListService) ListNetworkLists(opts ListNetworkListsOptions) (*[]AkamaiNetworkList, *ClientResponse, error) {

// 	apiURI := fmt.Sprintf("%s?listType=%s&extended=%t&includeDeprecated=%t&includeElements=%t",
// 		NetworkListPathV1,
// 		opts.TypeOflist,
// 		opts.Extended,
// 		opts.IncludeDeprecated,
// 		opts.IncludeElements)

// 	var k *AkamaiNetworkLists
// 	resp, err := nls.client.NewRequest("GET", apiURI, nil, &k)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return &k.NetworkLists, resp, err

// }

// // GetNetworkList Gets single Akamai network list
// //
// // Akamai API docs: https://developer.akamai.com/api/luna/network-list/resources.html#getanetworklist
// func (nls *NetworkListService) GetNetworkList(ListID string, opts ListNetworkListsOptions) (*AkamaiNetworkList, *ClientResponse, error) {

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
// func (nls *NetworkListService) CreateNetworkList(opts CreateNetworkListOptions) (*NetworkListResponse, *ClientResponse, error) {

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
// func (nls *NetworkListService) ModifyNetworkList(ListID string, opts AkamaiNetworkList) (*NetworkListResponse, *ClientResponse, error) {

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
// func (nls *NetworkListService) AddNetworkListItems(ListID string, opts CreateNetworkListOptions) (*NetworkListResponse, *ClientResponse, error) {

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
// func (nls *NetworkListService) AddNetworkListElement(ListID, ListElement string) (*NetworkListResponse, *ClientResponse, error) {

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
// func (nls *NetworkListService) RemoveNetworkListItem(ListID, ListItem string) (*NetworkListResponse, *ClientResponse, error) {

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
// func (nls *NetworkListService) SearchNetworkListItem(ListItem string, opts ListNetworkListsOptions) (*[]AkamaiNetworkList, *ClientResponse, error) {

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
// func (nls *NetworkListService) ActivateNetworkList(ListID string, targetEnvironment AkamaiEnvironment, opts ActivateNetworkListOptions) (*NetworkListResponse, *ClientResponse, error) {

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
// func (nls *NetworkListService) GetNetworkListActivationStatus(ListID string, targetEnvironment AkamaiEnvironment) (*ActivateNetworkListStatus, *ClientResponse, error) {

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
