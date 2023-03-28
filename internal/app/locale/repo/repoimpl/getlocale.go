package repoimpl

import (
	module "airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) GetLocale(ctx context.Context, code string) (locale *module.Locale, err error) {
	err = r.Gorm.DB.Where("code = ?", code).First(&locale).Error
	return
}
