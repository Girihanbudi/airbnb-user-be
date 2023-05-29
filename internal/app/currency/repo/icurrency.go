package repo

import (
	module "airbnb-user-be/internal/app/currency"
	"airbnb-user-be/internal/app/currency/preset/response"
	"context"
)

type ICurrency interface {
	GetCurrenciesWithTranslation(ctx context.Context, localeCode string) (currencies *[]response.CurrencyWithTranslation, err error)
	GetCurrency(ctx context.Context, code string) (currency *module.Currency, err error)
}
