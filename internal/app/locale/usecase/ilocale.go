package usecase

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

type ILocale interface {
	GetLocales(ctx context.Context) (res response.GetLocales, err *stderror.StdError)
	GetLocale(ctx context.Context, cmd request.GetLocale) (res response.GetLocale, err *stderror.StdError)
}
