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

// ListPropertyContracts This operation provides a read-only list of contract names and identifiers
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcontracts
func (pas *PropertyService) ListPropertyContracts() (*PropertyContracts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/contracts", PAPIPathV1)

	var k *PropertyContracts
	resp, err := pas.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListPropertyGroups This operation provides a read-only list of groups, which may contain properties.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getgroups
func (pas *PropertyService) ListPropertyGroups() (*PropertyGroups, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/groups", PAPIPathV1)

	var k *PropertyGroups
	resp, err := pas.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListPropertyCPCodes This operation lists CP codes available within your contract/group pairing.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *PropertyService) ListPropertyCPCodes(contractID, groupID string) (*PropertyCPCodes, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/cpcodes?contractId=%s&groupId=%s",
		PAPIPathV1,
		contractID,
		groupID)

	var k *PropertyCPCodes
	resp, err := pas.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListPropertyProducts ListPropertyProducts.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *PropertyService) ListPropertyProducts(contractId string) (*PropertyProducts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/products?contractId=%s",
		PAPIPathV1,
		contractId)

	var k *PropertyProducts
	resp, err := pas.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// NewPropertyCPcode Creates new CP Code
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#postcpcodes
func (pas *PropertyService) NewPropertyCPcode(newCPcode *PropertyCPCodeNew, contractID, groupID string) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/cpcodes?contractId=%s&groupId=%s",
		PAPIPathV1,
		contractID,
		groupID)

	resp, err := pas.client.NewRequest(http.MethodPost, apiURI, newCPcode, nil)
	if err != nil {
		return resp, err
	}

	return resp, err

}

// ListPropertyCPEdgehosts This lists all edge hostnames available under a contract..
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getedgehostnames
func (pas *PropertyService) ListPropertyCPEdgehosts(contractId, groupID string) (*PropertyCPEdgehosts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/edgehostnames?contractId=%s&groupId=%s&options=mapDetails",
		PAPIPathV1,
		contractId,
		groupID)

	var k *PropertyCPEdgehosts
	resp, err := pas.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListProperties This lists all properties available under a contract/group
//
// Akamai API docs: https://developer.akamai.com/api/core_features/property_manager/v1.html#getproperties
func (pas *PropertyService) ListPropertyProperties(contractId, groupID string) (*PropertyProps, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/properties?contractId=%s&groupId=%s",
		PAPIPathV1,
		contractId,
		groupID)

	var k *PropertyProps
	resp, err := pas.client.NewRequest(http.MethodGet, apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}
