package app

import (
	"airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/http/server"
	httprouter "airbnb-user-be/internal/pkg/http/server/router"
	"airbnb-user-be/internal/pkg/log"
	"airbnb-user-be/internal/pkg/validator"
	"context"
	"sync"

	"github.com/gin-gonic/gin"

	authrest "airbnb-user-be/internal/app/auth/api/rest"
	currencygql "airbnb-user-be/internal/app/currency/api/gql"
	localegql "airbnb-user-be/internal/app/locale/api/gql"
	authmid "airbnb-user-be/internal/app/middleware/auth"
	cookiemid "airbnb-user-be/internal/app/middleware/cookie"
	translation "airbnb-user-be/internal/app/translation/repo"
)

var Instance = "App"

type Options struct {
	HttpServer *server.Server

	Translation        translation.IErrTranslation
	AuthHandler        *authrest.Handler
	LocaleGqlHandler   *localegql.Handler
	CurrencyGqlHandler *currencygql.Handler
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
	log.Event(Instance, "Starting...")

	// init app validator
	validator.InitValidator()

	// init app cache
	auth.InitAuthCache()

	// recover from panic
	a.HttpServer.Router.Use(gin.Recovery())

	// GIN apply CORS setting
	a.HttpServer.Router.Use(httprouter.DefaultCORSSetting())

	// GIN bind all cookie
	a.HttpServer.Router.Use(cookiemid.BindAll())

	// GIN bind access token if any
	// bind access token in all route to adapt with graphql endpoint
	a.HttpServer.Router.Use(authmid.GinBindBearerAuthorization())

	// Register all routes
	a.registerHttpHandler()

	go func() {
		err := a.HttpServer.Start()
		if err != nil {
			log.Fatal(Instance, "failed to start http server", err)
		}
	}()

	<-ctx.Done()
}

func (a App) stopModules() {
	log.Event(Instance, "Stoping...")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.HttpServer.Stop()
		if err != nil {
			log.Fatal(Instance, "failed to stop http server", err)
		}
	}()

	wg.Wait()
}
