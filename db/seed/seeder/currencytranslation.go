package seeder

import (
	"airbnb-user-be/internal/app/currency"

	"github.com/Rhymond/go-money"
	"gorm.io/gorm"
)

func SeedCurrencyTranslation(db gorm.DB) error {

	data := []currency.CurrencyTranslation{
		// usd translation
		makeCurrencyTranslation(money.USD, "en-US", "United State Dollar"),
		makeCurrencyTranslation(money.USD, "id-ID", "Dolar Amerika Serikat"),
		// idr translation
		makeCurrencyTranslation(money.IDR, "en-US", "Indonesian Rupiah"),
		makeCurrencyTranslation(money.IDR, "id-ID", "Rupiah Indonesia"),
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func makeCurrencyTranslation(code, localeCode, name string) currency.CurrencyTranslation {
	return currency.CurrencyTranslation{
		Code:       code,
		LocaleCode: localeCode,
		Name:       name,
	}
}
