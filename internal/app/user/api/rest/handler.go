package rest

import (
	"airbnb-user-be/env/appcontext"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/app/user/preset/request"
	_ "airbnb-user-be/internal/app/user/preset/response"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/stderror"
	stdresponse "airbnb-user-be/internal/pkg/stdresponse/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) ChangeLocaleSetting(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx)
	var req request.ChangeLocaleSetting
	if bindErr := ctx.ShouldBindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	// bind user information if logged in
	claims := appcontext.GetUserClaims(ctx)
	if claims != nil {
		req.UserId = &claims.UserID
	}

	err := h.User.ChangeLocaleSetting(ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	// Set locale cookie
	ctx.SetCookie(
		appcontext.LocaleCode,
		req.Locale,
		appcontext.LocaleDuration,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	stdresponse.GinMakeHttpResponse(ctx, http.StatusCreated, nil, nil)
}

func (h Handler) ChangeCurrencySetting(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx)
	var req request.ChangeCurrencySetting
	if bindErr := ctx.ShouldBindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	// bind user information if logged in
	claims := appcontext.GetUserClaims(ctx)
	if claims != nil {
		req.UserId = &claims.UserID
	}

	err := h.User.ChangeCurrencySetting(ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	// Set currency cookie
	ctx.SetCookie(
		appcontext.CurrencyCode,
		req.Currency,
		appcontext.CurrencyDuration,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	stdresponse.GinMakeHttpResponse(ctx, http.StatusCreated, nil, nil)
}
