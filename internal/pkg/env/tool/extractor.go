package tool

import (
	credential "airbnb-user-be/internal/pkg/credential/config"
	"airbnb-user-be/internal/pkg/env/config"
	gorm "airbnb-user-be/internal/pkg/gorm/config"
	grpcserver "airbnb-user-be/internal/pkg/grpcserver/config"
	httpserver "airbnb-user-be/internal/pkg/http/server/config"
	kafka "airbnb-user-be/internal/pkg/kafka/config"
	kafkaconsumer "airbnb-user-be/internal/pkg/kafka/consumer/config"
	kafkarouter "airbnb-user-be/internal/pkg/kafka/router/config"
)

func ExtractCredsConfig(config config.Config) credential.Config {
	return config.Creds
}

func ExtractHttpServerConfig(config config.Config) httpserver.Config {
	return config.HttpServer
}

func ExtractGrpcServerConfig(config config.Config) grpcserver.Config {
	return config.GrpcServer
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
