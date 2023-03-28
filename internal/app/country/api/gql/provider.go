package gql

import (
	"airbnb-user-be/internal/app/country/usecase"
)

type Options struct {
	Country usecase.ICountry
}

type Handler struct {
	Options
}

func NewCountryHandler(options Options) *Handler {
	return &Handler{options}
}
