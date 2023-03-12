package config

import (
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpserver "airbnb-user-be/internal/pkg/http/server/config"
)

type Config struct {
	Stage      string            `mapstructure:"stage"`
	Domain     string            `mapstructure:"domain"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	DB         gorm.Config       `mapstructure:"db"`
}
