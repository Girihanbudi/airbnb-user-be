package wire

import (
	"airbnb-user-be/internal/app/translation/repo"
	"airbnb-user-be/internal/app/translation/repo/repoimpl"

	"github.com/google/wire"
)

var ModuleSet = wire.NewSet(
	repoSet,
)

var repoSet = wire.NewSet(
	wire.Struct(new(repoimpl.Options), "*"),
	repoimpl.NewErrTranslationRepo,
	wire.Bind(new(repo.IErrTranslation), new(*repoimpl.Repo)),
)
