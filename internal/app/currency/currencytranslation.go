package currency

import (
	"time"
)

type CurrencyTranslation struct {
	Code       string `json:"code" gorm:"primaryKey"`
	LocaleCode string `json:"locale_code" gorm:"primaryKey"`
	Name       string `json:"name"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
