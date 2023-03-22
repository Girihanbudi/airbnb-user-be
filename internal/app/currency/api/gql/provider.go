package gql

import (
	currency "airbnb-user-be/internal/app/currency/usecase"
)

type Options struct {
	Currency currency.ICurrency
}

type Handler struct {
	Options
}

func NewCurrencyHandler(options Options) *Handler {
	return &Handler{options}
}
