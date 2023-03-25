package seeder

import (
	currencymodule "airbnb-user-be/internal/app/currency"

	"github.com/Rhymond/go-money"
	"gorm.io/gorm"
)

func SeedCurrencyTranslation(db gorm.DB) error {

	data := []currencymodule.CurrencyTranslation{
		// usd translation
		makeCurrencyTranslation(money.USD, "en-US", "United State Dollar"),
		makeCurrencyTranslation(money.USD, "id-ID", "Dolar Amerika Serikat"),
		// idr translation
		makeCurrencyTranslation(money.IDR, "en-US", "Indonesian Rupiah"),
		makeCurrencyTranslation(money.IDR, "id-ID", "Rupiah Indonesia"),
	}

	var currencyTranslationRecords []currencymodule.CurrencyTranslation
	if err := db.Find(&currencyTranslationRecords).Error; err != nil {
		return err
	}

	if len(currencyTranslationRecords) > 0 {
		if err := db.Delete(&currencyTranslationRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func makeCurrencyTranslation(code, localeCode, name string) currencymodule.CurrencyTranslation {
	return currencymodule.CurrencyTranslation{
		Code:       code,
		LocaleCode: localeCode,
		Name:       name,
	}
}
