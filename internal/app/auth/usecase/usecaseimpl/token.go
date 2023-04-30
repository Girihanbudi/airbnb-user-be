package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/auth/preset/error"
	transutil "airbnb-user-be/internal/app/translation/util"
	usermodule "airbnb-user-be/internal/app/user"
	authcache "airbnb-user-be/internal/pkg/cache/auth"
	otpcache "airbnb-user-be/internal/pkg/cache/otp"
	"airbnb-user-be/internal/pkg/codegenerator"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/jwt"
	"airbnb-user-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

func (u Usecase) createAndStoreTokensPair(ctx *gin.Context, user usermodule.User) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)
	at, claims, createAtErr := jwt.GenerateToken(appcontext.AccessTokenDuration, nil)
	if createAtErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknGenerateFailed, clientLocale)
		return
	}

	AtClaims := authcache.DefaultClaims{
		UserID:     user.Id,
		FirstName:  user.FirstName,
		FullName:   user.FullName,
		Role:       user.Role,
		VerifiedAt: user.VerifiedAt,
	}

	storeAtErr := authcache.Set(claims["jti"].(string), AtClaims, appcontext.AccessTokenDuration)
	if storeAtErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknStoreFailed, clientLocale)
		return
	}
	rt, claims, createRtErr := jwt.GenerateToken(appcontext.RefreshTokenDuration, nil)
	if createRtErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknGenerateFailed, clientLocale)
		return
	}

	RtClaims := authcache.DefaultClaims{
		UserID: user.Id,
	}
	storeRtErr := authcache.Set(claims["jti"].(string), RtClaims, appcontext.RefreshTokenDuration)
	if storeRtErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknStoreFailed, clientLocale)
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

func (u Usecase) createAndStoreOtp(ctx *gin.Context, userId string) (otp string, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	otp = codegenerator.RandomEncodedNumbers(6)
	storeOtpErr := otpcache.Set(otp, userId, appcontext.OtpDuration)
	if storeOtpErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknStoreFailed, clientLocale)
		return
	}

	return
}

func (u Usecase) extractToken(ctx *gin.Context, token string) (jti string, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)
	tokenMetadata := jwt.ExtractTokenMetadata(token)
	if tokenMetadata == nil {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	claims := *tokenMetadata
	jti = claims["jti"].(string)

	return
}

func (u Usecase) deleteOldToken(ctx *gin.Context, name string) {
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
