package app

import (
	"airbnb-user-be/internal/pkg/http/server"
	"airbnb-user-be/internal/pkg/log"
	"airbnb-user-be/internal/pkg/validator"
	"context"
	"sync"

	localegql "airbnb-user-be/internal/app/locale/api/gql"
	translation "airbnb-user-be/internal/app/translation/repo"
)

var Instance = "App"

type Options struct {
	HttpServer *server.Server

	Translation      translation.IErrTranslation
	LocaleGqlHandler *localegql.Handler
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
