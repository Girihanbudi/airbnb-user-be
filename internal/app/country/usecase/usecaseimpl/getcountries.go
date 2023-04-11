package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/country/preset/error"
	"airbnb-user-be/internal/app/country/preset/request"
	"airbnb-user-be/internal/app/country/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetCountries(ctx context.Context, cmd request.GetCountries) (res response.GetCountries, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	countries, getCountriesErr := u.CountryRepo.GetCountries(ctx, &cmd.Pagination)
	if getCountriesErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	res.Countries = countries
	res.Pagination = &cmd.Pagination

	return
}
