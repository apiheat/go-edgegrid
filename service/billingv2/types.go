package billingv2

// BillingResp response from Akamai
type BillingResp []BillingRespElement

// BillingRespElement is item in billing response
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
