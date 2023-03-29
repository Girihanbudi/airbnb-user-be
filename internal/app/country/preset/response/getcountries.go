package response

import (
	module "airbnb-user-be/internal/app/country"
	"airbnb-user-be/internal/pkg/pagination"
)

type GetCountries struct {
	Countries  *[]module.Country     `json:"countries"`
	Pagination *pagination.SQLPaging `json:"pagination"`
}
