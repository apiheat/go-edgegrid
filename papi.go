package edgegrid

import (
	"fmt"
	"net/http"
	"time"
)

// PropertyService represents exposed services to manage properties
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi
type PropertyService struct {
	client *Client
}

type PropertyGroup struct {
	GroupName   string   `json:"groupName"`
	GroupID     string   `json:"groupId"`
	ContractIds []string `json:"contractIds"`
}

type PropertyGroups struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	Groups      struct {
		Items []PropertyGroup `json:"items"`
	} `json:"groups"`
}

type PropertyContract struct {
	ContractID       string `json:"contractId"`
	ContractTypeName string `json:"contractTypeName"`
}

type PropertyContracts struct {
	AccountID string `json:"accountId"`
	Contracts struct {
		Items []PropertyContract `json:"items"`
	} `json:"contracts"`
}

type PropertyProduct struct {
	ProductName string `json:"productName"`
	ProductID   string `json:"productId"`
}

type PropertyProducts struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	Products   struct {
		Items []PropertyProduct `json:"items"`
	} `json:"products"`
}

type PropertyCPCodeNew struct {
	ProductID  string `json:"productId"`
	CpcodeName string `json:"cpcodeName"`
}

type PropertyCPCode struct {
	CpcodeID    string    `json:"cpcodeId"`
	CpcodeName  string    `json:"cpcodeName"`
	ProductIds  []string  `json:"productIds"`
	CreatedDate time.Time `json:"createdDate"`
}

type PropertyCPCodes struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	GroupID    string `json:"groupId"`
	Cpcodes    struct {
		Items []PropertyCPCode `json:"items"`
	} `json:"cpcodes"`
}

type PropertyCPEdgehost struct {
	EdgeHostnameID         string `json:"edgeHostnameId"`
	EdgeHostnameDomain     string `json:"edgeHostnameDomain"`
	ProductID              string `json:"productId"`
	DomainPrefix           string `json:"domainPrefix"`
	DomainSuffix           string `json:"domainSuffix"`
	Secure                 bool   `json:"secure"`
	IPVersionBehavior      string `json:"ipVersionBehavior"`
	MapDetailsSerialNumber int    `json:"mapDetails:serialNumber"`
	MapDetailsMapDomain    string `json:"mapDetails:mapDomain"`
}

type PropertyCPEdgehosts struct {
	AccountID     string `json:"accountId"`
	ContractID    string `json:"contractId"`
	GroupID       string `json:"groupId"`
	EdgeHostnames struct {
		Items []PropertyCPEdgehost `json:"items"`
	} `json:"edgeHostnames"`
}

type PropertyProps struct {
	Properties struct {
		Items []struct {
			AccountID         string `json:"accountId"`
			ContractID        string `json:"contractId"`
			GroupID           string `json:"groupId"`
			PropertyID        string `json:"propertyId"`
			PropertyName      string `json:"propertyName"`
			LatestVersion     int    `json:"latestVersion"`
			StagingVersion    int    `json:"stagingVersion"`
			ProductionVersion int    `json:"productionVersion"`
			AssetID           string `json:"assetId"`
			Note              string `json:"note"`
		} `json:"items"`
	} `json:"properties"`
}

// QStrPropertyAPI includes query params used across calls for PAPI
type QStrPropertyAPI struct {
	ContractID string `url:"contractId,omitempty"`
	GroupID    string `url:"groupId,omitempty"`
	Options    string `url:"options,omitempty"`
}

// ListPropertyContracts This operation provides a read-only list of contract names and identifiers
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcontracts
func (pas *PropertyService) ListPropertyContracts() (*PropertyContracts, *ClientResponse, error) {
	qParams := QStrPropertyAPI{}
	path := fmt.Sprintf("%s/contracts", PAPIPathV1)

	var respStruct *PropertyContracts
	resp, err := pas.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)
	if err != nil {
		return nil, resp, err
	}

	return respStruct, resp, err

}

// ListPropertyGroups This operation provides a read-only list of groups, which may contain properties.
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getgroups
func (pas *PropertyService) ListPropertyGroups() (*PropertyGroups, *ClientResponse, error) {
	qParams := QStrPropertyAPI{}
	path := fmt.Sprintf("%s/groups", PAPIPathV1)

	var respStruct *PropertyGroups
	resp, err := pas.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)
	if err != nil {
		return nil, resp, err
	}

	return respStruct, resp, err

}

// ListPropertyCPCodes This operation lists CP codes available within your contract/group pairing.
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *PropertyService) ListPropertyCPCodes(contractID, groupID string) (*PropertyCPCodes, *ClientResponse, error) {
	qParams := QStrPropertyAPI{
		ContractID: contractID,
		GroupID:    groupID,
	}
	path := fmt.Sprintf("%s/cpcodes", PAPIPathV1)

	var respStruct *PropertyCPCodes
	resp, err := pas.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)
	if err != nil {
		return nil, resp, err
	}

	return respStruct, resp, err

}

// ListPropertyProducts ListPropertyProducts.
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *PropertyService) ListPropertyProducts(contractId string) (*PropertyProducts, *ClientResponse, error) {
	qParams := QStrPropertyAPI{
		ContractID: contractId,
	}
	path := fmt.Sprintf("%s/products", PAPIPathV1)

	var respStruct *PropertyProducts
	resp, err := pas.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)
	if err != nil {
		return nil, resp, err
	}

	return respStruct, resp, err

}

// NewPropertyCPcode Creates new CP Code
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#postcpcodes
func (pas *PropertyService) NewPropertyCPcode(newCPcode *PropertyCPCodeNew, contractID, groupID string) (*ClientResponse, error) {
	qParams := QStrPropertyAPI{
		ContractID: contractID,
		GroupID:    groupID,
	}

	path := fmt.Sprintf("%s/cpcodes", PAPIPathV1)
	resp, err := pas.client.makeAPIRequest(http.MethodPost, path, qParams, nil, newCPcode, nil)
	if err != nil {
		return resp, err
	}

	return resp, err

}

// ListPropertyCPEdgehosts This lists all edge hostnames available under a contract..
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getedgehostnames
func (pas *PropertyService) ListPropertyCPEdgehosts(contractId, groupID string) (*PropertyCPEdgehosts, *ClientResponse, error) {
	qParams := QStrPropertyAPI{
		ContractID: contractId,
		GroupID:    groupID,
		Options:    "mapDetails",
	}
	path := fmt.Sprintf("%s/edgehostnames", PAPIPathV1)

	var respStruct *PropertyCPEdgehosts
	resp, err := pas.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)
	if err != nil {
		return nil, resp, err
	}

	return respStruct, resp, err

}

// ListProperties This lists all properties available under a contract/group
// Akamai API docs: https://developer.akamai.com/api/core_features/property_manager/v1.html#getproperties
func (pas *PropertyService) ListPropertyProperties(contractId, groupID string) (*PropertyProps, *ClientResponse, error) {
	qParams := QStrPropertyAPI{
		ContractID: contractId,
		GroupID:    groupID,
	}
	path := fmt.Sprintf("%s/properties", PAPIPathV1)

	var respStruct *PropertyProps
	resp, err := pas.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)
	if err != nil {
		return nil, resp, err
	}

	return respStruct, resp, err

}
