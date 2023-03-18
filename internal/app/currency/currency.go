package currency

import (
	"time"
)

type Currency struct {
	Code   string `json:"code" gorm:"primaryKey"`
	Symbol string `json:"symbol"`

	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
