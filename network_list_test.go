package edgegrid

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListNetworkLists(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	// Set options for working with network lists
	opt := ListNetworkListsOptions{
		TypeOflist:        "IP",
		Extended:          true,
		IncludeDeprecated: false,
		IncludeElements:   false,
	}

	apiURI := fmt.Sprintf("%s?listType=%s&extended=%t&includeDeprecated=%t&includeElements=%t",
		apiPaths["network_list"],
		opt.TypeOflist,
		opt.Extended,
		opt.IncludeDeprecated,
		opt.IncludeElements)

	fmt.Println(fmt.Sprintf("API URL: %s", apiURI))

	mux.HandleFunc(apiURI, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{whatever}`)
	})

	fmt.Println(fmt.Sprintf("Base URL: %s", client.baseURL.String()))

	allLists, resp, err := client.NetworkLists.ListNetworkLists(opt)
	if err != nil {
		fmt.Println("error occured")
	}

	fmt.Println(fmt.Sprintf("Statuscode is : %v", resp.StatusCode))
	fmt.Println(fmt.Sprintf("Status is : %s", resp.Status))

	want := []AkamaiNetworkList{{UniqueID: "aa"}, {UniqueID: "bbb"}}
	if !reflect.DeepEqual(want, allLists) {
		t.Errorf("NetworkLists.ListNetworkLists returned %+v, want %+v", allLists, want)
	}
}
