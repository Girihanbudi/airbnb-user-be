package repo

import (
	"airbnb-user-be/internal/app/currency/preset/response"
	"context"
)

type ICurrency interface {
	GetCurrenciesWithTranslation(ctx context.Context, localeCode string) (currencies *[]response.CurrencyWithTranslation, err error)
}
