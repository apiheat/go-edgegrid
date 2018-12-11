package edgegrid

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type BillingService struct {
	client *Client
}

type BillingRespElement struct {
	Date      string  `json:"date"`
	Value     float64 `json:"value"`
	Statistic struct {
		Unit string `json:"unit"`
		Name string `json:"name"`
	} `json:"statistic"`
	ProductID  string `json:"productId"`
	ContractID string `json:"contractId"`
	Final      bool   `json:"final"`
}

type BillingMeasures struct {
	BillingDayOnly bool   `url:"billingDayOnly,omitempty"`
	FromMonth      int    `url:"fromMonth,omitempty"`
	FromYear       int    `url:"fromYear,omitempty"`
	Month          int    `url:"month,omitempty"`
	StatisticName  string `url:"statisticName,omitempty"`
	ToMonth        int    `url:"toMonth,omitempty"`
	ToYear         int    `url:"toYear,omitempty"`
	Year           int    `url:"year,omitempty"`
}

func prepareQueryParameters(params BillingMeasures) (queryString string, err error) {
	v, err := query.Values(params)

	if err != nil {
		return "", err
	}

	return v.Encode(), nil
}

func (nls *BillingService) ListContractUsage(contractID, productID string, params BillingMeasures) (*[]BillingRespElement, *ClientResponse, error) {
	queryParams, err := prepareQueryParameters(params)
	if err != nil {
		return nil, nil, err
	}
	apiURI := fmt.Sprintf("%s/contracts/%s/products/%s/measures?%s", BillingPathV2, contractID, productID, queryParams)

	var k *[]BillingRespElement
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}
