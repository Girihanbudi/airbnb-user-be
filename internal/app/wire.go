//go:build wireinject
// +build wireinject

package app

import (
	env "airbnb-user-be/internal/pkg/env/wire"
	gorm "airbnb-user-be/internal/pkg/gorm/wire"
	http "airbnb-user-be/internal/pkg/http/server/wire"

	locale "airbnb-user-be/internal/app/locale/wire"
	translation "airbnb-user-be/internal/app/translation/wire"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(Options), "*"),
	wire.Struct(new(App), "*"),
)

func ProvideApp() (*App, error) {
	panic(wire.Build(
		env.PackageSet,
		gorm.PackageSet,
		http.PackageSet,

		AppSet,

		translation.ModuleSet,
		locale.ModuleSet,
	))
}
