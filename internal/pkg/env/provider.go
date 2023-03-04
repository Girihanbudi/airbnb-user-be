package env

import (
	"airbnb-user-be/internal/pkg/gorm"
	httpServer "airbnb-user-be/internal/pkg/http/server"
	"airbnb-user-be/internal/pkg/log"

	"github.com/spf13/viper"
)

var Instance string = "Env"

func ProvideEnv() Config {
	log.Event(Instance, "reading config...")

	viper.AddConfigPath(".")
	// viper.SetConfigName("app")
	viper.SetConfigType("env")

	env := Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(Instance, "failed to read config", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(Instance, "failed to unmarshal config", err)
	}

	return env
}

func ExtractServerConfig(config Config) httpServer.Config {
	return config.HttpServer
}

func ExtractDBConfig(config Config) gorm.Config {
	return config.DB
}
