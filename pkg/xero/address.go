package xero

type Address struct {
	AddressType string `json:"AddressType"`
	City        string `json:"City"`
	Region      string `json:"Region"`
	PostalCode  string `json:"PostalCode"`
	Country     string `json:"Country"`
}
