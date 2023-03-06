package gql

import (
	"airbnb-user-be/internal/app/locale"
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

func (h Handler) GetLocale(ctx context.Context, code string) (*locale.Locale, error) {
	cmd := request.GetLocale{
		Code: code,
	}

	res, err := h.Locale.GetLocale(ctx, cmd)

	return res.Locale, err.Error
}

func (h Handler) GetLocales(ctx context.Context, limit, page int) (*[]locale.Locale, error) {
	paging := pagination.DefaultSQLPaging
	paging.Limit = limit
	paging.Page = page

	cmd := request.GetLocales{
		Pagination: paging,
	}

	res, err := h.Locale.GetLocales(ctx, cmd)

	return res.Locales, err.Error
}
