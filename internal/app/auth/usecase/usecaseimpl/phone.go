package usecaseimpl

import (
	"airbnb-user-be/env/appcontext"
	errpreset "airbnb-user-be/internal/app/auth/preset/error"
	"airbnb-user-be/internal/app/auth/preset/request"
	"airbnb-user-be/internal/app/auth/preset/response"
	transutil "airbnb-user-be/internal/app/translation/util"
	usermodule "airbnb-user-be/internal/app/user"
	otpcache "airbnb-user-be/internal/pkg/cache/otp"
	"airbnb-user-be/internal/pkg/json"
	msgpreset "airbnb-user-be/internal/pkg/messaging/preset"
	"airbnb-user-be/internal/pkg/stderror"
	"airbnb-user-be/internal/pkg/util"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (u Usecase) ContinueWithPhone(ctx *gin.Context, cmd request.ContinueWithPhone) (res response.ContinueWithPhone, err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	if _, getCountryErr := u.CountryRepo.GetCountryByPhoneCode(ctx, cmd.CountryCode); getCountryErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getCountryErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	// update or create user if not exist
	var user usermodule.User
	if recordUser, getUserErr := u.UserRepo.GetUserByPhone(ctx, cmd.CountryCode, cmd.PhoneNumber); getUserErr != nil {
		user.CountryCode = &cmd.CountryCode
		user.PhoneNumber = &cmd.PhoneNumber
		user.Role = usermodule.UserRole.String()

		// create user default setting
		var userDefaultSetting usermodule.UserDefaultSetting
		userDefaultSetting.UserId = user.Id
		userDefaultSetting.Locale = clientLocale
		userDefaultSetting.Currency = appcontext.GetCurrency(ctx)

		user.DefaultSetting = &userDefaultSetting

		// insert new user to database
		createUserErr := u.UserRepo.CreateUser(ctx, &user)
		if createUserErr != nil {
			err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
			return
		}
	} else {
		user = recordUser
	}

	otp, err := u.createAndStoreOtp(ctx, user.Id)
	if err != nil {
		return
	}

	userPhoneNumber := fmt.Sprintf("+%d%s", cmd.CountryCode, cmd.PhoneNumber)
	recipients := []string{userPhoneNumber}

	template, err := transutil.TranslateMessage(ctx, "otp", clientLocale)
	if err != nil {
		return
	}
	message := fmt.Sprintf(template, otp)

	payload := msgpreset.SendSmsPayload{
		Recipients: recipients,
		Body:       message,
	}

	msg := msgpreset.SendSms{
		Type:    "otp",
		Context: "signin",
		Payload: *json.Set(payload),
	}

	if _, _, produceEventErr := u.EventProducer.ProduceMessage("sms.send.init", msg); produceEventErr != nil {
		err = transutil.TranslateError(ctx, errpreset.EvtSendMsgFailed, clientLocale)
		return
	}

	res.IsVerified = user.VerifiedAt != nil
	return
}

func (u Usecase) CompletePhoneRegistration(ctx *gin.Context, cmd request.CompletePhoneRegistration) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	if valid, _ := cmd.Validate(); !valid {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	userId, extractOtpErr := otpcache.Get(cmd.Otp)
	if extractOtpErr != nil {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	user, getUserErr := u.UserRepo.GetUser(ctx, userId)
	if getUserErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getUserErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	if user.VerifiedAt != nil {
		err = transutil.TranslateError(ctx, errpreset.UscForbidden, clientLocale)
		return
	}

	user.FirstName = util.Case(cmd.FirstName, util.CaseLower, util.CaseTitle)
	user.FullName = util.Case(cmd.FirstName+" "+cmd.LastName, util.CaseLower, util.CaseTitle)
	user.Email = &cmd.Email
	user.DateOfBirth = cmd.ConvertedDateOfBirth()
	user.Role = usermodule.UserRole.String()

	if saveUserErr := u.UserRepo.CreateOrUpdateUser(ctx, &user); saveUserErr != nil {
		err = transutil.TranslateError(ctx, errpreset.DbServiceUnavailable, clientLocale)
		return
	}

	// delete old token
	u.deleteOldToken(ctx, appcontext.AccessTokenCode)
	u.deleteOldToken(ctx, appcontext.RefreshTokenCode)

	return u.createAndStoreTokensPair(ctx, user)
}

func (u Usecase) MakePhoneSession(ctx *gin.Context, cmd request.MakePhoneSession) (err *stderror.StdError) {
	clientLocale := appcontext.GetLocale(ctx)

	userId, extractOtpErr := otpcache.Get(cmd.Otp)
	if extractOtpErr != nil {
		err = transutil.TranslateError(ctx, errpreset.UscBadRequest, clientLocale)
		return
	}

	user, getUserErr := u.UserRepo.GetUser(ctx, userId)
	if getUserErr != nil {
		ec := errpreset.DbServiceUnavailable
		if errors.Is(getUserErr, gorm.ErrRecordNotFound) {
			ec = errpreset.DbRecordNotFound
		}
		err = transutil.TranslateError(ctx, ec, clientLocale)
		return
	}

	if user.VerifiedAt == nil {
		err = transutil.TranslateError(ctx, errpreset.UscForbidden, clientLocale)
		return
	}

	// delete old token
	u.deleteOldToken(ctx, appcontext.AccessTokenCode)
	u.deleteOldToken(ctx, appcontext.RefreshTokenCode)

	return u.createAndStoreTokensPair(ctx, user)
}
