package repoimpl

import (
	"airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) CreateLocale(ctx context.Context, locale *locale.Locale) (err error) {
	err = r.Gorm.DB.Create(locale).Error
	return
}
