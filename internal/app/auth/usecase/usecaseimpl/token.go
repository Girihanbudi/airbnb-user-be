package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/auth/preset/error"
	transutil "airbnb-user-be/internal/app/translation/util"
	authcache "airbnb-user-be/internal/pkg/cache/auth"
	otpcache "airbnb-user-be/internal/pkg/cache/otp"
	"airbnb-user-be/internal/pkg/codegenerator"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/jwt"
	"airbnb-user-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

func (u Usecase) createAndStoreTokensPair(ctx gin.Context, userId string) (err *stderror.StdError) {
	reqCtx := ctx.Request.Context()
	clientLocale := appcontext.GetLocale(reqCtx)
	at, claims, createAtErr := jwt.GenerateToken(appcontext.AccessTokenDuration, nil)
	if createAtErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknGenerateFailed, clientLocale)
		return
	}

	storeAtErr := authcache.Set(claims["jti"].(string), userId, appcontext.AccessTokenDuration)
	if storeAtErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknStoreFailed, clientLocale)
		return
	}

	rt, claims, createRtErr := jwt.GenerateToken(appcontext.RefreshTokenDuration, nil)
	if createRtErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknGenerateFailed, clientLocale)
		return
	}

	storeRtErr := authcache.Set(claims["jti"].(string), userId, appcontext.RefreshTokenDuration)
	if storeRtErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknStoreFailed, clientLocale)
		return
	}

	// set cookies
	ctx.SetCookie(
		appcontext.AccessTokenCode,
		at,
		appcontext.AccessTokenDuration,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	ctx.SetCookie(
		appcontext.RefreshTokenCode,
		rt,
		appcontext.RefreshTokenDuration,
		"/sessions",
		env.CONFIG.Domain,
		true,
		true,
	)

	ctx.SetCookie(
		appcontext.IsLoggedInCode,
		"true",
		appcontext.AccessTokenDuration,
		"/",
		env.CONFIG.Domain,
		true,
		false,
	)

	return
}

func (u Usecase) createAndStoreOtp(ctx gin.Context, userId string) (otp string, err *stderror.StdError) {
	reqCtx := ctx.Request.Context()
	clientLocale := appcontext.GetLocale(reqCtx)

	otp = codegenerator.RandomEncodedNumbers(6)
	storeOtpErr := otpcache.Set(otp, userId, appcontext.OtpDuration)
	if storeOtpErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.TknStoreFailed, clientLocale)
		return
	}

	return
}

func (u Usecase) extractToken(ctx gin.Context, token string) (jti string, err *stderror.StdError) {
	reqCtx := ctx.Request.Context()
	clientLocale := appcontext.GetLocale(reqCtx)
	tokenMetadata := jwt.ExtractTokenMetadata(token)
	if tokenMetadata == nil {
		err = transutil.TranslateError(reqCtx, errpreset.UscBadRequest, clientLocale)
		return
	}

	claims := *tokenMetadata
	jti = claims["jti"].(string)

	return
}

func (u Usecase) deleteOldToken(ctx gin.Context, name string) {
	token, readCookieErr := ctx.Cookie(name)
	if readCookieErr != nil {
		return
	}

	key, extractTokenErr := u.extractToken(ctx, token)
	if extractTokenErr != nil {
		return
	}

	if delOldTokenErr := authcache.Del(key); delOldTokenErr != nil {
		return
	}
}
