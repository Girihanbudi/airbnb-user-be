package usecaseimpl

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetLocale(ctx context.Context, cmd request.GetLocale) (res response.GetLocale, err *stderror.StdError) {

	Locale, getLocaleErr := u.LocaleRepo.GetLocale(ctx, cmd.Code)
	if getLocaleErr != nil {

	}

	res.Locale = Locale

	return
}
