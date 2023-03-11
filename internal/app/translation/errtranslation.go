package translation

import (
	"time"

	"gorm.io/gorm"
)

type ErrTranslation struct {
	Code       string    `json:"code" gorm:"primaryKey"`
	LocaleCode string    `json:"locale_code" gorm:"primaryKey"`
	Message    string    `json:"message" gorm:"not null"`
	HttpCode   int       `json:"http_code" gorm:"not null"`
	CreatedAt  time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"not null"`
}

func (e *ErrTranslation) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now

	return
}

func (e *ErrTranslation) BeforeSave(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return
}
