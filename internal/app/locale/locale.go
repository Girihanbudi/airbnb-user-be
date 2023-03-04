package locale

import "github.com/graphql-go/graphql"

type Region struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Region   string `json:"region"`
}

var RegionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Region",
	Fields: graphql.Fields{
		"code": &graphql.Field{
			Type: graphql.String,
		},
		"language": &graphql.Field{
			Type: graphql.String,
		},
		"region": &graphql.Field{
			Type: graphql.String,
		},
	},
})
