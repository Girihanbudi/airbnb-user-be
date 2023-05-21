package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	module "airbnb-user-be/internal/app/locale"
	errpreset "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/locale/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"

	"github.com/thoas/go-funk"
)

func (u Usecase) GetLocales(ctx context.Context) (res response.GetLocales, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	locales, getLocalesErr := u.LocaleRepo.GetLocales(ctx)
	if getLocalesErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	data := funk.Map(*locales, func(data module.Locale) response.Locale {
		var locale response.Locale

		locale.Code = data.Code
		locale.Name = data.Name
		locale.Local = data.Local
		locale.Location = data.Location
		locale.Lcid = data.Lcid
		locale.ISO639_1 = data.ISO639_1
		locale.ISO639_2 = data.ISO639_2
		locale.CreatedAt = data.CreatedAt
		locale.UpdatedAt = data.UpdatedAt

		return locale
	}).([]response.Locale)

	res.Locales = &data

	return
}
