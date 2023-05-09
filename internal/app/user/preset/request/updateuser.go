package request

import "time"

type UpdateUser struct {
	Id          string    `json:"id" validation:"required"`
	FirstName   string    `json:"firstName" validation:"required"`
	LastName    string    `json:"lastName" validation:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validation:"required"`

	DefaultSetting *UserDefaultSetting `json:"defaultSetting"`
}
