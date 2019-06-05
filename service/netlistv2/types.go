package netlistv2

import "time"

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
