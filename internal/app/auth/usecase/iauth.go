package usecase

import (
	"airbnb-user-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

type IAuth interface {
	ContinueWithGoogle(ctx gin.Context)
	OauthGoogleCallback(ctx gin.Context) (err *stderror.StdError)
}
