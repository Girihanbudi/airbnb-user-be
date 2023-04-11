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

func (u Usecase) DeleteLocale(ctx context.Context, cmd request.DeleteLocale) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	deleteLocaleErr := u.LocaleRepo.DeleteLocale(ctx, cmd.Code)
	if deleteLocaleErr != nil {
		rc := errpreset.DbServiceUnavailable
		if errors.Is(deleteLocaleErr, gorm.ErrRecordNotFound) {
			rc = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, rc, clientLocale)
		return
	}

	return
}
