package xero

import (
	"encoding/json"
	"fmt"
	"io"
)

type Connection struct {
	ID          string `json:"id"`
	TenantID    string `json:"tenantId"`
	AuthEventId string `json:"authEventId"`
	TenantType  string `json:"tenantType"`
	TenantName  string `json:"tenantName"`
}

type ConnectionsResponse struct {
	Connections []Connection `json:"connections"`
}

func (c *XeroClient) GetConnections() (*ConnectionsResponse, error) {
	req := c.SetupBaseRequest(GET, "/connections")

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("could not list connections, unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	// N.B. The unmarshalling here is done a little different, because the response
	// is a list of connections, not a single object pointing to connections, e.g.:
	// [{}] instead of { "connections": [{}] }
	var connectionsResponse []Connection

	err = json.Unmarshal(body, &connectionsResponse)

	if err != nil {
		return nil, err
	}

	return &ConnectionsResponse{
		Connections: connectionsResponse,
	}, nil
}
