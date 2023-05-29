package repoimpl

import (
	module "airbnb-user-be/internal/app/currency"
	"context"
)

func (r Repo) GetCurrency(ctx context.Context, code string) (currency *module.Currency, err error) {
	err = r.Gorm.DB.Where("code = ?", code).First(&currency).Error
	return
}
