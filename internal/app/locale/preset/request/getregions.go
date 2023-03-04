package request

import (
	"airbnb-user-be/internal/pkg/pagination"
	"airbnb-user-be/internal/pkg/validator"
)

type GetRegions struct {
	Pagination pagination.SQLPaging `json:"pagination"`
}

func (req *GetRegions) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
