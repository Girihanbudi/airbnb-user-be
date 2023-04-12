package rest

import (
	"airbnb-user-be/env/appcontext"
	"airbnb-user-be/internal/app/auth/preset/request"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/stderror"
	stdresponse "airbnb-user-be/internal/pkg/stdresponse/rest"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) ContinueWithGoogle(ctx *gin.Context) {
	h.Auth.ContinueWithGoogle(*ctx)
}

func (h Handler) ContinueWithFacebook(ctx *gin.Context) {
	h.Auth.ContinueWithFacebook(*ctx)
}

func (h Handler) ContinueWithPhone(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx.Request.Context())
	var req request.ContinueWithPhone
	if bindErr := ctx.BindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	res, err := h.Auth.ContinueWithPhone(*ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	stdresponse.GinMakeHttpResponse(ctx, 200, res, nil)
}

func (h Handler) CompletePhoneRegistration(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx.Request.Context())
	var req request.CompletePhoneRegistration
	if bindErr := ctx.BindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	err := h.Auth.CompletePhoneRegistration(*ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

func (h Handler) MakePhoneSession(ctx *gin.Context) {
	clientLocale := appcontext.GetLocale(ctx.Request.Context())
	var req request.MakePhoneSession
	if bindErr := ctx.BindJSON(&req); bindErr != nil {
		err := transutil.TranslateError(ctx, stderror.DEF_AUTH_401, clientLocale)
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}
	err := h.Auth.MakePhoneSession(*ctx, req)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

func (h Handler) OauthGoogleCallback(ctx *gin.Context) {
	err := h.Auth.OauthGoogleCallback(*ctx)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

func (h Handler) OauthFacebookCallback(ctx *gin.Context) {
	err := h.Auth.OauthFacebookCallback(*ctx)
	if err != nil {
		stdresponse.GinMakeHttpResponseErr(ctx, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}
