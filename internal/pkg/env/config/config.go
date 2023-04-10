package config

import (
	cache "airbnb-user-be/internal/pkg/cache/config"
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpserver "airbnb-user-be/internal/pkg/http/server/config"
	jwt "airbnb-user-be/internal/pkg/jwt/config"
	kafka "airbnb-user-be/internal/pkg/kafka/config"
	oauth "airbnb-user-be/internal/pkg/oauth/config"
)

type Config struct {
	Origins    []string          `mapstructure:"origins"`
	Stage      string            `mapstructure:"stage"`
	Domain     string            `mapstructure:"domain"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	DB         gorm.Config       `mapstructure:"db"`
	Oauth      oauth.Config      `mapstructure:"oauth"`
	Jwt        jwt.Config        `mapstructure:"jwt"`
	Cache      cache.Config      `mapstructure:"cache"`
	Kafka      kafka.Config      `mapstructure:"kafka"`
}
