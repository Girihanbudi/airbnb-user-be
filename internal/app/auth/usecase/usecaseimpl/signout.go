package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/auth/preset/error"
	"airbnb-user-be/internal/app/auth/preset/request"
	transutil "airbnb-user-be/internal/app/translation/util"
	authcache "airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

func (u Usecase) SignOut(ctx gin.Context, cmd request.SignOut) (err *stderror.StdError) {
	reqCtx := ctx.Request.Context()
	clientLocale := appcontext.GetLocale(reqCtx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(reqCtx, errpreset.TknInvalid, clientLocale)
		return
	}

	// Remove access token
	atKey, err := u.extractToken(ctx, cmd.AccessToken)
	if err != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknInvalid, clientLocale)
		return
	}
	if delAtKeyErr := authcache.Del(atKey); delAtKeyErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknInvalid, clientLocale)
		return
	}
	ctx.SetCookie(
		appcontext.AccessTokenCode,
		cmd.AccessToken,
		-1,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	// Remove refresh token
	rtKey, err := u.extractToken(ctx, cmd.AccessToken)
	if err != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknInvalid, clientLocale)
		return
	}
	if delRtKeyErr := authcache.Del(rtKey); delRtKeyErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknInvalid, clientLocale)
		return
	}
	ctx.SetCookie(
		appcontext.RefreshTokenCode,
		cmd.AccessToken,
		-1,
		"/sessions",
		env.CONFIG.Domain,
		true,
		true,
	)

	ctx.SetCookie(
		appcontext.IsLoggedInCode,
		"false",
		-1,
		"/",
		env.CONFIG.Domain,
		true,
		false,
	)

	return
}
