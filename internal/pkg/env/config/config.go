package config

import (
	cache "airbnb-user-be/internal/pkg/cache/config"
	credential "airbnb-user-be/internal/pkg/credential/config"
	elastic "airbnb-user-be/internal/pkg/elasticsearch/config"
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpserver "airbnb-user-be/internal/pkg/http/server/config"
	jwt "airbnb-user-be/internal/pkg/jwt/config"
	kafka "airbnb-user-be/internal/pkg/kafka/config"
)

type Config struct {
	Stage      string            `mapstructure:"stage"`
	Origins    []string          `mapstructure:"origins"`
	Domain     string            `mapstructure:"domain"`
	Creds      credential.Config `mapstructure:"creds"`
	HttpServer httpserver.Config `mapstructure:"httpserver"`
	DB         gorm.Config       `mapstructure:"db"`
	Jwt        jwt.Config        `mapstructure:"jwt"`
	Cache      cache.Config      `mapstructure:"cache"`
	Elastic    elastic.Config    `mapstructure:"elastic"`
	Kafka      kafka.Config      `mapstructure:"kafka"`
}
