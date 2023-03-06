package usecaseimpl

import (
	"airbnb-user-be/internal/app/locale/repo"
)

type Options struct {
	LocaleRepo repo.ILocale
}

type Usecase struct {
	Options
}

func NewLocaleUsecase(options Options) *Usecase {
	return &Usecase{options}
}
