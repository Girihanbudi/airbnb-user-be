package repoimpl

import (
	"airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) GetLocales(ctx context.Context) (Locales *[]locale.Locale, err error) {
	err = r.Gorm.DB.Find(&Locales).Error
	return
}
