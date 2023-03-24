package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	"airbnb-user-be/internal/app/locale"
	errpreset "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/locale/preset/request"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) CreateLocale(ctx context.Context, cmd request.CreateLocale) (err *stderror.StdError) {
	clientLocale := ctx.Value(appcontext.LocaleCode).(string)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.LOCALE_VAL_400, clientLocale)
		return
	}

	locale := locale.Locale{
		Code:     cmd.Code,
		Name:     cmd.Name,
		Local:    cmd.Local,
		Location: cmd.Location,
		Lcid:     cmd.Lcid,
		ISO639_2: cmd.ISO639_1,
		ISO639_1: cmd.ISO639_2,
	}

	createLocaleErr := u.LocaleRepo.CreateLocale(ctx, &locale)
	if createLocaleErr != nil {
		err = transutil.TranslateError(ctx, errpreset.LOCALE_CREATE_503, clientLocale)
		return
	}

	return
}
