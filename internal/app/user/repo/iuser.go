package repo

import (
	module "airbnb-user-be/internal/app/user"
	"context"
)

type IUser interface {
	GetUserByEmail(ctx context.Context, email string) (user module.User, err error)
	CreateUser(ctx context.Context, user *module.User) (err error)
}
