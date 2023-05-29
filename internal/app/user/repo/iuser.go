package repo

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

type IUser interface {
	GetUser(ctx context.Context, id string, preloads ...string) (user module.User, err error)
	GetUserByEmail(ctx context.Context, email string, preloads ...string) (user module.User, err error)
	GetUserByPhone(ctx context.Context, countryCode int, phoneNumber string, preloads ...string) (user module.User, err error)
	CreateUser(ctx context.Context, user *module.User) (err error)
	CreateOrUpdateUser(ctx context.Context, user *module.User) (err error)
	GetDefaultSettingByUser(ctx context.Context, userId string) (setting module.UserDefaultSetting, err error)
	CreateOrUpdateDefaultSetting(ctx context.Context, setting *module.UserDefaultSetting) (err error)
}
