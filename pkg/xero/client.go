package xero

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type XeroClient struct {
	client   *http.Client
	tenantId string
}

type OAuth2ClientCrendentials struct {
	ClientId     string
	ClientSecret string
}

type Params struct {
	TenantId string
}

func NewClient(ctx context.Context, credentials OAuth2ClientCrendentials, params Params, scopes []string) (*XeroClient, error) {
	config := &clientcredentials.Config{
		ClientID:     credentials.ClientId,
		ClientSecret: credentials.ClientSecret,
		TokenURL:     "https://identity.xero.com/connect/token",
		Scopes:       scopes,
	}

	// Sends the "client_id" and "client_secret" in the POST body
	// as application/x-www-form-urlencoded parameters.
	config.AuthStyle = 1

	client := config.Client(ctx)

	return &XeroClient{client: client, tenantId: params.TenantId}, nil
}
