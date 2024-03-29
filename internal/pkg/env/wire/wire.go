package wire

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/env/tool"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	env.ProvideEnv,
	tool.ExtractCredsConfig,
	tool.ExtractHttpServerConfig,
	tool.ExtractGrpcServerConfig,
	tool.ExtractDBConfig,
	tool.ExtractKafkaConfig,
	tool.ExtractKafkaConsumerConfig,
	tool.ExtractKafkaRouterConfig,
)
