package config

import (
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpserver "airbnb-user-be/internal/pkg/http/server/config"
	oauth "airbnb-user-be/internal/pkg/oauth/config"
)

type Config struct {
	Origins    []string          `mapstructure:"origins"`
	Stage      string            `mapstructure:"stage"`
	Domain     string            `mapstructure:"domain"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	DB         gorm.Config       `mapstructure:"db"`
	Oauth      oauth.Config      `mapstructure:"oauth"`
}
