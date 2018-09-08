package edgegrid

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// PropertyService represents exposed services to manage properties
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi
type ReportingAPIService struct {
	client *Client
}

type ReportingAPIGroup struct {
	GroupName   string   `json:"groupName"`
	GroupID     string   `json:"groupId"`
	ContractIds []string `json:"contractIds"`
}

type ReportingAPIGroups struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	Groups      struct {
		Items []ReportingAPIGroup `json:"items"`
	} `json:"groups"`
}

type ReportingAPIContract struct {
	ContractID       string `json:"contractId"`
	ContractTypeName string `json:"contractTypeName"`
}

type ReportingAPIContracts struct {
	AccountID string `json:"accountId"`
	Contracts struct {
		Items []ReportingAPIContract `json:"items"`
	} `json:"contracts"`
}

type ReportingAPIProduct struct {
	ProductName string `json:"productName"`
	ProductID   string `json:"productId"`
}

type ReportingAPIProducts struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	Products   struct {
		Items []ReportingAPIProduct `json:"items"`
	} `json:"products"`
}

type ReportingAPICPCodeNew struct {
	ProductID  string `json:"productId"`
	CpcodeName string `json:"cpcodeName"`
}

type ReportingAPICPCode struct {
	CpcodeID    string    `json:"cpcodeId"`
	CpcodeName  string    `json:"cpcodeName"`
	ProductIds  []string  `json:"productIds"`
	CreatedDate time.Time `json:"createdDate"`
}

type ReportingAPICPCodes struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	GroupID    string `json:"groupId"`
	Cpcodes    struct {
		Items []ReportingAPICPCode `json:"items"`
	} `json:"cpcodes"`
}

type ReportingAPICPEdgehost struct {
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

type ReportingAPICPEdgehosts struct {
	AccountID     string `json:"accountId"`
	ContractID    string `json:"contractId"`
	GroupID       string `json:"groupId"`
	EdgeHostnames struct {
		Items []ReportingAPICPEdgehost `json:"items"`
	} `json:"edgeHostnames"`
}

func (pas *ReportingAPIService) GenerateReportRequest() {
	req, err := http.NewRequest("GET", "http://api.themoviedb.org/3/tv/popular", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("api_key", "key_from_environment_or_flag")
	q.Add("another_thing", "foo & bar")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	// Output:
	// http://api.themoviedb.org/3/tv/popular?another_thing=foo+%26+bar&api_key=key_from_environment_or_flag
}

// ListReportingAPIContracts This operation provides a read-only list of contract names and identifiers
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcontracts
func (pas *ReportingAPIService) ListReportingAPIContracts() (*ReportingAPIContracts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/contracts", apiPaths["papi_v1"])

	var k *ReportingAPIContracts
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListReportingAPIGroups This operation provides a read-only list of groups, which may contain properties.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getgroups
func (pas *ReportingAPIService) ListReportingAPIGroups() (*ReportingAPIGroups, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/groups", apiPaths["papi_v1"])

	var k *ReportingAPIGroups
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListReportingAPICPCodes This operation lists CP codes available within your contract/group pairing.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *ReportingAPIService) ListReportingAPICPCodes(contractID, groupID string) (*ReportingAPICPCodes, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/cpcodes?contractId=%s&groupId=%s",
		apiPaths["papi_v1"],
		contractID,
		groupID)

	var k *ReportingAPICPCodes
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// ListReportingAPIProducts ListReportingAPIProducts.
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getcpcodes
func (pas *ReportingAPIService) ListReportingAPIProducts(contractId string) (*ReportingAPIProducts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/products?contractId=%s",
		apiPaths["papi_v1"],
		contractId)

	var k *ReportingAPIProducts
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}

// NewReportingAPICPcode Creates new CP Code
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#postcpcodes
func (pas *ReportingAPIService) NewReportingAPICPcode(newCPcode *ReportingAPICPCodeNew, contractID, groupID string) (*ClientResponse, error) {

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

// ListReportingAPICPEdgehosts This lists all edge hostnames available under a contract..
//
// Akamai API docs: https://developer.akamai.com/api/luna/papi/resources.html#getedgehostnames
func (pas *ReportingAPIService) ListReportingAPICPEdgehosts(contractId string) (*ReportingAPICPEdgehosts, *ClientResponse, error) {

	apiURI := fmt.Sprintf("%s/edgehostnames?contractId=%s&groupId=%s&options=mapDetails",
		apiPaths["papi_v1"],
		contractId)

	var k *ReportingAPICPEdgehosts
	resp, err := pas.client.NewRequest("GET", apiURI, nil, &k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, err

}
