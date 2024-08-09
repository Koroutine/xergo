package xero

import (
	"fmt"
	"net/http"
	"net/url"
)

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	PATCH  HTTPMethod = "PATCH"
	DELETE HTTPMethod = "DELETE"
)

func (m HTTPMethod) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	}
	return "unknown"
}

func (c *XeroClient) SetupBaseRequest(method HTTPMethod, entity string) http.Request {
	// Setup the endpoint URL:
	endpoint := url.URL{
		Scheme: c.baseURL.Scheme,
		Host:   c.baseURL.Host,
		Path:   fmt.Sprintf("%s%s", c.baseURL.Path, entity),
	}

	header := http.Header{
		"Accept":           []string{"application/json"},
		"Content-Type":     []string{"application/json"},
		"Xero-Tenant-Id":   []string{c.tenantId},
		"X-Xero-Tenant-Id": []string{c.tenantId},
	}

	// Setup the base request using the endpoint:
	req := http.Request{
		Method: method.String(),
		URL:    &endpoint,
		Header: header,
	}

	return req
}
