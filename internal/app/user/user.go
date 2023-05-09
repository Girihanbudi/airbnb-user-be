package user

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	Id          string     `json:"id" gorm:"primaryKey"`
	FirstName   string     `json:"first_name"`
	FullName    string     `json:"full_name"`
	Email       *string    `json:"email" gorm:"unique"`
	CountryCode *int       `json:"country_code"`
	PhoneNumber *string    `json:"phone_number" gorm:"unique"`
	Image       *string    `json:"image"`
	Password    *string    `json:"password"`
	Role        string     `json:"role" gorm:"not null"`
	DateOfBirth *time.Time `json:"date_of_birth"`

	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"not null"`
	VerifiedAt *time.Time     `json:"verified_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	DefaultSetting *UserDefaultSetting `json:"default_setting" gorm:"foreignKey:UserId"`
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {

	id, err := gonanoid.New()
	if err != nil {
		return
	}

	m.Id = id

	return
}
