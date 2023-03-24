package usecase

import "github.com/gin-gonic/gin"

type IAuth interface {
	ContinueWithGoogle(ctx gin.Context)
	OauthGoogleCallback(ctx gin.Context)
}
