package currency

import (
	"time"

	"gorm.io/gorm"
)

type CurrencyTranslation struct {
	Code       string `json:"code" gorm:"primaryKey"`
	LocaleCode string `json:"locale_code" gorm:"primaryKey"`
	Name       string `json:"name"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (c *CurrencyTranslation) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	return
}

func (c *CurrencyTranslation) BeforeSave(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}
