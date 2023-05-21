package rpc

import (
	"airbnb-user-be/internal/app/locale/usecase"
)

type Options struct {
	Locale usecase.ILocale
}

type Handler struct {
	Options
	UnimplementedLocaleServiceServer
}

func NewLocaleHandler(options Options) LocaleServiceServer {
	return Handler{
		Options: options,
	}
}
