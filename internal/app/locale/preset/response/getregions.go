package response

import (
	"airbnb-user-be/internal/app/locale"
	"airbnb-user-be/internal/pkg/pagination"
)

type GetRegions struct {
	Regions *[]locale.Region      `json:"regions"`
	Paging  *pagination.SQLPaging `json:"paging"`
}
