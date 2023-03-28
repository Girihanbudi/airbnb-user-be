package usecaseimpl

import (
	"airbnb-user-be/internal/app/country/repo"
)

type Options struct {
	CountryRepo repo.ICountry
}

type Usecase struct {
	Options
}

func NewCountryUsecase(options Options) *Usecase {
	return &Usecase{options}
}
