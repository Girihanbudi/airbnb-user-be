package wire

import (
	"airbnb-user-be/internal/pkg/oauth/facebook"
	"airbnb-user-be/internal/pkg/oauth/google"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	wire.Struct(new(google.Options), "*"),
	google.NewGoogleOauth,

	wire.Struct(new(facebook.Options), "*"),
	facebook.NewFacebookOauth,
)
