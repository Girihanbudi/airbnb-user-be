package seeder

import (
	"airbnb-user-be/internal/pkg/env"
)

var batchSize = 100

var envConfig = env.EnvConfig{
	Path:     "../../../env",
	FileName: "config",
	Ext:      "yaml",
}
