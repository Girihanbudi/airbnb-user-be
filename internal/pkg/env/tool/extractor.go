package tool

import (
	"airbnb-user-be/internal/pkg/env/config"
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpServer "airbnb-user-be/internal/pkg/http/server/config"
	kafka "airbnb-user-be/internal/pkg/kafka/config"
	kafkaconsumer "airbnb-user-be/internal/pkg/kafka/consumer/config"
	kafkarouter "airbnb-user-be/internal/pkg/kafka/router/config"
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

func ExtractKafkaConfig(config config.Config) kafka.Config {
	return config.Kafka
}

func ExtractKafkaConsumerConfig(config config.Config) kafkaconsumer.Config {
	return config.Kafka.Consumer
}

func ExtractKafkaRouterConfig(config config.Config) kafkarouter.Config {
	return config.Kafka.Consumer.Router
}
