package main

import (
	"fmt"

	edgegrid "github.com/RafPe/go-edgegrid"
)

func main() {

	apiClientOpts := &edgegrid.ClientOptions{}
	apiClientOpts.ConfigPath = "/path/to/.edgerc/"
	apiClientOpts.ConfigSection = "default"

	// create new Akamai API client
	api := edgegrid.NewClient(nil, apiClientOpts)

	// Set options for working with network lists
	opt := edgegrid.ListNetworkListsOptions{
		TypeOflist:        "IP",
		Extended:          true,
		IncludeDeprecated: false,
		IncludeElements:   false,
	}

	// List all network lists
	allLists, resp, err := api.NetworkLists.ListNetworkLists(opt)
	if err != nil {
		fmt.Println(err)
	}

	// We return response from our API call
	fmt.Println(resp)

	for i, c := range *allLists {
		fmt.Println(i, c.Name)
	}

	// List single network list
	singleList, resp, err := api.NetworkLists.GetNetworkList("LIST_ID", opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(singleList.Name, singleList.NumEntries, singleList.Account, singleList.ProductionActivationStatus, singleList.StagingActivationStatus)

	// create dummy network list
	newListItems := []string{"1.2.3.4/32", "5.6.7.8/32"}
	newListOpts := edgegrid.CreateNetworkListOptions{
		Name:        "dummy_delete_1",
		Type:        "IP",
		Description: "",
		List:        newListItems,
	}
	newList, resp, err := api.NetworkLists.CreateNetworkList(newListOpts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newList.UniqueID)

	// Modify existing network list
	newListItemsToMod := []string{"4.4.3.4/32"}
	editListOpts := edgegrid.CreateNetworkListOptions{
		List: newListItemsToMod,
	}

	modifyResp, resp, err := api.NetworkLists.AddNetworkListItems("LIST_ID", editListOpts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(modifyResp.Message)

	// Remove item from network list
	removeItemResp, resp, err := api.NetworkLists.RemoveNetworkListItem("LIST_ID", "4.4.3.4/32")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(removeItemResp.Message)

	// Activate a network list
	activateListOpts := edgegrid.ActivateNetworkListOptions{
		SiebelTicketID:         "test-01",
		NotificationRecipients: []string{},
		Comments:               "activated by new API client",
	}

	activateList, resp, err := api.NetworkLists.ActivateNetworkList("LIST_ID", edgegrid.Staging, activateListOpts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(activateList.ActivationStatus)

	activationStatus, resp, err := api.NetworkLists.GetNetworkListActivationStatus("LIST_ID", edgegrid.Staging)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(activationStatus.ActivationStatus)
}
