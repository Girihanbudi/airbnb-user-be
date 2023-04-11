package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/locale/preset/request"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u Usecase) UpdateLocale(ctx context.Context, cmd request.CreateLocale) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	locale, getLocaleErr := u.LocaleRepo.GetLocale(ctx, cmd.Code)
	if getLocaleErr != nil {
		rc := errpreset.DbServiceUnavailable
		if errors.Is(getLocaleErr, gorm.ErrRecordNotFound) {
			rc = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, rc, clientLocale)
		return
	}

	locale.Name = cmd.Name
	locale.Location = cmd.Location
	locale.Lcid = cmd.Lcid
	locale.ISO639_1 = cmd.ISO639_1
	locale.ISO639_2 = cmd.ISO639_2

	updateLocaleErr := u.LocaleRepo.UpdateLocale(ctx, locale)
	if updateLocaleErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	return
}
