package graph

import (
	"airbnb-user-be/graph/model"
	module "airbnb-user-be/internal/app/country"
	"context"

	"github.com/thoas/go-funk"
)

// Countries is the resolver for the countries field.
func (r *queryResolver) Countries(ctx context.Context, limit *int, page *int) ([]*model.Country, error) {

	data, err := r.Resolver.Country.GetCountries(ctx, limit, page)
	if err != nil {
		return nil, err
	}

	countries := funk.Map(*data.Countries, func(data module.Country) *model.Country {
		var country model.Country

		country.Iso = data.Iso
		country.Iso3 = data.Iso3
		country.Name = data.Name
		country.NumCode = data.NumCode
		country.PhoneCode = data.PhoneCode

		return &country
	}).([]*model.Country)

	return countries, nil
}
