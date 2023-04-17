package usecase

import (
	"airbnb-user-be/internal/app/auth/preset/request"
	"airbnb-user-be/internal/app/auth/preset/response"
	"airbnb-user-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

type IAuth interface {
	ContinueWithGoogle(ctx *gin.Context)
	ContinueWithFacebook(ctx *gin.Context)
	ContinueWithPhone(ctx *gin.Context, cmd request.ContinueWithPhone) (res response.ContinueWithPhone, err *stderror.StdError)
	CompletePhoneRegistration(ctx *gin.Context, cmd request.CompletePhoneRegistration) (err *stderror.StdError)
	MakePhoneSession(ctx *gin.Context, cmd request.MakePhoneSession) (err *stderror.StdError)
	OauthGoogleCallback(ctx *gin.Context) (err *stderror.StdError)
	OauthFacebookCallback(ctx *gin.Context) (err *stderror.StdError)
	RefreshToken(ctx *gin.Context, cmd request.RefreshToken) (err *stderror.StdError)
	SignOut(ctx *gin.Context, cmd request.SignOut) (err *stderror.StdError)
}
