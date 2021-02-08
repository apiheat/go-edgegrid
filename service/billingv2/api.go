package billingv2

import (
	"fmt"
)

// ListContractUsage returns billing measures per product in a given contract
func (bl *Billingv2) ListContractUsage(contractID, productID string, qStringParams map[string]string) (*BillingResp, error) {
	apiURI := fmt.Sprintf("%s/contracts/%s/products/%s/measures", basePath, contractID, productID)

	// Create and execute request
	resp, err := bl.Client.Rclient.R().
		SetResult(BillingResp{}).
		SetQueryParams(qStringParams).
		SetError(BillingErrorv2{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*BillingErrorv2)

		return nil, fmt.Errorf("Error requesting contract usage per product. code: %s, title: %s", e.Code, e.Title)
	}

	return resp.Result().(*BillingResp), nil
}
