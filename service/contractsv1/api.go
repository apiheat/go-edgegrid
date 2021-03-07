package contractsv1

import "fmt"

func depthAllowedValues(val string) (ok bool) {
	values := []string{"TOP", "ALL"}
	for i := range values {
		if ok = values[i] == val; ok {
			return
		}
	}
	return
}

// ListContracts gets the list of contracts that a user has access to.
// 'depth' returns a specific set of contracts.
// Select TOP to return only parent contracts or ALL to return both parent and child contracts.
func (c *Contractsv1) ListContracts(depth string) (*OutputContractIDs, error) {
	query := map[string]string{}
	if depth != "" {
		if depthAllowedValues(depth) {
			return nil, fmt.Errorf("Unsupported argument 'depth' value. Use 'TOP' or 'ALL' value")
		}

		query["depth"] = depth
	}

	apiURI := fmt.Sprintf("%s/contracts/identifiers", basePath)

	// Create and execute request
	resp, err := c.Client.Rclient.R().
		SetResult(OutputContractIDs{}).
		SetError(ContractsErrorv1{}).
		SetQueryParams(query).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*ContractsErrorv1)

		return nil, e
	}

	return resp.Result().(*OutputContractIDs), nil
}

// ListProductsPerContract gets the IDs and names of the products associated with a contract for the time frame selected.
// From - Ex: 2015-01-31. The start date, in UTC, to use when looking for products associated with a contract.
//        The search always begins at midnight (0:00) UTC of the specified date.
//        The default start date is 30 days prior to the current date.
//        For current contracts, you can select a date within the past 15 months of the current date.
//        For expired contracts, you are limited to a date range of 30 days within the 15 month window.
// To - Ex: 2016-03-31. The end date, in UTC, to use when looking for products associated with a contract.
//      The search always ends at 23:59:59 UTC of the specified date. The default end date is the current date.
func (c *Contractsv1) ListProductsPerContract(contractID, from, to string) (*OutputContracts, error) {
	query := map[string]string{}
	if contractID == "" {
		return nil, fmt.Errorf("Missing argument 'contractID'")
	}

	apiURI := fmt.Sprintf("%s/contracts/%s/products/summaries", basePath, contractID)

	if from != "" {
		query["from"] = from
	}

	if to != "" {
		query["to"] = to
	}

	// Create and execute request
	resp, err := c.Client.Rclient.R().
		SetResult(OutputContracts{}).
		SetError(ContractsErrorv1{}).
		SetQueryParams(query).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*ContractsErrorv1)

		return nil, e
	}

	return resp.Result().(*OutputContracts), nil
}

// ListReportingGroups gets the IDs of the Content Provider (CP) reporting groups that you have access to along with their names.
// To run this operation, your user account needs the CPCode Rep Group role.
func (c *Contractsv1) ListReportingGroups() (*OutputReportingGroups, error) {
	apiURI := fmt.Sprintf("%s/reportingGroups/", basePath)

	// Create and execute request
	resp, err := c.Client.Rclient.R().
		SetResult(OutputReportingGroups{}).
		SetError(ContractsErrorv1{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*ContractsErrorv1)

		return nil, e
	}

	return resp.Result().(*OutputReportingGroups), nil
}

// ListReportingGroupIDs gets the IDs of the Content Provider (CP) reporting groups that you have access to.
// To run this operation, your user account needs the CPCode Rep Group role.
func (c *Contractsv1) ListReportingGroupIDs() (*OutputReportingGroupIDs, error) {
	apiURI := fmt.Sprintf("%s/reportingGroups/identifiers", basePath)

	// Create and execute request
	resp, err := c.Client.Rclient.R().
		SetResult(OutputReportingGroupIDs{}).
		SetError(ContractsErrorv1{}).
		Get(apiURI)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*ContractsErrorv1)

		return nil, e
	}

	return resp.Result().(*OutputReportingGroupIDs), nil
}

// ListProductsPerReportingGroup gets the IDs and names of the products associated with the reporting group for the time frame selected.
// To run this operation, your user account needs the CPCode Rep Group role.
// When a request is successful, it may return either a 200 or a 300 response.
// The API returns a 200 (OK) response when the CP code reporting group is associated with only one contract.
// You’ll receive a 300 (Multiple Choices) response when the request returns a list of matching contracts because
// the CP code reporting group is associated with multiple contracts.
// To retrieve product information when you receive a 300 response code, make a new GET request to the hyperlinks provided in the response.
// From - Ex: 2015-01-31. The start date, in UTC, to use when looking for products associated with a contract.
//        The search always begins at midnight (0:00) UTC of the specified date.
//        The default start date is 30 days prior to the current date.
//        For current contracts, you can select a date within the past 15 months of the current date.
//        For expired contracts, you are limited to a date range of 30 days within the 15 month window.
// To - Ex: 2016-03-31. The end date, in UTC, to use when looking for products associated with a contract.
//      The search always ends at 23:59:59 UTC of the specified date. The default end date is the current date.
func (c *Contractsv1) ListProductsPerReportingGroup(reportingGroupID, from, to string) (*OutputProducts, *OutputContracts, error) {
	query := map[string]string{}
	if reportingGroupID == "" {
		return nil, nil, fmt.Errorf("Missing argument 'reportingGroupID'")
	}

	apiURI := fmt.Sprintf("%s/reportingGroups/%s/products/summaries", basePath, reportingGroupID)

	if from != "" {
		query["from"] = from
	}

	if to != "" {
		query["to"] = to
	}

	// Create and execute request
	resp, err := c.Client.Rclient.R().
		SetResult(OutputProducts{}).
		SetError(ContractsErrorv1{}).
		SetQueryParams(query).
		Get(apiURI)

	if err != nil {
		return nil, nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*ContractsErrorv1)

		return nil, nil, e
	}

	if resp.StatusCode() == 300 {
		return nil, resp.Result().(*OutputContracts), nil
	}

	return resp.Result().(*OutputProducts), nil, nil
}
