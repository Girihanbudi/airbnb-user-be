package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	transutil "airbnb-user-be/internal/app/translation/util"
	module "airbnb-user-be/internal/app/user"
	errpreset "airbnb-user-be/internal/app/user/preset/error"
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/app/user/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"airbnb-user-be/internal/pkg/util"
	"context"
)

func (u Usecase) CreateUser(ctx context.Context, cmd request.CreateUser) (res response.CreateUser, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	user := module.User{
		FirstName:   util.Case(cmd.FirstName, util.CaseLower, util.CaseTitle),
		FullName:    util.Case(cmd.FirstName+" "+cmd.LastName, util.CaseLower, util.CaseTitle),
		Email:       cmd.Email,
		CountryCode: cmd.CountryCode,
		PhoneNumber: cmd.PhoneNumber,
		Image:       cmd.Image,
		Role:        cmd.Role,
		DateOfBirth: cmd.DateOfBirth,
	}
	if cmd.DefaultSetting != nil {
		user.DefaultSetting = &module.UserDefaultSetting{
			Locale:   cmd.DefaultSetting.Locale,
			Currency: cmd.DefaultSetting.Currency,
		}
	}
	createUserErr := u.UserRepo.CreateUser(ctx, &user)
	if createUserErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}
	res.Id = user.Id

	return
}
