package gql

import (
	"airbnb-user-be/internal/app/locale"

	"github.com/graphql-go/graphql"
)

func (h Handler) QuerySchema() map[string]*graphql.Field {

	Locale := graphql.Field{
		Type:        locale.LocaleType,
		Description: "Get a Locale",
		Args: graphql.FieldConfigArgument{
			"code": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: h.GetLocale,
	}

	Locales := graphql.Field{
		Type:        graphql.NewList(locale.LocaleType),
		Description: "List of Locales",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: h.GetLocales,
	}

	return map[string]*graphql.Field{
		"Locale":  &Locale,
		"Locales": &Locales,
	}
}
