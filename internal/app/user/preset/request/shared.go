package request

import "time"

type UserDefaultSetting struct {
	Locale   string `json:"locale"`
	Currency string `json:"currency"`
}

type User struct {
	FirstName   string     `json:"firstName,omitempty"`
	LastName    string     `json:"lastName,omitempty"`
	FullName    string     `json:"fullName,omitempty"`
	Email       *string    `json:"email,omitempty"`
	CountryCode *int       `json:"countryCode,omitempty"`
	PhoneNumber *string    `json:"phoneNumber,omitempty"`
	Image       *string    `json:"image"`
	Role        string     `json:"role,omitempty"`
	DateOfBirth *time.Time `json:"dateOfBirth,omitempty"`

	DefaultSetting *UserDefaultSetting `json:"defaultSetting,omitempty"`
}

type Identifier struct {
	Id string `json:"id" validation:"required"`
}
