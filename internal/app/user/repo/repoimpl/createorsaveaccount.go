package repoimpl

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

func (r Repo) CreateOrUpdateUserAccount(ctx context.Context, account *module.Account) (err error) {
	err = r.Gorm.DB.Save(account).Error
	return
}
