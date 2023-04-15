package graph

import (
	country "airbnb-user-be/internal/app/country/api/gql"
	currency "airbnb-user-be/internal/app/currency/api/gql"
	locale "airbnb-user-be/internal/app/locale/api/gql"
	user "airbnb-user-be/internal/app/user/api/gql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Country  country.Handler
	Locale   locale.Handler
	Currency currency.Handler
	User     user.Handler
}
