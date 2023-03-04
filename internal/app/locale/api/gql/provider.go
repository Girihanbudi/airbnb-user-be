package gql

import (
	locale "airbnb-user-be/internal/app/locale/usecase"
)

type Options struct {
	Region locale.IRegion
}

type Handler struct {
	Options
}

func ProvideLocaleHandler(options Options) *Handler {
	return &Handler{options}
}
