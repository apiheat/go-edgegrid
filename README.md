# go-edgegrid

Golang based client for interaction with Akamai API services.

## Coverage

This go-edgegrid API client package covers in most cases complete APIs. Below you can see highlight of services supported by the client:

|            Resource                |            Coverage                  |
|------------------------------------|------------------------------------|
|  Adaptive Acceleration             |  partial  |
|  Network Lists                     | complete |
|  Property APIs ( PAPI )                    | partial |
|  Identity Management - user                   | partial |
|  Identity Management - API                   | partial |
|  Firewall Rule notifications                    | complete |
|  Siteshield                    | complete |

To add new/update existing features create a new PR

## Using go-edgegrid in your code

### client
To start using the client you just need to reference package within your code.

```go
import "github.com/apiheat/go-edgegrid"
```

Construct a new Akamai client, then use the various services on the client to
access different parts of the akamai API.

```go
apiClientOpts := &edgegrid.ClientOptions{}
apiClientOpts.ConfigPath =  "/path/to/.edgerc/"
apiClientOpts.ConfigSection = "default"
apiClientOpts.DebugLevel = "warn"


// create new Akamai API client
akamaiApi, err := edgegrid.NewClient(nil, apiClientOpts)
```

The debug property `apiClientOpts.DebugLevel` is *optional* and can be lower case string of `debug warn info error fatal panic`


Once created you will have access to exposed services on `akamaiApi` client object.

### Support for Account Switch Key ( manage multiple accounts )
Client in version starting from `v5.x.x` supports *account switch key* which allows you to manage multiple accounts with single credentials.

* Specify when initialising client

    ```golang
    // if using to manage multiple accounts
    apiClientOpts.AccountSwitchKey = "1-231-213123"
    ```

* Using exposed functions to control *account switch key*

    ```golang
    // EnableAccountSwitchKey instructs client to use ASK
    EnableAccountSwitchKey()

    // DisableAccountSwitchKey instructs client to not use ASK
    DisableAccountSwitchKey()

    // SetAccountSwitchKey instructs client to not use ASK
    SetAccountSwitchKey(accountSwitchKey string)
    ```

Current setup of the client client will use account switch key if value for it have been provided during the `init`.You can control if account switch key is enabled via exposed methods to enable or disable the account switch key

More information can be found under the following link https://learn.akamai.com/en-us/learn_akamai/getting_started_with_akamai_developers/developer_tools/accountSwitch.html

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



## go-edgegrid - refactored
```go
package main

import (
	"fmt"

	eauth "github.com/apiheat/go-edgegrid/edgegrid/edgeauth"
)

func main() {
	x, _ := eauth.NewCredentials().FromFile("/Users/rafpe/.edgerc").Section("abc")
	x, _ := eauth.NewCredentials().FromEnv()
	x, _ := eauth.NewCredentials().FromJSON(`{ "args":"xxx"}`)

	fmt.Println(x)
}
```


