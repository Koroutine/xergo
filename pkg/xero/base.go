package xero

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/oklog/ulid/v2"
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

func getIdempotencyKey(tenantId string, entity string) (string, error) {
	// Generate a ULID based on the current timestamp
	now := time.Now().UTC()
	// Generate a new entropy source:
	entropy := ulid.Monotonic(rand.Reader, 0)
	// Generate a ULID:
	id, err := ulid.New(ulid.Timestamp(now), entropy)

	// Oooh dear :(
	if err != nil {
		return "", err
	}

	// e.g., "tenant:invoices:01J4V6F2QYM7K0HTPNGZ7ZY32T"
	return fmt.Sprintf("%s:%s:%s", tenantId, entity, id.String()), nil
}

func (c *XeroClient) SetupBaseRequest(method HTTPMethod, entity string) http.Request {
	// Setup the endpoint URL:
	endpoint := url.URL{
		Scheme: "https",
		Host:   "api.xero.com",
		Path:   fmt.Sprintf("/api.xro/2.0/%s", entity),
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

	key, err := getIdempotencyKey(c.tenantId, entity)

	// Check that the req.Method is in the list of methods that require the "Idempotency-Key" header:
	if err != nil && req.Method != GET.String() {
		// Add the "Idempotency-Key" header to the request:
		// This allows the client to retry the request without the risk of
		// creating duplicate resources, in quick succession.
		// https://developer.xero.com/documentation/guides/idempotent-requests/idempotency
		req.Header.Add("Idempotency-Key", key)
	}

	return req
}
