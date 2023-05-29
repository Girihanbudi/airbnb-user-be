package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetDefaultSettingByUser(ctx context.Context, userId string) (setting module.UserDefaultSetting, err error) {
	err = r.Gorm.DB.Where("user_id = ?", userId).First(&setting).Error
	return
}
