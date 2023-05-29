package app

import (
	"airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/cache/otp"
	elastic "airbnb-user-be/internal/pkg/elasticsearch"
	grpcserver "airbnb-user-be/internal/pkg/grpcserver"
	"airbnb-user-be/internal/pkg/http/server"
	httprouter "airbnb-user-be/internal/pkg/http/server/router"
	kafkaconsumer "airbnb-user-be/internal/pkg/kafka/consumer"
	kafkaproducer "airbnb-user-be/internal/pkg/kafka/producer"
	"airbnb-user-be/internal/pkg/log"
	"context"
	"sync"

	"github.com/gin-gonic/gin"

	countrygql "airbnb-user-be/internal/app/country/api/gql"
	countryrpc "airbnb-user-be/internal/app/country/api/rpc"
	currencygql "airbnb-user-be/internal/app/currency/api/gql"
	localegql "airbnb-user-be/internal/app/locale/api/gql"
	localerpc "airbnb-user-be/internal/app/locale/api/rpc"
	authmid "airbnb-user-be/internal/app/middleware/auth"
	cookiemid "airbnb-user-be/internal/app/middleware/cookie"
	elasticmid "airbnb-user-be/internal/app/middleware/elastic"
	translation "airbnb-user-be/internal/app/translation/repo"
	usergql "airbnb-user-be/internal/app/user/api/gql"
	userrest "airbnb-user-be/internal/app/user/api/rest"
	userrpc "airbnb-user-be/internal/app/user/api/rpc"

	"airbnb-user-be/internal/pkg/credential"
)

var Instance = "App"

type Options struct {
	TlsCreds      credential.TlsCredentials
	HttpServer    *server.Server
	RpcServer     *grpcserver.Server
	EventListener *kafkaconsumer.Listener
	EventProducer *kafkaproducer.Producer
	Translation   translation.ITranslation

	UserRestHandler *userrest.Handler

	CountryGqlHandler  *countrygql.Handler
	LocaleGqlHandler   *localegql.Handler
	CurrencyGqlHandler *currencygql.Handler
	UserGqlHandler     *usergql.Handler

	UserRpcHandler    userrpc.UserServiceServer
	LocaleRpcHandler  localerpc.LocaleServiceServer
	CountryRpcHandler countryrpc.CountryServiceServer
}

type App struct {
	Options
}

// Run all the modules of the app.
func (a App) Run(ctx context.Context) {
	a.runModules(ctx)
	a.stopModules()
}

func (a App) runModules(ctx context.Context) {
	log.Event(Instance, "Starting service and connections...")

	// Init app cache
	auth.InitAuthCache()
	otp.InitOtpCache()

	// Init elasticsearch client
	elastic.InitElasticSearch()

	// Create required index in elastic
	elasticmid.CreateIndex()

	// Recover from panic
	a.HttpServer.Router.Use(gin.Recovery())

	// GIN apply CORS setting
	a.HttpServer.Router.Use(httprouter.DefaultCORSSetting())

	// GIN log request and response to elastic
	a.HttpServer.Router.Use(elasticmid.LogRequestToElastic())

	// GIN bind all cookie
	a.HttpServer.Router.Use(cookiemid.BindAll())

	// GIN bind access token if any
	// bind access token in all route to adapt with graphql endpoint
	a.HttpServer.Router.Use(authmid.GinBindAccessToken())

	// Register all routes
	a.registerRpcHandler()
	a.registerHttpHandler()

	go func() {
		a.HttpServer.Start()
	}()

	go func() {
		a.RpcServer.Start()
	}()

	<-ctx.Done()
}

func (a App) stopModules() {
	log.Event(Instance, "Stoping service and connections...")

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		a.EventProducer.Stop()
	}()

	go func() {
		defer wg.Done()
		a.RpcServer.Stop()
	}()

	go func() {
		defer wg.Done()
		a.HttpServer.Stop()
	}()

	wg.Wait()
	log.Event(Instance, "successfully stopped service and connections")
}
