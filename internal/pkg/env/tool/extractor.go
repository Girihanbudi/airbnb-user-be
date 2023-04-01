package tool

import (
	"airbnb-user-be/internal/pkg/env/config"
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpServer "airbnb-user-be/internal/pkg/http/server/config"
	oauthFacebook "airbnb-user-be/internal/pkg/oauth/facebook/config"
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

func ExtractOauthFacebookConfig(config config.Config) oauthFacebook.Config {
	return config.Oauth.Facebook
}
