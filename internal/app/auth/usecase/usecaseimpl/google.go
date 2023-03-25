package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	module "airbnb-user-be/internal/app/auth"
	usermodule "airbnb-user-be/internal/app/user"
	authcache "airbnb-user-be/internal/pkg/cache/auth"
	"airbnb-user-be/internal/pkg/codegenerator"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/jwt"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (u Usecase) ContinueWithGoogle(ctx gin.Context) {
	// Create oauthState cookie
	oauthState := codegenerator.RandomEncodedBytes(16)
	// appcontext.SetFromGinRouter(&ctx, appcontext.OauthCode, oauthState)
	ctx.SetCookie(
		appcontext.OauthCode,
		oauthState,
		appcontext.OauthDuration,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	link := u.GoogleOauth.AuthCodeURL(oauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, link)
}

func (u Usecase) OauthGoogleCallback(ctx gin.Context) {
	// Read oauthState from Cookie
	oauthState, _ := ctx.Cookie(appcontext.OauthCode)

	if ctx.Request.FormValue("state") != oauthState {
		ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
		return
	}

	data, err := u.extractGoogleUserData(ctx.Request.FormValue("code"))
	if err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
		return
	}

	reqCtx := ctx.Request.Context()

	var user usermodule.User

	// create user if not exist
	if user, err = u.UserRepo.GetUserByEmail(reqCtx, data.Email); err != nil {
		user.FirstName = data.GivenName
		user.FullName = data.Name
		user.Email = &data.Email
		user.Image = data.Picture
		user.Role = usermodule.UserRole.String()

		// get locale list for references
		locales, err := u.LocaleRepo.GetLocales(reqCtx)
		if err != nil {
			ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
		}

		// create user default setting
		var userDefaultSetting usermodule.UserDefaultSetting
		userDefaultSetting.UserId = user.Id
		// set user locale using google locale
		isLocaleFound := false
		if locales != nil {
			for _, locale := range *locales {
				if data.Locale == locale.Code {
					isLocaleFound = true
					userDefaultSetting.Locale = data.Locale
					break
				}
			}
		}
		// otherwise using current locale
		if !isLocaleFound {
			userDefaultSetting.Locale = appcontext.GetLocale(reqCtx)
			userDefaultSetting.Currency = appcontext.GetCurrency(reqCtx)
		}
		user.DefaultSetting = userDefaultSetting

		// insert new user to database
		err = u.UserRepo.CreateUser(ctx.Request.Context(), &user)
		if err != nil {
			ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
		}

	}

	atKey, err := gonanoid.New()
	if err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
	}

	if err := authcache.Set(atKey, user.Id, appcontext.AccessTokenDuration); err != nil {
		ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
	}

	claims := jwt.MapClaims{}
	claims["jti"] = atKey
	at := jwt.GenerateToken(claims, appcontext.AccessTokenDuration)

	// set user cookie
	ctx.SetCookie(
		appcontext.AccessTokenCode,
		*at,
		appcontext.AccessTokenDuration,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	ctx.Redirect(http.StatusPermanentRedirect, env.CONFIG.Oauth.RedirectUrl)
}

func (u Usecase) extractGoogleUserData(code string) (userInfo module.GoogleUserInfo, err error) {
	token, err := u.GoogleOauth.Exchange(context.Background(), code)
	if err != nil {
		err = fmt.Errorf("code exchange wrong: %s", err.Error())
		return
	}

	response, err := http.Get(u.GoogleOauth.UserInfoApi + token.AccessToken)
	if err != nil {
		err = fmt.Errorf("failed getting user info: %s", err.Error())
		return
	}

	// run to the closest return
	defer response.Body.Close()

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed read response: %s", err.Error())
		return
	}

	err = json.Unmarshal(contents, &userInfo)

	return
}
