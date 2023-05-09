package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetUserByEmail(ctx context.Context, email string, preloads ...string) (user module.User, err error) {
	query := r.Gorm.DB
	if len(preloads) > 0 {
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}
	err = query.Where("email = ?", email).First(&user).Error
	return
}
