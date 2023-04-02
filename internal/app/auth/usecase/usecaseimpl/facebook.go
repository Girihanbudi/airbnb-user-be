package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	module "airbnb-user-be/internal/app/auth"
	errpreset "airbnb-user-be/internal/app/auth/preset/error"
	transutil "airbnb-user-be/internal/app/translation/util"
	usermodule "airbnb-user-be/internal/app/user"
	"airbnb-user-be/internal/pkg/codegenerator"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u Usecase) ContinueWithFacebook(ctx gin.Context) {
	// Create CSRF token cookie
	oauthState := codegenerator.RandomEncodedBytes(16)

	ctx.SetCookie(
		appcontext.OauthCode,
		oauthState,
		appcontext.OauthDuration,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	link := u.FacebookOauth.AuthCodeURL(oauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, link)
}

func (u Usecase) OauthFacebookCallback(ctx gin.Context) (err *stderror.StdError) {
	reqCtx := ctx.Request.Context()
	clientLocale := appcontext.GetLocale(reqCtx)

	// Read CSRF token from Cookie
	oauthState, _ := ctx.Cookie(appcontext.OauthCode)

	if ctx.Request.FormValue("state") != oauthState {
		err = transutil.TranslateError(reqCtx, errpreset.AUTH_GET_401, clientLocale)
		return
	}

	data, account, extractDataErr := u.extractFacebookUserData(ctx.Request.FormValue("code"))
	if extractDataErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.AUTH_GET_502, clientLocale)
		return
	}

	// update or create user if not exist
	var user usermodule.User
	if recordUser, getUserErr := u.UserRepo.GetUserByEmail(reqCtx, data.Email); getUserErr != nil {
		user.FirstName = data.FirstName
		user.FullName = data.Name
		user.Email = &data.Email
		user.Image = data.Picture["data"]["url"].(string)
		user.Role = usermodule.UserRole.String()

		// get locale list for references
		locales, getLocalesErr := u.LocaleRepo.GetLocales(reqCtx)
		if getLocalesErr != nil {
			err = transutil.TranslateError(reqCtx, errpreset.AUTH_GET_503, clientLocale)
			return
		}

		// create user default setting
		var userDefaultSetting usermodule.UserDefaultSetting
		userDefaultSetting.UserId = user.Id
		// set user locale using facebook locale
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
			userDefaultSetting.Locale = clientLocale
			userDefaultSetting.Currency = appcontext.GetCurrency(reqCtx)
		}
		user.DefaultSetting = userDefaultSetting

		// insert new user to database
		createUserErr := u.UserRepo.CreateUser(ctx.Request.Context(), &user)
		if createUserErr != nil {
			err = transutil.TranslateError(reqCtx, errpreset.AUTH_GET_503, clientLocale)
			return
		}
	} else {
		user = recordUser
	}

	// update or create user account if not exist
	account.UserId = user.Id
	createAcountErr := u.UserRepo.CreateOrUpdateUserAccount(reqCtx, &account)
	if createAcountErr != nil {
		err = transutil.TranslateError(reqCtx, errpreset.AUTH_GET_503, clientLocale)
		return
	}

	return u.createAndStoreTokensPair(ctx, user.Id)
}

func (u Usecase) extractFacebookUserData(code string) (userInfo module.FacebookUserInfo, account usermodule.Account, err error) {
	token, err := u.FacebookOauth.Exchange(context.Background(), code)
	if err != nil {
		err = fmt.Errorf("code exchange wrong: %s", err.Error())
		return
	}

	fmt.Printf("%+v\n", token)

	// bind token info
	account.Provider = module.ProviderFacebook.String()
	account.AccessToken = token.AccessToken
	account.RefreshToken = token.RefreshToken
	account.ExpiredAt = token.Expiry
	account.TokenType = token.TokenType

	// get user info from facebook apis
	response, err := http.Get(u.FacebookOauth.UserInfoApi + token.AccessToken)
	if err != nil {
		err = fmt.Errorf("failed getting user info: %s", err.Error())
		return
	}

	// run to the closest return
	defer response.Body.Close()

	// read message
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed read response: %s", err.Error())
		return
	}

	// bind to user info struct
	err = json.Unmarshal(contents, &userInfo)
	if err != nil {
		return
	}

	if userInfo.Email == "" {
		err = errors.New("email not found")
		return
	}

	return
}