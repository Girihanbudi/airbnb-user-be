package wire

import (
	"airbnb-user-be/internal/app/auth/api/rest"
	"airbnb-user-be/internal/app/auth/usecase"
	"airbnb-user-be/internal/app/auth/usecase/usecaseimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	usecaseSet,
	apiSet,
)

var usecaseSet = wire.NewSet(
	wire.Struct(new(usecaseimpl.Options), "*"),
	usecaseimpl.NewAuthUsecase,
	wire.Bind(new(usecase.IAuth), new(*usecaseimpl.Usecase)),
)

var apiSet = wire.NewSet(
	wire.Struct(new(rest.Options), "*"),
	rest.NewAuthHandler,
)
