package wire

import (
	"airbnb-user-be/internal/app/user/repo"
	"airbnb-user-be/internal/app/user/repo/repoimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewUserRepo,
	wire.Bind(new(repo.IUser), new(*repoimpl.Repo)),
)
