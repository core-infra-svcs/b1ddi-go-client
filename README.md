# Overview

This library enables the management of BloxOne DDI resources.

The following BloxOne DDI APIs are supported:

- [IP Address Management (Ipamsvc API)](https://csp.infoblox.com/apidoc/?url=https://csp.infoblox.com/apidoc/docs/Ipamsvc)
- [DNS Configuration (DNSConfig API)](https://csp.infoblox.com/apidoc/?url=https://csp.infoblox.com/apidoc/docs/DnsConfig)
- [DNS Data (DnsData API)](https://csp.infoblox.com/apidoc/?url=https://csp.infoblox.com/apidoc/docs/DnsData)
- [Bloxone Cloud Infrastructure Management](https://csp.infoblox.com/apidoc?url=https:/csp.infoblox.com/apidoc/docs/Infrastructure)

# Installation

To install `b1ddi-go-client` use `go get` command:

```bash
go get github.com/infobloxopen/b1ddi-go-client
```

# Usage Guide

## Examples

The following program will print the subnet mask of each subnet in the BloxOne DDI environment:
```go
package main

import (
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	b1cli "github.com/infobloxopen/b1ddi-go-client/client"
)

func main() {
	// Create new go-swagger runtime client
	transport := httptransport.New(
		os.Getenv("B1DDI_HOST"), "api/ddi/v1", nil,
	)

	// Create default auth header for all API requests
	tokenAuth := b1cli.B1DDIAPIKey(os.Getenv("B1DDI_API_KEY"))
	transport.DefaultAuthentication = tokenAuth

	// Create the BloxOne API client
	client := b1cli.NewClient(transport, strfmt.Default)

	// List all subnets using IPAM API client
	subnetList, err := client.IPAddressManagementAPI.Subnet.SubnetList(nil, nil)
	if err != nil {
		panic(err)
	}

	// Print subnet mask for each subnet
	for _, subnet := range subnetList.Payload.Results {
		fmt.Printf("%s/%d\n", *subnet.Address, subnet.Cidr)
	}
}
```
