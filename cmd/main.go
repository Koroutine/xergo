package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michealroberts/xergo/pkg/xero"
)

func getEnvironmentVariable(name string) (string, error) {
	value := os.Getenv(name)

	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", name)
	}

	return value, nil
}

func main() {
	XERO_CLIENT_ID, err := getEnvironmentVariable("XERO_CLIENT_ID")

	if err != nil {
		fmt.Println(err)
		return
	}

	XERO_CLIENT_SECRET, err := getEnvironmentVariable("XERO_CLIENT_SECRET")

	if err != nil {
		fmt.Println(err)
		return
	}

	XERO_TENANT_ID, err := getEnvironmentVariable("XERO_TENANT_ID")

	if err != nil {
		fmt.Println(err)
		return
	}

	// https://developer.xero.com/documentation/guides/oauth2/scopes
	scopes := []string{"accounting.contacts", "accounting.attachments", "accounting.transactions"}

	ctx := context.Background()

	client, err := xero.NewClient(
		ctx,
		"https://api.xero.com",
		xero.OAuth2ClientCrendentials{
			ClientId:     XERO_CLIENT_ID,
			ClientSecret: XERO_CLIENT_SECRET,
		},
		xero.Params{
			TenantId: XERO_TENANT_ID,
		},
		scopes,
	)

	if err != nil {
		fmt.Println("Error creating Xero client:", err)
		return
	}

	// Get all invoices for the organisation:
	invoices, err := client.GetInvoices()

	if err != nil {
		fmt.Println("Error getting invoices:", err)
		return
	}

	for _, invoice := range invoices.Invoices {
		fmt.Println(invoice.InvoiceNumber)
	}
}
