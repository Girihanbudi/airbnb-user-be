package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetUserByEmail(ctx context.Context, email string) (user module.User, err error) {
	err = r.Gorm.DB.Where("email = ?", email).First(&user).Error
	return
}
