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
	repoimpl.NewLocaleRepo,
	wire.Bind(new(repo.ILocale), new(*repoimpl.Repo)),
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewLocaleUsecase,
	wire.Bind(new(usecase.ILocale), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(gql.Options), "*"),
	gql.NewLocaleHandler,
)
