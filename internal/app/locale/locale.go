package locale

import (
	"time"

	"gorm.io/gorm"
)

type Locale struct {
	Code     string  `json:"code" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Local    *string `json:"local"`
	Location *string `json:"location"`
	Lcid     *int    `json:"lcid"`
	ISO639_1 *string `json:"iso639_1"`
	ISO639_2 *string `json:"iso639_2"`

	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null"`
}

func (l *Locale) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	l.CreatedAt = now
	l.UpdatedAt = now

	return
}

func (l *Locale) BeforeSave(tx *gorm.DB) (err error) {
	l.UpdatedAt = time.Now()
	return
}
