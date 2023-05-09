package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	transutil "airbnb-user-be/internal/app/translation/util"
	module "airbnb-user-be/internal/app/user"
	errpreset "airbnb-user-be/internal/app/user/preset/error"
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/pkg/stderror"
	"airbnb-user-be/internal/pkg/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (u Usecase) UpdateUser(ctx context.Context, cmd request.UpdateUser) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	user, getUserErr := u.UserRepo.GetUser(ctx, cmd.Id, "DefaultSetting")
	if getUserErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getUserErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	user.FirstName = util.Case(cmd.FirstName, util.CaseLower, util.CaseTitle)
	user.FullName = util.Case(cmd.FirstName+" "+cmd.LastName, util.CaseLower, util.CaseTitle)
	user.DateOfBirth = &cmd.DateOfBirth

	if cmd.DefaultSetting != nil {
		defaultSetting := module.UserDefaultSetting{
			Locale:   cmd.DefaultSetting.Locale,
			Currency: cmd.DefaultSetting.Currency,
		}
		user.DefaultSetting = &defaultSetting
	}

	updateUserErr := u.UserRepo.CreateOrUpdateUser(ctx, &user)
	if updateUserErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	return
}
