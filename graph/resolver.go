package graph

import (
	locale "airbnb-user-be/internal/app/locale/api/gql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Locale locale.Handler
}
