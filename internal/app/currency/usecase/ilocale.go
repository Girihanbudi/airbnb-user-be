package usecase

import (
	"airbnb-user-be/internal/app/currency/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

type ICurrency interface {
	GetCurrenciesWithTranslation(ctx context.Context) (res response.GetCurrencyWithTranslation, err *stderror.StdError)
}
