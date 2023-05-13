// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"airbnb-user-be/internal/app/country/api/gql"
	repoimpl2 "airbnb-user-be/internal/app/country/repo/repoimpl"
	"airbnb-user-be/internal/app/country/usecase/usecaseimpl"
	gql3 "airbnb-user-be/internal/app/currency/api/gql"
	repoimpl4 "airbnb-user-be/internal/app/currency/repo/repoimpl"
	usecaseimpl3 "airbnb-user-be/internal/app/currency/usecase/usecaseimpl"
	gql2 "airbnb-user-be/internal/app/locale/api/gql"
	repoimpl3 "airbnb-user-be/internal/app/locale/repo/repoimpl"
	usecaseimpl2 "airbnb-user-be/internal/app/locale/usecase/usecaseimpl"
	"airbnb-user-be/internal/app/translation/repo/repoimpl"
	gql4 "airbnb-user-be/internal/app/user/api/gql"
	"airbnb-user-be/internal/app/user/api/rpc"
	repoimpl5 "airbnb-user-be/internal/app/user/repo/repoimpl"
	usecaseimpl4 "airbnb-user-be/internal/app/user/usecase/usecaseimpl"
	"airbnb-user-be/internal/pkg/credential"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/env/tool"
	"airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/grpc"
	"airbnb-user-be/internal/pkg/http/server"
	"airbnb-user-be/internal/pkg/http/server/router"
	"airbnb-user-be/internal/pkg/kafka"
	"airbnb-user-be/internal/pkg/kafka/consumer"
	"airbnb-user-be/internal/pkg/kafka/producer"
	router2 "airbnb-user-be/internal/pkg/kafka/router"
	"github.com/google/wire"
)

import (
	_ "airbnb-user-be/docs"
)

// Injectors from wire.go:

func NewApp() (*App, error) {
	config := env.ProvideEnv()
	configConfig := tool.ExtractServerConfig(config)
	engine := router.NewRouter()
	config2 := tool.ExtractCredsConfig(config)
	options := credential.Options{
		Config: config2,
	}
	tlsCredentials := credential.NewTLSCredentials(options)
	serverOptions := server.Options{
		Config: configConfig,
		Router: engine,
		Creds:  tlsCredentials,
	}
	serverServer := server.NewServer(serverOptions)
	grpcOptions := grpc.Options{
		Creds: tlsCredentials,
	}
	grpcServer := grpc.NewRpcServer(grpcOptions)
	config3 := tool.ExtractKafkaConsumerConfig(config)
	config4 := tool.ExtractKafkaConfig(config)
	config5 := tool.ExtractKafkaRouterConfig(config)
	routerOptions := router2.Options{
		Config: config5,
	}
	routerRouter := router2.NewRouter(routerOptions)
	kafkaOptions := kafka.Options{
		Config: config4,
		Router: routerRouter,
	}
	client := kafka.NewSaramaClient(kafkaOptions)
	consumerOptions := consumer.Options{
		Config: config3,
		Client: client,
		Router: routerRouter,
	}
	listener := consumer.NewEventListener(consumerOptions)
	producerOptions := producer.Options{
		Client: client,
	}
	producerProducer := producer.NewEventProducer(producerOptions)
	config6 := tool.ExtractDBConfig(config)
	gormOptions := gorm.Options{
		Config: config6,
	}
	gormEngine := gorm.NewORM(gormOptions)
	repoimplOptions := repoimpl.Options{
		Gorm: gormEngine,
	}
	repo := repoimpl.NewTranslationRepo(repoimplOptions)
	options2 := repoimpl2.Options{
		Gorm: gormEngine,
	}
	repoimplRepo := repoimpl2.NewCountryRepo(options2)
	usecaseimplOptions := usecaseimpl.Options{
		CountryRepo: repoimplRepo,
	}
	usecase := usecaseimpl.NewCountryUsecase(usecaseimplOptions)
	gqlOptions := gql.Options{
		Country: usecase,
	}
	handler := gql.NewCountryHandler(gqlOptions)
	options3 := repoimpl3.Options{
		Gorm: gormEngine,
	}
	repo2 := repoimpl3.NewLocaleRepo(options3)
	options4 := usecaseimpl2.Options{
		LocaleRepo: repo2,
	}
	usecaseimplUsecase := usecaseimpl2.NewLocaleUsecase(options4)
	options5 := gql2.Options{
		Locale: usecaseimplUsecase,
	}
	gqlHandler := gql2.NewLocaleHandler(options5)
	options6 := repoimpl4.Options{
		Gorm: gormEngine,
	}
	repo3 := repoimpl4.NewCurrencyRepo(options6)
	options7 := usecaseimpl3.Options{
		CurrencyRepo: repo3,
	}
	usecase2 := usecaseimpl3.NewCurrencyUsecase(options7)
	options8 := gql3.Options{
		Currency: usecase2,
	}
	handler2 := gql3.NewCurrencyHandler(options8)
	options9 := repoimpl5.Options{
		Gorm: gormEngine,
	}
	repo4 := repoimpl5.NewUserRepo(options9)
	options10 := usecaseimpl4.Options{
		UserRepo: repo4,
	}
	usecase3 := usecaseimpl4.NewUserUsecase(options10)
	options11 := gql4.Options{
		User: usecase3,
	}
	handler3 := gql4.NewUserHandler(options11)
	rpcOptions := rpc.Options{
		User: usecase3,
	}
	userServiceServer := rpc.NewUserHandler(rpcOptions)
	appOptions := Options{
		HttpServer:         serverServer,
		RpcServer:          grpcServer,
		EventListener:      listener,
		EventProducer:      producerProducer,
		Translation:        repo,
		CountryHandler:     handler,
		LocaleGqlHandler:   gqlHandler,
		CurrencyGqlHandler: handler2,
		UserGqlHandler:     handler3,
		UserRpcHandler:     userServiceServer,
	}
	app := &App{
		Options: appOptions,
	}
	return app, nil
}

// wire.go:

var AppSet = wire.NewSet(wire.Struct(new(Options), "*"), wire.Struct(new(App), "*"))
