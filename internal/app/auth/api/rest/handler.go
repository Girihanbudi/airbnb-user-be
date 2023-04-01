package rest

import (
	"airbnb-user-be/internal/pkg/env"
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
