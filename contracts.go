package edgegrid

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ContractsService struct {
	client *Client
}

type ContractsProductsParams struct {
	From string `url:"from,omitempty"`
	To   string `url:"to,omitempty"`
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

func prepareContractsQueryParameters(params ContractsProductsParams) (queryString string, err error) {
	v, err := query.Values(params)

	if err != nil {
		return "", err
	}

	return v.Encode(), nil
}

func (nls *ContractsService) List(depth string) (*[]string, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/contracts/identifiers", ContractsPath)
	if depth != "" {
		apiURI = fmt.Sprintf("%s/contracts/identifiers?depth=%s", ContractsPath, depth)
	}

	var k *[]string
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}

func (nls *ContractsService) ListContractProducts(contractID string, params ContractsProductsParams) (*ContractProductsResp, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/contracts/%s/products/summaries", ContractsPath, contractID)
	if params.From != "" || params.To != "" {
		queryParams, err := prepareContractsQueryParameters(params)
		if err != nil {
			return nil, nil, err
		}
		apiURI = fmt.Sprintf("%s/contracts/%s/products/summaries?%s", ContractsPath, contractID, queryParams)
	}

	var k *ContractProductsResp
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}
