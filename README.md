# xergo
A simple oauth2 client for the Xero accounting API

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

 
