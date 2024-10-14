package xero

type Item struct {
	ItemId string `json:"ItemId,omitempty"` // Xero generated identifier
	Name   string `json:"Name,omitempty"`   // User defined item code
	Code   string `json:"Code,omitempty"`   // The name of the item
}

type LineItem struct {
	LineItemId     string  `json:"LineItemId,omitempty"`     //
	Description    string  `json:"Description,omitempty"`    //
	Quantity       float32 `json:"Quantity,omitempty"`       //
	UnitAmount     float32 `json:"UnitAmount,omitempty"`     // Lineitem unit amount. By default, unit amount will be rounded to two decimal places. You can opt in to use four decimal places by adding the querystring parameter unitdp=4 to your query. See the Rounding in Xero guide for more information.
	ItemCode       string  `json:"ItemCode,omitempty"`       //
	AccountCode    string  `json:"AccountCode,omitempty"`    //
	AccountId      string  `json:"AccountId,omitempty"`      //
	TaxType        string  `json:"TaxType,omitempty"`        //
	TaxAmount      float32 `json:"TaxAmount,omitempty"`      // The tax amount is auto calculated as a percentage of the line amount based on the tax rate
	LineAmount     float32 `json:"LineAmount,omitempty"`     // The line amount reflects the discounted price if a DiscountRate has been used i.e LineAmount = Quantity * Unit Amount * ((100 – DiscountRate)/100)
	DiscountRate   float32 `json:"DiscountRate,omitempty"`   // Percentage discount being applied to a line item. Only supported on ACCREC invoices and quotes. ACCPAY invoices and credit notes in Xero do not support discounts
	DiscountAmount float32 `json:"DiscountAmount,omitempty"` // Discount amount being applied to a line item. Only supported on ACCREC invoices and quotes. ACCPAY invoices and credit notes in Xero do not support discounts
}
