//go:build wireinject
// +build wireinject

package app

import (
	env "airbnb-user-be/internal/pkg/env/wire"
	gorm "airbnb-user-be/internal/pkg/gorm/wire"
	http "airbnb-user-be/internal/pkg/http/server/wire"
	kafka "airbnb-user-be/internal/pkg/kafka/wire"

	country "airbnb-user-be/internal/app/country/wire"
	currency "airbnb-user-be/internal/app/currency/wire"
	locale "airbnb-user-be/internal/app/locale/wire"
	translation "airbnb-user-be/internal/app/translation/wire"
	user "airbnb-user-be/internal/app/user/wire"

	"github.com/google/wire"
)

var AppSet = wire.NewSet(
	wire.Struct(new(Options), "*"),
	wire.Struct(new(App), "*"),
)

func NewApp() (*App, error) {
	panic(
		wire.Build(
			env.PackageSet,
			gorm.PackageSet,
			http.PackageSet,
			kafka.PackageSet,

			AppSet,

			user.ModuleSet,
			country.ModuleSet,
			translation.ModuleSet,
			locale.ModuleSet,
			currency.ModuleSet,
		),
	)
}
