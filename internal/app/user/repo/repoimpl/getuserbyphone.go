package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetUserByPhone(ctx context.Context, countryCode int, phoneNumber string) (user module.User, err error) {
	err = r.Gorm.DB.
		Where("country_code = ?", countryCode).
		Where("phone_number = ?", phoneNumber).
		First(&user).Error
	return
}
