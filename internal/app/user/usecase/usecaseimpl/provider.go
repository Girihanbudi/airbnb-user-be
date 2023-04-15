package usecaseimpl

import (
	"airbnb-user-be/internal/app/user/repo"
)

type Options struct {
	UserRepo repo.IUser
}

type Usecase struct {
	Options
}

func NewUserUsecase(options Options) *Usecase {
	return &Usecase{options}
}
