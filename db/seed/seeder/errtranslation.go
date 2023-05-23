package seeder

import (
	countryerr "airbnb-user-be/internal/app/country/preset/error"
	currencyerr "airbnb-user-be/internal/app/currency/preset/error"
	localeerr "airbnb-user-be/internal/app/locale/preset/error"
	middlewareerr "airbnb-user-be/internal/app/middleware/preset/error"
	usererr "airbnb-user-be/internal/app/user/preset/error"
	"airbnb-user-be/internal/pkg/stderror"
	"net/http"

	translationmodule "airbnb-user-be/internal/app/translation"

	"gorm.io/gorm"
)

func SeedErrTranslation(db gorm.DB) error {

	data := []translationmodule.ErrTranslation{
		/*
			Default
		*/
		// En translation
		MakeErrTranslation(stderror.DEF_SERVER_500, "en", http.StatusInternalServerError, "Request aborted due to server error"),
		MakeErrTranslation(stderror.DEF_AUTH_401, "en", http.StatusUnauthorized, "Cannot authorize user"),
		MakeErrTranslation(stderror.DEF_DATA_400, "en", http.StatusBadRequest, "Requested data is not valid"),
		// Id translation
		MakeErrTranslation(stderror.DEF_SERVER_500, "id", http.StatusInternalServerError, "Permintaan dibatalkan karena terjadi kesalahan server"),
		MakeErrTranslation(stderror.DEF_AUTH_401, "id", http.StatusUnauthorized, "Tidak dapat mengotorisasi user"),
		MakeErrTranslation(stderror.DEF_DATA_400, "id", http.StatusBadRequest, "Permintaan tidak valid"),

		/*
			Middleware
		*/
		// En translation
		MakeErrTranslation(middlewareerr.TokenNotFound, "en", http.StatusUnauthorized, "Authorization not found"),
		MakeErrTranslation(middlewareerr.TokenNotValid, "en", http.StatusUnauthorized, "Token is not valid"),
		MakeErrTranslation(middlewareerr.UserAlreadyVerified, "en", http.StatusForbidden, "User already verified"),
		// Id translation
		MakeErrTranslation(middlewareerr.TokenNotFound, "id", http.StatusUnauthorized, "Otorisasi tidak ditemukan"),
		MakeErrTranslation(middlewareerr.TokenNotValid, "id", http.StatusUnauthorized, "Token tidak valid"),
		MakeErrTranslation(middlewareerr.UserAlreadyVerified, "id", http.StatusForbidden, "User telah terverifikasi"),

		/*
			Locale
		*/
		// En translation
		MakeErrTranslation(localeerr.DbServiceUnavailable, "en", http.StatusServiceUnavailable, "Failed to communicate with store server"),
		MakeErrTranslation(localeerr.DbRecordNotFound, "en", http.StatusNotFound, "Requested data not found"),
		MakeErrTranslation(localeerr.DbEmptyResult, "en", http.StatusNotFound, "Requested result nothing"),
		MakeErrTranslation(localeerr.UscBadRequest, "en", http.StatusBadRequest, "Requested data is not valid"),
		// Id translation
		MakeErrTranslation(localeerr.DbServiceUnavailable, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan server penyimpanan"),
		MakeErrTranslation(localeerr.DbRecordNotFound, "id", http.StatusNotFound, "Data tidak ditemukan"),
		MakeErrTranslation(localeerr.DbEmptyResult, "id", http.StatusNotFound, "Tidak ada hasil apapun"),
		MakeErrTranslation(localeerr.UscBadRequest, "id", http.StatusBadRequest, "Permintaan tidak valid"),

		/*
			Currency
		*/
		// En translation
		MakeErrTranslation(currencyerr.DbServiceUnavailable, "en", http.StatusServiceUnavailable, "Failed to communicate with store server"),
		MakeErrTranslation(currencyerr.DbRecordNotFound, "en", http.StatusNotFound, "Requested data not found"),
		MakeErrTranslation(currencyerr.DbEmptyResult, "en", http.StatusNotFound, "Requested result nothing"),
		// Id translation
		MakeErrTranslation(currencyerr.DbServiceUnavailable, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan server penyimpanan"),
		MakeErrTranslation(currencyerr.DbRecordNotFound, "id", http.StatusNotFound, "Data tidak ditemukan"),
		MakeErrTranslation(currencyerr.DbEmptyResult, "id", http.StatusNotFound, "Tidak ada hasil apapun"),

		/*
			Country
		*/
		// En translation
		MakeErrTranslation(countryerr.DbServiceUnavailable, "en", http.StatusServiceUnavailable, "Failed to communicate with store server"),
		MakeErrTranslation(countryerr.DbRecordNotFound, "en", http.StatusNotFound, "Requested data not found"),
		MakeErrTranslation(countryerr.DbEmptyResult, "en", http.StatusNotFound, "Requested result nothing"),
		// Id translation
		MakeErrTranslation(countryerr.DbServiceUnavailable, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan server penyimpanan"),
		MakeErrTranslation(countryerr.DbRecordNotFound, "id", http.StatusNotFound, "Data tidak ditemukan"),
		MakeErrTranslation(countryerr.DbEmptyResult, "id", http.StatusNotFound, "Tidak ada hasil apapun"),

		/*
			User
		*/
		// En translation
		MakeErrTranslation(usererr.DbServiceUnavailable, "en", http.StatusServiceUnavailable, "Failed to communicate with store server"),
		MakeErrTranslation(usererr.DbRecordNotFound, "en", http.StatusNotFound, "Requested data not found"),
		MakeErrTranslation(usererr.DbEmptyResult, "en", http.StatusNotFound, "Requested result nothing"),
		MakeErrTranslation(usererr.UscBadRequest, "en", http.StatusBadRequest, "Requested data is not valid"),
		// Id translation
		MakeErrTranslation(usererr.DbServiceUnavailable, "id", http.StatusServiceUnavailable, "Gagal berkomunikasi dengan server penyimpanan"),
		MakeErrTranslation(usererr.DbRecordNotFound, "id", http.StatusNotFound, "Data tidak ditemukan"),
		MakeErrTranslation(usererr.DbEmptyResult, "id", http.StatusNotFound, "Tidak ada hasil apapun"),
		MakeErrTranslation(usererr.UscBadRequest, "id", http.StatusBadRequest, "Permintaan tidak valid"),
	}

	var errTranslationRecords []translationmodule.ErrTranslation
	if err := db.Find(&errTranslationRecords).Error; err != nil {
		return err
	}

	if len(errTranslationRecords) > 0 {
		if err := db.Delete(&errTranslationRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func MakeErrTranslation(code, localeCode string, httpCode int, message string) translationmodule.ErrTranslation {
	return translationmodule.ErrTranslation{
		Code:       code,
		LocaleCode: localeCode,
		HttpCode:   httpCode,
		Message:    message,
	}
}
