package app

import (
	_ "airbnb-user-be/docs"
	"airbnb-user-be/graph"
	gqlcurrency "airbnb-user-be/internal/app/currency/api/gql"
	gqllocale "airbnb-user-be/internal/app/locale/api/gql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// Defining the Graphql handler
func graphqlHandler(localeHandler gqllocale.Handler, currencyHandler gqlcurrency.Handler) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{
			Locale:   localeHandler,
			Currency: currencyHandler,
		}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (a App) RegisterHttpHandler() {
	// register modules to graph solver handler
	a.HttpServer.Router.GET("/graph", graphqlHandler(
		*a.LocaleGqlHandler,
		*a.CurrencyGqlHandler,
	))
	// a.HttpServer.Router.GET("/", graphqlPlaygroundHandler())
	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}
