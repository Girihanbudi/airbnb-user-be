package gql

import (
	"airbnb-user-be/internal/app/user/usecase"
)

type Options struct {
	User usecase.IUser
}

type Handler struct {
	Options
}

func NewUserHandler(options Options) *Handler {
	return &Handler{options}
}
