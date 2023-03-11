package repoimpl

import (
	"airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) GetLocale(ctx context.Context, code string) (Locale *locale.Locale, err error) {
	err = r.Gorm.DB.Where("code = ?", code).First(&Locale).Error
	return
}
