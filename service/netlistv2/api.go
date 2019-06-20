package netlistv2

import (
	"fmt"
	"strconv"

	"github.com/apiheat/go-edgegrid/edgegrid"
)

// ListNetworkLists List all configured Network Lists for the authenticated user.
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlists
func (nls *Netlistv2) ListNetworkLists(opts ListNetworkListsOptionsv2) (*NetworkListsv2, error) {
	var networkListsv2 NetworkListsv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetQueryParams(map[string]string{
			"extended":        strconv.FormatBool(opts.Extended),
			"includeElements": strconv.FormatBool(opts.IncludeElements),
			"search":          opts.Search,
		}).
		SetResult(&networkListsv2).
		SetError(&e).
		Get(basePath)

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListsv2, nil
}

// CreateNetworkList Create a new network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *Netlistv2) CreateNetworkList(opts NetworkListsOptionsv2) (*NetworkListv2, error) {

	var networkListv2 NetworkListv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetResult(&networkListv2).
		SetError(&e).
		SetBody(opts).
		Post(basePath)

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListv2, nil
}

// GetNetworkList Gets a specific network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlist
func (nls *Netlistv2) GetNetworkList(ListID string, opts ListNetworkListsOptionsv2) (*NetworkListv2, error) {

	var networkListv2 NetworkListv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetQueryParams(map[string]string{
			"extended":        strconv.FormatBool(opts.Extended),
			"includeElements": strconv.FormatBool(opts.IncludeElements),
		}).
		SetResult(&networkListv2).
		SetError(&e).
		Get(fmt.Sprintf("%s/%s", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListv2, nil

}

// AddNetworkListElement Adds items to network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *Netlistv2) AddNetworkListElement(ListID string, opts NetworkListsOptionsv2) (*NetworkListv2, error) {

	var networkListv2 NetworkListv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetResult(&networkListv2).
		SetError(&e).
		SetBody(opts).
		Post(fmt.Sprintf("%s/%s/append", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListv2, nil
}

// RemoveNetworkListElement Removes network list element
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) RemoveNetworkListElement(ListID, element string) (*NetworkListv2, error) {

	var networkListv2 NetworkListv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetResult(&networkListv2).
		SetQueryParams(map[string]string{
			"element": element,
		}).
		SetError(&e).
		Delete(fmt.Sprintf("%s/%s/elements", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListv2, nil
}

// ActivateNetworkList Activates network list on specified network ( PRODUCTION or STAGING )
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) ActivateNetworkList(ListID string, targetEnv edgegrid.AkamaiEnvironment, opts NetworkListActivationOptsv2) (*NetworkListActivationStatusv2, error) {
	var networkListActivationStatusv2 NetworkListActivationStatusv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetBody(opts).
		SetResult(&networkListActivationStatusv2).
		SetError(&e).
		Post(fmt.Sprintf("%s/%s/environments/%s/activate", basePath, ListID, targetEnv))

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListActivationStatusv2, nil
}

// GetActivationStatus Gets activation network list status on specified network ( PRODUCTION or STAGING )
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) GetActivationStatus(ListID string, targetEnv edgegrid.AkamaiEnvironment) (*NetworkListActivationStatusv2, error) {
	var networkListActivationStatusv2 NetworkListActivationStatusv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetResult(&networkListActivationStatusv2).
		SetError(&e).
		Get(fmt.Sprintf("%s/%s/environments/%s/status", basePath, ListID, targetEnv))

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListActivationStatusv2, nil

}

// DeleteNetworkList Remove network list element
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) DeleteNetworkList(ListID string) (*NetworkListDeleteResponse, error) {
	var networkListDeleteResponse NetworkListDeleteResponse
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetResult(&networkListDeleteResponse).
		SetError(&e).
		Delete(fmt.Sprintf("%s/%s", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if e.Status != 0 {
		return nil, e
	}

	return &networkListDeleteResponse, nil
}

// NetworkListNotification Manage network list subscription
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) NetworkListNotification(action edgegrid.AkamaiSubscription, sub NetworkListSubscription) error {

	var networkListv2 NetworkListv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetResult(&networkListv2).
		SetError(&e).
		SetBody(sub).
		Post(fmt.Sprintf("/network-list/v2/notifications/%s", action))

	if err != nil {
		return err
	}

	if e.Status != 0 {
		return e
	}

	return nil

}
