package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) CreateOrUpdateDefaultSetting(ctx context.Context, setting *module.UserDefaultSetting) (err error) {
	err = r.Gorm.DB.Save(setting).Error
	return
}
