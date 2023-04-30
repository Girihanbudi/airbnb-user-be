package response

import (
	"airbnb-user-be/internal/pkg/pagination"
)

type GetCountries struct {
	Countries  *[]Country            `json:"countries"`
	Pagination *pagination.SQLPaging `json:"pagination"`
}
