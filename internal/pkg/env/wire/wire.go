package wire

import (
	"airbnb-user-be/internal/pkg/env"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	env.ProvideEnv,
	env.ExtractServerConfig,
	env.ExtractDBConfig,
)
