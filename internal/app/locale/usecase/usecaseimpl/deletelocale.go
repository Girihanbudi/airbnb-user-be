package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/locale/preset/request"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) DeleteLocale(ctx context.Context, cmd request.DeleteLocale) (err *stderror.StdError) {
	clientLocale := ctx.Value(appcontext.LocaleCode).(string)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.LOCALE_VAL_400, clientLocale)
		return
	}

	deleteLocaleErr := u.LocaleRepo.DeleteLocale(ctx, cmd.Code)
	if deleteLocaleErr != nil {
		err = transutil.TranslateError(ctx, errpreset.LOCALE_DELETE_500, clientLocale)
		return
	}

	return
}
