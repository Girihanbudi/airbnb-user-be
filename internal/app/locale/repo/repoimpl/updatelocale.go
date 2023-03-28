package repoimpl

import (
	module "airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) UpdateLocale(ctx context.Context, locale *module.Locale) (err error) {
	err = r.Gorm.DB.Save(locale).Error
	return
}
