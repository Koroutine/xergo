package xero

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type Invoice struct {
	InvoiceID     string `json:"InvoiceID"`
	InvoiceNumber string `json:"InvoiceNumber"`
}

type InvoicesResponse struct {
	Invoices []Invoice `json:"Invoices"`
}

type OnlineInvoice struct {
	OnlineInvoiceUrl string `json:"OnlineInvoiceUrl"`
}

type OnlineInvoicesResponse struct {
	OnlineInvoices []OnlineInvoice `json:"OnlineInvoices"`
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

func (c *XeroClient) SendInvoiceAsEmail(invoiceID string) error {
	req := c.SetupBaseRequest(POST, fmt.Sprintf("/api.xro/2.0/Invoices/%s/Email", invoiceID))

	response, err := c.client.Do(&req)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("failed to send invoice as email with ID: %s", invoiceID)
	}

	defer response.Body.Close()

	// Request Body: <Empty> so we don't need to read the response body,
	// we can just return nil
	return nil
}

// https://developer.xero.com/documentation/api/accounting/invoices#retrieving-the-online-invoice-url
func (c *XeroClient) GetInvoiceAsURL(invoiceID string) (*url.URL, error) {
	req := c.SetupBaseRequest(GET, fmt.Sprintf("/api.xro/2.0/Invoices/%s/OnlineInvoice", invoiceID))

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get invoice as url with ID: %s", invoiceID)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var onlineInvoicesResponse OnlineInvoicesResponse

	err = json.Unmarshal(body, &onlineInvoicesResponse)

	if err != nil {
		return nil, err
	}

	if len(onlineInvoicesResponse.OnlineInvoices) == 0 {
		return nil, fmt.Errorf("no invoice found with ID: %s", invoiceID)
	}

	// We are only interested in the first invoice in the response:
	onlineInvoice := onlineInvoicesResponse.OnlineInvoices[0]

	// Parse the OnlineInvoiceUrl to a URL:
	return url.Parse(onlineInvoice.OnlineInvoiceUrl)
}
