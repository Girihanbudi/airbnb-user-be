package request

import (
	"airbnb-user-be/internal/pkg/validator"
)

type GetLocale struct {
	Code string `json:"code" validate:"required"`
}

func (req *GetLocale) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
