package netlistv2

import (
	"fmt"
	"strconv"
)

// ListNetworkLists List all configured Network Lists for the authenticated user.
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlists
func (nls *Netlistv2) ListNetworkLists(opts ListNetworkListsOptionsv2) (*NetworkListsv2, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetQueryParams(map[string]string{
			"extended":        strconv.FormatBool(opts.Extended),
			"includeElements": strconv.FormatBool(opts.IncludeElements),
			"search":          opts.Search,
		}).
		SetResult(NetworkListsv2{}).
		SetError(NetworkListErrorv2{}).
		Get(basePath)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListsv2), nil
}

// CreateNetworkList Create a new network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *Netlistv2) CreateNetworkList(opts NetworkListsOptionsv2) (*NetworkListv2, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetResult(NetworkListv2{}).
		SetError(NetworkListErrorv2{}).
		SetBody(opts).
		Post(basePath)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListv2), nil
}

// GetNetworkList Gets a specific network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlist
func (nls *Netlistv2) GetNetworkList(ListID string, opts ListNetworkListsOptionsv2) (*NetworkListv2, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetQueryParams(map[string]string{
			"extended":        strconv.FormatBool(opts.Extended),
			"includeElements": strconv.FormatBool(opts.IncludeElements),
		}).
		SetResult(NetworkListv2{}).
		SetError(NetworkListErrorv2{}).
		Get(fmt.Sprintf("%s/%s", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListv2), nil

}

// AddNetworkListElement Adds items to network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
func (nls *Netlistv2) AddNetworkListElement(ListID string, opts NetworkListsOptionsv2) (*NetworkListv2, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetResult(NetworkListv2{}).
		SetError(NetworkListErrorv2{}).
		SetBody(opts).
		Post(fmt.Sprintf("%s/%s/append", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListv2), nil
}

// RemoveNetworkListElement Removes network list element
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) RemoveNetworkListElement(ListID, element string) (*NetworkListv2, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetResult(NetworkListv2{}).
		SetQueryParams(map[string]string{
			"element": element,
		}).
		SetError(NetworkListErrorv2{}).
		Delete(fmt.Sprintf("%s/%s/elements", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListv2), nil
}

// ActivateNetworkList Activates network list on specified network ( PRODUCTION or STAGING )
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) ActivateNetworkList(ListID string, targetEnv AkamaiEnvironment, opts NetworkListActivationOptsv2) (*NetworkListActivationStatusv2, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetBody(opts).
		SetResult(NetworkListActivationStatusv2{}).
		SetError(NetworkListErrorv2{}).
		Post(fmt.Sprintf("%s/%s/environments/%s/activate", basePath, ListID, targetEnv))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListActivationStatusv2), nil
}

// GetActivationStatus Gets activation network list status on specified network ( PRODUCTION or STAGING )
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) GetActivationStatus(ListID string, targetEnv AkamaiEnvironment) (*NetworkListActivationStatusv2, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetResult(NetworkListActivationStatusv2{}).
		SetError(NetworkListErrorv2{}).
		Get(fmt.Sprintf("%s/%s/environments/%s/status", basePath, ListID, targetEnv))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListActivationStatusv2), nil

}

// DeleteNetworkList Remove network list
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) DeleteNetworkList(ListID string) (*NetworkListDeleteResponse, error) {

	// Create and execute request
	resp, err := nls.Client.Rclient.R().
		SetResult(NetworkListDeleteResponse{}).
		SetError(NetworkListErrorv2{}).
		Delete(fmt.Sprintf("%s/%s", basePath, ListID))

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		e := resp.Error().(*NetworkListErrorv2)
		if e.Status != 0 {
			return nil, e
		}
	}

	return resp.Result().(*NetworkListDeleteResponse), nil
}

// NetworkListNotification Manage network list subscription
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
func (nls *Netlistv2) NetworkListNotification(action AkamaiSubscription, sub NetworkListSubscription) error {

	var networkListv2 NetworkListv2
	var e NetworkListErrorv2

	// Create and execute request
	_, err := nls.Client.Rclient.R().
		SetResult(&networkListv2).
		SetError(NetworkListErrorv2{}).
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
