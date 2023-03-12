package repo

import (
	"airbnb-user-be/internal/app/locale"
	"context"
)

type ILocale interface {
	GetLocales(ctx context.Context) (locales *[]locale.Locale, err error)
	GetLocale(ctx context.Context, code string) (locale *locale.Locale, err error)
	CreateLocale(ctx context.Context, locale *locale.Locale) (err error)
	UpdateLocale(ctx context.Context, locale *locale.Locale) (err error)
	DeleteLocale(ctx context.Context, code string) (err error)
}
