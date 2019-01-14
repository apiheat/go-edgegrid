package edgegrid

import (
	"fmt"
	"net/http"
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

type QStrBillingMeasures struct {
	BillingDayOnly bool   `url:"billingDayOnly,omitempty"`
	FromMonth      int    `url:"fromMonth,omitempty"`
	FromYear       int    `url:"fromYear,omitempty"`
	Month          int    `url:"month,omitempty"`
	StatisticName  string `url:"statisticName,omitempty"`
	ToMonth        int    `url:"toMonth,omitempty"`
	ToYear         int    `url:"toYear,omitempty"`
	Year           int    `url:"year,omitempty"`
}

// ListContractUsage Provides information
func (nls *BillingService) ListContractUsage(contractID, productID string, qStringParams QStrBillingMeasures) (*[]BillingRespElement, *ClientResponse, error) {

	path := fmt.Sprintf("%s/contracts/%s/products/%s/measures", BillingPathV2, contractID, productID)

	var respStruct *[]BillingRespElement
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qStringParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
