package util

import (
	"airbnb-user-be/internal/app/translation/repo/repoimpl"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
	"net/http"
)

var (
	defaultErr = stderror.New(http.StatusInternalServerError, "DEF_SERVER_500", "Failed to get translation")
)

func TranslateError(ctx context.Context, code, localeCode string) (err *stderror.StdError) {
	trans, getTransErr := repoimpl.ErrTranslationRepo.GetErrTranslation(ctx, code, localeCode)
	if getTransErr != nil {
		err = &defaultErr
		return
	}
	newErr := stderror.New(trans.HttpCode, trans.Code, trans.Message)
	err = &newErr
	return
}
