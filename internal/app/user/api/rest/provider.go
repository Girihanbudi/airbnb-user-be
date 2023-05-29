package rest

import (
	"github.com/gin-gonic/gin"

	userusecase "airbnb-user-be/internal/app/user/usecase"
)

type Options struct {
	Router *gin.Engine

	User userusecase.IUser
}

type Handler struct {
	Options
}

func NewUserHandler(options Options) *Handler {
	return &Handler{options}
}
