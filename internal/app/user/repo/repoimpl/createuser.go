package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) CreateUser(ctx context.Context, user *module.User) (err error) {
	err = r.Gorm.DB.Create(user).Error
	return
}
