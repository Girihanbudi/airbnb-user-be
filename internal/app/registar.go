package app

import (
	_ "airbnb-user-be/docs"
	"airbnb-user-be/graph"
	gqlcountry "airbnb-user-be/internal/app/country/api/gql"
	rpccountry "airbnb-user-be/internal/app/country/api/rpc"
	gqlcurrency "airbnb-user-be/internal/app/currency/api/gql"
	gqllocale "airbnb-user-be/internal/app/locale/api/gql"
	rpclocale "airbnb-user-be/internal/app/locale/api/rpc"
	gqluser "airbnb-user-be/internal/app/user/api/gql"
	rpcuser "airbnb-user-be/internal/app/user/api/rpc"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

// Defining the Graphql handler
func graphqlHandler(countryHandler gqlcountry.Handler, localeHandler gqllocale.Handler, currencyHandler gqlcurrency.Handler, userHandler gqluser.Handler) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(
		graph.Config{Resolvers: &graph.Resolver{
			Country:  countryHandler,
			Locale:   localeHandler,
			Currency: currencyHandler,
			User:     userHandler,
		}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (a App) registerRpcHandler() {
	rpcuser.RegisterUserServiceServer(a.RpcServer.Server, a.UserRpcHandler)
	rpclocale.RegisterLocaleServiceServer(a.RpcServer.Server, a.LocaleRpcHandler)
	rpccountry.RegisterCountryServiceServer(a.RpcServer.Server, a.CountryRpcHandler)
}

func (a App) registerHttpHandler() {

	// register modules to graph solver handler
	a.HttpServer.Router.GET("/graph", graphqlHandler(
		*a.CountryGqlHandler,
		*a.LocaleGqlHandler,
		*a.CurrencyGqlHandler,
		*a.UserGqlHandler,
	))

	a.HttpServer.Router.GET("/docs/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
}
