package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetUserByEmail(ctx context.Context, email string) (user module.User, err error) {
	err = r.Gorm.DB.First(&user).Error
	return
}
