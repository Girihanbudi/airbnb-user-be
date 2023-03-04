package main

import (
	"airbnb-user-be/internal/app"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title           Airbnb User Backend API
// @version         1.0
// @description     Airbnb User Backend Service API
// @termsOfService  https://airbnb.co.id

// @contact.name   API Support
// @contact.url    https://airbnb.co.id/support
// @contact.email  support@airbnb.co.id

// @host      localhost:8080
// @BasePath  /v1

// @securityDefinitions.basic BasicAuth
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	serviceApp, err := app.ProvideApp()
	if err != nil {
		log.Panic(err)
	}

	serviceApp.RegisterHttpHandler()
	serviceApp.Run(ctx)
	stop()
}
