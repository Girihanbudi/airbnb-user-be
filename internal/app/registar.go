package app

import (
	_ "airbnb-user-be/docs"
	"airbnb-user-be/graph"
	gqllocale "airbnb-user-be/internal/app/locale/api/gql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// Defining the Graphql handler
func graphqlHandler(localeHandler gqllocale.Handler) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{
			Locale: localeHandler,
		}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
// func graphqlPlaygroundHandler() gin.HandlerFunc {
// 	h := playground.Handler("GraphQL", "/query")
// 	return func(c *gin.Context) {
// 		h.ServeHTTP(c.Writer, c.Request)
// 	}
// }

func (a App) RegisterHttpHandler() {

	a.HttpServer.Router.POST("/graph", graphqlHandler(*a.LocaleGqlHandler))
	// a.HttpServer.Router.GET("/", graphqlPlaygroundHandler())
	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}
