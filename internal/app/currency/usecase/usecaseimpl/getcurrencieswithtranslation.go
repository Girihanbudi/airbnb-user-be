package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/currency/preset/error"
	"airbnb-user-be/internal/app/currency/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetCurrenciesWithTranslation(ctx context.Context) (res response.GetCurrencyWithTranslation, err *stderror.StdError) {
	clientLocale := ctx.Value(appcontext.LocaleCode).(string)

	currencies, getCurrenciesErr := u.CurrencyRepo.GetCurrenciesWithTranslation(ctx, clientLocale)
	if getCurrenciesErr != nil {
		err = transutil.TranslateError(ctx, errpreset.CURRENCY_GET_500, clientLocale)
		return
	}

	res.Currencies = currencies

	return
}
