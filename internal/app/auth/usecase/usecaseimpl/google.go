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
	"airbnb-user-be/internal/pkg/util"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// test
func (u Usecase) ContinueWithGoogle(ctx *gin.Context) {
	// Create CSRF token cookie
	oauthState := codegenerator.RandomEncodedBytes(16)

	// Set oauth CSRF code
	ctx.SetCookie(
		appcontext.OauthCode,
		oauthState,
		appcontext.OauthDuration,
		"/",
		env.CONFIG.Domain,
		true,
		true,
	)

	// Redirect to google oauth page
	link := u.GoogleOauth.AuthCodeURL(oauthState)
	ctx.Redirect(http.StatusTemporaryRedirect, link)
}

func (u Usecase) OauthGoogleCallback(ctx *gin.Context) (err *stderror.StdError) {
	// Get user locale code
	clientLocale := appcontext.GetLocale(ctx)

	// Read CSRF token from Cookie
	oauthState, _ := ctx.Cookie(appcontext.OauthCode)
	if ctx.Request.FormValue("state") != oauthState {
		err = transutil.TranslateError(ctx, errpreset.UscInvalidOauth, clientLocale)
		return
	}

	// Extract user data from google apis
	data, account, extractDataErr := u.extractGoogleUserData(ctx.Request.FormValue("code"))
	if extractDataErr != nil {
		err = transutil.TranslateError(ctx, errpreset.UscFailedExtractGoogleInfo, clientLocale)
		return
	}

	// Update or create user if not exist
	var user usermodule.User
	if recordUser, getUserErr := u.UserRepo.GetUserByEmail(ctx, data.Email); getUserErr != nil {
		currentTime := time.Now()
		user.FirstName = util.Case(data.GivenName, util.CaseLower, util.CaseTitle)
		user.FullName = util.Case(data.Name, util.CaseLower, util.CaseTitle)
		user.Email = &data.Email
		user.Image = data.Picture
		user.Role = usermodule.UserRole.String()
		user.VerifiedAt = &currentTime

		// Get locale list for references
		locales, getLocalesErr := u.LocaleRepo.GetLocales(ctx)
		if getLocalesErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}

		// Create user default setting
		var userDefaultSetting usermodule.UserDefaultSetting
		userDefaultSetting.UserId = user.Id
		// Set user locale using google locale
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
		// Otherwise using current locale
		if !isLocaleFound {
			userDefaultSetting.Locale = clientLocale
			userDefaultSetting.Currency = appcontext.GetCurrency(ctx)
		}
		user.DefaultSetting = &userDefaultSetting

		// Insert new user to database
		createUserErr := u.UserRepo.CreateUser(ctx, &user)
		if createUserErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}
	} else {
		user = recordUser
	}

	// Update or create user account if not exist
	account.UserId = user.Id
	createAcountErr := u.UserRepo.CreateOrUpdateUserAccount(ctx, &account)
	if createAcountErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	// Delete old tokens
	u.deleteOldToken(ctx, appcontext.AccessTokenCode)
	u.deleteOldToken(ctx, appcontext.RefreshTokenCode)

	// Create and store user access and refresh tokens in cache
	return u.createAndStoreTokensPair(ctx, user)
}

func (u Usecase) extractGoogleUserData(code string) (userInfo module.GoogleUserInfo, account usermodule.Account, err error) {
	// Convert authorization code into a token
	token, err := u.GoogleOauth.Exchange(context.Background(), code)
	if err != nil {
		err = fmt.Errorf("code exchange wrong: %s", err.Error())
		return
	}

	// Bind token info
	account.Provider = module.ProviderGoogle.String()
	account.AccessToken = token.AccessToken
	account.RefreshToken = token.RefreshToken
	account.ExpiredAt = token.Expiry
	account.TokenType = token.TokenType

	// Get user info from google apis
	response, err := http.Get(u.GoogleOauth.UserInfoApi + token.AccessToken)
	if err != nil {
		err = fmt.Errorf("failed getting user info: %s", err.Error())
		return
	}

	// Close response body to the closest return
	defer response.Body.Close()

	// Read message
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed read response: %s", err.Error())
		return
	}

	// Bind to user info struct
	err = json.Unmarshal(contents, &userInfo)
	if err != nil {
		return
	}

	// Check if email exist
	if userInfo.Email == "" {
		err = errors.New("email not found")
		return
	}

	return
}
