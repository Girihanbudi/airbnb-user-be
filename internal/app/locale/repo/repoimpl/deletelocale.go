package repoimpl

import (
	"airbnb-user-be/internal/app/locale"
	"context"
)

func (r Repo) DeleteLocale(ctx context.Context, code string) (err error) {
	err = r.Gorm.DB.Delete(&locale.Locale{}, code).Error
	return
}
