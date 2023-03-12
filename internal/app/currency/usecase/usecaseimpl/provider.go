package usecaseimpl

import (
	"airbnb-user-be/internal/app/currency/repo"
)

type Options struct {
	CurrencyRepo repo.ICurrency
}

type Usecase struct {
	Options
}

func NewCurrencyUsecase(options Options) *Usecase {
	return &Usecase{options}
}
