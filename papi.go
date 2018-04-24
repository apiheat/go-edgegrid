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
func (pas *PropertyAPIService) ListPropertyAPICPCodes(contractId, groupId string) (*PropertyAPICPCodes, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/cpcodes?contractId=%s&groupId=%s",
		apiPaths["papi_v1"],
		contractId,
		groupId)

	var k *PropertyAPICPCodes
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}
