package cookie

import (
	"airbnb-user-be/env/appcontext"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

func BindLocale(ctx *gin.Context) {
	locale, err := ctx.Cookie(appcontext.LocaleCode)
	if err != nil {
		CreateLocale(ctx, nil)
		SetLocale(ctx, nil)
		return
	}

	SetLocale(ctx, locale)
}

func CreateLocale(ctx *gin.Context, val *string) {
	if val == nil {
		newVal := appcontext.LocaleDefault
		val = &newVal
	}

	ctx.SetSameSite(http.DefaultSameSite())

	ctx.SetCookie(
		appcontext.LocaleCode,
		*val,
		appcontext.LocaleDuration,
		"/",
		env.CONFIG.Domain,
		env.CONFIG.Stage != string(env.StageLocal),
		false,
	)
}

func SetLocale(ctx *gin.Context, val interface{}) {
	if val == nil {
		appcontext.SetFromGinRouter(ctx, appcontext.LocaleCode, appcontext.LocaleDefault)
	} else {
		appcontext.SetFromGinRouter(ctx, appcontext.LocaleCode, val)
	}
}
