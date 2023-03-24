package rest

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) ContinueWithGoogle(ctx *gin.Context) {
	h.Auth.ContinueWithGoogle(*ctx)
}

// func (r Router) ContinueWithEmail(c *gin.Context) {
// 	var req request.SignIn
// 	err := c.BindJSON(&req)
// 	if err != nil {

// 	}

// 	r.Auth.SignIn(c, req)
// }

func (h Handler) OauthGoogleCallback(ctx *gin.Context) {

	h.Auth.OauthGoogleCallback(*ctx)
}
