package repo

import (
	module "airbnb-user-be/internal/app/locale"
	"context"
)

type ILocale interface {
	GetLocales(ctx context.Context) (locales *[]module.Locale, err error)
	GetLocale(ctx context.Context, code string) (locale *module.Locale, err error)
	CreateLocale(ctx context.Context, locale *module.Locale) (err error)
	UpdateLocale(ctx context.Context, locale *module.Locale) (err error)
	DeleteLocale(ctx context.Context, code string) (err error)
}
