package wire

import (
	"airbnb-user-be/internal/app/user/api/gql"
	"airbnb-user-be/internal/app/user/api/rpc"
	"airbnb-user-be/internal/app/user/repo"
	"airbnb-user-be/internal/app/user/repo/repoimpl"
	"airbnb-user-be/internal/app/user/usecase"
	"airbnb-user-be/internal/app/user/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
	usecaseSet,
	apiSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewUserRepo,
	wire.Bind(new(repo.IUser), new(*repoimpl.Repo)),
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewUserUsecase,
	wire.Bind(new(usecase.IUser), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(gql.Options), "*"),
	gql.NewUserHandler,
	wire.Struct(new(rpc.Options), "*"),
	rpc.NewUserHandler,
)
