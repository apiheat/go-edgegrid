package netlistv2

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/apiheat/go-edgegrid/edgegrid"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

//setupEdgeClient prepares and inits client for making all calls towards Akamai's APIs
func setupEdgeClient(baseURL string) *Netlistv2 {

	var targetURL string

	if baseURL == "" {
		targetURL = "http://test.local"
	} else {
		targetURL = baseURL
	}

	// Get credentials
	creds, err := edgegrid.NewCredentials().FromJSON(`{ "client_secret": "kljwekfjf", "host": "akab-k2112.31k23jl1k23.luna.akamaiapis.net", "access_token": "akab-l12h3iu123y923huk-4uc54n5xmwhqu4zh", "client_token": "akab-90821u3hkjbnmk-jkhg" }`)
	if err != nil {
		fmt.Println(err)
	}

	// Create configuration and specify some of the configuration items
	cfg := edgegrid.NewConfig().
		WithCredentials(creds).
		WithLogVerbosity("info").
		WithLocalTesting(true).
		WithScheme("http").
		WithTestingURL(targetURL)

	// Create new client
	client := New(cfg)

	return client
}

// To test:
// * list all [x]
// * list by id [x]
// * add network list element [x]
// * remove network list element [x]
// * create network list [x]
// * activate network list []

func TestListNetworkLists(t *testing.T) {
	//--Init API client
	apiClient := setupEdgeClient("")
	responseJSON := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`

	httpmock.ActivateNonDefault(apiClient.Client.Rclient.GetClient())
	defer httpmock.DeactivateAndReset()

	// mock APIs
	httpmock.RegisterResponder("GET", "http://test.local/network-list/v2/network-lists",
		func(req *http.Request) (*http.Response, error) {

			assert.Nil(t, req.Body, "Request body should be nil")
			assert.Equal(t, "GET", req.Method, "Request method should be GET")

			resp := httpmock.NewStringResponse(200, responseJSON)
			resp.Header.Add("Content-Type", "application/json")

			return resp, nil
		})

	listNetListOptsv2 := ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

	var expectedType *NetworkListsv2

	apiResp, err := apiClient.ListNetworkLists(listNetListOptsv2)
	if assert.NoError(t, err) {
		assert.IsType(t, expectedType, apiResp)
	}

}

func TestGetNetworkListById(t *testing.T) {
	//--Init API client
	apiClient := setupEdgeClient("")
	responseJSON := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`

	httpmock.ActivateNonDefault(apiClient.Client.Rclient.GetClient())
	defer httpmock.DeactivateAndReset()

	// mock to list out the add network list elements
	httpmock.RegisterResponder("GET", "http://test.local/network-list/v2/network-lists/123_TEST",
		func(req *http.Request) (*http.Response, error) {

			assert.Nil(t, req.Body, "Request body should be nil")
			assert.Equal(t, "GET", req.Method, "Request method should be GET")
			assert.Contains(t, req.URL.String(), "123_TEST", "Request URL should contain list ID")

			resp := httpmock.NewStringResponse(200, responseJSON)
			resp.Header.Add("Content-Type", "application/json")

			return resp, nil
		})

	listNetListOptsv2 := ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

	var expectedType *NetworkListv2

	apiResp, err := apiClient.GetNetworkList("123_TEST", listNetListOptsv2)
	if assert.NoError(t, err) {
		assert.IsType(t, expectedType, apiResp)
	}
}

func TestAddNetworNetworkListElement(t *testing.T) {

	//--Init API client
	apiClient := setupEdgeClient("")
	responseJSON := `{"name":"Ec2 Akamai Network List","uniqueId":"345_BOTLIST","syncPoint":65,"type":"IP","networkListType":"networkListResponse","account":"Kona Security Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","elementCount":13,"readOnly":true,"list":["1.2.3.4/32","13.126.0.0/15","13.210.0.0/15","13.228.0.0/15","13.230.0.0/15","13.232.0.0/14","13.236.0.0/14","13.250.0.0/15","13.54.0.0/15","13.56.0.0/16","13.57.0.0/16","13.58.0.0/15","174.129.0.0/16"],"links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}}`

	httpmock.ActivateNonDefault(apiClient.Client.Rclient.GetClient())
	defer httpmock.DeactivateAndReset()

	// mock to list out the add network list elements
	httpmock.RegisterResponder("POST", "http://test.local/network-list/v2/network-lists/345_BOTLIST/append",
		func(req *http.Request) (*http.Response, error) {

			// Decoding may return an error - hence we assert it
			body, err := ioutil.ReadAll(req.Body)

			if assert.NoError(t, err) {
				assert.NotEmpty(t, string(body), "Request body should not be empty")
				assert.Contains(t, req.URL.String(), "345_BOTLIST", "Request URL should contain list ID")
				assert.Equal(t, "POST", req.Method, "Method should be POST")
				assert.Equal(t, "{\"list\":[\"1.2.3.4/32\"]}", string(body), "Request body should contain list of addresses to activate")
			}

			resp := httpmock.NewStringResponse(201, responseJSON)
			resp.Header.Add("Content-Type", "application/json")

			return resp, nil
		})

	//--Modify existing network list
	itemsToAdd := []string{"1.2.3.4/32"}
	editListOpts := NetworkListsOptionsv2{
		List: itemsToAdd,
	}

	var expectedType *NetworkListv2

	addItemRequest, err := apiClient.AddNetworkListElement("345_BOTLIST", editListOpts)
	if assert.NoError(t, err) {
		assert.IsType(t, expectedType, addItemRequest)
	}

}

