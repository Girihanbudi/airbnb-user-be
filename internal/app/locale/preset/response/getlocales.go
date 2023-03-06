package response

import (
	"airbnb-user-be/internal/app/locale"
	"airbnb-user-be/internal/pkg/pagination"
)

type GetLocales struct {
	Locales *[]locale.Locale      `json:"locales"`
	Paging  *pagination.SQLPaging `json:"paging"`
}
