package repoimpl

import (
	module "airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) GetLocales(ctx context.Context) (locales *[]module.Locale, err error) {
	err = r.Gorm.DB.Find(&locales).Error
	return
}
