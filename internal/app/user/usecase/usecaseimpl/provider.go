package usecaseimpl

import (
	currency "airbnb-user-be/internal/app/currency/repo"
	locale "airbnb-user-be/internal/app/locale/repo"
	user "airbnb-user-be/internal/app/user/repo"
)

type Options struct {
	UserRepo     user.IUser
	LocaleRepo   locale.ILocale
	CurrencyRepo currency.ICurrency
}

type Usecase struct {
	Options
}

func NewUserUsecase(options Options) *Usecase {
	return &Usecase{options}
}
