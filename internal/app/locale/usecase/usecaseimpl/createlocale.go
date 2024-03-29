package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	module "airbnb-user-be/internal/app/locale"
	errpreset "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/locale/preset/request"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) CreateLocale(ctx context.Context, cmd request.CreateLocale) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	locale := module.Locale{
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
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	return
}
