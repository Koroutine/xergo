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
	case DELETE:
		return "DELETE"
	}
	return "unknown"
}

func SetupBaseRequest(method HTTPMethod, entity string) http.Request {
	// Setup the endpoint URL:
	endpoint := url.URL{
		Scheme: "https",
		Host:   "api.xero.com",
		Path:   fmt.Sprintf("/api.xro/2.0/%s", entity),
	}

	header := http.Header{
		"Accept":       []string{"application/json"},
		"Content-Type": []string{"application/json"},
	}

	// Setup the base request using the endpoint:
	req := http.Request{
		Method: method.String(),
		URL:    &endpoint,
		Header: header,
	}

	return req
}
