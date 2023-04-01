package facebook

import (
	"airbnb-user-be/internal/pkg/oauth/facebook/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type Oauth struct {
	oauth2.Config
	UserInfoApi string
}

func NewFacebookOauth(config config.Config) Oauth {
	var oauth Oauth
	oauth.ClientID = config.ClientId
	oauth.ClientSecret = config.ClientSecret
	oauth.Endpoint = facebook.Endpoint
	oauth.UserInfoApi = config.UserInfoApi
	oauth.RedirectURL = config.RedirectUrl
	oauth.Scopes = config.Scopes

	return oauth
}
