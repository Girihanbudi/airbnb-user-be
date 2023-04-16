package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/auth/preset/error"
	"airbnb-user-be/internal/app/auth/preset/request"
	transutil "airbnb-user-be/internal/app/translation/util"
	authcache "airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/stderror"

	"github.com/gin-gonic/gin"
)

func (u Usecase) RefreshToken(ctx gin.Context, cmd request.RefreshToken) (err *stderror.StdError) {
	reqCtx := ctx.Request.Context()
	clientLocale := appcontext.GetLocale(reqCtx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(reqCtx, errpreset.TknInvalid, clientLocale)
		return
	}

	key, err := u.extractToken(ctx, cmd.Token)
	if err != nil {
		return
	}

	userId, _ := authcache.Get(key)
	if userId == "" {
		err = transutil.TranslateError(reqCtx, errpreset.TknInvalid, clientLocale)
		return
	}

	// delete old token
	u.deleteOldToken(ctx, appcontext.AccessTokenCode)
	u.deleteOldToken(ctx, appcontext.RefreshTokenCode)

	if err = u.createAndStoreTokensPair(ctx, userId); err != nil {
		return
	}

	return
}
