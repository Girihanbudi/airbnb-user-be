package repoimpl

import (
	"airbnb-user-be/internal/app/currency"
	"airbnb-user-be/internal/app/currency/preset/response"
	"context"
)

func (r Repo) GetCurrenciesWithTranslation(ctx context.Context, localeCode string) (currencies *[]response.CurrencyWithTranslation, err error) {
	err = r.Gorm.DB.
		Model(&currency.CurrencyTranslation{}).
		Select("currencies.code, currencies.symbol, currency_translations.name").
		Joins("left join currencies on currencies.code = currency_translations.code").
		Where("currency_translations.locale_code = ?", localeCode).
		Scan(&currencies).Error
	return
}
