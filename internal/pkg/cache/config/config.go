package config

import (
	auth "airbnb-user-be/internal/pkg/cache/auth/config"
)

type Config struct {
	Auth auth.Config `mapstructure:"auth"`
}
