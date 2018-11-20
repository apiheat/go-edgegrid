package edgegrid

import (
	"fmt"
	"net/http"
)

type FirewallRulesNotificationsService struct {
	client *Client
}

// AkamaiFRNServices data representation
type AkamaiFRNServices []struct {
	AkamaiFRNService
}

type AkamaiFRNService struct {
	ServiceID   int    `json:"serviceId"`
	ServiceName string `json:"serviceName"`
	Description string `json:"description"`
}

func (nls *FirewallRulesNotificationsService) ListServices() (*AkamaiFRNServices, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/services", FRNPathV1)

	var k *AkamaiFRNServices
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}

func (nls *FirewallRulesNotificationsService) ListService(id string) (*AkamaiFRNService, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/services/%s", FRNPathV1, id)

	var k *AkamaiFRNService
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}

// AkamaiFRNSubscriptions data representation
type AkamaiFRNSubscriptions struct {
	Subscriptions []AkamaiFRNSubscription `json:"subscriptions"`
}

// AkamaiFRNSubscription data representation
type AkamaiFRNSubscription struct {
	ServiceID   int    `json:"serviceId"`
	ServiceName string `json:"serviceName,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email"`
	SignupDate  string `json:"signupDate,omitempty"`
}

func (nls *FirewallRulesNotificationsService) ListSubscriptions() (*AkamaiFRNSubscriptions, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/subscriptions", FRNPathV1)

	var k *AkamaiFRNSubscriptions
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}

func (nls *FirewallRulesNotificationsService) UpdateSubscriptions(services []int, email string) (*AkamaiFRNSubscriptions, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/subscriptions", FRNPathV1)

	var obj AkamaiFRNSubscriptions
	for _, s := range services {
		service := AkamaiFRNSubscription{
			ServiceID: s,
			Email:     email,
		}
		obj.Subscriptions = append(obj.Subscriptions, service)
	}

	var k *AkamaiFRNSubscriptions
	resp, err := nls.client.NewRequest(http.MethodPut, apiURI, obj, &k)

	return k, resp, err
}

// AkamaiFRNCidrs data representation
type AkamaiFRNCidrs []struct {
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

func (nls *FirewallRulesNotificationsService) ListCIDRBlocks(filterStr string) (*AkamaiFRNCidrs, *ClientResponse, error) {
	apiURI := fmt.Sprintf("%s/cidr-blocks", FRNPathV1)
	if filterStr != "" {
		apiURI = fmt.Sprintf("%s/cidr-blocks%s", FRNPathV1, filterStr)
	}

	var k *AkamaiFRNCidrs
	resp, err := nls.client.NewRequest(http.MethodGet, apiURI, nil, &k)

	return k, resp, err
}
