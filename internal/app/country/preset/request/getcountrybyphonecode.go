package request

import (
	"airbnb-user-be/internal/pkg/validator"
)

type GetCountryByPhoneCode struct {
	PhoneCode int `json:"phoneCode" validation:"required"`
}

func (req *GetCountryByPhoneCode) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
