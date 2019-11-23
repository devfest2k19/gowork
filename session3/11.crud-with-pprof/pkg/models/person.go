package models

import "fmt"

type Person struct {
	ID          int    `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	ContactInfo `json:"contactinfo,omitempty"`
}

type ContactInfo struct {
	City     string `json:"city,omitempty"`
	AreaCode string `json:"areacode,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

// GetFullNumber retrns the dialing number
func (p *Person) GetFullNumber() string {
	return fmt.Sprintf("%s-%s", p.AreaCode, p.Phone)
}
