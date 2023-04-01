package wire

import (
	"airbnb-user-be/internal/pkg/oauth/facebook"
	"airbnb-user-be/internal/pkg/oauth/google"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	google.NewGoogleOauth,
	facebook.NewFacebookOauth,
)
