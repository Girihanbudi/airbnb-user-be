package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	transutil "airbnb-user-be/internal/app/translation/util"
	errpreset "airbnb-user-be/internal/app/user/preset/error"
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (u Usecase) VerifiedUser(ctx context.Context, cmd request.Identifier) (err *stderror.StdError) {
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

	currentTime := time.Now()
	user.VerifiedAt = &currentTime

	updateUserErr := u.UserRepo.CreateOrUpdateUser(ctx, &user)
	if updateUserErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	return
}
