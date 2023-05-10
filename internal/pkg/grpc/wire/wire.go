package wire

import (
	"airbnb-user-be/internal/pkg/grpc"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(grpc.Options), "*"),
	grpc.NewRpcServer,
)
