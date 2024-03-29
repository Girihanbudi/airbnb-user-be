package request

import (
	"airbnb-user-be/internal/pkg/pagination"
	"airbnb-user-be/internal/pkg/validator"
)

type GetCountries struct {
	Pagination pagination.SQLPaging `json:"pagination"`
}

func (req *GetCountries) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
