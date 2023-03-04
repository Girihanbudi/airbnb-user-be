package wire

import (
	"airbnb-user-be/internal/app/locale/api/gql"
	"airbnb-user-be/internal/app/locale/repo"
	"airbnb-user-be/internal/app/locale/repo/repoimpl"
	"airbnb-user-be/internal/app/locale/usecase"
	"airbnb-user-be/internal/app/locale/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
	usecaseSet,
	apiSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewRegionRepo,
	wire.Bind(new(repo.IRegion), new(*repoimpl.Repo)),
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewRegionUsecase,
	wire.Bind(new(usecase.IRegion), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(gql.Options), "*"),
	gql.ProvideLocaleHandler,
)
