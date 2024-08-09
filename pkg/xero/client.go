package xero

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/oauth2/clientcredentials"
)

type XeroClient struct {
	client   *http.Client
	baseURL  *url.URL
	tenantId string
}

type OAuth2ClientCrendentials struct {
	ClientId     string
	ClientSecret string
}

type Params struct {
	TenantId string
}

func NewClient(ctx context.Context, baseURL string, credentials OAuth2ClientCrendentials, params Params, scopes []string) (*XeroClient, error) {
	base, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}

	// Add the /api.xro/2.0/ path to the base URL:
	endpoint := url.URL{
		Scheme: base.Scheme,
		Host:   base.Host,
	}

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

	return &XeroClient{
		client:   client,
		baseURL:  &endpoint,
		tenantId: params.TenantId,
	}, nil
}
