package models

import (
	"fmt"
)

type Person struct {
	Id          int64  `json:"id,omitempty"`
	FirstName   string `json:"firstname,omitempty"`
	LastName    string `json:"lastname,omitempty"`
	ContactInfo `json:"contactinfo,omitempty"`
}

type ContactInfo struct {
	City     string `json:"city,omitempty"`
	AreaCode string `json:"areacode,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

func (p Person) GetFullNumber() string {
	return fmt.Sprintf("%s-%s", p.AreaCode, p.Phone)
}
