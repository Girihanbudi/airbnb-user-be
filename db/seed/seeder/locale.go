package seeder

import (
	localemodule "airbnb-user-be/internal/app/locale"

	"gorm.io/gorm"
)

func SeedLocale(db gorm.DB) error {

	data := []localemodule.Locale{
		makeLocale("en-US", "English", "English", "United States", 1033, "eng", "en"),
		makeLocale("id-ID", "Indonesian", "Bahasa Indonesia", "Indonesia", 1057, "ind", "id"),
	}

	var localeRecords []localemodule.Locale
	if err := db.Find(&localeRecords).Error; err != nil {
		return err
	}

	if len(localeRecords) > 0 {
		if err := db.Delete(&localeRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func makeLocale(code, name, local, location string, lcid int, ISO639_1, ISO639_2 string) localemodule.Locale {
	return localemodule.Locale{
		Code:     code,
		Name:     name,
		Local:    &local,
		Location: &location,
		Lcid:     &lcid,
		ISO639_1: &ISO639_1,
		ISO639_2: &ISO639_2,
	}
}
