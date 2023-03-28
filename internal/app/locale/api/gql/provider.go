package gql

import (
	"airbnb-user-be/internal/app/locale/usecase"
)

type Options struct {
	Locale usecase.ILocale
}

type Handler struct {
	Options
}

func NewLocaleHandler(options Options) *Handler {
	return &Handler{options}
}
