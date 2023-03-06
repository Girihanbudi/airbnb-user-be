package repo

import (
	"airbnb-user-be/internal/app/locale"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

type ILocale interface {
	GetLocales(ctx context.Context, page *pagination.SQLPaging) (Locales *[]locale.Locale, paging *pagination.SQLPaging, err error)
	GetLocale(ctx context.Context, code string) (Locale *locale.Locale, err error)
}