func TestCreateNetworkList(t *testing.T) {
	//--Init API client
	apiClient := setupEdgeClient("")
	responseJSON := `{"name":"name-of-netlist","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","syncPoint":65,"type":"IP","networkListType":"networkListResponse","account":"Kona Security Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","elementCount":13,"readOnly":true,"list":["13.125.0.0/16","13.126.0.0/15","13.210.0.0/15","13.228.0.0/15","13.230.0.0/15","13.232.0.0/14","13.236.0.0/14","13.250.0.0/15","13.54.0.0/15","13.56.0.0/16","13.57.0.0/16","13.58.0.0/15","174.129.0.0/16"],"links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}}`

	httpmock.ActivateNonDefault(apiClient.Client.Rclient.GetClient())
	defer httpmock.DeactivateAndReset()

	// mock APIs
	httpmock.RegisterResponder("POST", "http://test.local/network-list/v2/network-lists",
		func(req *http.Request) (*http.Response, error) {

			body, err := ioutil.ReadAll(req.Body)
			if assert.NoError(t, err) {
				assert.Equal(t, "{\"name\":\"name-of-netlist\",\"type\":\"IP\",\"description\":\"desc-by-test\"}", string(body), "Request body should be empty")
				assert.Equal(t, "POST", req.Method, "Request method should be POST")
			}

			resp := httpmock.NewStringResponse(201, responseJSON)
			resp.Header.Add("Content-Type", "application/json")

			return resp, nil

		})

	//--Create options
	newNetworkListOpst := NetworkListsOptionsv2{}
	newNetworkListOpst.Description = "desc-by-test"
	newNetworkListOpst.Name = "name-of-netlist"
	newNetworkListOpst.Type = "IP"

	var expectedType *NetworkListv2

	//--Call api
	apiResp, err := apiClient.CreateNetworkList(newNetworkListOpst)
	if assert.NoError(t, err) {
		assert.IsType(t, expectedType, apiResp)
	}
}

func TestRemoveNetworkListItem(t *testing.T) {
	//--Init API client
	apiClient := setupEdgeClient("")
	responseJSON := `{"name":"Ec2 Akamai Network List","uniqueId":"345_BOTLIST","syncPoint":65,"type":"IP","networkListType":"networkListResponse","account":"Kona Security Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","elementCount":13,"readOnly":true,"list":["13.126.0.0/15","13.210.0.0/15","13.228.0.0/15","13.230.0.0/15","13.232.0.0/14","13.236.0.0/14","13.250.0.0/15","13.54.0.0/15","13.56.0.0/16","13.57.0.0/16","13.58.0.0/15","174.129.0.0/16"],"links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}}`

	httpmock.ActivateNonDefault(apiClient.Client.Rclient.GetClient())
	defer httpmock.DeactivateAndReset()

	// mock APIs
	httpmock.RegisterResponder("DELETE", "http://test.local/network-list/v2/network-lists/345_BOTLIST/elements",
		func(req *http.Request) (*http.Response, error) {

			assert.Equal(t, "DELETE", req.Method, "Request method should be DELETE")
			assert.Contains(t, req.URL.String(), "345_BOTLIST/elements", "Request URL should contain list ID")

			// // body, err := ioutil.ReadAll(req.Body)
			// body := ""
			// if assert.NoError(t, err) {

			// }

			resp := httpmock.NewStringResponse(201, responseJSON)
			resp.Header.Add("Content-Type", "application/json")

			return resp, nil

		})

	//--Expect result type
	var expectedType *NetworkListv2

	//--Init API client
	apiResp, err := apiClient.RemoveNetworkListElement("345_BOTLIST", "1.2.3.4/32")
	if err != nil {
		fmt.Println(err)
	}

	if assert.NoError(t, err) {
		assert.IsType(t, expectedType, apiResp)
	}
}

func TestActivateNetworkList(t *testing.T) {
	//--Init API client
	apiClient := setupEdgeClient("")
	responseJSON := `{"activationId":12345,"activationComments":"test-activation","activationStatus":"PENDING_ACTIVATION","syncPoint":5,"uniqueId":"345_BOTLIST","fast":false,"dispatchCount":1,"links":{"appendItems":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/append","method":"POST"},"retrieve":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"syncPointHistory":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST/sync-points/5/history"},"update":{"href":"/networklist-api/rest/v2/network-lists/25614_GENERALLIST","method":"PUT"},"activationDetails":{"href":"/network-list/v2/network-lists/activations/12345/"}}}`

	httpmock.ActivateNonDefault(apiClient.Client.Rclient.GetClient())
	defer httpmock.DeactivateAndReset()

	// mock APIs
	httpmock.RegisterResponder("POST", "http://test.local/network-list/v2/network-lists/345_BOTLIST/environments/production/activate",
		func(req *http.Request) (*http.Response, error) {

			body, err := ioutil.ReadAll(req.Body)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, string(body), "Request body should not be empty")
				assert.Equal(t, `{"comments":"test-activation","notificationRecipients":["dummy@mailinator.com"],"fast":true}`, string(body), "Request body should contain list of addresses to activate")
				assert.Equal(t, "POST", req.Method, "Request method should be POST")
				assert.Contains(t, req.URL.String(), "/network-list/v2/network-lists/345_BOTLIST/environments/production/activate", "Request URL should target activation env with list ID")

			}

			resp := httpmock.NewStringResponse(201, responseJSON)
			resp.Header.Add("Content-Type", "application/json")

			return resp, nil

		})

	var expectedType *NetworkListActivationStatusv2

	//--Init API client
	actNetworkListOpts := NetworkListActivationOptsv2{
		Comments:               "test-activation",
		Fast:                   true,
		NotificationRecipients: []string{"dummy@mailinator.com"},
	}
	apiResp, err := apiClient.ActivateNetworkList("345_BOTLIST", edgegrid.Production, actNetworkListOpts)
	if assert.NoError(t, err) {
		assert.IsType(t, expectedType, apiResp)
	}

}
