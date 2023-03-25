package rest

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) ContinueWithGoogle(ctx *gin.Context) {
	h.Auth.ContinueWithGoogle(*ctx)
}

func (h Handler) OauthGoogleCallback(ctx *gin.Context) {

	h.Auth.OauthGoogleCallback(*ctx)
}
