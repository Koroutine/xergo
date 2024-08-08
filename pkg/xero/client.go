package xero

import (
	"context"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
)

type XeroClient struct {
	client *http.Client
}

type OAuth2ClientCrendentials struct {
	ClientId     string
	ClientSecret string
}

func NewClient(ctx context.Context, credentials OAuth2ClientCrendentials, scopes []string) (*XeroClient, error) {
	config := &clientcredentials.Config{
		ClientID:     credentials.ClientId,
		ClientSecret: credentials.ClientSecret,
		TokenURL:     "https://identity.xero.com/connect/token",
		Scopes:       scopes,
	}

	client := config.Client(ctx)

	return &XeroClient{client: client}, nil
}
