package seeder

import (
	localeErr "airbnb-user-be/internal/app/locale/preset/error"
	"airbnb-user-be/internal/app/translation"

	"gorm.io/gorm"
)

func SeedErrTranslation(db gorm.DB) error {

	data := []translation.ErrTranslation{
		// En
		MakeErrTranslation(localeErr.LOCALE_VAL_400, "en-US", 400, "Failed to validate request"),
		MakeErrTranslation(localeErr.LOCALE_GET_404, "en-US", 404, "Locale record not found"),
		MakeErrTranslation(localeErr.LOCALE_GET_500, "en-US", 500, "Internal server error"),
		MakeErrTranslation(localeErr.LOCALE_CREATE_500, "en-US", 500, "Failed to create locale due to internal server error"),
		MakeErrTranslation(localeErr.LOCALE_UPDATE_500, "en-US", 500, "Failed to update locale due to internal server error"),
		MakeErrTranslation(localeErr.LOCALE_DELETE_500, "en-US", 500, "Failed to delete locale due to internal server error"),
		// Id
		MakeErrTranslation(localeErr.LOCALE_VAL_400, "id-ID", 400, "Gagal melakukan validasi request"),
		MakeErrTranslation(localeErr.LOCALE_GET_404, "id-ID", 404, "Rekaman tidak ditemukan"),
		MakeErrTranslation(localeErr.LOCALE_GET_500, "id-ID", 500, "Terjadi kesalahan server"),
		MakeErrTranslation(localeErr.LOCALE_CREATE_500, "id-ID", 500, "Gagal membuat locale karena terjadi kesalahan server"),
		MakeErrTranslation(localeErr.LOCALE_UPDATE_500, "id-ID", 500, "Gagal mengupdate locale karena terjadi kesalahan server"),
		MakeErrTranslation(localeErr.LOCALE_DELETE_500, "id-ID", 500, "Gagal menghapus locale karena terjadi kesalahan server"),
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func MakeErrTranslation(code, localeCode string, httpCode int, message string) translation.ErrTranslation {
	return translation.ErrTranslation{
		Code:       code,
		LocaleCode: localeCode,
		HttpCode:   httpCode,
		Message:    message,
	}
}
