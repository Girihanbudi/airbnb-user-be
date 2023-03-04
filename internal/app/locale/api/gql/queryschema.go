package gql

import (
	"airbnb-user-be/internal/app/locale"

	"github.com/graphql-go/graphql"
)

func (h Handler) QuerySchema() map[string]*graphql.Field {

	region := graphql.Field{
		Type:        locale.RegionType,
		Description: "Get a region",
		Args: graphql.FieldConfigArgument{
			"code": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: h.GetRegion,
	}

	regions := graphql.Field{
		Type:        graphql.NewList(locale.RegionType),
		Description: "List of regions",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"page": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: h.GetRegions,
	}

	return map[string]*graphql.Field{
		"region":  &region,
		"regions": &regions,
	}
}
