package request

import "airbnb-user-be/internal/pkg/validator"

type ChangeLocaleSetting struct {
	UserId *string
	Locale string `validation:"required"`
}

func (req *ChangeLocaleSetting) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
