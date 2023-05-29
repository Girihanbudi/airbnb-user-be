package request

import "airbnb-user-be/internal/pkg/validator"

type ChangeCurrencySetting struct {
	UserId   *string
	Currency string `validation:"required"`
}

func (req *ChangeCurrencySetting) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
