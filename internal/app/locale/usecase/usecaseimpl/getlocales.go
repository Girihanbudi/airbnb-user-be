package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/locale/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetLocales(ctx context.Context) (res response.GetLocales, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	locales, getLocalesErr := u.LocaleRepo.GetLocales(ctx)
	if getLocalesErr != nil {
		err = transutil.TranslateError(ctx, errpreset.LOCALE_GET_503, clientLocale)
		return
	}

	res.Locales = locales

	return
}
