# go-edgegrid

Golang based client for interaction with Akamai API services.

## Step by step using the client
The client has been created with simplicity in mind. After several iterations the best approach seems to be allow modularity by seperating credentials/config and services and combining them into `apiClient` object therefore allowing full control of customisation

### Imports
Import client and all services you plan to use in your code. In example below we import network list service along with client
```go
import (
	"github.com/apiheat/go-edgegrid/edgegrid"
	
	// ..... all other services u may also need
	"github.com/apiheat/go-edgegrid/service/netlistv2"
)
```

### Credentials
Create new credentials object using on of the below methods:
* credentials file
	```go
	creds, err := edgegrid.NewCredentials().FromFile("/Users/rafpe/.edgerc").Section("sample")
	```
* JSON string
	```go
	creds, err := edgegrid.NewCredentials().FromJSON(`{ "args":"xxx"}`)
	```
* ENV variables
	```go
	creds, err := eauth.NewCredentials().FromEnv()
	```

### Config
Create config object which defines client behaviour. Define options which u require.
```go
	config := edgegrid.NewConfig().
		WithCredentials(creds). 					// Required
		WithLogVerbosity("info").					// Optional
		WithLocalTesting(true).						// Optional
		WithScheme("http").							// Optional
		WithTestingURL("http://localhost.test").	// Optional
		WithRequestDebug(true)						// Optional
```
### Support for Account Switch Key ( manage multiple accounts )
Client in version starting from `v5.x.x` supports *account switch key* which allows you to manage multiple accounts with single credentials.

* Specify when initialising client

    ```go
	// 2 - Config using credentials
	config := edgegrid.NewConfig().
		WithCredentials(creds).
    	WithAccountSwitchKey("1-231-213123")
    ```


More information can be found under the following link https://learn.akamai.com/en-us/learn_akamai/getting_started_with_akamai_developers/developer_tools/accountSwitch.html

### Example 
Below is minimalistic example of all the steps required to get the client running.

```go
package main

import (
	"fmt"

	"github.com/apiheat/go-edgegrid/edgegrid"
	"github.com/apiheat/go-edgegrid/service/netlistv2"
)

func main() {

	// 1 - Credentials - multiple way to obrain credentials
	creds, err := edgegrid.NewCredentials().FromFile("/Users/rafpe/.edgerc").Section("sample")

	// 2 - Config using credentials
	config := edgegrid.NewConfig().
		WithCredentials(creds)

	listNetListOptsv2 := netlistv2.ListNetworkListsOptionsv2{}
	listNetListOptsv2.Search = ""

	// 3 - Service using config
	apiNetlistv2 := netlistv2.New(config)

	// 4 - Actions using service
	res, err := apiNetlistv2.ListNetworkLists(listNetListOptsv2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}

```

### Debugging
The debug use `WithLogVerbosity(<level>)` ( *optional* part of config object )  where `<level>` can be lower case string of `debug` | `warn` |  `info` | `error` | `fatal` | `panic`


## Development
 - More info to come 

### Tests

- The biggest thing this package still needs is tests :disappointed:

### Issues

- If you have an issue: report it on the [issue tracker](https://github.com/apiheat/go-edgegrid/issues)



## Authors

* RafPe [https://github.com/rafpe]
* Petr Artamonov [https://github.com/partamonov/]

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

## Inspired by
* https://github.com/Comcast/go-edgegrid







