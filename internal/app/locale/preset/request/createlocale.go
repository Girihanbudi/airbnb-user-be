package request

import (
	"airbnb-user-be/internal/pkg/validator"
)

type CreateLocale struct {
	Code     string  `json:"code" validation:"required"`
	Name     string  `json:"name" validation:"required"`
	Local    *string `json:"local"`
	Location *string `json:"location"`
	Lcid     *int    `json:"lcid"`
	ISO639_2 *string `json:"iso639_2"`
	ISO639_1 *string `json:"iso639_1"`
}

func (req *CreateLocale) Validate() (bool, error) {
	err := validator.ValidateStruct(req)
	if err != nil {
		return false, err
	}

	return true, nil
}
