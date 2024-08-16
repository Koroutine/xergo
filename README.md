# xergo

![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Koroutine/xergo/main?filename=go.mod&label=Go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Koroutine/xergo)](https://goreportcard.com/report/github.com/Koroutine/xergo)
[![xero/ci](https://github.com/Koroutine/xergo/actions/workflows/ci.yaml/badge.svg)](https://github.com/Koroutine/xergo/actions/workflows/ci.yaml)
[![xero/test](https://github.com/Koroutine/xergo/actions/workflows/test.yaml/badge.svg)](https://github.com/Koroutine/xergo/actions/workflows/test.yaml)

A simple oauth2 client for the Xero accounting API, handling authorised and authenticated HTTP requests, as specified in RFC 6749. 

It can additionally grant authorization with Bearer JWT and handle the automatic refreshing of bearer tokens.

It also handles idempotency, sending the `Idempotency-Key` header along with POST, PUT and PACTH requests, ensuring that resources are not duplicated.

### Installation

```
go get github.com/Koroutine/xergo@latest
```

### Usage

```go

import (
  "github.com/Koroutine/xergo/pkg/xero"
)

func main() {
  	// Setup your scopes, matching you OAuth client setup.
	// https://developer.xero.com/documentation/guides/oauth2/scopes
	scopes := []string{"accounting.contacts", "accounting.attachments", "accounting.transactions"}

	ctx := context.Background()

	client, err := xero.NewClient(
		ctx,
		"https://api.xero.com",
		xero.OAuth2ClientCrendentials{
			ClientId:     "<YOUR_XERO_CLIENT_ID>",
			ClientSecret: "<YOUR_XERO_CLIENT_SECRET>",
		},
		xero.Params{
			TenantId: "<YOUR_XERO_TENANT_ID>",
		},
		scopes,
	)
  
  	// Oh, no.
	if err != nil {
		fmt.Println("Error creating Xero client:", err)
		return
	}

	// Get all invoices for the organisation attached to your OAuth client (or Tenant Id).
	_, err = client.GetInvoices()

	if err != nil {
		fmt.Println("Error getting invoices:", err)
		return
	}
}
```

 
