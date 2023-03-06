package response

import (
	"airbnb-user-be/internal/app/locale"
)

type GetLocale struct {
	Locale *locale.Locale `json:"locale"`
}
