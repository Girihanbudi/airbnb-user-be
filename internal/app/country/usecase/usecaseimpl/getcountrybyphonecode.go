package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/country/preset/error"
	"airbnb-user-be/internal/app/country/preset/request"
	"airbnb-user-be/internal/app/country/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u Usecase) GetCountryByPhoneCode(ctx context.Context, cmd request.GetCountryByPhoneCode) (res response.GetCountryByPhoneCode, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	country, getCountryErr := u.CountryRepo.GetCountryByPhoneCode(ctx, cmd.PhoneCode)
	if getCountryErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getCountryErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	res.Country = (response.Country)(country)

	return
}
