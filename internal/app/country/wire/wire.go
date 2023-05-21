package wire

import (
	"airbnb-user-be/internal/app/country/api/gql"
	"airbnb-user-be/internal/app/country/api/rpc"
	"airbnb-user-be/internal/app/country/repo"
	"airbnb-user-be/internal/app/country/repo/repoimpl"
	"airbnb-user-be/internal/app/country/usecase"
	"airbnb-user-be/internal/app/country/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
	usecaseSet,
	apiSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewCountryRepo,
	wire.Bind(new(repo.ICountry), new(*repoimpl.Repo)),
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewCountryUsecase,
	wire.Bind(new(usecase.ICountry), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(gql.Options), "*"),
	gql.NewCountryHandler,
	wire.Struct(new(rpc.Options), "*"),
	rpc.NewCountryHandler,
)
