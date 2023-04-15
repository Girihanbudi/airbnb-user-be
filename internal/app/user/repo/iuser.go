package repo

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

type IUser interface {
	GetUser(ctx context.Context, id string, preloads *[]string) (user module.User, err error)
	GetUserByEmail(ctx context.Context, email string) (user module.User, err error)
	GetUserByPhone(ctx context.Context, countryCode int, phoneNumber string) (user module.User, err error)
	CreateUser(ctx context.Context, user *module.User) (err error)
	CreateOrUpdateUser(ctx context.Context, user *module.User) (err error)
	CreateOrUpdateUserAccount(ctx context.Context, account *module.Account) (err error)
}
