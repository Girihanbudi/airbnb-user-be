package rpc

import (
	"airbnb-user-be/internal/app/country/usecase"
)

type Options struct {
	Country usecase.ICountry
}

type Handler struct {
	Options
	UnimplementedCountryServiceServer
}

func NewCountryHandler(options Options) CountryServiceServer {
	return Handler{
		Options: options,
	}
}
