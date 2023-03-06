package usecaseimpl

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetLocales(ctx context.Context, cmd request.GetLocales) (res response.GetLocales, err *stderror.StdError) {

	Locales, paging, getLocalesErr := u.LocaleRepo.GetLocales(ctx, &cmd.Pagination)
	if getLocalesErr != nil {

	}

	res.Locales = Locales
	res.Paging = paging

	return
}
