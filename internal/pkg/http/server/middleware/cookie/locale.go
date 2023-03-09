package cookie

import (
	"airbnb-user-be/internal/pkg/appcontext"
	"airbnb-user-be/internal/pkg/env"

	"github.com/gin-gonic/gin"
)

func BindLocale(ctx *gin.Context) {

	Locale, err := ctx.Request.Cookie(appcontext.LocaleCode)
	if err != nil {
		CreateLocale(ctx, nil)
		SetLocale(ctx, nil)
		return
	}
	SetLocale(ctx, Locale.Value)
}

func CreateLocale(ctx *gin.Context, val *string) {
	if val == nil {
		newVal := appcontext.LocaleDefault
		val = &newVal
	}

	useSecureConnection := env.CONFIG.Stage != string(env.StageLocal)
	ctx.SetCookie(
		appcontext.LocaleCode,
		*val, appcontext.LocaleDuration,
		"/",
		env.CONFIG.Domain,
		useSecureConnection,
		true,
	)
}

func SetLocale(ctx *gin.Context, val interface{}) {
	if val == nil {
		newVal := appcontext.LocaleDefault
		val = &newVal
	}

	appcontext.SetFromGinRouter(ctx, appcontext.LocaleCode, val)
}
