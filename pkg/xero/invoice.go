package xero

import (
	"encoding/json"
	"fmt"
	"io"
)

type Invoice struct {
	InvoiceID     string `json:"InvoiceID"`
	InvoiceNumber string `json:"InvoiceNumber"`
}

type InvoicesResponse struct {
	Invoices []Invoice `json:"Invoices"`
}

func (c *XeroClient) GetInvoices() (*InvoicesResponse, error) {
	req := c.SetupBaseRequest(GET, "/api.xro/2.0/Invoices")

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var invoicesResponse InvoicesResponse

	err = json.Unmarshal(body, &invoicesResponse)

	if err != nil {
		return nil, err
	}

	return &invoicesResponse, nil
}

func (c *XeroClient) GetInvoiceByID(invoiceID string) (*Invoice, error) {
	req := c.SetupBaseRequest(GET, fmt.Sprintf("/api.xro/2.0/Invoices/%s", invoiceID))

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("no invoice found with ID: %s", invoiceID)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var invoicesResponse InvoicesResponse

	err = json.Unmarshal(body, &invoicesResponse)

	if err != nil {
		return nil, err
	}

	if len(invoicesResponse.Invoices) == 0 {
		return nil, fmt.Errorf("no invoice found with ID: %s", invoiceID)
	}

	return &invoicesResponse.Invoices[0], nil
}
