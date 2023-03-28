package graph

import (
	"airbnb-user-be/graph/model"
	currecyresp "airbnb-user-be/internal/app/currency/preset/response"
	"context"

	"github.com/thoas/go-funk"
)

// Currencies is the resolver for the currencies field.
func (r *queryResolver) Currencies(ctx context.Context) ([]*model.Currency, error) {
	data, err := r.Resolver.Currency.GetCurrenciesWithTranslation(ctx)
	if err != nil {
		return nil, err
	}

	currencies := funk.Map(*data.Currencies, func(data currecyresp.CurrencyWithTranslation) *model.Currency {
		var currency model.Currency

		currency.Code = data.Code
		currency.Symbol = data.Symbol
		currency.Name = data.Name

		return &currency
	}).([]*model.Currency)

	return currencies, nil
}
