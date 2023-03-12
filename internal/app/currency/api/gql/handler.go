package gql

import (
	"airbnb-user-be/internal/app/currency/preset/response"
	"context"
)

func (h Handler) GetCurrenciesWithTranslation(ctx context.Context) (*response.GetCurrencyWithTranslation, error) {

	res, err := h.Currency.GetCurrenciesWithTranslation(ctx)
	if err != nil {
		return nil, err.Error
	}

	return &res, nil
}
