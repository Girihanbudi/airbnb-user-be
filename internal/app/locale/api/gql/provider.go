package gql

import (
	locale "airbnb-user-be/internal/app/locale/usecase"
)

type Options struct {
	Locale locale.ILocale
}

type Handler struct {
	Options
}

func NewLocaleHandler(options Options) *Handler {
	return &Handler{options}
}
