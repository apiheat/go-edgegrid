# go-edgegrid

An Akamai API client enabling Go programs to interact with Akamai in a simple and uniform way

## NOTE

n.v.t

## Coverage

This API client package covers most used parts the existing akamai API calls and is updated regularly. Currently the following services are supported:

- [x] Network Lists
- [ ] Firewall Rule notifications
- [ ] Siteshield
- [ ] Netstorage
- [ ] Certificate management

To add new/update existing features create a new PR

## Usage

```go
import "github.com/RafPe/go-edgegrid"
```

Construct a new Akamai client, then use the various services on the client to
access different parts of the akamai API.

```go
apiClientOpts := &edgegrid.ClientOptions{}
apiClientOpts.ConfigPath =  "/path/to/.edgerc/"
apiClientOpts.ConfigSection = "default"

// create new Akamai API client
akamaiApi := edgegrid.NewClient(nil, apiClientOpts)
```

Passing `nil` into client options will cause it to try and initiate using `ENV VARS`

Some API methods have optional parameters that can be passed.

```go


// create new API client - using ENV VARS
// * AKAMAI_EDGERC_CONFIG
// * AKAMAI_EDGERC_SECTION
akamaiApi := edgegrid.NewClient(nil,nil)

// Set options for working with network lists
opt := edgegrid.ListNetworkListsOptions{
	TypeOflist:        "IP",
	Extended:          true,
	IncludeDeprecated: false,
	IncludeElements:   false,
}

// List all network lists
netLists, resp, err := apiClient.NetworkLists.ListNetworkLists(opt)

if err != nil {
	return err
}
```

### Examples

The [examples](https://github.com/RafPe/go-edgegrid/tree/master/examples) directory
contains a couple for clear examples.

```go
package main

import (
	"log"

	"github.com/RafPe/go-edgegrid"
)

func main() {

	// create new Akamai API client
	akamaiApi := edgegrid.NewClient(nil, "/path/to/.edgerc/", "section-name")

	// create options for new list creation
	newListItems := []string{"1.2.3.4/32", "5.6.7.8/32"}
	newListOpts := edgegrid.CreateNetworkListsOptions{
		Name:        "dummy_delete_1",
		Type:        "IP",
		Description: "",
		List:        newListItems,
	}
	
	newList, err := api.NetworkLists.CreateNetworkList(newListOpts)
	if err != nil {
		log.Fatal(err)
	}
}

```
## ToDo

- The biggest thing this package still needs is tests :disappointed:

## Issues

- If you have an issue: report it on the [issue tracker](https://github.com/RafPe/go-edgegrid/issues)

## Author

* RafPe [https://github.com/rafpe]
* Petr Artamonov [https://github.com/partamonov/]

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

## Inspired by
* https://github.com/Comcast/go-edgegrid
