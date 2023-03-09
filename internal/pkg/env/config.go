package env

import (
	"airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/http/server"
)

type Config struct {
	Stage      string        `mapstructure:"stage"`
	Domain     string        `mapstructure:"domain"`
	HttpServer server.Config `mapstructure:"httpserver"`
	DB         gorm.Config   `mapstructure:"db"`
}
