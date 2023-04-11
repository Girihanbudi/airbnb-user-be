package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) CreateOrUpdateUser(ctx context.Context, user *module.User) (err error) {
	err = r.Gorm.DB.Save(user).Error
	return
}
