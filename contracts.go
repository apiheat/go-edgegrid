package edgegrid

import (
	"fmt"
	"net/http"
)

type ContractsService struct {
	client *Client
}

type QStrContractsProducts struct {
	From  string `url:"from,omitempty"`
	To    string `url:"to,omitempty"`
	Depth string `url:"depth,omitempty"`
}

type ContractProductsResp struct {
	Products struct {
		ContractID        string `json:"contractId"`
		MarketingProducts []struct {
			MarketingProductID   string `json:"marketingProductId"`
			MarketingProductName string `json:"marketingProductName"`
		} `json:"marketing-products"`
	} `json:"products"`
}

// List Lists contracts
func (nls *ContractsService) List(depth string) (*[]string, *ClientResponse, error) {
	qParams := QStrContractsProducts{
		Depth: depth,
	}
	path := fmt.Sprintf("%s/contracts/identifiers", ContractsPath)

	var respStruct *[]string
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// ListContractProducts Lists products
func (nls *ContractsService) ListContractProducts(contractID, qFrom, qTo string) (*ContractProductsResp, *ClientResponse, error) {
	qParams := QStrContractsProducts{
		From: qFrom,
		To:   qTo,
	}
	path := fmt.Sprintf("%s/contracts/%s/products/summaries", ContractsPath, contractID)

	var respStruct *ContractProductsResp
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
