package response

import module "airbnb-user-be/internal/app/country"

type GetCountries struct {
	Countries *[]module.Country `json:"countries"`
}
