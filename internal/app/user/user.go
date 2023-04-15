package user

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	FirstName   string    `json:"first_name" gorm:"not null"`
	FullName    string    `json:"full_name" gorm:"not null"`
	Email       *string   `json:"email" gorm:"unique"`
	CountryCode *int      `json:"country_code"`
	PhoneNumber *string   `json:"phone_number" gorm:"unique"`
	Image       string    `json:"image"`
	Password    *string   `json:"password"`
	Role        string    `json:"role"`
	DateOfBirth time.Time `json:"date_of_birth"`

	CreatedAt  time.Time      `json:"created_at" gorm:"not null"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"not null"`
	VerifiedAt *time.Time     `json:"verified_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	DefaultSetting *UserDefaultSetting `json:"default_setting" gorm:"foreignKey:UserId"`
	Accounts       *[]Account          `json:"accounts"`
	// sessions      Session[]
}

func (m *User) BeforeCreate(tx *gorm.DB) (err error) {

	id, err := gonanoid.New()
	if err != nil {
		return
	}

	m.Id = id

	return
}

// model Session {
//   id           String   @id @default(cuid())
//   createdAt    DateTime @default(now()) @map("created_at")
//   sessionToken String   @unique @map("session_token")
//   userId       String   @map("user_id")
//   expires      DateTime
//   user         User     @relation(fields: [userId], references: [id], onDelete: Cascade)
//   @@map("sessions")
// }
