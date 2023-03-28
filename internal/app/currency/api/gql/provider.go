package gql

import (
	"airbnb-user-be/internal/app/currency/usecase"
)

type Options struct {
	Currency usecase.ICurrency
}

type Handler struct {
	Options
}

func NewCurrencyHandler(options Options) *Handler {
	return &Handler{options}
}
