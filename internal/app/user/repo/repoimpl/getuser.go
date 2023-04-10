package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetUserById(ctx context.Context, id string) (user module.User, err error) {
	err = r.Gorm.DB.First(&user, id).Error
	return
}
