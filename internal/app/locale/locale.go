package locale

import "github.com/graphql-go/graphql"

type Locale struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Local    string `json:"local"`
	Location string `json:"location"`
	LCID     int    `json:"lcid"`
	ISO639_2 string `json:"iso639_2"`
	ISO639_1 string `json:"iso639_1"`
}

var LocaleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Locale",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"local": &graphql.Field{
			Type: graphql.String,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"lcid": &graphql.Field{
			Type: graphql.Int,
		},
		"iso639_2": &graphql.Field{
			Type: graphql.String,
		},
		"iso639_1": &graphql.Field{
			Type: graphql.String,
		},
	},
})
