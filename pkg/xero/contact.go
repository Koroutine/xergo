package xero

import (
	"encoding/json"
	"fmt"
	"io"
)

type Contact struct {
	ContactID    string `json:"ContactID"`
	Name         string `json:"Name"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	EmailAddress string `json:"EmailAddress"`
}

type ContactsResponse struct {
	Contacts []Contact `json:"Contacts"`
}

func (c *XeroClient) GetContacts() (*ContactsResponse, error) {
	req := c.SetupBaseRequest(GET, "/api.xro/2.0/Contacts")

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("could not list contacts, unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var contactsResponse ContactsResponse

	err = json.Unmarshal(body, &contactsResponse)

	if err != nil {
		return nil, err
	}

	return &contactsResponse, nil
}
