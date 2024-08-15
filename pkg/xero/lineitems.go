package xero

type Item struct {
	ItemId string `json:"ItemId,omitempty"` // Xero generated identifier
	Name   string `json:"Name"`             // User defined item code
	Code   string `json:"Code"`             // The name of the item
}

type LineItem struct {
	LineItemId     string  `json:"LineItemId,omitempty"` //
	Description    string  `json:"Description"`          //
	Quantity       float32 `json:"Quantity"`             //
	UnitAmount     float32 `json:"UnitAmount"`           // Lineitem unit amount. By default, unit amount will be rounded to two decimal places. You can opt in to use four decimal places by adding the querystring parameter unitdp=4 to your query. See the Rounding in Xero guide for more information.
	ItemCode       string  `json:"ItemCode"`             //
	AccountCode    string  `json:"AccountCode"`          //
	AccountId      string  `json:"AccountId"`            //
	TaxType        string  `json:"TaxType"`              //
	TaxAmount      float32 `json:"TaxAmount"`            // The tax amount is auto calculated as a percentage of the line amount based on the tax rate
	LineAmount     float32 `json:"LineAmount"`           // The line amount reflects the discounted price if a DiscountRate has been used i.e LineAmount = Quantity * Unit Amount * ((100 – DiscountRate)/100)
	DiscountRate   float32 `json:"DiscountRate"`         // Percentage discount being applied to a line item. Only supported on ACCREC invoices and quotes. ACCPAY invoices and credit notes in Xero do not support discounts
	DiscountAmount float32 `json:"DiscountAmount"`       // Discount amount being applied to a line item. Only supported on ACCREC invoices and quotes. ACCPAY invoices and credit notes in Xero do not support discounts
}
