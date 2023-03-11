package response

import (
	"airbnb-user-be/internal/app/locale"
)

type GetLocales struct {
	Locales *[]locale.Locale `json:"locales"`
}
