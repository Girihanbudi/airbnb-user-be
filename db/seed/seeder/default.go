package seeder

import (
	"airbnb-user-be/internal/pkg/env"
)

var batchSize = 100

var envOps = env.Options{
	Path:     "../../../env",
	FileName: "config",
	Ext:      "yaml",
}
