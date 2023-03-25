package seeder

import (
	currencymodule "airbnb-user-be/internal/app/currency"

	"github.com/Rhymond/go-money"
	"gorm.io/gorm"
)

func SeedCurrency(db gorm.DB) error {

	data := []currencymodule.Currency{
		makeCurrency(money.USD, "$"),
		makeCurrency(money.IDR, "Rp"),
	}

	var currencyRecords []currencymodule.Currency
	if err := db.Find(&currencyRecords).Error; err != nil {
		return err
	}

	if len(currencyRecords) > 0 {
		if err := db.Delete(&currencyRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func makeCurrency(code, symbol string) currencymodule.Currency {
	return currencymodule.Currency{
		Code:   code,
		Symbol: symbol,
	}
}
