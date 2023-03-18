package cookie

import (
	"airbnb-user-be/internal/pkg/appcontext"
	"airbnb-user-be/internal/pkg/env"

	"github.com/gin-gonic/gin"
)

func BindCurrency(ctx *gin.Context) {
	currency, err := ctx.Cookie(appcontext.CurrencyCode)
	if err != nil {
		CreateCurrency(ctx, nil)
		SetCurrency(ctx, nil)
		return
	}

	SetCurrency(ctx, currency)
}

func CreateCurrency(ctx *gin.Context, val *string) {
	if val == nil {
		newVal := appcontext.CurrencyDefault
		val = &newVal
	}

	ctx.SetCookie(
		appcontext.CurrencyCode,
		*val, appcontext.CurrencyDuration,
		"/",
		env.CONFIG.Domain,
		env.CONFIG.Stage != string(env.StageLocal),
		false,
	)
}

func SetCurrency(ctx *gin.Context, val interface{}) {
	if val == nil {
		appcontext.SetFromGinRouter(ctx, appcontext.CurrencyCode, appcontext.CurrencyDefault)
	} else {
		appcontext.SetFromGinRouter(ctx, appcontext.CurrencyCode, val)
	}
}
