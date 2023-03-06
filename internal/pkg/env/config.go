package env

import (
	"airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/http/server"
)

type Config struct {
	HttpServer server.Config `mapstructure:"httpserver"`
	DB         gorm.Config   `mapstructure:"db"`
}
