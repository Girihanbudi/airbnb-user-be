package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	transutil "airbnb-user-be/internal/app/translation/util"
	errpreset "airbnb-user-be/internal/app/user/preset/error"
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/app/user/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) Me(ctx context.Context, cmd request.Me) (res response.Me, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	user, getUserErr := u.UserRepo.GetUser(ctx, cmd.UserId, "DefaultSetting", "Accounts")
	if getUserErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

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
