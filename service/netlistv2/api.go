package netlistv2

import (
	"fmt"
	"strconv"
)

// ListNetworkLists List all configured Network Lists for the authenticated user.
// Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlists
func (nls *Netlistv2) ListNetworkLists(opts ListNetworkListsOptionsv2) (*NetworkListsv2, error) {

	// Create new request
	req := nls.REST.NewRequest()

	response, err := req.SetQueryParams(map[string]string{
		"includeElements": strconv.FormatBool(opts.IncludeElements),
		"extended":        strconv.FormatBool(opts.Extended),
		"search":          opts.Search,
	}).SetHeaders(
		map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   nls.Config.UserAgent,
		}).SetResult(&NetworkListsv2{}).SetError(&NetworkListErrorv2{}).Get("/network-list/v2/network-lists")

	fmt.Println("Error is .... ")
	fmt.Println(err)

	fmt.Println("xxx...")
	return response.Result().(*NetworkListsv2), response.Error().(*NetworkListErrorv2)

}

// // CreateNetworkList Create a new network list
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
// func (nls *Netlistv2) CreateNetworkList(opts NetworkListsOptionsv2) (*NetworkListv2, *ClientResponse, error) {

// 	qParams := QStrNetworkList{}
// 	path := NetworkListPathV2

// 	var respStruct *NetworkListv2
// 	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, opts, nil)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return respStruct, resp, netListError
// 	}

// 	return respStruct, resp, err
// }

// // GetNetworkList Gets a specific network list
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#getlist
// func (nls *Netlistv2) GetNetworkList(ListID string, opts ListNetworkListsOptionsv2) (*NetworkListv2, *ClientResponse, error) {

// 	qParams := QStrNetworkList{
// 		Extended:        opts.Extended,
// 		IncludeElements: opts.IncludeElements,
// 	}

// 	path := fmt.Sprintf("%s/%s", NetworkListPathV2, ListID)

// 	var respStruct *NetworkListv2
// 	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return respStruct, resp, netListError
// 	}

// 	return respStruct, resp, err

// }

// // AppendListNetworkList Adds items to network list
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html#postlists
// func (nls *Netlistv2) AppendListNetworkList(ListID string, opts NetworkListsOptionsv2) (*NetworkListv2, *ClientResponse, error) {

// 	qParams := QStrNetworkList{}
// 	path := fmt.Sprintf("%s/%s/append", NetworkListPathV2, ListID)

// 	var respStruct *NetworkListv2
// 	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, opts, nil)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return respStruct, resp, netListError
// 	}

// 	return respStruct, resp, err
// }

// // RemoveNetworkListElement Removes network list element
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
// func (nls *Netlistv2) RemoveNetworkListElement(ListID, element string) (*NetworkListv2, *ClientResponse, error) {

// 	qParams := QStrNetworkList{
// 		Element: element,
// 	}
// 	path := fmt.Sprintf("%s/%s/elements", NetworkListPathV2, ListID)

// 	var respStruct *NetworkListv2
// 	resp, err := nls.client.makeAPIRequest(http.MethodDelete, path, qParams, &respStruct, nil, nil)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return respStruct, resp, netListError
// 	}

// 	return respStruct, resp, err
// }

// // ActivateNetworkList Activates network list on specified network ( PRODUCTION or STAGING )
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
// func (nls *Netlistv2) ActivateNetworkList(ListID string, targetEnv AkamaiEnvironment, opts NetworkListActivationOptsv2) (*NetworkListActivationStatusv2, *ClientResponse, error) {

// 	qParams := QStrNetworkList{}
// 	path := fmt.Sprintf("%s/%s/environments/%s/activate", NetworkListPathV2, ListID, targetEnv)

// 	var respStruct *NetworkListActivationStatusv2
// 	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, &respStruct, opts, nil)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return respStruct, resp, netListError
// 	}

// 	return respStruct, resp, err
// }

// // GetNetworkListActStatus Gets activation network list status on specified network ( PRODUCTION or STAGING )
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
// func (nls *Netlistv2) GetNetworkListActStatus(ListID string, targetEnv AkamaiEnvironment) (*NetworkListActivationStatusv2, *ClientResponse, error) {

// 	qParams := QStrNetworkList{}
// 	path := fmt.Sprintf("%s/%s/environments/%s/status", NetworkListPathV2, ListID, targetEnv)

// 	var respStruct *NetworkListActivationStatusv2
// 	resp, err := nls.client.makeAPIRequest(http.MethodGet, path, qParams, &respStruct, nil, nil)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return respStruct, resp, netListError
// 	}

// 	return respStruct, resp, err
// }

// // DeleteNetworkList Remove network list element
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
// func (nls *Netlistv2) DeleteNetworkList(ListID string) (*NetworkListDeleteResponse, *ClientResponse, error) {

// 	qParams := QStrNetworkList{}
// 	path := fmt.Sprintf("%s/%s", NetworkListPathV2, ListID)

// 	var respStruct *NetworkListDeleteResponse
// 	resp, err := nls.client.makeAPIRequest(http.MethodDelete, path, qParams, &respStruct, nil, nil)

// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return respStruct, resp, netListError
// 	}

// 	return respStruct, resp, err
// }

// // NetworkListNotification Manage network list subscription
// // Akamai API docs: https://developer.akamai.com/api/cloud_security/network_lists/v2.html
// func (nls *Netlistv2) NetworkListNotification(action AkamaiSubscription, sub NetworkListSubscription) (*ClientResponse, error) {

// 	qParams := QStrNetworkList{}
// 	path := fmt.Sprintf("/network-list/v2/notifications/%s", action)

// 	resp, err := nls.client.makeAPIRequest(http.MethodPost, path, qParams, nil, sub, nil)
// 	if err != nil {
// 		return resp, err
// 	}

// 	if resp.Response.StatusCode >= http.StatusBadRequest {
// 		netListError := &NetworkListErrorv2{}
// 		err := json.Unmarshal([]byte(resp.Body), &netListError)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		return resp, err
// 	}

// 	return resp, err
// }
