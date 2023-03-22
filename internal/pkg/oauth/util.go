package oauth

import (
	"airbnb-user-be/internal/pkg/appcontext"
	"airbnb-user-be/internal/pkg/env"
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func CreateOauthCookie(ctx *gin.Context, state string) {
	ctx.SetCookie(
		appcontext.OauthCode,
		state,
		appcontext.OauthDuration,
		"/",
		env.CONFIG.Domain,
		env.CONFIG.Stage != string(env.StageLocal),
		false)
}

func GenerateOauthState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
