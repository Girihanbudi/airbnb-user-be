package auth

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/middleware/preset/error"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/jwt"
	stdresponse "airbnb-user-be/internal/pkg/stdresponse/rest"
	"context"

	"github.com/gin-gonic/gin"
)

func GinBindBearerAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientLocale := ctx.Request.Context().Value(appcontext.LocaleCode).(string)
		schema := "bearer"
		value := ctx.GetHeader("Authorization")
		if len(value) <= len(schema) {
			err := transutil.TranslateError(ctx.Request.Context(), errpreset.AUTH_MID_001, clientLocale)
			stdresponse.GinMakeHttpResponseErr(ctx, err)
			return
		}

		appcontext.SetFromGinRouter(ctx, appcontext.AccessTokenCode, value)
		ctx.Next()
	}
}

func GqlValidateJwtToken(ctx context.Context) (err error) {
	accessToken := ctx.Value(appcontext.AccessTokenCode).(string)
	clientLocale := ctx.Value(appcontext.LocaleCode).(string)
	if accessToken == "" {
		err := transutil.TranslateError(ctx, errpreset.AUTH_MID_001, clientLocale)
		return err.Error
	}
	claims := *jwt.ExtractTokenMetadata(accessToken)
	userId, _ := auth.Cache.Get(claims["jti"].(string)).Result()
	if userId == "" {
		err := transutil.TranslateError(ctx, errpreset.AUTH_MID_001, clientLocale)
		return err.Error
	}

	appcontext.SetFromDefaultRouter(&ctx, appcontext.UserCode, userId)

	return
}

func GinValidateJwtToken(ctx *gin.Context) {
	accessToken := ctx.Value(appcontext.AccessTokenCode).(string)
	clientLocale := ctx.Value(appcontext.LocaleCode).(string)
	if accessToken == "" {
		err := transutil.TranslateError(ctx.Request.Context(), errpreset.AUTH_MID_001, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	claims := *jwt.ExtractTokenMetadata(accessToken)
	userId, _ := auth.Cache.Get(claims["jti"].(string)).Result()
	if userId == "" {
		err := transutil.TranslateError(ctx.Request.Context(), errpreset.AUTH_MID_001, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	appcontext.SetFromGinRouter(ctx, appcontext.UserCode, userId)

	ctx.Next()
}
