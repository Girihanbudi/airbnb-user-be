package request

import (
	"airbnb-user-be/internal/pkg/pagination"
	"airbnb-user-be/internal/pkg/validator"
)

type GetLocales struct {
	Pagination pagination.SQLPaging `json:"pagination"`
}

func (req *GetLocales) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
