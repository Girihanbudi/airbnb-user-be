package config

import (
	google "airbnb-user-be/internal/pkg/oauth/google/config"
)

type Config struct {
	Google google.Config `mapstructure:"google"`
}
