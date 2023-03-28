package repoimpl

import (
	module "airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) CreateLocale(ctx context.Context, locale *module.Locale) (err error) {
	err = r.Gorm.DB.Create(locale).Error
	return
}
