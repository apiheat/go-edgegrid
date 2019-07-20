package netlistv2

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
		WithLogVerbosity("debug").
		WithLocalTesting(true).
		WithScheme("http").
		WithTestingURL(targetURL)

	fmt.Println(creds, cfg)

	// Create new client
	client := New(cfg)

	return client
}

//TestListNetworkLists checks if listing all network lists works
func TestListNetworkLists(t *testing.T) {

	response := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		body, _ := ioutil.ReadAll(r.Body)
		assert.Empty(t, string(body), "Request body should be empty")
		assert.Equal(t, "GET", r.Method, "Method should be GET")

		fmt.Fprintln(w, response)
	}))

	// Init API client
	apiClient := setupEdgeClient(server.URL)

	listNetListOptsv2 := ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

	apiResp, reqErr := apiClient.ListNetworkLists(listNetListOptsv2)

	if reqErr != nil {
		fmt.Println(reqErr)
	}

	var expectedType *NetworkListsv2

	assert.IsType(t, expectedType, apiResp)
	defer server.Close()
}

// //TestGetNetworkListById checks if listing specific network list by-id works
// func TestGetNetworkListById(t *testing.T) {

// 	response := `{"networkLists":[{"networkListType":"networkListResponse","accessControlGroup":"KSD\nwith ION 3-13H1234","name":"General List","elementCount":3011,"syncPoint":22,"type":"IP","uniqueId":"25614_GENERALLIST","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/25614_GENERALLIST"},"statusInProduction":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/25614_GENERALLIST/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/25614_GENERALLIST","method":"PUT"}}},{"networkListType":"networkListResponse","account":"Kona\nSecurity Engineering","accessControlGroup":"Top-Level Group: 3-12DAF123","name":"Ec2 Akamai Network List","elementCount":235,"readOnly":true,"syncPoint":65,"type":"IP","uniqueId":"1024_AMAZONELASTICCOMPUTECLOU","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU"},"statusInProduction":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/1024_AMAZONELASTICCOMPUTECLOU","method":"PUT"}}},{"networkListType":"networkListResponse","accessControlGroup":"KSD\nTest - 3-13H5523","name":"GeoList_1913New","elementCount":16,"syncPoint":2,"type":"GEO","uniqueId":"26732_GEOLIST1913","links":{"activateInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/activate","method":"POST"},"activateInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/activate","method":"POST"},"appendItems":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/append","method":"POST"},"retrieve":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913"},"statusInProduction":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/PRODUCTION/status"},"statusInStaging":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913/environments/STAGING/status"},"update":{"href":"/network-list/v2/network-lists/26732_GEOLIST1913","method":"PUT"}}}],"links":{"create":{"href":"/network-list/v2/network-lists/","method":"POST"}}}`
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		body, _ := ioutil.ReadAll(r.Body)
// 		assert.Empty(t, string(body), "Request body should be empty")
// 		assert.Equal(t, "GET", r.Method, "Request method should be GET")
// 		assert.Contains(t, r.URL.String(), "123_TEST", "Request URL should contain list ID")
// 		fmt.Fprintln(w, response)
// 	}))

// 	// Init API client
// 	apiClient := setupEdgeClient(server.URL)

// 	listNetListOptsv2 := ListNetworkListsOptionsv2{}
// 	listNetListOptsv2.Search = "" // Since we are listing all we do not filter results

// 	apiResp, reqErr := apiClient.GetNetworkList("123_TEST", listNetListOptsv2)

// 	if reqErr != nil {
// 		fmt.Println(reqErr)
// 	}

// 	var expectedType *NetworkListv2

// 	assert.IsType(t, expectedType, apiResp)
// 	defer server.Close()
// }

//TestAddNetworNetworkListElement adds network elements to list
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
			bodyFromHeader, err := base64.StdEncoding.DecodeString(req.Header["X-Test-Body"][0])
			if assert.NoError(t, err) {
				assert.NotEmpty(t, string(bodyFromHeader), "Request body should not be empty")
				assert.Contains(t, req.URL.String(), "345_BOTLIST", "Request URL should contain list ID")
				assert.Equal(t, "POST", req.Method, "Method should be POST")
				assert.Equal(t, "{\"list\":[\"1.2.3.4/32\"]}", string(bodyFromHeader), "Request body should contain list of addresses to activate")
			}

			// body, err := ioutil.ReadAll(r.Body)
			// if err != nil {
			// 	log.Printf("Error reading body: %v", err)
			// 	http.Error(w, "can't read body", http.StatusBadRequest)
			// 	return
			// }
			// Work / inspect body. You may even modify it!

			// // And now set a new body, which will simulate the same data we read:
			// r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			// // body, err := ioutil.ReadAll(req.Body)

			// x, _ := httputil.DumpRequest(req, true)
			// assert.Equal(t, "{\"list\":[\"1.2.3.5/32\"]}", string(x), "Request body should contain list of addresses to activate")

			// // if assert.NoError(t, err) {
			// 	// assert.Empty(t, string(body), "Request body should be empty")
			// 	assert.Contains(t, req.URL.String(), "345_BOTLIST", "Request URL should contain list ID")
			// 	assert.Equal(t, "POST", req.Method, "Method should be POST")

			// // 	assert.Equal(t, "{\"list\":[\"1.2.3.5/32\"]}", string(x), "Request body should contain list of addresses to activate")
			// assert.Equal(t, "{\"list\":[\"1.2.3.5/32\"]}", string(bodyBytes), "Request body should contain list of addresses to activate")
			// // }

			resp := httpmock.NewStringResponse(201, responseJSON)

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
