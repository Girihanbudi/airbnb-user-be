package env

import (
	"airbnb-user-be/internal/pkg/gorm"
	httpServer "airbnb-user-be/internal/pkg/http/server"
	"airbnb-user-be/internal/pkg/log"

	"github.com/spf13/viper"
)

const Instance string = "Env"

// global env declaration
var CONFIG Config

type EnvConfig struct {
	Path     string
	FileName string
	Ext      string
}

func ProvideDefaultEnvConf() EnvConfig {
	return EnvConfig{
		Path:     "./env",
		FileName: "config",
		Ext:      "yaml",
	}
}

func ProvideEnv(conf EnvConfig) Config {
	log.Event(Instance, "reading config...")

	viper.AddConfigPath(conf.Path)
	viper.SetConfigName(conf.FileName)
	viper.SetConfigType(conf.Ext)

	env := Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(Instance, "failed to read config", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal(Instance, "failed to unmarshal config", err)
	}

	CONFIG = env

	return env
}

func ExtractServerConfig(config Config) httpServer.Config {
	return config.HttpServer
}

func ExtractDBConfig(config Config) gorm.Config {
	return config.DB
}
