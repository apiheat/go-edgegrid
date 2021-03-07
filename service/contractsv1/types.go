package contractsv1

// OutputContractIDs lists the contract ids the requesting user has access to.
type OutputContractIDs []string

// OutputContracts lists the contracts that the requesting user has access to.
type OutputContracts struct {
	Contracts []OutputContractElement `json:"contracts"`
}

// OutputContractElement provides contract details
type OutputContractElement struct {
	// The unique identifier for a contract.
	ContractID string `json:"contractId"`
	// The URL that accesses product information for the contractId.
	Href string `json:"href"`
}

// OutputProducts lists the products associated with specific contracts.
type OutputProducts struct {
	// Object that lists the products associated with the contract specified.
	Products OutputProduct `json:"products"`
}

// OutputProduct Object that lists the products associated with the contract specified.
type OutputProduct struct {
	// The unique identifier for a contract.
	ContractID string `json:"contractId"`
	// The identifiers and names for each product included on a contract.
	MarketingProducts []OutputMarketingProduct `json:"marketing-products"`
}

// OutputMarketingProduct The identifier and names for product included on a contract.
type OutputMarketingProduct struct {
	// The unique identifier for a product.
	MarketingProductID string `json:"marketingProductId"`
	// The formal name of a product.
	MarketingProductName string `json:"marketingProductName"`
}

// OutputReportingGroups ReportingGroup members
type OutputReportingGroups []OutputReportingGroup

// OutputReportingGroupIDs ReportingGroup IDs
type OutputReportingGroupIDs []string

// OutputReportingGroup A logical grouping of content provider (CP) codes.
type OutputReportingGroup struct {
	// Unique identifier for each reporting group.
	ID int `json:"id"`
	// The descriptive name you supply for each reporting group.
	Name string `json:"name"`
}
