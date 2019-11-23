package models

type Person struct {
	ID          int64  `json:"id,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	ContactInfo `json:"contactinfo,omitempty"`
}

type ContactInfo struct {
	City     string `json:"city,omitempty"`
	AreaCode string `json:"areacode,omitempty"`
	Phone    string    `json:"phone,omitempty"`
}
