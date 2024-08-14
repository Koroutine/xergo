package xero

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ContactBase struct {
	Name                      string    `json:"Name"`
	FirstName                 string    `json:"FirstName"`
	LastName                  string    `json:"LastName"`
	EmailAddress              string    `json:"EmailAddress"`
	BankAccountDetails        string    `json:"BankAccountDetails,omitempty"`
	CompanyNumber             string    `json:"CompanyNumber,omitempty"`
	TaxNumber                 string    `json:"TaxNumber,omitempty"`
	AccountsReceivableTaxType string    `json:"AccountsReceivableTaxType,omitempty"`
	AccountsPayableTaxType    string    `json:"AccountsPayableTaxType,omitempty"`
	Addresses                 []Address `json:"Addresses,omitempty"`
	Phones                    []Phone   `json:"Phones,omitempty"`
	DefaultCurrency           string    `json:"DefaultCurrency,omitempty"`
	XeroNetworkKey            string    `json:"XeroNetworkKey,omitempty"`
}

type Contact struct {
	ContactID                   string        `json:"ContactID,omitempty"`
	ContactStatus               string        `json:"ContactStatus,omitempty"`
	Name                        string        `json:"Name,omitempty"`
	FirstName                   string        `json:"FirstName"`
	LastName                    string        `json:"LastName"`
	EmailAddress                string        `json:"EmailAddress"`
	BankAccountDetails          string        `json:"BankAccountDetails,omitempty"`
	CompanyNumber               string        `json:"CompanyNumber,omitempty"`
	TaxNumber                   string        `json:"TaxNumber,omitempty"`
	AccountsReceivableTaxType   string        `json:"AccountsReceivableTaxType,omitempty"`
	AccountsPayableTaxType      string        `json:"AccountsPayableTaxType,omitempty"`
	Addresses                   []Address     `json:"Addresses,omitempty"`
	Phones                      []Phone       `json:"Phones,omitempty"`
	ContactGroups               []interface{} `json:"ContactGroups,omitempty"`
	IsSupplier                  bool          `json:"IsSupplier,omitempty"`
	IsCustomer                  bool          `json:"IsCustomer,omitempty"`
	DefaultCurrency             string        `json:"DefaultCurrency,omitempty"`
	XeroNetworkKey              string        `json:"XeroNetworkKey,omitempty"`
	SalesTrackingCategories     []interface{} `json:"SalesTrackingCategories,omitempty"`
	PurchasesTrackingCategories []interface{} `json:"PurchasesTrackingCategories,omitempty"`
	ContactPersons              []interface{} `json:"ContactPersons,omitempty"`
	HasValidationErrors         bool          `json:"HasValidationErrors,omitempty"`
	UpdatedDateUTC              string        `json:"UpdatedDateUTC,omitempty"`
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

func (c *XeroClient) GetContactById(contactID string) (*Contact, error) {
	req := c.SetupBaseRequest(http.MethodGet, fmt.Sprintf("/api.xro/2.0/Contacts/%s", contactID))

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, fmt.Errorf("error making API call: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("could not get contact, unexpected status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var contactResponse struct {
		Contacts []Contact `json:"Contacts"`
	}

	err = json.Unmarshal(body, &contactResponse)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	if len(contactResponse.Contacts) == 0 {
		return nil, fmt.Errorf("no contact found with ID %s", contactID)
	}

	return &contactResponse.Contacts[0], nil
}
