package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetLocale(ctx context.Context, cmd request.GetLocale) (res response.GetLocale, err *stderror.StdError) {
	clientLocale := ctx.Value(appcontext.LocaleCode).(string)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.LOCALE_VAL_400, clientLocale)
		return
	}

	Locale, getLocaleErr := u.LocaleRepo.GetLocale(ctx, cmd.Code)
	if getLocaleErr != nil {
		err = transutil.TranslateError(ctx, errpreset.LOCALE_GET_404, clientLocale)
		return
	}

	res.Locale = Locale

	return
}
