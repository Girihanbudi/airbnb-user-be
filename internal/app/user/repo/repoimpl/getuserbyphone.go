package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetUserByPhone(ctx context.Context, countryCode int, phoneNumber string, preloads ...string) (user module.User, err error) {
	query := r.Gorm.DB
	if len(preloads) > 0 {
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	err = query.
		Where("country_code = ?", countryCode).
		Where("phone_number = ?", phoneNumber).
		First(&user).Error
	return
}
