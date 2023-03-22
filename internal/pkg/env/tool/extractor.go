package tool

import (
	"airbnb-user-be/internal/pkg/env/config"
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpServer "airbnb-user-be/internal/pkg/http/server/config"
	oauthGoogle "airbnb-user-be/internal/pkg/oauth/google/config"
)

func ExtractServerConfig(config config.Config) httpServer.Config {
	return config.HttpServer
}

func ExtractDBConfig(config config.Config) gorm.Config {
	return config.DB
}

func ExtractOauthGoogleConfig(config config.Config) oauthGoogle.Config {
	return config.Oauth.Google
}
