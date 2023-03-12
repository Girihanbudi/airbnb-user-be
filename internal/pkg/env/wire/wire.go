package wire

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/env/tool"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	env.ProvideDefaultEnvConf,
	env.ProvideEnv,
	tool.ExtractServerConfig,
	tool.ExtractDBConfig,
)
