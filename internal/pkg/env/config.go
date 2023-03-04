package env

import (
	"airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/http/server"
)

type Config struct {
	DB         gorm.Config
	HttpServer server.Config
}
