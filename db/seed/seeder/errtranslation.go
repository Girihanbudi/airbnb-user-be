package seeder

import (
	currencyerr "airbnb-user-be/internal/app/currency/preset/error"
	localeerr "airbnb-user-be/internal/app/locale/preset/error"
	middlewareerr "airbnb-user-be/internal/app/middleware/preset/error"
	"net/http"

	translationmodule "airbnb-user-be/internal/app/translation"

	"gorm.io/gorm"
)

func SeedErrTranslation(db gorm.DB) error {

	data := []translationmodule.ErrTranslation{
		// En translation
		// Middleware
		MakeErrTranslation(middlewareerr.AUTH_MID_001, "en-US", http.StatusUnauthorized, "Authorization not found"),
		// Locale
		MakeErrTranslation(localeerr.LOCALE_VAL_400, "en-US", http.StatusBadRequest, "Failed to validate request"),
		MakeErrTranslation(localeerr.LOCALE_GET_404, "en-US", http.StatusNotFound, "Locale record not found"),
		MakeErrTranslation(localeerr.LOCALE_GET_503, "en-US", http.StatusServiceUnavailable, "Internal server error"),
		MakeErrTranslation(localeerr.LOCALE_CREATE_503, "en-US", http.StatusServiceUnavailable, "Failed to create locale due to internal server error"),
		MakeErrTranslation(localeerr.LOCALE_UPDATE_503, "en-US", http.StatusServiceUnavailable, "Failed to update locale due to internal server error"),
		MakeErrTranslation(localeerr.LOCALE_DELETE_503, "en-US", http.StatusServiceUnavailable, "Failed to delete locale due to internal server error"),
		// Currency
		MakeErrTranslation(currencyerr.CURRENCY_VAL_400, "en-US", http.StatusBadRequest, "Failed to validate request"),
		MakeErrTranslation(currencyerr.CURRENCY_GET_404, "en-US", http.StatusNotFound, "Currency record not found"),
		MakeErrTranslation(currencyerr.CURRENCY_GET_503, "en-US", http.StatusServiceUnavailable, "Internal server error"),
		MakeErrTranslation(currencyerr.CURRENCY_CREATE_503, "en-US", http.StatusServiceUnavailable, "Failed to create currency due to internal server error"),
		MakeErrTranslation(currencyerr.CURRENCY_UPDATE_503, "en-US", http.StatusServiceUnavailable, "Failed to update currency due to internal server error"),
		MakeErrTranslation(currencyerr.CURRENCY_DELETE_503, "en-US", http.StatusServiceUnavailable, "Failed to delete currency due to internal server error"),

		// Id translation
		// Middleware
		MakeErrTranslation(middlewareerr.AUTH_MID_001, "id-ID", http.StatusUnauthorized, "Otorisasi tidak ditemukan"),
		// Locale
		MakeErrTranslation(localeerr.LOCALE_VAL_400, "id-ID", http.StatusBadRequest, "Gagal melakukan validasi request"),
		MakeErrTranslation(localeerr.LOCALE_GET_404, "id-ID", http.StatusNotFound, "Rekaman lokal tidak ditemukan"),
		MakeErrTranslation(localeerr.LOCALE_GET_503, "id-ID", http.StatusServiceUnavailable, "Terjadi kesalahan server"),
		MakeErrTranslation(localeerr.LOCALE_CREATE_503, "id-ID", http.StatusServiceUnavailable, "Gagal membuat lokal karena terjadi kesalahan server"),
		MakeErrTranslation(localeerr.LOCALE_UPDATE_503, "id-ID", http.StatusServiceUnavailable, "Gagal mengupdate lokal karena terjadi kesalahan server"),
		MakeErrTranslation(localeerr.LOCALE_DELETE_503, "id-ID", http.StatusServiceUnavailable, "Gagal menghapus lokal karena terjadi kesalahan server"),
		// Currency
		MakeErrTranslation(currencyerr.CURRENCY_VAL_400, "id-ID", http.StatusBadRequest, "Gagal melakukan validasi request"),
		MakeErrTranslation(currencyerr.CURRENCY_GET_404, "id-ID", http.StatusNotFound, "Rekaman mata uang tidak ditemukan"),
		MakeErrTranslation(currencyerr.CURRENCY_GET_503, "id-ID", http.StatusServiceUnavailable, "Terjadi kesalahan server"),
		MakeErrTranslation(currencyerr.CURRENCY_CREATE_503, "id-ID", http.StatusServiceUnavailable, "Gagal membuat mata uang karena terjadi kesalahan server"),
		MakeErrTranslation(currencyerr.CURRENCY_UPDATE_503, "id-ID", http.StatusServiceUnavailable, "Gagal mengupdate mata uang karena terjadi kesalahan server"),
		MakeErrTranslation(currencyerr.CURRENCY_DELETE_503, "id-ID", http.StatusServiceUnavailable, "Gagal menghapus mata uang karena terjadi kesalahan server"),
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
