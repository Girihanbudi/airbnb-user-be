package gql

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	"context"
)

func (h Handler) GetLocale(ctx context.Context, code string) (*response.GetLocale, error) {
	cmd := request.GetLocale{
		Code: code,
	}

	res, err := h.Locale.GetLocale(ctx, cmd)
	if err != nil {
		return nil, err.Error
	}

	return &res, nil
}

func (h Handler) GetLocales(ctx context.Context) (*response.GetLocales, error) {

	res, err := h.Locale.GetLocales(ctx)
	if err != nil {
		return nil, err.Error
	}

	return &res, nil
}
