# go-edgegrid

Golang based client for interaction with Akamai API services.

## Coverage

This API client package covers most used parts the existing akamai API calls and is updated regularly. Currently the following services are supported:

- [x] Network Lists
- [x] Property APIs
- [ ] Firewall Rule notifications
- [ ] Siteshield
- [ ] Certificate management

To add new/update existing features create a new PR

## Usage
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

// create new Akamai API client
akamaiApi, err := edgegrid.NewClient(nil, apiClientOpts)
```

Once created you will have access to exposed services on `akamaiApi` client object.

## ToDo

- The biggest thing this package still needs is tests :disappointed:

## Issues

- If you have an issue: report it on the [issue tracker](https://github.com/apiheat/go-edgegrid/issues)

## Authors

* RafPe [https://github.com/rafpe]
* Petr Artamonov [https://github.com/partamonov/]

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

## Inspired by
* https://github.com/Comcast/go-edgegrid
