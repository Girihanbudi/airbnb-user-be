package seeder

import (
	"airbnb-user-be/internal/app/locale"

	"gorm.io/gorm"
)

func SeedLocale(db gorm.DB) error {

	data := []locale.Locale{
		makeLocale("en-US", "English", "English", "United States", 1033, "eng", "en"),
		makeLocale("id-ID", "Indonesian", "Bahasa Indonesia", "Indonesia", 1057, "ind", "id"),
	}

	return db.CreateInBatches(&data, 100).Error
}

func makeLocale(code, name, local, location string, lcid int, ISO639_1, ISO639_2 string) locale.Locale {
	return locale.Locale{
		Code:     code,
		Name:     name,
		Local:    &local,
		Location: &location,
		Lcid:     &lcid,
		ISO639_1: &ISO639_1,
		ISO639_2: &ISO639_2,
	}
}
