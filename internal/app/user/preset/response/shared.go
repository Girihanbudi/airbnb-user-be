package response

import "time"

type UserDefaultSetting struct {
	Locale   string `json:"locale"`
	Currency string `json:"currency"`
}

type User struct {
	Id          string     `json:"id"`
	FirstName   string     `json:"firstName"`
	FullName    string     `json:"fullName"`
	Email       *string    `json:"email"`
	CountryCode *int       `json:"countryCode"`
	PhoneNumber *string    `json:"phoneNumber"`
	Image       *string    `json:"image"`
	Role        string     `json:"role"`
	DateOfBirth *time.Time `json:"dateOfBirth"`

	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	VerifiedAt *time.Time `json:"verifiedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`

	DefaultSetting *UserDefaultSetting `json:"defaultSetting"`
}
