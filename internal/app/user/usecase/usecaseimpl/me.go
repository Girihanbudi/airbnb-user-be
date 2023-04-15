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

	"github.com/thoas/go-funk"
)

func (u Usecase) Me(ctx context.Context, cmd request.Me) (res response.Me, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	preloads := []string{"DefaultSetting", "Accounts"}
	user, getUserErr := u.UserRepo.GetUser(ctx, clientLocale, &preloads)
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

	if user.Accounts != nil {
		res.Accounts = funk.Map(*user.Accounts, func(acc module.Account) response.Account {
			return response.Account{
				Provider:     acc.Provider,
				AccessToken:  acc.AccessToken,
				RefreshToken: acc.RefreshToken,
				ExpiredAt:    acc.ExpiredAt,
				TokenType:    acc.TokenType,
				CreatedAt:    acc.CreatedAt,
				UpdatedAt:    acc.UpdatedAt,
			}
		}).(*[]response.Account)
	}

	return
}
