package tool

import (
	credential "airbnb-user-be/internal/pkg/credential/config"
	"airbnb-user-be/internal/pkg/env/config"
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	httpServer "airbnb-user-be/internal/pkg/http/server/config"
	kafka "airbnb-user-be/internal/pkg/kafka/config"
	kafkaconsumer "airbnb-user-be/internal/pkg/kafka/consumer/config"
	kafkarouter "airbnb-user-be/internal/pkg/kafka/router/config"
)

func ExtractCredsConfig(config config.Config) credential.Config {
	return config.Creds
}

func ExtractServerConfig(config config.Config) httpServer.Config {
	return config.HttpServer
}

func ExtractDBConfig(config config.Config) gorm.Config {
	return config.DB
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
