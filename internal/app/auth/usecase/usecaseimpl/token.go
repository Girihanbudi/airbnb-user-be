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

// Create access and refresh tokens and store them in cache
func (u Usecase) createAndStoreTokensPair(ctx *gin.Context, user usermodule.User) (err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Generate access token
	at, claims, createAtErr := jwt.GenerateToken(appcontext.AccessTokenDuration, nil)
	if createAtErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknGenerateFailed, clientLocale)
		return
	}

	// Create and store access token claims metadata in cache
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

	// Generate refresh token
	rt, claims, createRtErr := jwt.GenerateToken(appcontext.RefreshTokenDuration, nil)
	if createRtErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknGenerateFailed, clientLocale)
		return
	}

	// Create and store refresh token claims metadata in cache
	RtClaims := authcache.DefaultClaims{
		UserID: user.Id,
	}
	storeRtErr := authcache.Set(claims["jti"].(string), RtClaims, appcontext.RefreshTokenDuration)
	if storeRtErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknStoreFailed, clientLocale)
		return
	}

	// Set cookies
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

// Create An OTP and store it into a cache
func (u Usecase) createAndStoreOtp(ctx *gin.Context, userId string) (otp string, err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Generate OTP using 6 digit random number
	otp = codegenerator.RandomEncodedNumbers(6)

	// Store generated OTP in cache
	storeOtpErr := otpcache.Set(otp, userId, appcontext.OtpDuration)
	if storeOtpErr != nil {
		err = transutil.TranslateError(ctx, errpreset.TknStoreFailed, clientLocale)
		return
	}

	return
}

// Extract a token into a key
func (u Usecase) extractToken(ctx *gin.Context, token string) (jti string, err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Extract a token and get metadata from  it
	tokenMetadata := jwt.ExtractTokenMetadata(token)
	if tokenMetadata == nil {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	// Set JTI value from token metadata
	claims := *tokenMetadata
	jti = claims["jti"].(string)

	return
}

// Remove existing token in cache
func (u Usecase) deleteOldToken(ctx *gin.Context, name string) {
	// Read token from cookie
	token, readCookieErr := ctx.Cookie(name)
	if readCookieErr != nil {
		return
	}

	// Extract the token and get JTI key
	key, extractTokenErr := u.extractToken(ctx, token)
	if extractTokenErr != nil {
		return
	}

	// Delete token in cache
	if delOldTokenErr := authcache.Del(key); delOldTokenErr != nil {
		return
	}
}
