package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	module "airbnb-user-be/internal/app/country"
	errpreset "airbnb-user-be/internal/app/country/preset/error"
	"airbnb-user-be/internal/app/country/preset/request"
	"airbnb-user-be/internal/app/country/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"

	"github.com/thoas/go-funk"
)

func (u Usecase) GetCountries(ctx context.Context, cmd request.GetCountries) (res response.GetCountries, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	countries, getCountriesErr := u.CountryRepo.GetCountries(ctx, &cmd.Pagination)
	if getCountriesErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	data := funk.Map(*countries, func(data module.Country) response.Country {
		var country response.Country

		country.Iso = data.Iso
		country.Iso3 = data.Iso3
		country.Name = data.Name
		country.NumCode = data.NumCode
		country.PhoneCode = data.PhoneCode

		return country
	}).([]response.Country)

	res.Countries = &data
	res.Pagination = &cmd.Pagination

	return
}
