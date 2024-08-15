package xero

import (
	"encoding/json"
	"fmt"
)

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

func (pt *PhoneType) MarshalJSON() ([]byte, error) {
	var s string
	switch *pt {
	case DEFAULT:
		s = "DEFAULT"
	case DDI:
		s = "DDI"
	case MOBILE:
		s = "MOBILE"
	case FAX:
		s = "FAX"
	default:
		return nil, fmt.Errorf("unknown PhoneType: %d", *pt)
	}
	return json.Marshal(s)
}

func (pt *PhoneType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "DEFAULT":
		*pt = DEFAULT
	case "DDI":
		*pt = DDI
	case "MOBILE":
		*pt = MOBILE
	case "FAX":
		*pt = FAX
	default:
		return fmt.Errorf("unknown PhoneType: %s", s)
	}

	return nil
}

type Phone struct {
	PhoneType        PhoneType `json:"PhoneType"`
	PhoneNumber      string    `json:"PhoneNumber"`
	PhoneAreaCode    string    `json:"PhoneAreaCode"`
	PhoneCountryCode string    `json:"PhoneCountryCode"`
}
