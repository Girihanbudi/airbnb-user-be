package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	transutil "airbnb-user-be/internal/app/translation/util"
	errpreset "airbnb-user-be/internal/app/user/preset/error"
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u Usecase) ChangeCurrencySetting(ctx context.Context, cmd request.ChangeCurrencySetting) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	// Get selected currency
	_, getCurrencyErr := u.CurrencyRepo.GetCurrency(ctx, cmd.Currency)
	if getCurrencyErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getCurrencyErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	// Change currency in user setting
	if cmd.UserId != nil {
		setting, getSettingErr := u.UserRepo.GetDefaultSettingByUser(ctx, *cmd.UserId)
		if getSettingErr != nil {
			ec := errpreset.DbServiceUnavailable
			if errors.Is(getSettingErr, gorm.ErrRecordNotFound) {
				ec = errpreset.DbRecordNotFound
			}
			err = transutil.TranslateError(ctx, ec, clientLocale)
			return
		}

		setting.Currency = cmd.Currency

		updateSettingErr := u.UserRepo.CreateOrUpdateDefaultSetting(ctx, &setting)
		if updateSettingErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}
	}

	return
}
