package auth

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/middleware/preset/error"
	transutil "airbnb-user-be/internal/app/translation/util"
	authcache "airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/jwt"
	stdresponse "airbnb-user-be/internal/pkg/stdresponse/rest"
	"context"

	"github.com/gin-gonic/gin"
)

func GinBindBearerAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		schema := "Bearer"
		minKeyLength := len(schema) + 1
		value := ctx.GetHeader("Authorization")
		if len(value) > minKeyLength {
			appcontext.SetFromGinRouter(ctx, appcontext.AccessTokenCode, value[minKeyLength:])
		}

		ctx.Next()
	}
}

func GqlValidateJwtToken(ctx context.Context) (err error) {
	accessToken := appcontext.GetAccessToken(ctx)
	clientLocale := appcontext.GetLocale(ctx)
	if accessToken == nil {
		err := transutil.TranslateError(ctx, errpreset.AUTH_MID_001, clientLocale)
		return err.Error
	}
	tokenMetadata := jwt.ExtractTokenMetadata(*accessToken)
	if tokenMetadata == nil {
		err := transutil.TranslateError(ctx, errpreset.AUTH_MID_002, clientLocale)
		return err.Error
	}

	claims := *tokenMetadata
	userId, _ := authcache.Get(claims["jti"].(string))
	if userId == "" {
		err := transutil.TranslateError(ctx, errpreset.AUTH_MID_001, clientLocale)
		return err.Error
	}

	appcontext.SetFromDefaultRouter(&ctx, appcontext.UserCode, userId)

	return
}

func GinValidateJwtToken(ctx *gin.Context) {
	accessToken := appcontext.GetAccessToken(ctx.Request.Context())
	clientLocale := appcontext.GetLocale(ctx.Request.Context())
	if accessToken == nil {
		err := transutil.TranslateError(ctx.Request.Context(), errpreset.AUTH_MID_001, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	tokenMetadata := jwt.ExtractTokenMetadata(*accessToken)
	if tokenMetadata == nil {
		err := transutil.TranslateError(ctx.Request.Context(), errpreset.AUTH_MID_002, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	claims := *tokenMetadata
	userId, _ := authcache.Get(claims["jti"].(string))
	if userId == "" {
		err := transutil.TranslateError(ctx.Request.Context(), errpreset.AUTH_MID_001, clientLocale)
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
