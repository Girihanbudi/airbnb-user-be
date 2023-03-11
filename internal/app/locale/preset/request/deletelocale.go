package request

import (
	"airbnb-user-be/internal/pkg/validator"
)

type DeleteLocale struct {
	Code string `json:"code"`
}

func (req *DeleteLocale) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
