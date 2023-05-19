package wire

import (
	"airbnb-user-be/internal/pkg/grpcserver"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(grpcserver.Options), "*"),
	grpcserver.NewRpcServer,
)
