package request

import (
	"errors"
)

type GetUser struct {
	Id          *string `json:"id"`
	CountryCode *int    `json:"countryCode"`
	PhoneNumber *string `json:"phoneNumber"`
	Email       *string `json:"email"`
}

func (req *GetUser) Validate() (bool, error) {
	if req.Id != nil {
		return true, nil
	} else if req.CountryCode != nil && req.PhoneNumber != nil {
		return true, nil
	} else if req.Email != nil {
		return true, nil
	} else {
		return false, errors.New("parameter not found, need one of the following: id, phone number or email")
	}
}
