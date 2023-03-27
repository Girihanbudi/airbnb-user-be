package country

type Country struct {
	Iso       string  `json:"iso" gorm:"primaryKey"`
	Iso3      *string `json:"iso_3"`
	Name      string  `json:"name" gorm:"not null"`
	NumCode   *int    `json:"num_code"`
	PhoneCode int     `json:"phone_code" gorm:"not null"`
}
