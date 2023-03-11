package util

import (
	"airbnb-user-be/internal/app/translation/repo/repoimpl"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func TranslateError(ctx context.Context, code, localeCode string) (err *stderror.StdError) {
	trans, getTransErr := repoimpl.ErrTranslationRepo.GetErrTranslation(ctx, code, localeCode)
	if getTransErr != nil {
		err = stderror.DEF_SERVER_500.ErrorMsg(getTransErr)
		return
	}
	newErr := stderror.New(trans.HttpCode, trans.Code, trans.Message)
	err = &newErr
	return
}
