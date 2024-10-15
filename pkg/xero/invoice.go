package xero

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"
)

type InvoiceType int

const (
	ACCPAY InvoiceType = iota // A bill – commonly known as an Accounts Payable or supplier invoice
	ACCREC                    // A sales invoice – commonly known as an Accounts Receivable or customer invoice
)

func (it InvoiceType) String() string {
	types := [...]string{"ACCPAY", "ACCREC"}
	if int(it) < 0 || int(it) >= len(types) {
		return "UNKNOWN"
	}
	return types[it]
}

func (it *InvoiceType) MarshalJSON() ([]byte, error) {
	var s string
	switch *it {
	case ACCPAY:
		s = "ACCPAY"
	case ACCREC:
		s = "ACCREC"
	default:
		return nil, fmt.Errorf("unknown InvoiceType: %d", *it)
	}
	return json.Marshal(s)
}

func (it *InvoiceType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "ACCPAY":
		*it = ACCPAY
	case "ACCREC":
		*it = ACCREC
	default:
		return fmt.Errorf("unknown InvoiceType: %s", s)
	}

	return nil
}

type InvoiceStatus int

const (
	Draft      InvoiceStatus = iota // The default status if this element is not provided with your API call
	Submitted                       // Useful if there is an approval process required
	Deleted                         // An invoice with a status of DELETED will not be included in the Aged Receivables Report
	Authorised                      // The "approved" state of an invoice ready for sending to a customer
	Paid                            // Once an invoice is fully paid, the status will change to PAID
	Voided                          // If an invoice is no longer required, it can be voided
)

func (is InvoiceStatus) String() string {
	types := [...]string{"DRAFT", "SUBMITTED", "DELETED", "AUTHORISED", "PAID", "VOIDED"}
	if int(is) < 0 || int(is) >= len(types) {
		return "UNKNOWN"
	}
	return types[is]
}

func (is *InvoiceStatus) MarshalJSON() ([]byte, error) {
	var s string
	switch *is {
	case Draft:
		s = "DRAFT"
	case Submitted:
		s = "SUBMITTED"
	case Deleted:
		s = "DELETED"
	case Authorised:
		s = "AUTHORISED"
	case Paid:
		s = "PAID"
	case Voided:
		s = "VOIDED"
	default:
		return nil, fmt.Errorf("unknown InvoiceStatus: %d", *is)
	}
	return json.Marshal(s)
}

func (is *InvoiceStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "DRAFT":
		*is = Draft
	case "SUBMITTED":
		*is = Submitted
	case "DELETED":
		*is = Deleted
	case "AUTHORISED":
		*is = Authorised
	case "PAID":
		*is = Paid
	case "VOIDED":
		*is = Voided
	default:
		return fmt.Errorf("unknown InvoiceStatus: %s", s)
	}

	return nil
}

