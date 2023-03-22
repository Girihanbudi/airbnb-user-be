package google

import (
	"airbnb-user-be/internal/pkg/oauth/google/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Oauth struct {
	oauth2.Config
	UserInfoApi string
}

func NewGoogleOauth(config config.Config) Oauth {
	var oauth Oauth
	oauth.ClientID = config.ClientId
	oauth.ClientSecret = config.ClientSecret
	oauth.RedirectURL = config.RedirectUrl
	oauth.Endpoint = google.Endpoint
	oauth.UserInfoApi = config.UserInfoApi

	return oauth
}
