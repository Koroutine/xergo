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

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("could not list invoices, unexpected status code: %d", response.StatusCode)
	}

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
