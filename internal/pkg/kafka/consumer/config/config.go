package config

import (
	router "airbnb-user-be/internal/pkg/kafka/router/config"
)

type Config struct {
	Group       string        `mapstructure:"group"`
	Assigner    string        `mapstructure:"assigner"`
	IsUseOldest bool          `mapstructure:"isuseoldest"`
	Router      router.Config `mapstructure:"router"`
}
