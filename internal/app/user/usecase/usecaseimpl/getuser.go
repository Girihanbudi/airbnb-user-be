package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	transutil "airbnb-user-be/internal/app/translation/util"
	module "airbnb-user-be/internal/app/user"
	errpreset "airbnb-user-be/internal/app/user/preset/error"
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/app/user/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetUser(ctx context.Context, cmd request.GetUser) (res response.GetUser, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	var user module.User
	var getUserErr error
	preloads := []string{"DefaultSetting"}
	if cmd.Id != nil {
		user, getUserErr = u.UserRepo.GetUser(ctx, *cmd.Id, preloads...)
		if getUserErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}
	} else if cmd.CountryCode != nil && cmd.PhoneNumber != nil {
		user, getUserErr = u.UserRepo.GetUserByPhone(ctx, *cmd.CountryCode, *cmd.PhoneNumber, preloads...)
		if getUserErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}
	} else if cmd.Email != nil {
		user, getUserErr = u.UserRepo.GetUserByEmail(ctx, *cmd.Email, preloads...)
		if getUserErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}
	}

	res.Id = user.Id
	res.FirstName = user.FirstName
	res.FullName = user.FullName
	res.Email = user.Email
	res.CountryCode = user.CountryCode
	res.PhoneNumber = user.PhoneNumber
	res.Image = user.Image
	res.Role = user.Role
	res.DateOfBirth = user.DateOfBirth
	res.CreatedAt = user.CreatedAt
	res.UpdatedAt = user.UpdatedAt
	res.VerifiedAt = user.VerifiedAt
	res.DeletedAt = &user.DeletedAt.Time

	userDefaultSetting := user.DefaultSetting
	if userDefaultSetting != nil {
		resDefaultSetting := response.UserDefaultSetting{
			Locale:   userDefaultSetting.Locale,
			Currency: userDefaultSetting.Currency,
		}
		res.DefaultSetting = &resDefaultSetting
	}

	return
}
