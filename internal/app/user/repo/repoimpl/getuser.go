package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) GetUser(ctx context.Context, id string, preloads ...string) (user module.User, err error) {
	query := r.Gorm.DB
	if len(preloads) > 0 {
		for _, preload := range preloads {
			query = query.Preload(preload)
		}
	}

	err = query.Where("id = ?", id).First(&user).Error
	return
}
