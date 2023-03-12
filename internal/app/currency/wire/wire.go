package wire

import (
	"airbnb-user-be/internal/app/currency/api/gql"
	"airbnb-user-be/internal/app/currency/repo"
	"airbnb-user-be/internal/app/currency/repo/repoimpl"
	"airbnb-user-be/internal/app/currency/usecase"
	"airbnb-user-be/internal/app/currency/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
	usecaseSet,
	apiSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewCurrencyRepo,
	wire.Bind(new(repo.ICurrency), new(*repoimpl.Repo)),
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewCurrencyUsecase,
	wire.Bind(new(usecase.ICurrency), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(gql.Options), "*"),
	gql.ProvideCurrencyHandler,
)
