package currency

import (
	"time"

	"gorm.io/gorm"
)

type Currency struct {
	Code   string `json:"code" gorm:"primaryKey"`
	Symbol string `json:"symbol"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}

func (c *Currency) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now

	return
}

func (c *Currency) BeforeSave(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}
