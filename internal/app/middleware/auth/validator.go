package auth

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/middleware/preset/error"
	transutil "airbnb-user-be/internal/app/translation/util"
	authcache "airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/jwt"
	"airbnb-user-be/internal/pkg/stderror"
	stdresponse "airbnb-user-be/internal/pkg/stdresponse/rest"
	"context"

	"github.com/gin-gonic/gin"
)

func GinBindAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		at, _ := ctx.Cookie(appcontext.AccessTokenCode)
		appcontext.SetFromGinRouter(ctx, appcontext.AccessTokenCode, at)
		ctx.Next()
	}
}

func GqlValidateAccessToken(ctx context.Context) (err error) {
	accessToken := appcontext.GetAccessToken(ctx)
	clientLocale := appcontext.GetLocale(ctx)
	if accessToken == nil {
		err = transutil.TranslateError(ctx, errpreset.AUTH_MID_001, clientLocale).Error
		return
	}

	userId, validateErr := validateJwtToken(ctx, *accessToken)
	if validateErr != nil {
		err = validateErr.Error
		return
	}

	appcontext.SetFromDefaultRouter(&ctx, appcontext.UserCode, userId)

	return
}

func GinValidateAccessToken(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()
	accessToken := appcontext.GetAccessToken(reqCtx)
	clientLocale := appcontext.GetLocale(reqCtx)
	if accessToken == nil {
		err := transutil.TranslateError(reqCtx, errpreset.AUTH_MID_001, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	userId, err := validateJwtToken(reqCtx, *accessToken)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	appcontext.SetFromGinRouter(ctx, appcontext.UserCode, userId)

	ctx.Next()
}

func GinValidateNoJwtTokenFound(ctx *gin.Context) {
	accessToken := appcontext.GetAccessToken(ctx.Request.Context())
	clientLocale := appcontext.GetLocale(ctx.Request.Context())
	if accessToken != nil {
		err := transutil.TranslateError(ctx.Request.Context(), errpreset.AUTH_MID_003, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Next()
}

func validateJwtToken(ctx context.Context, accessToken string) (userId string, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)
	tokenMetadata := jwt.ExtractTokenMetadata(accessToken)
	if tokenMetadata == nil {
		err = transutil.TranslateError(ctx, errpreset.AUTH_MID_002, clientLocale)
		return
	}

	claims := *tokenMetadata
	userId, _ = authcache.Get(claims["jti"].(string))
	if userId == "" {
		err = transutil.TranslateError(ctx, errpreset.AUTH_MID_001, clientLocale)
		return
	}

	return
}
