package repoimpl

import (
	module "airbnb-user-be/internal/app/country"
	"context"
)

func (r Repo) GetCountryByPhoneCode(ctx context.Context, phoneCode int) (country module.Country, err error) {

	err = r.Gorm.DB.
		Where("phone_code = ?", phoneCode).
		First(&country).Error

	return
}