type InvoiceBase struct {
	Type            InvoiceType   `json:"Type,omitempty"`
	InvoiceNumber   string        `json:"InvoiceNumber,omitempty"`
	Reference       string        `json:"Reference,omitempty"`
	AmountDue       float64       `json:"AmountDue,omitempty"`
	AmountPaid      float64       `json:"AmountPaid,omitempty"`
	SentToContact   bool          `json:"SentToContact,omitempty"`
	CurrencyRate    float64       `json:"CurrencyRate,omitempty"`
	IsDiscounted    bool          `json:"IsDiscounted,omitempty"`
	HasErrors       bool          `json:"HasErrors,omitempty"`
	Contact         Contact       `json:"Contact"`
	DateString      string        `json:"DateString,omitempty"`
	Date            string        `json:"Date,omitempty"`
	DueDateString   string        `json:"DueDateString,omitempty"`
	DueDate         string        `json:"DueDate,omitempty"`
	BrandingThemeID string        `json:"BrandingThemeID,omitempty"`
	Status          InvoiceStatus `json:"Status,omitempty"`
	SubTotal        float64       `json:"SubTotal,omitempty"`
	TotalTax        float64       `json:"TotalTax,omitempty"`
	Total           float64       `json:"Total,omitempty"`
	UpdatedDateUTC  string        `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode    string        `json:"CurrencyCode"`
	Reason          string        `json:"Reason,omitempty"`
	OrderRef        string        `json:"OrderRef,omitempty"`
	LineItems       []LineItem    `json:"LineItems"`
}

type Invoice struct {
	InvoiceID       string        `json:"InvoiceID,omitempty"`
	Type            InvoiceType   `json:"Type,omitempty"`
	InvoiceNumber   string        `json:"InvoiceNumber,omitempty"`
	Reference       string        `json:"Reference,omitempty"`
	AmountDue       float64       `json:"AmountDue,omitempty"`
	AmountPaid      float64       `json:"AmountPaid,omitempty"`
	SentToContact   bool          `json:"SentToContact,omitempty"`
	CurrencyRate    float64       `json:"CurrencyRate,omitempty"`
	IsDiscounted    bool          `json:"IsDiscounted,omitempty"`
	HasErrors       bool          `json:"HasErrors,omitempty"`
	Contact         Contact       `json:"Contact"`
	DateString      string        `json:"DateString,omitempty"`
	Date            string        `json:"Date,omitempty"`
	DueDateString   string        `json:"DueDateString,omitempty"`
	DueDate         string        `json:"DueDate,omitempty"`
	BrandingThemeID string        `json:"BrandingThemeID"`
	Status          InvoiceStatus `json:"Status,omitempty"`
	SubTotal        float64       `json:"SubTotal,omitempty"`
	TotalTax        float64       `json:"TotalTax,omitempty"`
	Total           float64       `json:"Total,omitempty"`
	UpdatedDateUTC  string        `json:"UpdatedDateUTC,omitempty"`
	CurrencyCode    string        `json:"CurrencyCode,omitempty"`
	Reason          string        `json:"Reason,omitempty"`
	OrderRef        string        `json:"OrderRef,omitempty"`
	LineItems       []LineItem    `json:"LineItems"`
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

	if response.StatusCode != 204 {
		body, _ := io.ReadAll(response.Body)

		return fmt.Errorf("failed to send invoice as email with ID: %s: %s", invoiceID, string(body))
	}

	defer response.Body.Close()

	// Request Body: <Empty> so we don't need to read the response body,
	// we can just return nil
	return nil
}

func (c *XeroClient) GetInvoiceAsPDF(invoiceID string) ([]byte, error) {
	req := c.SetupBaseRequest(GET, fmt.Sprintf("/api.xro/2.0/Invoices/%s", invoiceID))

	// Get the PDF
	req.Header.Set("Accept", "application/pdf")

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	// Read response to file
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get invoice as PDF with ID: %s", invoiceID)
	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
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

func (c *XeroClient) PayInvoice(invoiceID string, accountID string, amount float64) error {
	_, err := c.CreatePayment(&Payment{
		Invoice: Invoice{
			InvoiceID: invoiceID,
		},
		Account: Account{
			AccountID: accountID,
		},
		Date:   time.Now().Format(time.DateOnly),
		Amount: amount,
	})

	if err != nil {
		return fmt.Errorf("could not pay invoice: %w", err)
	}

	return nil
}

func (c *XeroClient) CreateInvoice(invoice *InvoiceBase) (*Invoice, error) {
	req := c.SetupBaseRequest(POST, "/api.xro/2.0/Invoices")

	body, err := json.Marshal(invoice)
	if err != nil {
		return nil, err
	}

	req.Body = io.NopCloser(bytes.NewReader(body))
	req.ContentLength = int64(len(body))

	response, err := c.client.Do(&req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {

		body, _ := io.ReadAll(response.Body)

		return nil, fmt.Errorf("could not create invoice unexpected status code: %d: %s", response.StatusCode, string(body))
	}

	body, err = io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if c.debug {
		fmt.Println("xero response: ", string(body))
	}

	var invoiceResponse InvoicesResponse

	err = json.Unmarshal(body, &invoiceResponse)

	if err != nil {
		return nil, err
	}

	if len(invoiceResponse.Invoices) == 0 {
		return nil, fmt.Errorf("no invoice returned")
	}

	return &invoiceResponse.Invoices[0], nil
}
