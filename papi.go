package edgegrid

import (
	"fmt"
	"time"
)

// PropertyService represents exposed services to manage properties
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi
type PropertyAPIService struct {
	client *Client
}

type PropertyAPIGroup struct {
	GroupName   string   `json:"groupName"`
	GroupID     string   `json:"groupId"`
	ContractIds []string `json:"contractIds"`
}

type PropertyAPIGroups struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	Groups      struct {
		Items []PropertyAPIGroup `json:"items"`
	} `json:"groups"`
}

type PropertyAPIContract struct {
	ContractID       string `json:"contractId"`
	ContractTypeName string `json:"contractTypeName"`
}

type PropertyAPIContracts struct {
	AccountID string `json:"accountId"`
	Contracts struct {
		Items []PropertyAPIContract `json:"items"`
	} `json:"contracts"`
}

type PropertyAPIProduct struct {
	ProductName string `json:"productName"`
	ProductID   string `json:"productId"`
}

type PropertyAPIProducts struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	Products   struct {
		Items []PropertyAPIProduct `json:"items"`
	} `json:"products"`
}

type PropertyAPICPCodeNew struct {
	ProductID  string `json:"productId"`
	CpcodeName string `json:"cpcodeName"`
}

type PropertyAPICPCode struct {
	CpcodeID    string    `json:"cpcodeId"`
	CpcodeName  string    `json:"cpcodeName"`
	ProductIds  []string  `json:"productIds"`
	CreatedDate time.Time `json:"createdDate"`
}

type PropertyAPICPCodes struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	GroupID    string `json:"groupId"`
	Cpcodes    struct {
		Items []PropertyAPICPCode `json:"items"`
	} `json:"cpcodes"`
}

type PropertyAPICPEdgehost struct {
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

type PropertyAPICPEdgehosts struct {
	AccountID     string `json:"accountId"`
	ContractID    string `json:"contractId"`
	GroupID       string `json:"groupId"`
	EdgeHostnames struct {
		Items []PropertyAPICPEdgehost `json:"items"`
	} `json:"edgeHostnames"`
}

type PropertyAPIProps struct {
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

// ListPropertyAPIContracts This operation provides a read-only list of contract names and identifiers
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcontracts
func (pas *PropertyAPIService) ListPropertyAPIContracts() (*PropertyAPIContracts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/contracts", apiPaths["papi_v1"])

	var k *PropertyAPIContracts
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListPropertyAPIGroups This operation provides a read-only list of groups, which may contain properties.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getgroups
func (pas *PropertyAPIService) ListPropertyAPIGroups() (*PropertyAPIGroups, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/groups", apiPaths["papi_v1"])

	var k *PropertyAPIGroups
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListPropertyAPICPCodes This operation lists CP codes available within your contract/group pairing.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *PropertyAPIService) ListPropertyAPICPCodes(contractID, groupID string) (*PropertyAPICPCodes, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/cpcodes?contractId=%s&groupId=%s",
		apiPaths["papi_v1"],
		contractID,
		groupID)

	var k *PropertyAPICPCodes
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListPropertyAPIProducts ListPropertyAPIProducts.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *PropertyAPIService) ListPropertyAPIProducts(contractId string) (*PropertyAPIProducts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/products?contractId=%s",
		apiPaths["papi_v1"],
		contractId)

	var k *PropertyAPIProducts
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// NewPropertyAPICPcode Creates new CP Code
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#postcpcodes
func (pas *PropertyAPIService) NewPropertyAPICPcode(newCPcode *PropertyAPICPCodeNew, contractID, groupID string) (*ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/cpcodes?contractId=%s&groupId=%s",
		apiPaths["papi_v1"],
		contractID,
		groupID)

	resp, err := pas.client.NewRequest("POST", apiURI, newCPcode, nil)
	if err != nil {
		return resp, err
	}

	return resp, err

}

// ListPropertyAPICPEdgehosts This lists all edge hostnames available under a contract..
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getedgehostnames
func (pas *PropertyAPIService) ListPropertyAPICPEdgehosts(contractId, groupID string) (*PropertyAPICPEdgehosts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/edgehostnames?contractId=%s&groupId=%s&options=mapDetails",
		apiPaths["papi_v1"],
		contractId,
		groupID)

	var k *PropertyAPICPEdgehosts
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListProperties This lists all properties available under a contract/group
//
// Akamai API docs: https://developer.akamai.com/api/core_features/property_manager/v1.html#getproperties
func (pas *PropertyAPIService) ListPropertyAPIProperties(contractId, groupID string) (*PropertyAPIProps, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/properties?contractId=%s&groupId=%s",
		apiPaths["papi_v1"],
		contractId,
		groupID)

	var k *PropertyAPIProps
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}
