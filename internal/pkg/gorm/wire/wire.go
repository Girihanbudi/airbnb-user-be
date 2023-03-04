package wire

import (
	"airbnb-user-be/internal/pkg/gorm"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	gorm.ProvideORM,
)
