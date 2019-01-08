package edgegrid

import (
	"fmt"
	"net/http"
)

type FirewallRulesNotificationsService struct {
	client *Client
}

// FRNServices data representation
type FRNServices []struct {
	FRNService
}

type FRNService struct {
	ServiceID   int    `json:"serviceId"`
	ServiceName string `json:"serviceName"`
	Description string `json:"description"`
}

// FRNSubscriptions data representation
type FRNSubscriptions struct {
	Subscriptions []FRNSubscription `json:"subscriptions"`
}

// FRNSubscription data representation
type FRNSubscription struct {
	ServiceID   int    `json:"serviceId"`
	ServiceName string `json:"serviceName,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email"`
	SignupDate  string `json:"signupDate,omitempty"`
}

// FRNCidrs data representation
type FRNCidrs []struct {
	CidrID        int         `json:"cidrId"`
	ServiceID     int         `json:"serviceId"`
	ServiceName   string      `json:"serviceName"`
	Description   string      `json:"description"`
	Cidr          string      `json:"cidr"`
	CidrMask      string      `json:"cidrMask"`
	Port          string      `json:"port"`
	CreationDate  string      `json:"creationDate"`
	EffectiveDate string      `json:"effectiveDate"`
	ChangeDate    interface{} `json:"changeDate"`
	MinIP         string      `json:"minIp"`
	MaxIP         string      `json:"maxIp"`
	LastAction    string      `json:"lastAction"`
}

// QStrFRN includes query params used across firewall network rules
type QStrFRN struct{}

// ListServices provides list of services to which it is possible to subscribe
func (nls *FirewallRulesNotificationsService) ListServices() (*FRNServices, *ClientResponse, error) {
	qParams := QStrFRN{}
	path := fmt.Sprintf("%s/services", FRNPathV1)

	var respStruct *FRNServices
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// ListService provides details of service specified by its unique ID
func (nls *FirewallRulesNotificationsService) ListService(id string) (*FRNService, *ClientResponse, error) {
	qParams := QStrFRN{}
	path := fmt.Sprintf("%s/services/%s", FRNPathV1, id)

	var respStruct *FRNService
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

// ListSubscriptions provides list of services to which we are subscribed
func (nls *FirewallRulesNotificationsService) ListSubscriptions() (*FRNSubscriptions, *ClientResponse, error) {
	qParams := QStrFRN{}
	path := fmt.Sprintf("%s/subscriptions", FRNPathV1)

	var respStruct *FRNSubscriptions
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}

//UpdateSubscriptions updates current subscription
func (nls *FirewallRulesNotificationsService) UpdateSubscriptions(services []int, email string) (*FRNSubscriptions, *ClientResponse, error) {
	qParams := QStrFRN{}
	path := fmt.Sprintf("%s/subscriptions", FRNPathV1)

	var requestStruct FRNSubscriptions
	for _, s := range services {
		service := FRNSubscription{
			ServiceID: s,
			Email:     email,
		}
		requestStruct.Subscriptions = append(requestStruct.Subscriptions, service)
	}

	var respStruct *FRNSubscriptions
	resp, err := nls.client.makeAPIRequest(http.MethodPut, path, qParams, &respStruct, requestStruct, nil)

	return respStruct, resp, err
}

// ListCIDRBlocks provides information about CIDR blocks
func (nls *FirewallRulesNotificationsService) ListCIDRBlocks(filterStr string) (*FRNCidrs, *ClientResponse, error) {
	qParams := QStrFRN{}
	path := fmt.Sprintf("%s/cidr-blocks", FRNPathV1)
	if filterStr != "" {
		path = fmt.Sprintf("%s/cidr-blocks%s", FRNPathV1, filterStr)
	}

	var respStruct *FRNCidrs
	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

	return respStruct, resp, err
}
