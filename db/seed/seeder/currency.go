package seeder

import (
	"airbnb-user-be/internal/app/currency"

	"github.com/Rhymond/go-money"
	"gorm.io/gorm"
)

func SeedCurrency(db gorm.DB) error {

	data := []currency.Currency{
		makeCurrency(money.USD, "$"),
		makeCurrency(money.IDR, "Rp"),
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func makeCurrency(code, symbol string) currency.Currency {
	return currency.Currency{
		Code:   code,
		Symbol: symbol,
	}
}
