package main

import (
	"fmt"

	edgegrid "github.com/RafPe/go-edgegrid"
)

func main() {

	// create new API client
	api := edgegrid.NewClient(nil, "~/.edgerc", "dummy")

	// Set options for working with network lists
	opt := edgegrid.ListNetworkListsOptions{
		TypeOflist:        "IP",
		Extended:          true,
		IncludeDeprecated: false,
		IncludeElements:   false,
	}

	// List all network lists
	allLists, _ := api.NetworkLists.ListNetworkLists(opt)
	for i, c := range allLists {
		fmt.Println(i, c.Name)
	}

	// List single network list
	singleList, _ := api.NetworkLists.GetNetworkList("LIST_ID", opt)
	fmt.Println(singleList.Name, singleList.NumEntries, singleList.Account, singleList.ProductionActivationStatus, singleList.StagingActivationStatus)

	// create dummy network list
	newListItems := []string{"1.2.3.4/32", "5.6.7.8/32"}
	newListOpts := edgegrid.CreateNetworkListsOptions{
		Name:        "dummy_delete_1",
		Type:        "IP",
		Description: "",
		List:        newListItems,
	}
	newList, _ := api.NetworkLists.CreateNetworkList(newListOpts)

	// Modify existing network list
	newListItems := []string{"4.4.3.4/32"}
	editListOpts := edgegrid.CreateNetworkListOptions{
		List: newListItems,
	}

	api.NetworkLists.AddNetworkListItems("LIST_ID", editListOpts)

	// Remove item from network list
	api.NetworkLists.RemoveNetworkListItem("LIST_ID", "4.4.3.4/32")

	// Activate a network list
	activateListOpts := edgegrid.ActivateNetworkListOptions{
		SiebelTicketID:         "test-01",
		NotificationRecipients: []string{},
		Comments:               "activated by new API client",
	}
	activateList, _ := api.NetworkLists.ActivateNetworkList("LIST_ID", edgegrid.Staging, activateListOpts)

	activationStatus, _ := api.NetworkLists.GetNetworkListActivationStatus("LIST_ID", edgegrid.Staging)
	fmt.Println(activationStatus.ActivationStatus)
}
