package user

import (
	"airbnb-user-be/env/appcontext"

	"gorm.io/gorm"
)

type UserDefaultSetting struct {
	gorm.Model
	UserId   string `json:"user_id"`
	Locale   string `json:"locale"`
	Currency string `json:"currency"`
}

func (m *UserDefaultSetting) BeforeCreate(tx *gorm.DB) (err error) {
	// set user locale to en-US if empty
	if m.Locale == "" {
		m.Locale = appcontext.LocaleDefault
	}

	// set user currency to USD if empty
	if m.Currency == "" {
		m.Currency = appcontext.CurrencyDefault
	}

	return
}
