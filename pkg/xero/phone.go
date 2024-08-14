package xero

type PhoneType int

const (
	DEFAULT PhoneType = iota
	DDI
	MOBILE
	FAX
)

func (pt PhoneType) String() string {
	types := [...]string{"DEFAULT", "DDI", "MOBILE", "FAX"}
	if int(pt) < 0 || int(pt) >= len(types) {
		return "UNKNOWN"
	}
	return types[pt]
}

type Phone struct {
	PhoneType        PhoneType `json:"PhoneType"`
	PhoneNumber      string    `json:"PhoneNumber"`
	PhoneAreaCode    string    `json:"PhoneAreaCode"`
	PhoneCountryCode string    `json:"PhoneCountryCode"`
}
