package config

import (
	google "airbnb-user-be/internal/pkg/oauth/google/config"
)

type Config struct {
	RedirectUrl string        `mapstructure:"redirecturl"`
	Google      google.Config `mapstructure:"google"`
}
