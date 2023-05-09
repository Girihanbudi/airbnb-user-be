package request

import "errors"

type CreateUser struct {
	User
}

func (req *CreateUser) Validate() (bool, error) {
	if req.CountryCode != nil && req.PhoneNumber != nil {
		return true, nil
	} else if req.Email != nil {
		return true, nil
	} else {
		return false, errors.New("parameter not found, need one of the following: phone number or email")
	}
}
